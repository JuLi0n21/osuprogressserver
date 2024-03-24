package api

import (
	"context"
	"net/http"
	"osuprogressserver/types"
	"osuprogressserver/views"
)

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	//todo
	player := types.User{
		Username: "JuLi0n_",
	}
	component := views.Index(player)

	component.Render(context.Background(), w)
}
