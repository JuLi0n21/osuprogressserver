package api

import (
	"osuprogressserver/types"
	"osuprogressserver/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (s *Server) Index(c *fiber.Ctx) error {
	//todo
	player := types.User{
		Username:    "JuLi0n_",
		UserId:      14100399,
		Banner:      "https://assets.ppy.sh/user-profile-covers/14100399/cd1600936d7a56115cd147f47169addcf8f812133861b667c7dd3d177ca5068d.jpeg",
		Avatar:      "https://a.ppy.sh/14100399?1672009368.jpeg",
		GlobalRank:  "57497",
		LocalRank:   "2928",
		Country:     "Germany",
		Countrycode: "de",
	}

	stats := types.Stats{
		Time:   "9h",
		Status: "Afk",
		Count:  8234,
		Screen: "Mainmenu",
	}

	score := []types.Ext_ScoreData{
		{
			ScoreData: types.ScoreData{
				ScoreId:      "3432",
				Title:        "VANESSA",
				Version:      "Normal",
				Date:         "2024-03-31",
				BeatmapSetId: 123456,
				Playtype:     "Ranked",
				Ar:           9.5,
				Cs:           4.0,
				Hp:           7.0,
				Od:           8.0,
				SR:           4.25,
				Bpm:          180.0,
				Userid:       14100399,
				ACC:          98.5,
				Score:        950000,
				Combo:        1500,
				Maxcombo:     1500,
				Hit50:        5,
				Hit100:       50,
				Hit300:       1350,
				Ur:           200.0,
				HitMiss:      5,
				Mode:         0,
				Mods:         "dt,hd,hr",
				Time:         240,
				PP:           345.5,
				AIM:          150.0,
				SPEED:        200.0,
				ACCURACYATT:  9.8,
				Grade:        "S",
				FCPP:         400.5,
			},
			BeatmapSet: types.BeatmapSet{
				BeatmapSetId: 123456,
				Artist:       "Suzaku",
				Tags:         "st, stsl , stsrr, st",
				Cover:        "https://assets.ppy.sh/beatmaps/1208907/covers/cover@2x.jpg?1629681318",
				CoverList:    "https://assets.ppy.sh/beatmaps/1208907/covers/list@2x.jpg?1629681318",
				Preview:      "soundcloudlink",
				Rankedstatus: "Ranked",
			},
		},
	}

	s.store.SaveScore(score[len(score)-1].ScoreData)

	s.store.SaveBeatmapSet(score[len(score)-1].BeatmapSet)

	component := views.Index(player, stats, score)

	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)
}
