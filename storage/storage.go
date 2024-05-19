package storage

import "osuprogressserver/types"

type Storage interface {
	//userid, offset, limit
	GetScore(int, int, int) ([]types.ScoreData, error)

	//Show random Scores
	GetRandomScores(int) ([]types.Ext_ScoreData, error)

	//userid, includes beatmapset info
	GetExtScore(string, int, int, int) ([]types.Ext_ScoreData, error)

	//Scoreid
	GetExtScoreById(int) ([]types.Ext_ScoreData, error)

	//Start, Endtime
	GetBanchoTime(string, string) ([]types.BanchoTime, error)

	//Start, Endtime
	GetScreenTime(string, string) ([]types.ScreenTime, error)

	//Userid
	GetUser(int) (types.User, error)

	GetApiUser(int) (types.ApiUser, error)

	//User
	SaveUser(types.User) error

	SaveApiUser(types.ApiUser) error

	SaveExtendedScore(types.Ext_ScoreData) error

	SaveScore(types.ScoreData) error

	SaveBeatmap(types.Beatmap) error

	SaveBeatmapSet(types.BeatmapSet) error

	SaveBanchoTime(types.BanchoTime) error

	SaveScreenTime(types.ScreenTime) error

	MockScores(int) error
}
