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

	var scores []types.Ext_ScoreData

	scores, err := s.store.GetExtScore(q.Query, 14100399, q.Offset, q.Limit)
	if err != nil {
		return err
	}

	fmt.Println(len(scores))

	component := cmp.ScoreContainer(scores, q.Offset, q.Limit)

	handler := adaptor.HTTPHandler(templ.Handler(component))

	return handler(c)

}
