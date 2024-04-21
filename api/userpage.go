package api

import (
	"context"
	"fmt"
	"osuprogressserver/cmp"
	"osuprogressserver/types"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (s *Server) Userpage(c *fiber.Ctx) error {
	//todo
	player := c.Locals("User").(types.UserContext)
	fmt.Println(player.User.Username)

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

	scores, err := s.store.GetRandomScores(10)
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.WithValue(context.Background(), "player", player)

	component := cmp.View_Userpage(stats, scores)

	handler := adaptor.HTTPHandler(C(templ.Handler(component), ctx))

	return handler(c)
}
