package legacy

/*
#include "GAME4_3.h"

static int nox_call_objectType_parseUse_go(int (*fnc)(char*, void*), char* arg1, void* arg2) { return fnc(arg1, arg2); }
*/
import "C"
import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/opennox/opennox/v1/server"
)

var (
	Nox_xxx_useConsume_53EE10      func(obj1, obj2 *server.Object) int
	Nox_xxx_useMushroom_53ECE0     func(obj1, obj2 *server.Object) int
	Nox_xxx_useCiderConfuse_53EF00 func(obj1, obj2 *server.Object) int
	Nox_xxx_useEnchant_53ED60      func(obj1, obj2 *server.Object) int
	Nox_xxx_useCast_53ED90         func(obj1, obj2 *server.Object) int
	Nox_xxx_usePotion_53EF70       func(obj1, obj2 *server.Object) int
)

func init() {
	server.RegisterObjectUse("ConsumeUse", C.nox_xxx_useConsume_53EE10, unsafe.Sizeof(server.ConsumeUseData{}))
	server.RegisterObjectUse("ConsumeConfuseUse", C.nox_xxx_useCiderConfuse_53EF00, unsafe.Sizeof(server.ConsumeUseData{}))
	server.RegisterObjectUse("CastUse", C.nox_xxx_useCast_53ED90, unsafe.Sizeof(server.CastUseData{}))
	server.RegisterObjectUse("EnchantUse", C.nox_xxx_useEnchant_53ED60, unsafe.Sizeof(server.EnchantUseData{}))
	server.RegisterObjectUse("MushroomUse", C.nox_xxx_useMushroom_53ECE0, 0)
	server.RegisterObjectUse("PotionUse", C.nox_xxx_usePotion_53EF70, unsafe.Sizeof(server.PotionUseData{}))

	server.RegisterObjectUse("FireWandUse", C.nox_xxx_useFireWand_53F670, 0)
	server.RegisterObjectUse("ReadUse", C.nox_xxx_useRead_53F7C0, 260)
	server.RegisterObjectUse("WarpReadUse", C.sub_53F830, 260)
	server.RegisterObjectUse("WandUse", C.nox_xxx_useLesserFireballStaff_53F290, 116)
	server.RegisterObjectUse("WandCastUse", C.nox_xxx_useWandCastSpell_53F4F0, 116)
	server.RegisterObjectUse("SpellRewardUse", C.nox_xxx_useSpellReward_53F9E0, 1)
	server.RegisterObjectUse("AbilityRewardUse", C.nox_xxx_useAbilityReward_53FAE0, 1)
	server.RegisterObjectUse("FieldGuideUse", C.sub_53F930, 64)

	server.RegisterObjectUseParse("WandUse", wrapObjectUseParseC(C.sub_536260))
	server.RegisterObjectUseParse("WandCastUse", wrapObjectUseParseC(C.sub_5361B0))
}

func wrapObjectUseParseC(ptr unsafe.Pointer) server.ObjectParseFunc {
	return func(objt *server.ObjectType, args []string) error {
		if Nox_call_objectType_parseUse_go(ptr, strings.Join(args, " "), objt.UseData) == 0 {
			return fmt.Errorf("cannot parse use data for %q", objt.ID())
		}
		return nil
	}
}

func Nox_call_objectType_parseUse_go(a1 unsafe.Pointer, a2 string, a3 unsafe.Pointer) int {
	cstr := CString(a2)
	defer StrFree(cstr)
	return int(C.nox_call_objectType_parseUse_go((*[0]byte)(a1), cstr, a3))
}

//export nox_xxx_useMushroom_53ECE0
func nox_xxx_useMushroom_53ECE0(cobj1 *nox_object_t, cobj2 *nox_object_t) int {
	return Nox_xxx_useMushroom_53ECE0(asObjectS(cobj1), asObjectS(cobj2))
}

//export nox_xxx_useCiderConfuse_53EF00
func nox_xxx_useCiderConfuse_53EF00(cobj1 *nox_object_t, cobj2 *nox_object_t) int {
	return Nox_xxx_useCiderConfuse_53EF00(asObjectS(cobj1), asObjectS(cobj2))
}

//export nox_xxx_useEnchant_53ED60
func nox_xxx_useEnchant_53ED60(cobj1 *nox_object_t, cobj2 *nox_object_t) int {
	return Nox_xxx_useEnchant_53ED60(asObjectS(cobj1), asObjectS(cobj2))
}

//export nox_xxx_useCast_53ED90
func nox_xxx_useCast_53ED90(cobj1 *nox_object_t, cobj2 *nox_object_t) int {
	return Nox_xxx_useCast_53ED90(asObjectS(cobj1), asObjectS(cobj2))
}

//export nox_xxx_useConsume_53EE10
func nox_xxx_useConsume_53EE10(cobj1 *nox_object_t, cobj2 *nox_object_t) int {
	return Nox_xxx_useConsume_53EE10(asObjectS(cobj1), asObjectS(cobj2))
}

//export nox_xxx_usePotion_53EF70
func nox_xxx_usePotion_53EF70(cobj1 *nox_object_t, cobj2 *nox_object_t) int {
	return Nox_xxx_usePotion_53EF70(asObjectS(cobj1), asObjectS(cobj2))
}
