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

func (s *SQLite) GetExtScore(query string, userid int, limit int, offset int) ([]types.Ext_ScoreData, error) {

	if limit > 100 {
		limit = 100
	}

	stmt, err := s.DB.Prepare(`SELECT * FROM ScoreData
	INNER JOIN Beatmapset on BeatmapSet.BeatmapSetId = Scoredata.BeatmapSetID
	WHERE Userid = ? 
	AND (Title LIKE ? OR Version LIKE ?)
	ORDER BY Date
	LIMIT ? 
	OFFSET ?`)
	if err != nil {
		fmt.Println(err.Error())
		return []types.Ext_ScoreData{}, nil
	}

	var scores []types.Ext_ScoreData

	rows, err := stmt.Query(userid, query, query, offset, limit)
	if err != nil {
		fmt.Println(err.Error())
		return scores, nil
	}
	defer rows.Close()

	for rows.Next() {
		var score types.Ext_ScoreData
		if err := rows.Scan(&score); err != nil {

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

func (s *SQLite) SaveScore(score types.ScoreData) error {
	return nil
}

func (s *SQLite) SaveBeatmapSet(score types.BeatmapSet) error {
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
		Version TEXT,
		Date TEXT,
		BeatmapSetID INTEGER,
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
		MaxCombo INTEGER,
		Hit50 INTEGER,
		Hit100 INTEGER,
		Hit300 INTEGER,
		Ur REAL,
		HitMiss INTEGER,
		Mode INTEGER,
		Mods INTEGER,
		Time INTEGER,
		PP REAL,
		AIM REAL,
		SPEED REAL,
		ACCURACYATT REAL,
		Grade TEXT,
		FCPP REAL
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
		BeatmapSetId INT NOT NULL PRIMARY KEY,
		Artist       TEXT,
		Tags         TEXT,
		CoverList    TEXT,
		Cover        TEXT,
		Preview      TEXT,
		Rankedstatus TEXT
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

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_screentime_date ON ScreenTime(Userid, Date)`)
	if err != nil {
		log.Fatal(err)
	}
}
