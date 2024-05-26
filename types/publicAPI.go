package types

type ApiScore struct {
	Date         string
	BeatmapSetID int
	BeatmapID    int
	OsuFilename  string
	FolderName   string
	Replay       string
	PlayType     string
	AR           float64
	CS           float64
	HP           float64
	OD           float64
	Status       string
	SR           float64
	BPM          float64
	Artist       string
	Creator      string
	Username     string
	Userid       int
	ACC          float64
	MaxCombo     int
	Score        int
	Combo        int
	Hit50        int
	Hit100       int
	Hit300       int
	UR           float64
	HitMiss      int
	Mode         int
	Mods         int
	Version      string
	Tags         string
	CoverList    string
	Cover        string
	Preview      string
	Time         int
	PP           float64
	AIM          float64
	SPEED        float64
	AccuracyAtt  float64
	Grade        string
	FCPP         float64
	Title        string
}
