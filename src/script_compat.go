package opennox

import (
	"image"
	"time"

	"github.com/opennox/libs/console"
	"github.com/opennox/libs/object"
	"github.com/opennox/libs/script"
	"github.com/opennox/libs/types"

	"github.com/opennox/opennox/v1/legacy"
	"github.com/opennox/opennox/v1/server"
)

// TODO: This is only for old LUA compatibility.

type scrPlayer struct {
	*server.Player
}

func (p scrPlayer) SetPos(pos types.Pointf) {
	noxServer.PlayerSetPos(p.Player, pos)
}

func (p scrPlayer) IsHost() bool {
	return noxServer.Players.IsHost(p.Player)
}

func (p scrPlayer) Print(text string) {
	noxServer.PlayerPrint(p.Player, text)
}

func (p scrPlayer) Blind(v bool) {
	noxServer.PlayerBlind(p.Player, v)
}

func (p scrPlayer) Cinema(v bool) {
	noxServer.PlayerCinema(p.Player, v)
}

func (p scrPlayer) Unit() script.Unit {
	if p.PlayerUnit == nil {
		return nil
	}
	return scrObject{asObjectS(p.PlayerUnit)}
}

func (p scrPlayer) GetObject() script.Object {
	u := p.Unit()
	if u == nil {
		return nil
	}
	return u
}

type scrObject struct {
	*Object
}

func (obj scrObject) LookAngle(angle int) {
	obj.LookWithAngle(angle)
}

func (obj scrObject) GetObject() script.Object {
	if obj.Object == nil {
		return nil
	}
	return obj
}

func (obj scrObject) SetOwner(owner script.ObjectWrapper) {
	if owner == nil {
		obj.Object.SetOwner(nil)
		return
	}
	own := server.ToObject(owner.GetObject().(server.Obj))
	obj.SObj().SetOwner(own)
}

func (obj scrObject) Owner() script.Object {
	if obj.Object == nil {
		return nil
	}
	p := obj.Object.Owner()
	if p == nil {
		return nil
	}
	return scrObject{asObjectS(p)}
}

func (obj scrObject) CanSee(obj2 script.Object) bool {
	if obj.Object == nil || obj2 == nil {
		return false
	}
	return obj.Object.CanSee(server.ToObject(obj2.(server.Obj)))
}

func (obj scrObject) Guard() {
	panic("implement me")
}

func (obj scrObject) Destroy() {
	panic("implement me")
	obj.Delete()
}

func (obj scrObject) Say(text string, dur script.Duration) {
	panic("implement me")
}

func (obj scrObject) Mute() {
	panic("implement me")
}

func (obj scrObject) Follow(targ script.Positioner) {
	obj.Object.Follow(targ)
}

func (obj scrObject) Flee(target script.Positioner, dt script.Duration) {
	obj.Object.Flee(target, dt)
}

func (obj scrObject) Attack(targ script.Positioner) {
	obj.Object.Attack(targ)
}

func (obj scrObject) OnUnitDeath(fnc func()) {
	obj.SObj().OnUnitDeath(fnc)
}

func (obj scrObject) OnUnitIdle(fnc func()) {
	obj.SObj().OnUnitIdle(fnc)
}

func (obj scrObject) OnUnitDone(fnc func()) {
	obj.SObj().OnUnitDone(fnc)
}

func (obj scrObject) OnUnitAttack(fnc func(targ script.Unit)) {
	obj.SObj().OnUnitAttack(fnc)
}

func (obj scrObject) OnUnitSeeEnemy(fnc func(targ script.Unit)) {
	obj.SObj().OnUnitSeeEnemy(fnc)
}

func (obj scrObject) OnUnitLostEnemy(fnc func(targ script.Unit)) {
	obj.SObj().OnUnitLostEnemy(fnc)
}

func (obj scrObject) OnTriggerActivate(fnc func(u script.Object)) {
	obj.SObj().OnTriggerActivate(fnc)
}

func (obj scrObject) OnTriggerDeactivate(fnc func()) {
	obj.SObj().OnTriggerDeactivate(fnc)
}

type noxScriptImpl struct {
	s *Server
}

func (s noxScriptImpl) Frame() int {
	return int(s.s.Frame())
}

func (s noxScriptImpl) Time() time.Duration {
	sec := float64(s.s.Frame()) / float64(s.s.TickRate())
	return time.Duration(sec) * time.Second
}

func (noxScriptImpl) BlindPlayers(blind bool) {
	BlindPlayers(blind)
}

func (noxScriptImpl) CinemaPlayers(v bool) {
	noxServer.CinemaPlayers(v)
}

func (s noxScriptImpl) Players() []script.Player {
	list := s.s.Players.List()
	out := make([]script.Player, 0, len(list))
	for _, p := range list {
		out = append(out, scrPlayer{p})
	}
	return out
}

func (noxScriptImpl) HostPlayer() script.Player {
	return scrPlayer{noxServer.Players.Host()}
}

func (s noxScriptImpl) OnPlayerJoin(fnc func(p script.Player)) {
	s.s.OnPlayerJoinLegacy(fnc)
}

func (s noxScriptImpl) OnPlayerLeave(fnc func(p script.Player)) {
	s.s.OnPlayerLeaveLegacy(fnc)
}

type noxScriptObjType struct {
	*server.ObjectType
}

func (t noxScriptObjType) CreateObject(p types.Pointf) script.Object {
	return scrObject{asObjectS(noxServer.createObject(t.ObjectType, p))}
}

func (s noxScriptImpl) ObjectTypeByID(id string) script.ObjectType {
	tp := s.s.Types.ByID(id)
	if tp == nil {
		return nil
	}
	return noxScriptObjType{tp}
}

func (s noxScriptImpl) ObjectByID(id string) script.Object {
	obj := s.s.Objs.GetObjectByID(id)
	if obj == nil {
		return nil
	}
	if obj.Class().HasAny(object.MaskUnits) {
		return scrObject{asObjectS(obj)}
	}
	return scrObject{asObjectS(obj)}
}

func (s noxScriptImpl) ObjectGroupByID(id string) *script.ObjectGroup {
	g := s.s.getObjectGroupByID(id)
	if g == nil {
		return nil
	}
	return g
}

func (s noxScriptImpl) WaypointByID(id string) script.Waypoint {
	w := s.s.WPs.ByID(id)
	if w == nil {
		return nil
	}
	return w
}

func (s noxScriptImpl) WaypointGroupByID(id string) *script.WaypointGroup {
	g := s.s.getWaypointGroupByID(id)
	if g == nil {
		return nil
	}
	return g
}

func (s noxScriptImpl) WallAt(pos types.Pointf) script.Wall {
	w := s.s.Walls.GetWallAt(pos)
	if w == nil {
		return nil
	}
	return asWallS(w)
}

func (s noxScriptImpl) WallNear(pos types.Pointf) script.Wall {
	w := s.s.Walls.GetWallNear(pos)
	if w == nil {
		return nil
	}
	return asWallS(w)
}

func (s noxScriptImpl) WallAtGrid(pos image.Point) script.Wall {
	w := s.s.Walls.GetWallAtGrid(pos)
	if w == nil {
		return nil
	}
	return asWallS(w)
}

func (s noxScriptImpl) WallGroupByID(id string) *script.WallGroup {
	return s.s.getWallGroupByID(id)
}

type scriptConsole console.Color

func (c scriptConsole) Print(text string) {
	noxConsole.Print(console.Color(c), text)
}

func (noxScriptImpl) Console(error bool) script.Printer {
	if error {
		return scriptConsole(console.ColorLightRed)
	}
	return scriptConsole(console.ColorYellow)
}

func (noxScriptImpl) AudioEffect(name string, pos script.Positioner) {
	// FIXME: trigger audio event
}

type scriptGlobalPrint struct{}

func (scriptGlobalPrint) Print(text string) {
	legacy.PrintToPlayers(text)
}

func (noxScriptImpl) Global() script.Printer {
	return scriptGlobalPrint{}
}
