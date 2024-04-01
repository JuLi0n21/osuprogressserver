package cmp

type helper struct {
	Name string
	Link string
}

// used to host static files on diffrent services
var files = []helper{
	//grades
	{"XH", "/assets/img/XH.png"},
	{"X", "/assets/img/X.png"},
	{"SS", "/assets/img/SS.png"},
	{"S", "/assets/img/S.png"},
	{"A", "/assets/img/.png"},
	{"B", "/assets/img/.png"},
	{"C", "/assets/img/C.png"},
	{"D", "/assets/img/D.png"},
	{"E", "/assets/img/E.png"},
	{"F", "/assets/img/F.png"},
	//transparent-grades

	//mods
	{"ap", "/assets/img/Autopilot.png"},
	{"at", "/assets/img/Autoplay.png"},
	{"cn", "/assets/img/Cinema.png"},
	{"dt", "/assets/img/DoubleTime.png"},
	{"ez", "/assets/img/Easy.png"},
	{"fl", "/assets/img/Flashlight.png"},
	{"ht", "/assets/img/HalfTime.png"},
	{"hr", "/assets/img/HardRock.png"},
	{"hd", "/assets/img/Hidden.png"},
	{"nc", "/assets/img/NightCore.png"},
	{"nf", "/assets/img/Nofail.png"},
	{"pf", "/assets/img/Perfect.png"},
	{"rx", "/assets/img/Relax.png"},
	{"v2", "/assets/img/Scorev2.png"},
	{"sd", "/assets/img/SuddenDeath.png"},
	{"so", "/assets/img/SpunOut.png"},
	//additional files
}

func Img(name string) string {
	for _, file := range files {
		if file.Name == name {
			return file.Link
		}
	}
	return "404.png"
}
