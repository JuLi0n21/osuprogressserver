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

	if c.Locals("userid") != nil {
		q.userid = c.Locals("userid").(int)

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
