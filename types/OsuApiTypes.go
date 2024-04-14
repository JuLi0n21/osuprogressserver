package types

import "time"

type AuthUser struct {
	Tokentype    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TimeStamp    time.Time
}

type ApiUser struct {
	AvatarURL     string    `json:"avatar_url"`
	CountryCode   string    `json:"country_code"`
	DefaultGroup  string    `json:"default_group"`
	ID            int       `json:"id"`
	IsActive      bool      `json:"is_active"`
	IsBot         bool      `json:"is_bot"`
	IsDeleted     bool      `json:"is_deleted"`
	IsOnline      bool      `json:"is_online"`
	IsSupporter   bool      `json:"is_supporter"`
	LastVisit     time.Time `json:"last_visit"`
	PmFriendsOnly bool      `json:"pm_friends_only"`
	ProfileColour string    `json:"profile_colour"`
	Username      string    `json:"username"`
	CoverURL      string    `json:"cover_url"`
	Discord       string    `json:"discord"`
	HasSupported  bool      `json:"has_supported"`
	Interests     any       `json:"interests"`
	JoinDate      time.Time `json:"join_date"`
	Kudosu        struct {
		Total     int `json:"total"`
		Available int `json:"available"`
	} `json:"kudosu"`
	Location     any      `json:"location"`
	MaxBlocks    int      `json:"max_blocks"`
	MaxFriends   int      `json:"max_friends"`
	Occupation   any      `json:"occupation"`
	Playmode     string   `json:"playmode"`
	Playstyle    []string `json:"playstyle"`
	PostCount    int      `json:"post_count"`
	ProfileOrder []string `json:"profile_order"`
	Title        any      `json:"title"`
	Twitter      string   `json:"twitter"`
	Website      string   `json:"website"`
	Country      struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"country"`
	Cover struct {
		CustomURL string `json:"custom_url"`
		URL       string `json:"url"`
		ID        any    `json:"id"`
	} `json:"cover"`
	IsRestricted           bool  `json:"is_restricted"`
	AccountHistory         []any `json:"account_history"`
	ActiveTournamentBanner any   `json:"active_tournament_banner"`
	Badges                 []struct {
		AwardedAt   time.Time `json:"awarded_at"`
		Description string    `json:"description"`
		Image2XURL  string    `json:"image@2x_url"`
		ImageURL    string    `json:"image_url"`
		URL         string    `json:"url"`
	} `json:"badges"`
	FavouriteBeatmapsetCount int `json:"favourite_beatmapset_count"`
	FollowerCount            int `json:"follower_count"`
	GraveyardBeatmapsetCount int `json:"graveyard_beatmapset_count"`
	Groups                   []struct {
		ID          int    `json:"id"`
		Identifier  string `json:"identifier"`
		Name        string `json:"name"`
		ShortName   string `json:"short_name"`
		Description string `json:"description"`
		Colour      string `json:"colour"`
	} `json:"groups"`
	LovedBeatmapsetCount int `json:"loved_beatmapset_count"`
	MonthlyPlaycounts    []struct {
		StartDate string `json:"start_date"`
		Count     int    `json:"count"`
	} `json:"monthly_playcounts"`
	Page struct {
		HTML string `json:"html"`
		Raw  string `json:"raw"`
	} `json:"page"`
	PendingBeatmapsetCount int   `json:"pending_beatmapset_count"`
	PreviousUsernames      []any `json:"previous_usernames"`
	RankedBeatmapsetCount  int   `json:"ranked_beatmapset_count"`
	ReplaysWatchedCounts   []struct {
		StartDate string `json:"start_date"`
		Count     int    `json:"count"`
	} `json:"replays_watched_counts"`
	ScoresFirstCount int `json:"scores_first_count"`
	Statistics       struct {
		Level struct {
			Current  int `json:"current"`
			Progress int `json:"progress"`
		} `json:"level"`
		Pp                     float64 `json:"pp"`
		GlobalRank             int     `json:"global_rank"`
		RankedScore            int     `json:"ranked_score"`
		HitAccuracy            float64 `json:"hit_accuracy"`
		PlayCount              int     `json:"play_count"`
		PlayTime               int     `json:"play_time"`
		TotalScore             int     `json:"total_score"`
		TotalHits              int     `json:"total_hits"`
		MaximumCombo           int     `json:"maximum_combo"`
		ReplaysWatchedByOthers int     `json:"replays_watched_by_others"`
		IsRanked               bool    `json:"is_ranked"`
		GradeCounts            struct {
			Ss  int `json:"ss"`
			SSH int `json:"ssh"`
			S   int `json:"s"`
			Sh  int `json:"sh"`
			A   int `json:"a"`
		} `json:"grade_counts"`
		Rank struct {
			Global  int `json:"global"`
			Country int `json:"country"`
		} `json:"rank"`
	} `json:"statistics"`
	SupportLevel     int `json:"support_level"`
	UserAchievements []struct {
		AchievedAt    time.Time `json:"achieved_at"`
		AchievementID int       `json:"achievement_id"`
	} `json:"user_achievements"`
	RankHistory struct {
		Mode string `json:"mode"`
		Data []int  `json:"data"`
	} `json:"rank_history"`
}
