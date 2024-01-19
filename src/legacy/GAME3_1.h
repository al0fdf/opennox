#ifndef NOX_PORT_GAME3_1
#define NOX_PORT_GAME3_1

#include "defs.h"

int sub_4B9470(const char** a1);
int sub_4B94E0(nox_drawable* dr);
uint32_t* sub_4B95D0(nox_drawable* dr);
uint32_t* sub_4B9650(int a1);
uint32_t* sub_4B96F0(nox_drawable* dr);
void sub_4BA670(int a1, int a2, int a3, int a4, int a5);
int nox_xxx_prepareLightningEffects_4BAB30();
int sub_4BC720(int a1);
int nox_xxx_updDrawMonsterGen_4BC920();
uint32_t* sub_4BD280(int a1, int a2);
void sub_4BD2D0(void* lpMem);
uint32_t* sub_4BD2E0(uint32_t** a1);
int sub_4BD300(uint32_t* a1, int a2);
uint32_t* sub_4BD340(int a1, int a2, int a3, int a4);
void sub_4BD3C0(void* lpMem);
uint32_t* sub_4BD420(int a1, int a2);
uint32_t* sub_4BD470(uint32_t** a1, int a2);
int sub_4BD600(int a1);
int sub_4BD650(int a1);
int sub_4BD660(int a1);
int sub_4BD680(int a1);
int sub_4BD690(int a1);
int sub_4BD710(int a1);
uint32_t* sub_4BD720(int a1);
void sub_4BD7A0(void* lpMem);
uint32_t* sub_4BD7C0(uint32_t* a1);
int sub_4BD840(int a3);
int sub_4BD8C0(int a1);
int sub_4BD940(int a1);
int sub_4BD9B0(uint32_t* a2);
void sub_4BDA60(void* lpMem);
int sub_4BDA80(int a1);
int sub_4BDB20(int a1);
int sub_4BDB30(int a1);
int sub_4BDB40(int a2);
void sub_4BDB90(uint32_t* a1, uint32_t* a2);
int sub_4BDC00(int a1);
int nox_xxx_loadAdvancedWnd_4BDC10(int* a1);
int sub_4BDC70(int* a1);
int sub_4BDD10();
int sub_4BDDA0();
int nox_xxx_windowAdvancedServProc_4BDDB0(int a1, int a2, int* a3, int a4);
int sub_4BDF30();
int sub_4BDF70(int* a1);
int sub_4BDF90(int* a1);
int sub_4BDFD0();
int sub_4BE120(int a1);
int sub_4BE320();
int sub_4BE330(int a1, unsigned int a2, int* a3, int a4);
int sub_4BE610();
void nox_video_drawAnimatedImageOrCursorAt_4BE6D0(int a1, int a2, int a3);
int sub_4BE800(int a1);
char sub_4BE810(int a1, int a2, int a3, char a4);
void sub_4BEAD0(int2* a1, int2* a2, int2* a3, int2* a4, int a5, int a6);
void sub_4BEDE0(int2* a1, int2* a2, int2* a3, int2* a4, int a5, float a6, int a7, int a8);
int nox_xxx_clientReportSecondaryWeapon_4BF010(int a1);
short sub_4BF7E0(uint32_t* a1);
short sub_4BF9F0(int a1, int a2, int a3, int a4, int a5, int a6, int a7);
int sub_4BFAD0();
void sub_4BFB70(int a1);
void sub_4BFBB0(uint32_t* a1);
int sub_4BFBF0();
int sub_4BFC70();
int sub_4BFC90();
int sub_4BFCD0(int a1, int a2, int* a3, int a4);
void sub_4BFD10();
int sub_4BFD30();
void sub_4BFD40();
int sub_4BFDD0(uint32_t* a1, int a2, unsigned int a3);
int sub_4BFE40();
int nox_gui_itemAmount_init_4BFEF0();
int sub_4C0030(int a1);
int sub_4C01C0(int a1, int a2, int* a3, int a4);
void nox_gui_itemAmount_free_4C03E0();
int nox_gui_itemAmountDialog_4C0430(wchar2_t* title, int x, int y, int a4, int a5, const void* a6, int a7, int a8,
									void* accept, void* cancel);
int sub_4C0560(int a1, int a2);
int sub_4C05F0(int a1, int a2);
int nox_xxx_func_4C0610();
int sub_4C0630(int a1, unsigned int a2, unsigned int a3);
int nox_xxx_clientTrade_0_4C08E0(int a1);
char sub_4C0910(int2* a1);
int sub_4C0C90(int a1, int a2, int* a3, int a4);
int nox_xxx_clientTrade_4C0CE0();
int sub_4C0D00();
int sub_4C1120(int a1, int a2, wchar2_t** a3);
char sub_4C11E0(uint32_t* a1);
int nox_xxx_closeP2PTradeWnd_4C12A0();
int sub_4C12C0();
int nox_xxx_showP2PTradeWnd_4C12D0();
int nox_xxx_netP2PStartTrade_4C1320(int a1);
int sub_4C1410();
int sub_4C1590();
int sub_4C1710(int a1, int a2);
int sub_4C1760(int a1, int a2);
char* nox_xxx_tradeClientAddItem_4C1790(int a1);
int sub_4C18E0(int a1, uint32_t* a2);
char* sub_4C1910(int a1);
char* sub_4C19C0(int a1);
int sub_4C1B50(int a1);
int sub_4C1BC0(int a1);
int nox_xxx_prepareP2PTrade_4C1BF0();
int sub_4C1CA0(int a1);
int nox_xxx_guiDrawSummonBox_4C1FE0(uint32_t* a1);
int nox_xxx_wndSummonGet_4C2410(int2* a1);
int nox_xxx_guiDrawSummon_4C2440(int a1);
int nox_xxx_guiHideSummonWindow_4C2470();
int sub_4C24A0();
int nox_xxx_wndSummonBigButtonProc_4C24B0(int a1, int a2, unsigned int a3);
int sub_4C26F0(void* yTop);
int sub_4C2A00(int a1, int a2, int a3, int a4, short* a5);
int nox_xxx_clientOrderCreature_4C2A60(int a1, unsigned int a2);
int nox_xxx_wndSummonProc_4C2B10(uint32_t* a1, unsigned int a2, unsigned int a3);
int sub_4C2BD0();
int sub_4C2BE0();
int* sub_4C2BF0();
int sub_4C2C20(uint32_t* a1, int a2, unsigned int a3);
int sub_4C2C60(uint32_t* a1, int2* a2);
char* sub_4C2D60();
char* sub_4C2D90(int a1);
int sub_4C2DD0(int a1);
int sub_4C2E00();
char nox_xxx_cliSummonCreat_4C2E50(int a1, int a2, int a3);
int sub_4C2EF0(int a1);
char* sub_4C2F20();
char* sub_4C2F70();
int sub_4C2FD0(int a1);
int sub_4C3030(int* a1, int a2, int a3);
int sub_4C30C0(int* a1, int a2);
void nox_xxx_cliSummonOnDieOrBanish_4C3140(int a1, void* a2);
char* sub_4C31D0(int a1);
int sub_4C3210(int a1);
int nox_xxx_sprite_4C3220(nox_drawable* a1);
int sub_4C3260();
void nox_video_drawCircleColored_4C3270(int a1, int a2, int a3, int a4);
int nox_xxx_spriteDrawCircleMB_4C32A0(int a1, int a2, int a3, int a4);
int sub_4C3390();
int sub_4C3410(int* a1);
int sub_4C3460(int a1);
int sub_4C34A0();
int sub_4C3500();
int sub_4C35B0(int a1);
int sub_4C3A60(uint32_t* a1, unsigned int a2, unsigned int a3, int a4);
int sub_4C3A90(int a1, int a2, int* a3, int a4);
void sub_4C3B70();
int sub_4C3EB0(int a1, int a2, unsigned int a3, int a4);
int sub_4C3FC0(unsigned int a1);
int sub_4C4100(unsigned int a1);
int sub_4C4220();
void sub_4C4260();
int sub_4C4280();
int sub_4C42A0(int2* a1, int2* a2, int* a3, int* a4);
void nox_xxx_drawObject_4C4770_draw(nox_draw_viewport_t* vp, nox_drawable* dr, void* img);
char sub_4C4EC0(uint32_t* a1, int a2);
short nox_xxx_drawShinySpot_4C4F40(nox_draw_viewport_t* vp, nox_drawable* dr);
int nox_xxx_colorInit_4C4FD0();
int sub_4C5020(int a1);
void sub_4C5050();
int sub_4C5060(nox_draw_viewport_t* a1p);
int sub_4C51D0(int2* a1, int2* a2);
int sub_4C5630(int a1, int a2, int a3);
int nox_xxx_sprite_4CA540(uint32_t* a1, int a2);
int sub_4CA650(int a1, int a2);
int sub_4CA720(int a1, int a2);
nox_window* nox_gui_newProgressBar_4CAF10(int a1, int a2, int a3, int a4, int a5, int a6, uint32_t* a7);
int sub_4CAF80(int a1, int a2, int a3, int a4);
int sub_4CAFB0(int a1);
int sub_4CAFF0(uint32_t* a1, uint32_t* a2);
int sub_4CB1A0(uint32_t* a1, int a2);
void nox_client_advVideoOptsLoad_4CB330();
int nox_client_advVideoOpts_New_4CB590(nox_window* a1);
int sub_4CB880();
int sub_4CBB70();
int sub_4CBBB0();
void sub_4CBBF0();
int sub_4CBE70(int a1, int a2, int* a3, int a4);
int sub_4CC140(uint32_t* a1, unsigned int a2, unsigned int a3, int a4);
int sub_4CC170(int a1, int a2, char* a3, int a4);
int sub_4CC280(unsigned int a1);
int sub_4CC3C0(unsigned int a1);
int nox_xxx_updDrawUndeadKiller_4CCCF0();
int sub_4CCD00(int a1, int a2);
int nox_xxx_updDrawFist_4CCDB0(int a1, int a2);
int sub_4CCE70(int a1, uint32_t* a2);
int sub_4CD090(int a1, uint32_t* a2);
int sub_4CD0C0(int a1, uint32_t* a2);
int sub_4CD0F0(int a1, uint32_t* a2);
int sub_4CD120(int a1, uint32_t* a2);
int sub_4CD400(uint32_t* a1, int a2);

#endif // NOX_PORT_GAME3_1
