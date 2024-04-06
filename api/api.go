package api

import (
	"fmt"
	"osuprogressserver/cmp"
	"osuprogressserver/types"

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

	if userid, ok := c.Locals("userid").(int); ok {
		q.userid = userid
	} else {
		q.userid = c.QueryInt("userid")
	}

	var scores []types.Ext_ScoreData

	scores, err := s.store.GetExtScore(q.query, 14100399, q.limit, q.offset)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if h := c.GetReqHeaders()["Hx-Request"]; len(h) > 0 {
		component := cmp.ScoreContainer(scores, limit, offset+len(scores))

		handler := adaptor.HTTPHandler(templ.Handler(component))

		return handler(c)
	}

	return c.JSON(scores)
}
