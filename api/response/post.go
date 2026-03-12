package response

import "time"

type PostCategoryRes struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PostTagRes struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type PostMetaRes struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Summary     *string    `json:"summary"`
	Cover       *string    `json:"cover,omitempty"`
	ReadTime    uint       `json:"read_time_minutes"`
	ViewCount   uint       `json:"view_count"`
	Status      string    `json:"status"`
	PublishedAt *time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Author      string    `json:"author"`

	Categories []*PostCategoryRes `json:"categories"`
	Tags       []*PostTagRes      `json:"tags"`
}

type PostDetailDTO struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`
	Cover       string    `json:"cover,omitempty"`
	ReadTime    int       `json:"read_time_minutes"`
	ViewCount   int       `json:"view_count"`
	Status      string    `json:"status"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Author      string    `json:"author"`

	Categories []PostCategoryRes `json:"categories"`
	Tags       []PostTagRes      `json:"tags"`
}
