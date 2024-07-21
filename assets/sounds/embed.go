package sounds

import (
	_ "embed"
)

var (
	//go:embed bounces/bounce.wav
	B_bounce []byte

	//go:embed bounces/silencer.wav
	B_silencer []byte

	//go:embed mines/flyin_off.wav
	M_flyin_off []byte
	//go:embed mines/odd_beep.wav
	M_odd_beep []byte
	//go:embed mines/torpedo_mine.wav
	M_torpedo_mine []byte

	//go:embed explosions/1_bomb.wav
	E_1_bomb []byte
	//go:embed explosions/boom.wav
	E_boom []byte
	//go:embed explosions/grenade.wav
	E_grenade []byte
	//go:embed explosions/mortar.wav
	E_mortar []byte

	//go:embed grunts/turtle_mating.mp3
	G_turtle_mating []byte
	//go:embed grunts/grunt1.mp3
	G_grunt1 []byte
	//go:embed grunts/grunt2.mp3
	G_grunt2 []byte
	//go:embed grunts/banana.mp3
	G_banana []byte

	//go:embed ouches/1_burp.wav
	O_1_burp []byte
	//go:embed ouches/2_gargle.wav
	O_2_gargle []byte
	//go:embed ouches/3_grunt.wav
	O_3_grunt []byte
	//go:embed ouches/3_sheep.wav
	O_3_sheep []byte
	//go:embed ouches/4_heavy.wav
	O_4_heavy []byte
	//go:embed ouches/5_no.wav
	O_5_no []byte
	//go:embed ouches/6_oops.wav
	O_6_oops []byte
	//go:embed ouches/beback.wav
	O_beback []byte
	//go:embed ouches/glass.wav
	O_glass []byte
	//go:embed ouches/ouch.wav
	O_ouch []byte
	//go:embed ouches/scream.wav
	O_scream []byte
	//go:embed ouches/glass_shatter.mp3
	O_glass_shatter []byte
	//go:embed ouches/oof.mp3
	O_oof []byte
	//go:embed ouches/duck.mp3
	O_duck []byte
	//go:embed ouches/death.mp3
	O_death []byte

	//go:embed other/connected.mp3
	Ot_connected []byte
	//go:embed other/disconnected.mp3
	Ot_disconnected []byte

	//go:embed shots/barrett.wav
	S_barrett []byte
	//go:embed shots/m1_garand.wav
	S_m1_garand []byte
	//go:embed shots/m4a1.wav
	S_m4a1 []byte
	//go:embed shots/nitro_express.wav
	S_nitro_express []byte
	//go:embed shots/rifle.wav
	S_rifle []byte
	//go:embed shots/shotgun.wav
	S_shotgun []byte

	//go:embed shield/dodge_again.mp3
	Sh_dodge_again []byte
	//go:embed shield/eat.mp3
	Sh_eat []byte
	//go:embed shield/immune_to_physical_damage.mp3
	Sh_immune_to_physical_damage []byte
	//go:embed shield/magic_wand.mp3
	Sh_magic_wand []byte
	//go:embed shield/nope.mp3
	Sh_nope []byte
	//go:embed shield/sheep.mp3
	Sh_sheep []byte
	//go:embed shield/sike.mp3
	Sh_sike []byte
	//go:embed shield/squeak_2.mp3
	Sh_squeak_2 []byte
	//go:embed shield/squeak.mp3
	Sh_squeak []byte
	//go:embed shield/stop_it.mp3
	Sh_stop_it []byte
	//go:embed shield/taunt.mp3
	Sh_taunt []byte

	//go:embed victories/aaaah.wav
	V_aaaah []byte
	//go:embed victories/cheer.wav
	V_cheer []byte
	//go:embed victories/fireworks_win.wav
	V_fireworks_win []byte
	//go:embed victories/kiss.wav
	V_kiss []byte
	//go:embed victories/laugh.wav
	V_laugh []byte
	//go:embed victories/laugh3.wav
	V_laugh3 []byte
	//go:embed victories/laugh4.wav
	V_laugh4 []byte
	//go:embed victories/yahoo.wav
	V_yahoo []byte
	//go:embed victories/yippee.wav
	V_yippee []byte
	//go:embed victories/victory_fanfare.mp3
	V_victory_fanfare []byte
	//go:embed victories/dun_dun_dunnn.mp3
	V_dun_dun_dunnn []byte
	//go:embed victories/turtle_sex.mp3
	V_turtle_sex []byte
	//go:embed victories/sad_violin.mp3
	V_sad_violin []byte
	//go:embed victories/mmmm.mp3
	V_mmmm []byte
	//go:embed victories/mission_passed.mp3
	V_mission_passed []byte
	//go:embed victories/herewego.mp3
	V_herewego []byte
	//go:embed victories/hallelujah.mp3
	V_hallelujah []byte
	//go:embed victories/haha.mp3
	V_haha []byte
	//go:embed victories/ha_gotem.mp3
	V_ha_gotem []byte
	//go:embed victories/flush.mp3
	V_flush []byte
	//go:embed victories/fart_asmr.mp3
	V_fart_asmr []byte
	//go:embed victories/fail.mp3
	V_fail []byte
	//go:embed victories/aww.mp3
	V_aww []byte
	//go:embed victories/drdisrespect_getthisidiotkidoutofhere.mp3
	V_drdisrespect_getthisidiotkidoutofhere []byte
	//go:embed victories/drdisrespect_neverplaythisshitgameagain.mp3
	V_drdisrespect_neverplaythisshitgameagain []byte
	//go:embed victories/Keanu_Breathtaking.mp3
	V_Keanu_Breathtaking []byte
	//go:embed victories/noob.mp3
	V_noob []byte
	//go:embed victories/tdft2.mp3
	V_tdft2 []byte
	//go:embed victories/gay.mp3
	V_gay []byte
	// //go:embed victories/meme.mp3
	// V_meme []byte
)
