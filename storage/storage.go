package storage

import "osuprogressserver/types"

type Storage interface {
	//userid, offset, limit
	GetScore(int, int, int) ([]types.ScoreData, error)

	//userid, includes beatmapset info
	GetExtScore(int, int, int) ([]types.Ext_ScoreData, error)

	//Start, Endtime
	GetBanchoTime(string, string) ([]types.BanchoTime, error)

	//Start, Endtime
	GetScreenTime(string, string) ([]types.ScreenTime, error)
}
