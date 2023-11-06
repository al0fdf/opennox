#ifndef NOX_PORT_GAME4
#define NOX_PORT_GAME4

#include "defs.h"

int nox_xxx_XFerSpellReward_4F5F30(int* a1);
int nox_xxx_XFerAbilityReward_4F6240(int* a1);
int nox_xxx_XFerFieldGuide_4F6390(int* a1);
int nox_xxx_XFerWeapon_4F64A0(int a1);
int nox_xxx_XFerArmor_4F6860(int a1);
int nox_xxx_XFerAmmo_4F6B20(int* a1);
int nox_xxx_XFerTeam_4F6D20(int* a1);
int nox_xxx_XFerGold_4F6EC0(int a1);
int nox_xxx_XFerObelisk_4F6F60(int* a1);
int nox_xxx_XFerToxicCloud_4F70A0(int a1);
int nox_xxx_XFerMonsterGen_4F7130(int* a1);
int nox_xxx_XFerRewardMarker_4F74D0(int* a1);
int nox_xxx_equipedItemByCode_4F7920(int a1, int a2);
void sub_4F7950(nox_object_t* a1);
void nox_xxx_playerSetCustomWP_4F79A0(int a1, int a2, int a3);
int nox_xxx_playerConfusedGetDirection_4F7A40(nox_object_t* a1);
void nox_xxx_mapFindPlayerStart_4F7AB0(float2* a1, nox_object_t* a2p);
int sub_4F7CE0(int a1, int a2);
int nox_xxx_playerSubStamina_4F7D30(nox_object_t* a1, int a2);
void sub_4F7DB0(int a1, char a2);
int nox_xxx_checkWinkFlags_4F7DF0(nox_object_t* a1);
int nox_xxx_weaponGetStaminaByType_4F7E80(int a1);
short nox_xxx_playerRespawn_4F7EF0(nox_object_t* a1);
int sub_4F80C0(int a1, float2* a3);
void nox_xxx_updatePlayer_4F8100(nox_object_t* a1);
int sub_4F9A80(nox_object_t* a1);
int sub_4F9AB0(nox_object_t* a1);
int nox_xxx_playerCanMove_4F9BC0(nox_object_t* a1);
int nox_xxx_playerCanAttack_4F9C40(nox_object_t* a1);
void nox_xxx_playerInputAttack_4F9C70(nox_object_t* a1);
int nox_xxx_playerAimsAtEnemy_4F9DC0(int a1);
int sub_4F9E10(nox_object_t* a1);
void nox_xxx_animPlayerGetFrameRange_4F9F90(int a1, int* a2, int* a3);
int nox_xxx_unitGetStrength_4F9FD0(int a1);
int nox_xxx_playerSetState_4FA020(nox_object_t* a1, int a2);
int sub_4FA280(int a1);
int nox_common_mapPlrActionToStateId_4FA2B0(nox_object_t* a1);
int nox_xxx_checkInversionEffect_4FA4F0(int a1, int a2);
uint32_t* nox_xxx_playerAddGold_4FA590(int a1, int a2);
uint32_t* nox_xxx_playerSubGold_4FA5D0(int a1, unsigned int a2);
void nox_object_setGold_4FA620(nox_object_t* a1, int a2);
int nox_xxx_playerGetGold_4FA6B0(int a1);
int nox_object_getGold_4FA6D0(nox_object_t* a1);
int nox_xxx_playerBotCreate_4FA700(nox_object_t* a1);
char nox_xxx_mobMorphFromPlayer_4FAAC0(uint32_t* a1);
char nox_xxx_mobMorphToPlayer_4FAAF0(uint32_t* a1);
int nox_xxx_updatePlayerMonsterBot_4FAB20(uint32_t* a1);
char nox_xxx_monsterActionToPlrState_4FABC0(int a1);
int nox_xxx_respawnPlayerBot_4FAC70(int a1);
int nox_xxx_netSendRewardNotify_4FAD50(int a1, int a2, int a3, char a4);
void sub_4FADD0(int a1, const char* a2, char a3);
int sub_4FB000(int a1, int a2);
int sub_4FB050(int a1, int a2, int* a3);
int nox_xxx_playerDoSchedSpell_4FB0E0(nox_object_t* a1, nox_object_t* a2);
int nox_xxx_playerDoSchedSpellQueue_4FB1D0(nox_object_t* a1, nox_object_t* a2);
void nox_xxx_playerExecuteAbil_4FBB70(nox_object_t* a1, int a2);
int sub_4FBE60(void* a1, int a2);
void sub_4FBEA0(void* a1, int a2, int a3);
int sub_4FC030(nox_object_t* a1, int a2);
void sub_4FC070(nox_object_t* a1, int a2, int a3);
void sub_4FC0B0(nox_object_t* a1, int a2);
void nox_xxx_playerCancelAbils_4FC180(nox_object_t* a1);
int nox_common_playerIsAbilityActive_4FC250(nox_object_t* a1, int a2);
void sub_4FC300(nox_object_t* a1, int a2);
int nox_xxx_probablyWarcryCheck_4FC3E0(nox_object_t* a1, int a2);
void sub_4FC440(nox_object_t* a1, int a2);
void sub_4FC670(int a1);
int sub_4FC960(int a1, char a2);
int nox_xxx_Fn_4FCAC0(int a1, int a2);
void nox_xxx_spellCastByBook_4FCB80();
int sub_4FCEB0(int a1);
int nox_xxx_spellCheckSmth_4FCEF0(int a1, int* a2, int a3);
int sub_4FCF90(nox_object_t* a1, int a2, int a3);
unsigned short sub_4FD030(int a1, short a2);
void nox_xxx_teleportAllPixies_4FD090(nox_object_t* a1);
int sub_4FD0E0(nox_object_t* a1, int a2);
int nox_xxx_checkPlrCantCastSpell_4FD150(nox_object_t* a1, int a2, int a3);
int nox_xxx_spellAccept_4FD400(int a1, nox_object_t* a2, nox_object_t* a3p, nox_object_t* a4p, void* a5p, int a6);
int nox_xxx_gameCaptureMagic_4FDC10(int a1, nox_object_t* a2);
int nox_xxx_castSpellByUser_4FDD20(int a1, nox_object_t* a2, void* a3);
uint32_t* nox_xxx_createSpellFly_4FDDA0(nox_object_t* a1, nox_object_t* a2, int a3);
void nox_xxx_collide_4FDF90(int a1, int a2);
int nox_xxx_spellGetPhoneme_4FE1C0(int a1, char a2);
int nox_xxx_spellByBookInsert_4FE340(int a1, int* a2, int a3, int a4, int a5);
void nox_xxx_spell_4FE680(nox_object_t* a1, float a2);
int nox_xxx_spellGetPower_4FE7B0(int a1, nox_object_t* a2);
void sub_4FE8A0(int a1);
void* nox_xxx_spellCastedFirst_4FE930();
void* nox_xxx_spellCastedNext_4FE940(void* a1);
void sub_4FE980(void* a1);
void nox_xxx_spellCancelSpellDo_4FE9D0(void* a1);
int sub_4FEA70(int a1, float2* a2);
int nox_xxx_playerCancelSpells_4FEAE0(nox_object_t* a1);
void nox_xxx_spellCancelDurSpell_4FEB10(int a1, nox_object_t* a2);
void sub_4FEB60(int a1, int a2);
void nox_xxx_cancelAllSpells_4FEE90(nox_object_t* a1);
void nox_xxx_spellCastByPlayer_4FEEF0();
void nox_xxx_netStopRaySpell_4FEF90(void* a1, nox_object_t* a2);
char* nox_xxx_netStartDurationRaySpell_4FF130(int a1);
int sub_4FF2D0(int a1, int a2);
void sub_4FF310(nox_object_t* a1);
int nox_xxx_testUnitBuffs_4FF350(nox_object_t* unit, char buff);
void nox_xxx_buffApplyTo_4FF380(nox_object_t* unit, int buff, short dur, char power);
int nox_xxx_unitGetBuffTimer_4FF550(nox_object_t* unit, int buff);
char nox_xxx_buffGetPower_4FF570(nox_object_t* unit, int buff);
void nox_xxx_unitClearBuffs_4FF580(nox_object_t* unit);
int nox_xxx_spellBuffOff_4FF5B0(nox_object_t* a1, int a2);
void nox_xxx_updateUnitBuffs_4FF620(nox_object_t* a1);
char* nox_xxx_journalQuestSet_500540(char* a1, int a2);
char* nox_xxx_scriptGetJournal_5005E0(char* a1);
char* nox_xxx_journalQuestSetBool_5006B0(char* a1, int a2);
int sub_500750(char* a1);
double sub_500770(char* a1);
void sub_500790(void* lpMem);
char* sub_5007E0(char* a1);
unsigned int sub_5009B0(char* a1);
int sub_500A60();
int sub_500B70();
int nox_xxx_orderUnitLocal_500C70(int owner, int orderType);
int sub_500CA0(int a1, int a2);
int nox_xxx_creatureIsMonitored_500CC0(nox_object_t* a1, nox_object_t* a2);
bool nox_xxx_checkSummonedCreaturesLimit_500D70(nox_object_t* a1, int a2);
int nox_xxx_summonStart_500DA0(int a1);
int sub_500F40(int a1, float a2);
int nox_xxx_summonFinish_5010D0(int a1);
void nox_xxx_summonCancel_5011C0(int a1);
int nox_xxx_charmCreature1_5011F0(int* a1);
int nox_xxx_charmCreatureFinish_5013E0(int* a1);
int nox_xxx_charmCreature2_501690(int a1);
void nox_xxx_banishUnit_5017F0(int unit);
int nox_xxx_getSevenDwords3_501940(int a1);
void nox_xxx_aud_501960(int a1, nox_object_t* a2, int a3, int a4);
void nox_xxx_audCreate_501A30(int a1, float2* a2, int a3, int a4);
void nox_xxx_gameSetAudioFadeoutMb_501AC0(int a1);
char sub_501C00(float* a1, nox_object_t* a2);
void nox_xxx_netUpdateRemotePlr_501CA0(nox_object_t* a1);
int nox_xxx_mapgenMakeScript_502790(FILE* a1, char* a2);
void nox_xxx_mapReset_5028E0();
int sub_5029A0(char* a1);
int sub_5029F0(int a1);
int sub_502A20();
int sub_502A50(char* a1);
int sub_502AB0(char* a1);
int sub_502B10();
int sub_502D70(int a1);
FILE* sub_502DA0(char* a1);
FILE* sub_502DF0();
FILE* sub_502E10(int a1);
double sub_502E70(int a1);
double sub_502EA0(int a1);
int nox_xxx_mapgenSaveMap_503830(int a1);
int sub_503B30(float2* a1);
int sub_503EC0(int a1, float* a2);
void nox_xxx_free_503F40();
uint32_t* nox_xxx_tileAllocTileInCoordList_5040A0(int a1, int a2, float a3);
int nox_xxx_tileInit_504150(int a1, int a2);
uint32_t* sub_504290(char a1, char a2);
uint32_t* nox_xxx_cliWallGet_5042F0(int a1, int a2);
int sub_504330(int a1, int a2);
uint32_t* sub_5044B0(int a1, float a2, float a3);
int sub_504560(int a1, int a2);
void sub_504600(char* a1, unsigned int a2, unsigned char a3);
int sub_5046A0(uint32_t* a1, unsigned int a2);
int sub_504720(unsigned int a1, unsigned int a2);
uint32_t* nox_xxx_unitAddToList_5048A0(int a1);
int sub_504910(int a1, int a2);
int sub_504980();
int sub_5049C0(int a1);
void* sub_5049D0();
int sub_5049E0(int a1);
int sub_504A10(int a1);
void* sub_505060();
int nox_server_mapRWMapIntro_505080();
int nox_server_mapRWGroupData_505C30();
int nox_server_mapRWWaypoints_506260(uint32_t* a1);
int nox_xxx_allocVoteArray_5066D0();
int sub_506720();
int sub_506740(nox_object_t* a1);
void sub_5067B0(int a1);
int sub_506810(int a1);
int nox_xxx_netSendVote_506840(int a1);
char sub_506870(int a1, int a2, wchar2_t* a3);
char sub_5068E0(int a1, int a2, wchar2_t* a3);
uint32_t* sub_506A20(int a1, int a2);
int nox_xxx_voteAddMB_506AD0(int a1);
uint32_t* sub_506B00(int a1, int a2);
uint32_t* sub_506B80(int a1, int a2, wchar2_t* a3);
void sub_506C90(int a1, int a2, wchar2_t* a3);
void sub_506D00(int a1, wchar2_t* a2);
void sub_506DE0(int a1);
void sub_506E50(int a1, wchar2_t* a2);
void sub_506F80(int a1);
int sub_507000(int a1);
void sub_507090(int a1);
void sub_507100(int a1);
int sub_507190(int a1, char a2);
int sub_5071C0();
void sub_509120(uint32_t* a1, int a2, const char* a3);
int sub_5095E0();
int sub_5096F0();

void nox_server_scriptExecuteFnForEachGroupObj_502670(unsigned char* groupPtr, int expectedType, void (*a3)(int, int),
													  int a4);

#endif // NOX_PORT_GAME4
