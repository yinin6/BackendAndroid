package models

type APIResponse struct {
	Status    string `json:"status"`
	Data      Data   `json:"data"`
	Token     string `json:"token"`
	IPAddress string `json:"ipAddress"`
}

type Data struct {
	ID                string   `json:"id"`
	Content           string   `json:"content"`
	Popularity        int      `json:"popularity"`
	Origin            Origin   `json:"origin"`
	MatchTags         []string `json:"matchTags"`
	RecommendedReason string   `json:"recommendedReason"`
	CacheAt           string   `json:"cacheAt"`
}

type Origin struct {
	Title     string   `json:"title"`
	Dynasty   string   `json:"dynasty"`
	Author    string   `json:"author"`
	Content   []string `json:"content"`
	Translate []string `json:"translate"`
}

type UserFavorite struct {
	UserID    string `json:"user_id" db:"user_id"`       // 用户ID
	PoemID    string `json:"poem_id" db:"poem_id"`       // 古诗ID
	CreatedAt string `json:"created_at" db:"created_at"` // 收藏时间
}
