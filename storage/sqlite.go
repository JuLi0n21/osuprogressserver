package storage

import (
	"database/sql"
	"fmt"
	"log"
	"osuprogressserver/types"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	filename string
	DB       *sql.DB
}

func NewSQLite(filename string) (*SQLite, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	createTables(db)

	return &SQLite{
		filename: filename,
		DB:       db,
	}, nil
}

func (s *SQLite) GetScore(userid int, limit int, offset int) ([]types.ScoreData, error) {

	if limit > 100 {
		limit = 100
	}

	rows, err := s.DB.Query("SELECT * FROM ScoreData WHERE UserID = ? LIMIT ? OFFSET ?", userid, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []types.ScoreData

	for rows.Next() {
		var score types.ScoreData
		err := rows.Scan(&score)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return scores, nil
}

func (s *SQLite) GetExtScoreById(id int) ([]types.Ext_ScoreData, error) {

	stmt, err := s.DB.Prepare(`SELECT ScoreData.ROWID as Scoreid, * FROM ScoreData
	LEFT JOIN Beatmap on Beatmap.BeatmapID = Scoredata.BeatmapID
	LEFT JOIN BeatmapSet on BeatmapSet.BeatmapSetID = Beatmap.BeatmapSetID
	WHERE Scoreid = ? `)

	if err != nil {
		fmt.Println(err.Error())
		return []types.Ext_ScoreData{}, nil
	}

	var scores []types.Ext_ScoreData
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Println(err.Error())
		return scores, nil
	}
	defer rows.Close()

	for rows.Next() {
		var score types.Ext_ScoreData
		if err := rows.Scan(
			&score.ScoreId,
			&score.Title,
			&score.Date,
			&score.ScoreData.BeatmapID,
			&score.Playtype,
			&score.Ar,
			&score.Cs,
			&score.Hp,
			&score.Od,
			&score.SR,
			&score.Bpm,
			&score.Userid,
			&score.ACC,
			&score.Score,
			&score.Combo,
			&score.Hit50,
			&score.Hit100,
			&score.Hit300,
			&score.Ur,
			&score.HitMiss,
			&score.Mode,
			&score.Mods,
			&score.Time,
			&score.PP,
			&score.AIM,
			&score.SPEED,
			&score.ACCURACYATT,
			&score.Grade,
			&score.FCPP,
			&score.Beatmap.BeatmapID,
			&score.Beatmap.BeatmapSetID,
			&score.Beatmap.Maxcombo,
			&score.Beatmap.Version,
			&score.BeatmapSet.BeatmapSetID,
			&score.Artist,
			&score.Creator,
			&score.Tags,
			&score.CoverList,
			&score.Cover,
			&score.Preview,
			&score.Rankedstatus,
		); err != nil {

			fmt.Println(err.Error())
			return scores, nil
		}
		scores = append(scores, score)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return scores, nil
	}

	return scores, nil
}

func (s *SQLite) GetExtScore(query string, userid int, limit int, offset int) ([]types.Ext_ScoreData, error) {

	if limit > 100 || limit == 0 {
		limit = 100
	}

	stmt, err := s.DB.Prepare(`SELECT ScoreData.ROWID as Scoreid, * FROM ScoreData
	LEFT JOIN Beatmap on Beatmap.BeatmapID = Scoredata.BeatmapID
	LEFT JOIN BeatmapSet on BeatmapSet.BeatmapSetID = Beatmap.BeatmapSetID
	WHERE Userid = ? 
	AND (Title LIKE ? OR Version LIKE ?)
	ORDER BY date
	LIMIT ? 
	OFFSET ?`)
	if err != nil {
		fmt.Println(err.Error())
		return []types.Ext_ScoreData{}, nil
	}

	var scores []types.Ext_ScoreData
	var q = "%" + query + "%"
	//fmt.Println(q, limit, offset, userid)
	rows, err := stmt.Query(userid, q, q, limit, offset)
	if err != nil {
		fmt.Println(err.Error())
		return scores, nil
	}
	defer rows.Close()

	for rows.Next() {
		var score types.Ext_ScoreData
		if err := rows.Scan(
			&score.ScoreId,
			&score.Title,
			&score.Date,
			&score.ScoreData.BeatmapID,
			&score.Playtype,
			&score.Ar,
			&score.Cs,
			&score.Hp,
			&score.Od,
			&score.SR,
			&score.Bpm,
			&score.Userid,
			&score.ACC,
			&score.Score,
			&score.Combo,
			&score.Hit50,
			&score.Hit100,
			&score.Hit300,
			&score.Ur,
			&score.HitMiss,
			&score.Mode,
			&score.Mods,
			&score.Time,
			&score.PP,
			&score.AIM,
			&score.SPEED,
			&score.ACCURACYATT,
			&score.Grade,
			&score.FCPP,
			&score.Beatmap.BeatmapID,
			&score.Beatmap.BeatmapSetID,
			&score.Beatmap.Maxcombo,
			&score.Beatmap.Version,
			&score.BeatmapSet.BeatmapSetID,
			&score.Artist,
			&score.Creator,
			&score.Tags,
			&score.CoverList,
			&score.Cover,
			&score.Preview,
			&score.Rankedstatus,
		); err != nil {

			fmt.Println(err.Error())
			return scores, nil
		}
		scores = append(scores, score)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return scores, nil
	}

	return scores, nil
}

func (s *SQLite) GetBanchoTime(start string, end string) ([]types.BanchoTime, error) {

	return []types.BanchoTime{}, nil

}
func (s *SQLite) GetScreenTime(start string, end string) ([]types.ScreenTime, error) {

	return []types.ScreenTime{}, nil
}
func (s *SQLite) GetUser(Userid int) (types.User, error) {
	return types.User{}, nil
}

func (s *SQLite) SaveUser(User types.User) error {
	return nil
}

func (s *SQLite) SaveExtendedScore(score types.Ext_ScoreData) error {

	err := s.SaveScore(score.ScoreData)
	if err != nil {
		return err
	}

	err = s.SaveBeatmapSet(score.BeatmapSet)
	if err != nil {
		return err
	}

	err = s.SaveBeatmap(score.Beatmap)
	if err != nil {
		return err
	}
	return nil

}

func (s *SQLite) SaveScore(score types.ScoreData) error {

	stmt, err := s.DB.Prepare(`
	INSERT INTO ScoreData (Title, Date, BeatmapID, Playtype, Ar, Cs, Hp, Od, SR, Bpm, Userid, ACC, Score, Combo, Hit50, Hit100, Hit300, Ur, HitMiss, Mode, Mods, Time, PP, AIM, SPEED, ACCURACYATT, Grade, FCPP)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		score.Title,
		score.Date,
		score.BeatmapID,
		score.Playtype,
		score.Ar,
		score.Cs,
		score.Hp,
		score.Od,
		score.SR,
		score.Bpm,
		score.Userid,
		score.ACC,
		score.Score,
		score.Combo,
		score.Hit50,
		score.Hit100,
		score.Hit300,
		score.Ur,
		score.HitMiss,
		score.Mode,
		score.Mods,
		score.Time,
		score.PP,
		score.AIM,
		score.SPEED,
		score.ACCURACYATT,
		score.Grade,
		score.FCPP,
	)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (s *SQLite) SaveBeatmap(beatmap types.Beatmap) error {
	stmt, err := s.DB.Prepare(`
	INSERT INTO Beatmap (BeatmapID, BeatmapSetID, Maxcombo, Version)
	VALUES(?, ?, ?, ?)`)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		beatmap.BeatmapID,
		beatmap.BeatmapSetID,
		beatmap.Maxcombo,
		beatmap.Version,
	)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (s *SQLite) SaveBeatmapSet(beatmapset types.BeatmapSet) error {

	stmt, err := s.DB.Prepare(`
	INSERT INTO BeatmapSet (BeatmapSetID, Artist, Creator, Tags, CoverList, Cover, Preview, Rankedstatus)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		beatmapset.BeatmapSetID,
		beatmapset.Artist,
		beatmapset.Creator,
		beatmapset.Tags,
		beatmapset.CoverList,
		beatmapset.Cover,
		beatmapset.Preview,
		beatmapset.Rankedstatus,
	)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (s *SQLite) SaveBanchoTime(score types.BanchoTime) error {
	return nil
}

func (s *SQLite) SaveScreenTime(score types.ScreenTime) error {
	return nil
}

func createTables(db *sql.DB) {

	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS ScoreData (
		Title TEXT,
		Date TEXT,
		BeatmapID INTEGER,
		Playtype TEXT,
		Ar REAL,
		Cs REAL,
		Hp REAL,
		Od REAL,
		SR REAL,
		Bpm REAL,
		Userid INTEGER,
		ACC REAL,
		Score INTEGER,
		Combo INTEGER,
		Hit50 INTEGER,
		Hit100 INTEGER,
		Hit300 INTEGER,
		Ur REAL,
		HitMiss INTEGER,
		Mode INTEGER,
		Mods TEXT,
		Time INTEGER,
		PP REAL,
		AIM REAL,
		SPEED REAL,
		ACCURACYATT REAL,
		Grade TEXT,
		FCPP REAL,
		FOREIGN KEY (Userid) REFERENCES Users(Userid),
		UNIQUE(Userid, BeatmapID, Combo, Hit300, Hit100, Hit50)
		)`)
	if err != nil {
		log.Fatal(err)
	}

	// Create indexes on Date and Username columns
	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_username_date ON ScoreData(Userid, Date)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS BeatmapSet (
		BeatmapSetID INT NOT NULL PRIMARY KEY,
		Artist       TEXT,
		Creator		 TEXT,
		Tags         TEXT,
		CoverList    TEXT,
		Cover        TEXT,
		Preview      TEXT,
		Rankedstatus TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Beatmap (
		BeatmapID INT NOT NULL PRIMARY KEY,
		BeatmapSetID INTEGER,
		MaxCombo INTEGER,
		Version TEXT,
		FOREIGN KEY (BeatmapSetID) REFERENCES BeatmapSet(BeatmapSetID)
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS BanchoTime (
		Date TEXT,
		Screen TEXT,
		Time INTEGER,
		Userid INTEGER
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_banchotime_date ON BanchoTime(Userid, Date)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS ScreenTime (
		Date TEXT,
		Screen TEXT,
		Time INTEGER,
		Userid INTEGER
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Users (
		Userid INTEGER NOT NULL PRIMARY KEY,
		Username    TEXT,
		Banner      TEXT,
		Avatar      TEXT,
		GlobalRank  TEXT,
		LocalRank   TEXT,
		Country     TEXT,
		Countrycode TEXT,
		Mode		TEXT,
		Tokentype TEXT,
		Expiresin INTEGER,
		Accesstoken TEXT,
		Refreshtoken TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_screentime_date ON ScreenTime(Userid, Date)`)
	if err != nil {
		log.Fatal(err)
	}
}
