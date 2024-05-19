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
	AvatarURL     string    `json:"avatar_url,omitempty"`
	CountryCode   string    `json:"country_code,omitempty"`
	DefaultGroup  string    `json:"default_group,omitempty"`
	ID            int       `json:"id,omitempty"`
	IsActive      bool      `json:"is_active,omitempty"`
	IsBot         bool      `json:"is_bot,omitempty"`
	IsDeleted     bool      `json:"is_deleted,omitempty"`
	IsOnline      bool      `json:"is_online,omitempty"`
	IsSupporter   bool      `json:"is_supporter,omitempty"`
	LastVisit     time.Time `json:"last_visit,omitempty"`
	PmFriendsOnly bool      `json:"pm_friends_only,omitempty"`
	ProfileColour string    `json:"profile_colour,omitempty"`
	Username      string    `json:"username,omitempty"`
	CoverURL      string    `json:"cover_url,omitempty"`
	Discord       string    `json:"discord,omitempty"`
	HasSupported  bool      `json:"has_supported,omitempty"`
	Interests     any       `json:"interests,omitempty"`
	JoinDate      time.Time `json:"join_date,omitempty"`
	Kudosu        struct {
		Total     int `json:"total,omitempty"`
		Available int `json:"available,omitempty"`
	} `json:"kudosu,omitempty"`
	Location     any      `json:"location,omitempty"`
	MaxBlocks    int      `json:"max_blocks,omitempty"`
	MaxFriends   int      `json:"max_friends,omitempty"`
	Occupation   any      `json:"occupation,omitempty"`
	Playmode     string   `json:"playmode,omitempty"`
	Playstyle    []string `json:"playstyle,omitempty"`
	PostCount    int      `json:"post_count,omitempty"`
	ProfileOrder []string `json:"profile_order,omitempty"`
	Title        any      `json:"title,omitempty"`
	Twitter      string   `json:"twitter,omitempty"`
	Website      string   `json:"website,omitempty"`
	Country      Country  `json:"country,omitempty"`
	Cover        struct {
		CustomURL string `json:"custom_url,omitempty"`
		URL       string `json:"url,omitempty"`
		ID        any    `json:"id,omitempty"`
	} `json:"cover,omitempty"`
	IsRestricted           bool  `json:"is_restricted,omitempty"`
	AccountHistory         []any `json:"account_history,omitempty"`
	ActiveTournamentBanner any   `json:"active_tournament_banner,omitempty"`
	Badges                 []struct {
		AwardedAt   time.Time `json:"awarded_at,omitempty"`
		Description string    `json:"description,omitempty"`
		Image2XURL  string    `json:"image@2x_url,omitempty"`
		ImageURL    string    `json:"image_url,omitempty"`
		URL         string    `json:"url,omitempty"`
	} `json:"badges,omitempty"`
	FavouriteBeatmapsetCount int `json:"favourite_beatmapset_count,omitempty"`
	FollowerCount            int `json:"follower_count,omitempty"`
	GraveyardBeatmapsetCount int `json:"graveyard_beatmapset_count,omitempty"`
	Groups                   []struct {
		ID          int    `json:"id,omitempty"`
		Identifier  string `json:"identifier,omitempty"`
		Name        string `json:"name,omitempty"`
		ShortName   string `json:"short_name,omitempty"`
		Description string `json:"description,omitempty"`
		Colour      string `json:"colour,omitempty"`
	} `json:"groups,omitempty"`
	LovedBeatmapsetCount int `json:"loved_beatmapset_count,omitempty"`
	MonthlyPlaycounts    []struct {
		StartDate string `json:"start_date,omitempty"`
		Count     int    `json:"count,omitempty"`
	} `json:"monthly_playcounts,omitempty"`
	Page struct {
		HTML string `json:"html,omitempty"`
		Raw  string `json:"raw,omitempty"`
	} `json:"page,omitempty"`
	PendingBeatmapsetCount int   `json:"pending_beatmapset_count,omitempty"`
	PreviousUsernames      []any `json:"previous_usernames,omitempty"`
	RankedBeatmapsetCount  int   `json:"ranked_beatmapset_count,omitempty"`
	ReplaysWatchedCounts   []struct {
		StartDate string `json:"start_date,omitempty"`
		Count     int    `json:"count,omitempty"`
	} `json:"replays_watched_counts,omitempty"`
	ScoresFirstCount int `json:"scores_first_count,omitempty"`
	Statistics       struct {
		Level struct {
			Current  int `json:"current,omitempty"`
			Progress int `json:"progress,omitempty"`
		} `json:"level,omitempty"`
		Pp                     float64 `json:"pp,omitempty"`
		GlobalRank             int     `json:"global_rank,omitempty"`
		RankedScore            int     `json:"ranked_score,omitempty"`
		HitAccuracy            float64 `json:"hit_accuracy,omitempty"`
		PlayCount              int     `json:"play_count,omitempty"`
		PlayTime               int     `json:"play_time,omitempty"`
		TotalScore             int     `json:"total_score,omitempty"`
		TotalHits              int     `json:"total_hits,omitempty"`
		MaximumCombo           int     `json:"maximum_combo,omitempty"`
		ReplaysWatchedByOthers int     `json:"replays_watched_by_others,omitempty"`
		IsRanked               bool    `json:"is_ranked,omitempty"`
		GradeCounts            struct {
			Ss  int `json:"ss,omitempty"`
			SSH int `json:"ssh,omitempty"`
			S   int `json:"s,omitempty"`
			Sh  int `json:"sh,omitempty"`
			A   int `json:"a,omitempty"`
		} `json:"grade_counts,omitempty"`
		Rank struct {
			Global  int `json:"global,omitempty"`
			Country int `json:"country,omitempty"`
		} `json:"rank,omitempty"`
	} `json:"statistics,omitempty"`
	SupportLevel     int `json:"support_level,omitempty"`
	UserAchievements []struct {
		AchievedAt    time.Time `json:"achieved_at,omitempty"`
		AchievementID int       `json:"achievement_id,omitempty"`
	} `json:"user_achievements,omitempty"`
	RankHistory struct {
		Mode string `json:"mode,omitempty"`
		Data []int  `json:"data,omitempty"`
	} `json:"rank_history,omitempty"`
}

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type UserContext struct {
	User     User
	ApiUser  ApiUser
	Cookieid string
}
