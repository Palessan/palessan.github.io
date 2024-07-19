package fonts

import (
	_ "embed"
)

var (
	//go:embed bulletBounces\bounce.wav
	B_bounce []byte

	//go:embed bulletBounces/silencer.wav
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
	//go:embed ouches/scary_scream.wav
	O_scary_scream []byte
	//go:embed ouches/scream.wav
	O_scream []byte

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

	//go:embed grunts/turtle_mating.mp3
	G_turtle_mating []byte
	//go:embed grunts/angry-grunt-103204.mp3
	G_grunt1 []byte
	//go:embed grunts/grunt-106134.mp3
	G_grunt2 []byte

	//go:embed other/connected.mp3
	O_connected []byte
	//go:embed other/disconnected.mp3
	O_disconnected []byte
)
