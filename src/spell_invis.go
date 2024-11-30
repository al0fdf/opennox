package opennox

import (
	"github.com/opennox/libs/noxnet"
	"github.com/opennox/libs/spell"

	"github.com/opennox/opennox/v1/server"
)

func castInvisibility(spellID spell.ID, _, _, _ *server.Object, args *server.SpellAcceptArg, lvl int) int {
	return castBuffSpell(spellID, server.ENCHANT_INVISIBLE, lvl, args.Obj, spellBuffConf{
		DurOpt: "InvisibilityEnchantDuration", Defensive: true, PointFX: noxnet.MSG_FX_SMOKE_BLAST,
	})
}
