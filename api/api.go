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
	Query  string
	From   string
	To     string
	Limit  int
	Offset int
}

func (s *Server) ScoreSearch(c *fiber.Ctx) error {

	q := new(SearchQuery)
	if err := c.QueryParser(q); err != nil {
		fmt.Println(err.Error())
		return err
	}

	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)
	var scores []types.Ext_ScoreData

	scores, err := s.store.GetExtScore(q.Query, 14100399, limit, offset)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	component := cmp.ScoreContainer(scores, limit, offset+len(scores))

	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)

}
