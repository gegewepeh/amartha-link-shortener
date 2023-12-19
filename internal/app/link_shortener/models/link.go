package models

type Link struct {
	ID        string `json:"id,omitempty"`
	UserID    string `json:"userId,omitempty"`
	FullLink  string `json:"fullLink,omitempty"`
	Slug      string `json:"slug,omitempty"`
	Visit     int    `json:"visit"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt int    `json:"updatedAt,omitempty"`
}

type CreateSlugRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullLink string `json:"fullLink"`
}

type UpdateSlugRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Slug string `json:"slug"`
}