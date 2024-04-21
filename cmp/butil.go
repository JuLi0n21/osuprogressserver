package cmp

import (
	"context"
	"osuprogressserver/types"
	"strconv"
)

func I(num int) string {
	return strconv.Itoa(num)
}

func F(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func theme(ctx context.Context) types.Theme {
	if theme, ok := ctx.Value("theme").(types.Theme); ok {
		return theme
	}
	return DefaultTheme()
}

func player(ctx context.Context) types.UserContext {
	if userc, ok := ctx.Value("player").(types.UserContext); ok {
		return userc
	}
	return types.UserContext{}
}

func DefaultTheme() types.Theme {
	return types.Theme{
		Dark:         "backdrop--dark",
		Medium_dark:  "backdrop--medium--dark",
		Medium:       "backdrop--medium",
		Medium_light: "backdrop--medium--light",
		Light:        "backdrop--light",
	}
}
