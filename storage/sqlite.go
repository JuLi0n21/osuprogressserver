package storage

import (
	"database/sql"
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
		err := rows.Scan(
			&score.Title,
			&score.Version,
			&score.Date,
			&score.BeatmapSetId,
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
			&score.Maxcombo,
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
		)
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

func (s *SQLite) GetExtScore(userid int, limit int, offset int) ([]types.Ext_ScoreData, error) {

	return []types.Ext_ScoreData{}, nil
}
func (s *SQLite) GetBanchoTime(start string, end string) ([]types.BanchoTime, error) {

	return []types.BanchoTime{}, nil

}
func (s *SQLite) GetScreenTime(start string, end string) ([]types.ScreenTime, error) {

	return []types.ScreenTime{}, nil
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
