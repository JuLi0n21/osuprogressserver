package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"osuprogressserver/types"

	"github.com/brianvoe/gofakeit/v7"
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
	LEFT JOIN ApiUsers on ApiUsers.Userid = Scoredata.Userid
	INNER JOIN Beatmap on Beatmap.BeatmapID = Scoredata.BeatmapID
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

		var apiuser string
		var score types.Ext_ScoreData
		if err := rows.Scan(
			&score.ScoreId,
			&score.ScoreData.Title,
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
			&score.ApiUser.ID,
			&apiuser,
			&score.Beatmap.BeatmapID,
			&score.Beatmap.BeatmapSetID,
			&score.Beatmap.Maxcombo,
			&score.Beatmap.Version,
			&score.BeatmapSet.BeatmapSetID,
			&score.Artist,
			&score.Creator,
			&score.Tags,
			&score.CoverList,
			&score.BeatmapSet.Cover,
			&score.Preview,
			&score.Rankedstatus,
		); err != nil {

			fmt.Println(err.Error())
			return scores, nil
		}

		json.Unmarshal([]byte(apiuser), &score.ApiUser)

		scores = append(scores, score)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
		return scores, nil
	}

	return scores, nil
}

func (s *SQLite) GetRandomScores(limit int) ([]types.Ext_ScoreData, error) {

	stmt, err := s.DB.Prepare(`SELECT ScoreData.ROWID as Scoreid, * FROM ScoreData
	LEFT JOIN ApiUsers on ApiUsers.Userid = Scoredata.Userid
	INNER JOIN Beatmap on Beatmap.BeatmapID = Scoredata.BeatmapID
	LEFT JOIN BeatmapSet on BeatmapSet.BeatmapSetID = Beatmap.BeatmapSetID
	ORDER BY RANDOM()
	LIMIT ?`)

	if err != nil {
		fmt.Println(err.Error())
		return []types.Ext_ScoreData{}, nil
	}

	var scores []types.Ext_ScoreData
	//fmt.Println(q, limit, offset, userid)
	rows, err := stmt.Query(limit)
	if err != nil {
		fmt.Println(err.Error())
		return scores, nil
	}
	defer rows.Close()

	for rows.Next() {
		var apiuser string
		var score types.Ext_ScoreData
		if err := rows.Scan(
			&score.ScoreId,
			&score.ScoreData.Title,
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
			&score.ApiUser.ID,
			&apiuser,
			&score.Beatmap.BeatmapID,
			&score.Beatmap.BeatmapSetID,
			&score.Beatmap.Maxcombo,
			&score.Beatmap.Version,
			&score.BeatmapSet.BeatmapSetID,
			&score.Artist,
			&score.Creator,
			&score.Tags,
			&score.CoverList,
			&score.BeatmapSet.Cover,
			&score.Preview,
			&score.Rankedstatus,
		); err != nil {

			fmt.Println(err.Error())
			return scores, nil
		}

		json.Unmarshal([]byte(apiuser), &score.ApiUser)
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
	LEFT JOIN ApiUsers on ApiUsers.Userid = Scoredata.Userid
	INNER JOIN Beatmap on Beatmap.BeatmapID = Scoredata.BeatmapID
	LEFT JOIN BeatmapSet on BeatmapSet.BeatmapSetID = Beatmap.BeatmapSetID
	WHERE Scoredata.Userid = ? 
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
		var apiuser string
		var score types.Ext_ScoreData
		if err := rows.Scan(
			&score.ScoreId,
			&score.ScoreData.Title,
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
			&score.ApiUser.ID,
			&apiuser,
			&score.Beatmap.BeatmapID,
			&score.Beatmap.BeatmapSetID,
			&score.Beatmap.Maxcombo,
			&score.Beatmap.Version,
			&score.BeatmapSet.BeatmapSetID,
			&score.Artist,
			&score.Creator,
			&score.Tags,
			&score.CoverList,
			&score.BeatmapSet.Cover,
			&score.Preview,
			&score.Rankedstatus,
		); err != nil {

			fmt.Println(err.Error())
			return scores, nil
		}

		json.Unmarshal([]byte(apiuser), &score.ApiUser)
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

func (s *SQLite) SaveApiUser(User types.ApiUser) error {

	us, err := json.Marshal(User)
	if err != nil {
		return err
	}

	stmt, err := s.DB.Prepare(`INSERT OR REPLACE INTO ApiUsers (USERID, APIUSER)
	VALUES(?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(User.ID, us)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLite) GetApiUser(Userid int) (types.ApiUser, error) {

	stmt, err := s.DB.Prepare(`SELECT ApiUser FROM ApiUsers WHERE USERID = ?`)
	if err != nil {
		return types.ApiUser{}, err
	}
	defer stmt.Close()

	var us string

	err = stmt.QueryRow(Userid).Scan(&us)
	if err != nil {
		return types.ApiUser{}, err
	}

	var u types.ApiUser

	err = json.Unmarshal([]byte(us), &u)
	if err != nil {
		return types.ApiUser{}, err
	}

	return u, nil
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
		Creator      TEXT,
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

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_screentime_date ON ScreenTime(Userid, Date)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Users (
		Userid INTEGER NOT NULL PRIMARY KEY,
		Username    TEXT,
		Mode		TEXT,
		Tokentype TEXT,
		Expiresin INTEGER,
		Accesstoken TEXT,
		Refreshtoken TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_users ON Users(Userid)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS ApiUsers (
		Userid  INTEGER NOT NULL PRIMARY KEY,
		ApiUser TEXT
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_apiusers ON ApiUsers(Userid)`)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *SQLite) MockScores(length int) error {
	userids := []int{14100399, 11705938, 2, 18767550, 124493}
	usernames := []string{"JuLi0n_", "Nikurax", "Peppy", "Seru", "Chocomint"}
	banners := []string{"https://assets.ppy.sh/user-profile-covers/14100399/cd1600936d7a56115cd147f47169addcf8f812133861b667c7dd3d177ca5068d.jpeg",
		"https://assets.ppy.sh/user-profile-covers/11705938/c8ccd57ec0ba71939052918bde270fe396061a8b33ad2076a410b24fb23a74e7.png",
		"https://assets.ppy.sh/user-profile-covers/2/baba245ef60834b769694178f8f6d4f6166c5188c740de084656ad2b80f1eea7.jpeg",
		"https://assets.ppy.sh/user-profile-covers/18767550/0a1991467c0b417840b2b2ba2b5d6e26a18cab42cd89f9a2fb9948ab3be37af4.gif",
		"https://osu.ppy.sh/images/headers/profile-covers/c6.jpg",
	}
	countrys := []string{"Germany", "Germany", "Australia", "Saudi Arabia", "South Korea"}
	countrycodes := []string{"DE", "DE", "AU", "SA", "KR"}

	for i := range userids {
		u := types.UserContext{
			ApiUser: types.ApiUser{
				ID:        userids[i],
				Username:  usernames[i],
				AvatarURL: fmt.Sprintf("https://a.ppy.sh/%d", userids[i]),
				CoverURL:  banners[i],

				Country: types.Country{
					Code: countrycodes[i],
					Name: countrys[i],
				},
				IsActive: false,
				IsOnline: false,
			},
			User: types.User{
				Username: usernames[i],
				UserId:   userids[i],
			},
		}

		err := s.SaveApiUser(u.ApiUser)
		if err != nil {
			panic(err.Error())
		}
	}

	for i := range length {

		beatmapid := gofakeit.Int()
		beatmapsetid := gofakeit.IntN(1000000)
		fcpp := gofakeit.Float64Range(50, 800)
		maxcombo := gofakeit.IntRange(200, 1000)

		scoreData := types.ScoreData{
			ScoreId:     fmt.Sprint(gofakeit.Int() + i),
			Title:       gofakeit.BeerName(),
			Date:        fmt.Sprint(gofakeit.PastDate()),
			BeatmapID:   beatmapid,
			Playtype:    gofakeit.RandomString([]string{"Standard", "Taiko", "Catch the Beat", "Mania"}),
			Ar:          gofakeit.Float64Range(7, 10),
			Cs:          gofakeit.Float64Range(2.5, 6.5),
			Hp:          gofakeit.Float64Range(7, 10),
			Od:          gofakeit.Float64Range(7, 10),
			SR:          gofakeit.Float64Range(2, 10),
			Bpm:         gofakeit.Float64Range(50, 280),
			Userid:      gofakeit.RandomInt(userids),
			ACC:         gofakeit.Float64Range(85, 100),
			Score:       gofakeit.Int(),
			Combo:       gofakeit.IntRange(10, maxcombo),
			Hit50:       gofakeit.IntN(15),
			Hit100:      gofakeit.IntN(200),
			Hit300:      gofakeit.IntRange(300, 3000),
			Ur:          gofakeit.Float64Range(80, 150),
			HitMiss:     gofakeit.IntN(10),
			Mode:        gofakeit.Number(0, 3),
			Mods:        gofakeit.RandomString([]string{"", "", "", "", "", "", "hd,hr", "dt,hd", "hd", "hr", "dt", "ez", "ht", "nc", "nf", "fl"}),
			Time:        gofakeit.Number(60, 300),
			PP:          gofakeit.Float64Range(0, fcpp),
			AIM:         gofakeit.Float64Range(80, 100),
			SPEED:       gofakeit.Float64Range(80, 100),
			ACCURACYATT: gofakeit.Float64Range(80, 100),
			Grade:       gofakeit.RandomString([]string{"SS", "S", "A", "B", "C", "D"}),
			FCPP:        fcpp,
		}

		beatmap := types.Beatmap{
			BeatmapID:    beatmapid,
			BeatmapSetID: beatmapsetid,
			Maxcombo:     maxcombo,
			Version:      gofakeit.RandomString([]string{"Easy", "Normal", "Hard", "Expert"}),
		}

		// Generate random data for BeatmapSet
		beatmapSet := types.BeatmapSet{
			BeatmapSetID: beatmapsetid,
			Artist:       gofakeit.Name(),
			Creator:      gofakeit.Username(),
			Tags:         gofakeit.RandomString([]string{"music", "game", "fun"}),
			CoverList:    gofakeit.RandomString([]string{"https://assets.ppy.sh/beatmaps/1800953/covers/list@2x.jpg?1686054624", "https://assets.ppy.sh/beatmaps/628368/covers/cover@2x.jpg?1650651127"}),
			Cover:        gofakeit.RandomString([]string{"https://assets.ppy.sh/beatmaps/1800953/covers/cover@2x.jpg?1686054624", "https://assets.ppy.sh/beatmaps/628368/covers/cover@2x.jpg?1650651127", "https://assets.ppy.sh/beatmaps/1895624/covers/cover@2x.jpg?1672353044", "https://assets.ppy.sh/beatmaps/426890/covers/cover@2x.jpg?1683195501", "https://assets.ppy.sh/beatmaps/1174754/covers/cover@2x.jpg?1650697061"}),
			Preview:      gofakeit.RandomString([]string{"https://b.ppy.sh/preview/628368.mp3", "https://b.ppy.sh/preview/1800953.mp3"}),
			Rankedstatus: gofakeit.RandomString([]string{"Ranked", "Approved", "Qualified", "Pending"}),
		}

		// Compose them into an instance of ExtScoreData
		extScoreData := types.Ext_ScoreData{
			ScoreData:  scoreData,
			Beatmap:    beatmap,
			BeatmapSet: beatmapSet,
		}

		//		fmt.Println(extScoreData)
		s.SaveExtendedScore(extScoreData)
	}

	return nil
}
