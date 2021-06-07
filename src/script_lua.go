package main

import (
	"os"
	"path/filepath"
	"strings"

	"nox/v1/client/system/parsecmd"
	"nox/v1/common/maps"
	"nox/v1/server/script"
	"nox/v1/server/script/lua"
)

var (
	scriptLUA struct {
		curmap string
		vm     *lua.VM
	}
)

func init() {
	parseCmd.Register(&parsecmd.Command{
		Token: "lua", HelpID: "luahelp",
		Help:  "execute LUA command",
		Flags: parsecmd.Server | parsecmd.Cheat,
		Func: func(c *parsecmd.Console, tokens []string) bool {
			if len(tokens) == 0 {
				return false
			}
			if scriptLUA.vm == nil {
				c.Printf(parsecmd.ColorRed, "LUA is not running")
				return true
			}
			code := strings.Join(tokens, " ")
			if err := scriptLUA.vm.Exec(code); err != nil {
				c.Printf(parsecmd.ColorRed, "LUA error: %v", err)
			}
			return true
		},
	})
}

func luaScriptTick() {
	if scriptLUA.vm == nil {
		return
	}
	scriptLUA.vm.OnFrame()
}

func luaMaybeInitMap() {
	mp := nox_server_currentMapGetFilename_409B30()
	if mp == scriptLUA.curmap {
		return
	}
	luaShutdown()
	lua.Log.Printf("loading script for map %q", mp)
	scriptLUA.curmap = mp
	mp = strings.TrimSuffix(mp, maps.Ext)
	mapDir := getDataPath(maps.Dir, mp)
	scriptLUA.vm = lua.NewVM(noxScript{}, mapDir)
	lname := mp + ".lua"
	err := scriptLUA.vm.ExecFile(filepath.Join(mapDir, lname))
	if os.IsNotExist(err) {
		lua.Log.Printf("no lua script for map: %q", lname)
		scriptLUA.vm.InitDefault()
		return
	} else if err != nil {
		lua.Log.Printf("error opening script %q: %v (%T)", filepath.Join(maps.Dir, mp, lname), err, err)
		return
	}
	lua.Log.Printf("map script loaded: %q", lname)
}

func luaShutdown() {
	if scriptLUA.vm != nil {
		lua.Log.Printf("stopping script for map %q", scriptLUA.curmap)
		_ = scriptLUA.vm.Close()
		scriptLUA.vm = nil
	}
}

func luaOnEvent(event script.EventType) {
	switch event {
	case script.EventMapInitialize,
		script.EventMapEntry:
		luaMaybeInitMap()
	}
	if scriptLUA.vm != nil {
		scriptLUA.vm.OnEvent(event)
	}
	switch event {
	// TODO: change to EventMapShutdown and make sure it triggers
	//       actually, EventMapShutdown is called when saving game when the map _isn't_ shutting down
	//       so probably worth adding a new event that triggers at the right time
	case script.EventMapExit:
		luaShutdown()
	}
}