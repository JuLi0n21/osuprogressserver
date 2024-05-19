package types

type ScoreData struct {
	ScoreId     string
	Title       string
	Date        string
	BeatmapID   int
	Playtype    string
	Ar          float64
	Cs          float64
	Hp          float64
	Od          float64
	SR          float64
	Bpm         float64
	Userid      int
	ACC         float64
	Score       int
	Combo       int
	Hit50       int
	Hit100      int
	Hit300      int
	Ur          float64
	HitMiss     int
	Mode        int
	Mods        string
	Time        int
	PP          float64
	AIM         float64
	SPEED       float64
	ACCURACYATT float64
	Grade       string
	FCPP        float64
}

type Beatmap struct {
	BeatmapID    int
	BeatmapSetID int
	Maxcombo     int
	Version      string
}

type BeatmapSet struct {
	BeatmapSetID int
	Artist       string
	Creator      string
	Tags         string
	CoverList    string
	Cover        string
	Preview      string
	Rankedstatus string
}

type Ext_ScoreData struct {
	ScoreData
	Beatmap
	BeatmapSet
	ApiUser
}

type ScreenTime struct {
	Date   string
	Screen string
	Time   int
	Userid int
}

type BanchoTime struct {
	Date   string
	Screen string
	Time   int
	Userid int
}

type User struct {
	UserId   int
	Username string
	Mode     string
	Auth     AuthUser
}
