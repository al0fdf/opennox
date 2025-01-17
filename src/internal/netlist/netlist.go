package netlist

import (
	"github.com/opennox/libs/common"

	"github.com/opennox/opennox/v1/common/ntype"
)

const (
	bufSize    = 2048
	maxPackets = 512
	hostIndex  = common.MaxPlayers - 1
)

type Kind byte

const (
	Kind0 = Kind(iota)
	Kind1
	Kind2
)

func New() *List {
	return &List{
		buffersList: [3][]buffer{
			make([]buffer, common.MaxPlayers),
			make([]buffer, common.MaxPlayers),
			make([]buffer, common.MaxPlayers),
		},
	}
}

type List struct {
	buffersList [3][]buffer
	messageList [3][common.MaxPlayers]*MsgList
}

type buffer struct {
	buf [bufSize]byte
	cur int
}

func (s *List) Init() {
	for i := 0; i < common.MaxPlayers; i++ {
		s.messageList[0][i] = nil
		s.messageList[1][i] = newMsgList()
		s.messageList[2][i] = newMsgList()
	}
	s.messageList[0][hostIndex] = newMsgList()
}

func (s *List) Free() {
	for i := 0; i < common.MaxPlayers; i++ {
		for j := 0; j < 3; j++ {
			if l := s.messageList[j][i]; l != nil {
				l.Free()
				s.messageList[j][i] = nil
			}
		}
	}
}

func (s *List) checkSizesA(ind ntype.PlayerInd, kind Kind, sz int) bool {
	if kind == Kind1 {
		psz := s.ByInd(ind, Kind1).Size()
		if psz+sz+s.ByInd(ind, Kind2).Size() > bufSize {
			return false
		}
	} else {
		if sz+s.ByInd(ind, kind).Size() > bufSize {
			return false
		}
	}
	return s.ByInd(ind, kind).Count() < maxPackets
}

func (s *List) checkSizesB(ind ntype.PlayerInd, kind Kind, sz, sz2 int) bool {
	l := s.ByInd(ind, kind)
	psz := l.Size()
	if psz+sz+sz2 > bufSize {
		return false
	}
	return l.Count() < maxPackets
}

func (s *List) checkSizesC(ind ntype.PlayerInd, sz int) bool {
	l := s.ByInd(ind, Kind2)
	if (sz + l.Size()) > bufSize {
		return false
	}
	return l.Count() < maxPackets
}

func (s *List) ResetByInd(ind ntype.PlayerInd, kind Kind) {
	l := s.ByInd(ind, kind)
	if l == nil {
		return
	}
	l.Reset()
	s.buffersList[kind+1][ind].cur = 0
}

func (s *List) InitByInd(ind ntype.PlayerInd) {
	l := s.ByInd(ind, Kind2)
	l.Reset()
	s.buffersList[0][ind].cur = 0
}

func (s *List) ResetAllInd(kind Kind) {
	for i := ntype.PlayerInd(0); i < common.MaxPlayers; i++ {
		s.ResetByInd(i, kind)
	}
}

func (s *List) ResetAll() {
	for i := ntype.PlayerInd(0); i < common.MaxPlayers; i++ {
		s.ResetByInd(i, Kind1)
		s.ResetByInd(i, Kind0)
		s.InitByInd(i)
	}
}

func (s *List) allocBufferRaw(ind ntype.PlayerInd, kind Kind, buf []byte) []byte {
	if len(buf) == 0 {
		return nil
	}
	p := &s.buffersList[kind][ind]
	i := p.cur
	if i+len(buf) > bufSize {
		return nil
	}
	copy(p.buf[i:], buf)
	p.cur += len(buf)
	return p.buf[i : i+len(buf)]
}

func (s *List) allocBuffer(ind ntype.PlayerInd, kind Kind, buf []byte) []byte {
	return s.allocBufferRaw(ind, kind+1, buf)
}

func (s *List) allocBuffer0(ind ntype.PlayerInd, buf []byte) []byte {
	return s.allocBufferRaw(ind, Kind0, buf)
}

func (s *List) AddToMsgListCli(ind ntype.PlayerInd, kind Kind, buf []byte) bool {
	l := s.ByInd(ind, kind)
	if len(buf) == 0 {
		return true
	}
	if !s.checkSizesA(ind, kind, len(buf)) {
		return false
	}
	out := s.allocBuffer(ind, kind, buf)
	if out == nil {
		return false
	}
	l.add(out[:len(buf)], true)
	return true
}

func (s *List) CopyPacketsA(ind ntype.PlayerInd, kind Kind) []byte {
	list := s.ByInd(ind, kind)
	out := make([]byte, 0, bufSize)
	for {
		buf := list.Get()
		if len(buf) == 0 {
			return out
		} else if len(out)+len(buf) > cap(out) {
			// TODO: is it okay that it discards data?
			return out
		}
		out = append(out, buf...)
	}
}

func (s *List) HandlePacketsA(ind ntype.PlayerInd, kind Kind, fnc func(data []byte)) {
	list := s.ByInd(ind, kind)

	out := make([]byte, bufSize)
	out = out[:0]
	for {
		buf := list.Get()
		if len(buf) == 0 {
			break
		} else if len(out)+len(buf) > cap(out) {
			// TODO: is it okay that it discards data?
			break
		}
		out = append(out, buf...)
	}

	fnc(out)

	s.ResetByInd(ind, kind)
}

func (s *List) CopyPacketsB(ind ntype.PlayerInd) []byte {
	l := s.ByInd(ind, Kind2)
	cnt := 0
	sbuf := make([]byte, bufSize)
	for {
		src := l.Get()
		if len(src) == 0 {
			break // no more data
		}
		if cnt+len(src) > len(sbuf) {
			// we cannot store it, so put it back
			l.add(src, false)
			break
		}
		copy(sbuf[cnt:], src)
		cnt += len(src)
	}
	return sbuf[:cnt]
}

func (s *List) ClientSend0(ind ntype.PlayerInd, kind Kind, buf []byte, sz2 int) bool {
	l := s.ByInd(ind, kind)
	if len(buf) == 0 {
		return true
	}
	if !s.checkSizesB(ind, kind, len(buf), sz2) {
		return false
	}
	out := s.allocBuffer(ind, kind, buf)
	if out == nil {
		return false
	}
	l.add(out[:len(buf)], true)
	return true
}

func (s *List) AddToMsgListSrv(ind ntype.PlayerInd, buf []byte, flush func(ind ntype.PlayerInd)) bool {
	if len(buf) == 0 {
		return true
	}
	// If there are too many updates, then we may run out of space in a single
	// packet. Instead of fragmenting, we can instead send additional packets.
	var out []byte
	l := s.ByInd(ind, Kind2)
	if !s.checkSizesC(ind, len(buf)) || !(func() bool { out = s.allocBuffer0(ind, buf); return out != nil }()) {
		p := &s.buffersList[0][ind]

		// The new update packet needs to have correct bytes at the
		// beginning. Save the length of the first two queued datas so we
		// can replay them.
		len1 := len(l.first.buf)
		len2 := len(l.first.next.buf)

		// Flush old data to network.
		flush(ind)

		// Set buffer length and re-queue updates.
		p.cur = len1 + len2
		l.add(p.buf[0:len1], true)
		l.add(p.buf[len1:len2], true)

		// Retry original allocation.
		out = s.allocBuffer0(ind, buf)
	}
	if out == nil {
		return false
	}
	l.add(out[:len(buf)], true)
	return true
}

func (s *List) ByInd(ind ntype.PlayerInd, kind Kind) *MsgList {
	return s.messageList[kind][ind]
}

func newMsgList() *MsgList {
	return &MsgList{}
}

type netListItem struct {
	buf  []byte
	prev *netListItem
	next *netListItem
}

type MsgList struct {
	first *netListItem
	last  *netListItem
	count int
	size  int
}

func (l *MsgList) Free() {
	*l = MsgList{}
}

func (l *MsgList) Reset() {
	*l = MsgList{}
}

func (l *MsgList) freeItem(item *netListItem) {
	*item = netListItem{}
}

func (l *MsgList) Count() int {
	if l == nil {
		return 0
	}
	return l.count
}

func (l *MsgList) Size() int {
	if l == nil {
		return 0
	}
	return l.size
}

func (l *MsgList) Get() []byte {
	if l == nil || l.first == nil {
		return nil
	}
	item := l.first

	l.count--
	l.size -= len(item.buf)

	if next := item.next; next != nil {
		next.prev = item.prev
	} else {
		l.last = item.prev
	}
	if prev := item.prev; prev != nil {
		prev.next = item.next
	} else {
		l.first = item.next
	}

	buf := item.buf
	l.freeItem(item)
	return buf
}

func (l *MsgList) add(buf []byte, appnd bool) bool {
	it := &netListItem{buf: buf}
	if appnd {
		it.prev = l.last
		it.next = nil
		if l.last != nil {
			l.last.next = it
		} else {
			l.first = it
		}
		l.last = it
	} else {
		it.prev = nil
		it.next = l.first
		if l.first != nil {
			l.first.prev = it
		} else {
			l.last = it
		}
		l.first = it
	}
	l.count++
	l.size += len(buf)
	return true
}

func (l *MsgList) FindAndFreeBuf(buf []byte) {
	if l.first == nil {
		return
	}
	buf = buf[:1]
	var item *netListItem
	for p := l.first; p != nil; p = p.next {
		if &p.buf[0] == &buf[0] {
			item = p
			break
		}
	}
	if item == nil {
		return
	}
	l.count--
	l.size -= len(item.buf)

	if next := item.next; next != nil {
		next.prev = item.prev
	} else {
		l.last = item.prev
	}

	if prev := item.prev; prev != nil {
		prev.next = item.next
	} else {
		l.first = item.next
	}
	*item = netListItem{}
}

func (l *MsgList) Each(fnc func(b []byte) bool) {
	if fnc == nil {
		return
	}
	for p := l.first; p != nil; p = p.next {
		if fnc(p.buf) {
			break
		}
	}
}
