package client

import (
	"errors"

	"github.com/opennox/libs/noxnet"
	"github.com/opennox/libs/player"

	noxflags "github.com/opennox/opennox/v1/common/flags"
	"github.com/opennox/opennox/v1/common/ntype"
	"github.com/opennox/opennox/v1/internal/netlist"
	"github.com/opennox/opennox/v1/internal/netstr"
	"github.com/opennox/opennox/v1/server"
)

var (
	ErrLobbyNoSocket = errors.New("no broadcast socket")
)

type CurPlayerInfo struct {
	Connected bool
	Class     player.Class
}

func (c *Client) Nox_xxx_netSendClientReady_43C9F0() int {
	var data [1]byte
	data[0] = byte(noxnet.MSG_CLIENT_READY)
	c.Conn.SendReliable(data[:])
	return 1
}

func (c *Client) Nox_xxx_netKeepAliveSocket_43CA20() int {
	var data [1]byte
	data[0] = byte(noxnet.MSG_KEEP_ALIVE)
	c.Conn.SendUnreliable(data[:], true)
	return 1
}

func (c *Client) Nox_xxx_netRequestMap_43CA50() int {
	var data [1]byte
	data[0] = byte(noxnet.MSG_REQUEST_MAP)
	c.Conn.SendReliable(data[:])
	return 1
}

func (c *Client) Nox_xxx_netMapReceived_43CA80() int {
	var data [1]byte
	data[0] = byte(noxnet.MSG_RECEIVED_MAP)
	c.Conn.SendReliable(data[:])
	return 1
}

func (c *Client) Nox_xxx_cliSendCancelMap_43CAB0() int {
	var data [1]byte
	data[0] = byte(noxnet.MSG_CANCEL_MAP)
	seq, _ := c.Conn.SendReliable(data[:])
	if c.Conn.WaitServerResponse(seq, 20, netstr.RecvNoHooks|netstr.RecvJustOne) != 0 {
		return 0
	}
	c.Server.NetList.ResetByInd(server.HostPlayerIndex, netlist.Kind0)
	return 1
}

func (c *Client) Nox_xxx_netSendIncomingClient_43CB00() int {
	var data [1]byte
	data[0] = byte(noxnet.MSG_INCOMING_CLIENT)
	seq, _ := c.Conn.SendReliable(data[:])
	if c.Conn.WaitServerResponse(seq, 20, netstr.RecvNoHooks|netstr.RecvJustOne) != 0 {
		return 0
	}
	c.Server.NetList.ResetByInd(server.HostPlayerIndex, netlist.Kind0)
	return 1
}

func (c *Client) Nox_xxx_cliSendOutgoingClient_43CB50() int {
	var data [1]byte
	data[0] = byte(noxnet.MSG_OUTGOING_CLIENT)
	seq, _ := c.Conn.SendReliable(data[:])
	if c.Conn.WaitServerResponse(seq, 20, netstr.RecvNoHooks|netstr.RecvJustOne) != 0 {
		return 0
	}
	c.Conn.RecvLoop(true)
	c.Server.NetList.ResetByInd(server.HostPlayerIndex, netlist.Kind0)
	return 1
}

func (c *Client) OnClientPacketOpSub(pli ntype.PlayerInd, op noxnet.Op, data []byte, localFrame *uint32, localFrame16 *uint16, cur CurPlayerInfo) (int, bool, error) {
	switch op {
	case noxnet.MSG_SERVER_CLOSE_ACK:
		return 1, true, nil
	case noxnet.MSG_STAT_MULTIPLIERS:
		var p noxnet.MsgStatMult
		n, err := p.Decode(data[1:])
		if err != nil {
			return 0, false, err
		}
		c.Server.OnClassStats(cur.Class, server.ClassStats{
			Health:   p.Health,
			Mana:     p.Mana,
			Strength: p.Strength,
			Speed:    p.Speed,
		})
		return 1 + n, true, nil
	case noxnet.MSG_DESTROY_WALL:
		var p noxnet.MsgWallDestroy
		n, err := p.Decode(data[1:])
		if err != nil {
			return 0, false, err
		}
		if cur.Connected && !noxflags.HasGame(noxflags.GameHost) {
			c.Server.Walls.BreakByID(p.ID)
		}
		return 1 + n, true, nil
	case noxnet.MSG_FX_JIGGLE:
		var p noxnet.MsgFxJiggle
		n, err := p.Decode(data[1:])
		if err != nil {
			return 0, false, err
		}
		if cur.Connected {
			c.Viewport().Jiggle12 = int(p.Val) / 3
		}
		return 1 + n, true, nil
	case noxnet.MSG_XFER_MSG:
		var p noxnet.MsgXfer
		n, err := p.Decode(data[1:])
		if err != nil {
			return 0, false, err
		}
		c.Server.NetXfer.Handle(server.XferConn{c.Conn}, c.Server.FrameTS(), p.Msg)
		return 1 + n, true, nil
	}
	return 0, false, nil
}
