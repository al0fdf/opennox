package nox

/*
#include "GAME1_3.h"
#include "GAME3_2.h"
extern uint32_t dword_5d4594_1563080;
extern uint32_t dword_5d4594_1563084;
extern uint32_t dword_5d4594_1563088;
extern uint32_t dword_5d4594_1563092;
extern uint32_t dword_5d4594_1563096;
extern unsigned int dword_5d4594_251744;
*/
import "C"
import (
	noxflags "nox/v1/common/flags"
)

//export sub_413A00
func sub_413A00(a1 C.int) {
	if noxflags.HasGame(2048) {
		if a1 != 0 {
			noxflags.SetGame(0x40000)
		} else {
			if C.dword_5d4594_251744 == 0 {
				noxflags.UnsetGame(0x40000)
				nox_ticks_reset_416D40()
			}
		}
	}
}

//export sub_448640
func sub_448640() { C.sub_44A400() }

//export sub_4DB170
func sub_4DB170(a1, a2, a3 C.int) {
	C.dword_5d4594_1563092 = C.uint(a3)
	C.dword_5d4594_1563088 = C.uint(gameFrame())
	C.dword_5d4594_1563084 = C.uint(a2)
	C.dword_5d4594_1563080 = C.uint(a1)
	C.dword_5d4594_1563096 = C.uint(bool2int(a2 != 0))
	if a1 == 0 {
		C.sub_4DCBD0(0)
	}
}