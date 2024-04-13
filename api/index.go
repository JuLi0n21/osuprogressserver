package api

import (
	"fmt"
	"osuprogressserver/cmp"
	"osuprogressserver/types"

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

	scoreData := types.ScoreData{
		ScoreId:     "123abc",
		Title:       "Example Beatmap",
		Date:        "2024-04-05",
		BeatmapID:   1,
		Playtype:    "Standard",
		Ar:          9.0,
		Cs:          4.0,
		Hp:          8.0,
		Od:          9.0,
		SR:          6.5,
		Bpm:         180.0,
		Userid:      14100399,
		ACC:         98.5,
		Score:       1000000,
		Combo:       1000,
		Hit50:       5,
		Hit100:      50,
		Hit300:      300,
		Ur:          200.0,
		HitMiss:     0,
		Mode:        0,
		Mods:        "hd,hr",
		Time:        120,
		PP:          500.0,
		AIM:         98.0,
		SPEED:       97.0,
		ACCURACYATT: 96.0,
		Grade:       "SS",
		FCPP:        510.0,
	}

	beatmap := types.Beatmap{
		BeatmapID:    1,
		BeatmapSetID: 100,
		Maxcombo:     500,
		Version:      "Expert",
	}

	// Initialize BeatmapSet
	beatmapSet := types.BeatmapSet{
		BeatmapSetID: 100,
		Artist:       "Artist Name",
		Creator:      "Creator Name",
		Tags:         "tags, music, game",
		CoverList:    "https://assets.ppy.sh/beatmaps/1953712/covers/list@2x.jpg?1679351388",
		Cover:        "https://assets.ppy.sh/beatmaps/1953712/covers/cover@2x.jpg?1679351388",
		Preview:      "preview.mp3",
		Rankedstatus: "Ranked",
	}

	// Compose them into an instance of ExtScoreData
	extScoreData := types.Ext_ScoreData{
		ScoreData:  scoreData,
		Beatmap:    beatmap,
		BeatmapSet: beatmapSet,
	}

	// _ = extScoreData
	s.store.SaveExtendedScore(extScoreData)

	scores, err := s.store.GetExtScore("", 14100399, 10, 0)
	if err != nil {
		fmt.Println(err)
	}

	component := cmp.View_Index(player, stats, scores)

	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)
}
