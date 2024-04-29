package api

import (
	"context"
	"fmt"
	"osuprogressserver/cmp"
	"osuprogressserver/types"
	"strings"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type SearchQuery struct {
	query  string
	from   string
	to     string
	limit  int
	offset int
	userid int
}

func (s *Server) ScoreSearch(c *fiber.Ctx) error {

	q := new(SearchQuery)
	if err := c.QueryParser(q); err != nil {
		fmt.Println(err.Error())
		return err
	}

	q.limit = c.QueryInt("limit", 10)
	q.offset = c.QueryInt("offset", 0)
	q.query = c.Query("query", "")

	//fmt.Println(q)

	if c.Locals("User").(types.UserContext).User.UserId != 0 {
		q.userid = c.Locals("User").(types.UserContext).User.UserId

	} else {
		q.userid = c.QueryInt("userid")
	}

	var scores []types.Ext_ScoreData

	scores, err := s.store.GetExtScore(q.query, q.userid, q.limit, q.offset)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if h := c.GetReqHeaders()["Hx-Request"]; len(h) == 0 {
		return c.JSON(scores)
	}

	themes := cmp.DefaultTheme()

	if strings.Contains(c.Get("Referer"), "/score/") {
		themes = types.Theme{
			Dark:         "score-backdrop--dark",
			Medium_dark:  "score-backdrop--medium--dark",
			Medium:       "score-backdrop--medium",
			Medium_light: "score-backdrop--medium--light",
			Light:        "score-backdrop--light",
		}

	}

	ctx := context.WithValue(context.Background(), "theme", themes)

	component := cmp.ScoreContainer(scores, q.limit, q.offset+len(scores))

	handler := adaptor.HTTPHandler(C(templ.Handler(component), ctx))

	return handler(c)
}

func (s *Server) Score(c *fiber.Ctx) error {

	if t := c.GetReqHeaders()["Authorization"]; len(t) == 0 {
		return c.SendStatus(401)
	}

	var apiscore types.ApiScore

	if err := c.BodyParser(&apiscore); err != nil {
		return err
	}

	score := types.Ext_ScoreData{
		ScoreData: types.ScoreData{
			Title:       apiscore.Title,
			Date:        apiscore.Date,
			Playtype:    apiscore.PlayType,
			BeatmapID:   apiscore.BeatmapID,
			Ar:          apiscore.AR,
			Cs:          apiscore.CS,
			Hp:          apiscore.HP,
			Od:          apiscore.OD,
			SR:          apiscore.SR,
			Bpm:         apiscore.BPM,
			Userid:      apiscore.Username,
			ACC:         apiscore.ACC,
			Score:       apiscore.Score,
			Combo:       apiscore.Combo,
			Hit50:       apiscore.Hit50,
			Hit100:      apiscore.Hit100,
			Hit300:      apiscore.Hit300,
			Ur:          apiscore.UR,
			HitMiss:     apiscore.HitMiss,
			Mode:        apiscore.Mode,
			Mods:        apiscore.Mods,
			Time:        apiscore.Time,
			PP:          apiscore.PP,
			AIM:         apiscore.AIM,
			SPEED:       apiscore.SPEED,
			ACCURACYATT: apiscore.AccuracyAtt,
			Grade:       apiscore.Grade,
			FCPP:        apiscore.FCPP,
		},
		Beatmap: types.Beatmap{
			BeatmapID:    apiscore.BeatmapID,
			BeatmapSetID: apiscore.BeatmapSetID,
			Maxcombo:     apiscore.MaxCombo,
			Version:      apiscore.Version,
		},
		BeatmapSet: types.BeatmapSet{
			BeatmapSetID: apiscore.BeatmapSetID,
			Artist:       apiscore.Artist,
			Creator:      apiscore.Creator,
			Tags:         apiscore.Tags,
			CoverList:    apiscore.CoverList,
			Cover:        apiscore.Cover,
			Preview:      apiscore.Preview,
			Rankedstatus: apiscore.Status,
		},
	}

	s.store.SaveExtendedScore(score)

	return c.SendStatus(200)
}
