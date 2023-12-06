package handler

type AddRequest struct {
	Artikel string `json:"artikel" form:"artikel" validate:"required"`
	Gambar  string `json:"gambar" form:"gambar"`
}

type AddResponse struct {
	Artikel string `json:"artikel"`
	Gambar  string `json:"gambar"`
}

type GetResponse struct {
	ID       uint          `json:"id"`
	Artikel  string        `json:"artikel"`
	Gambar   string        `json:"gambar"`
	UserID   uint          `json:"user_id"`
	Username string        `json:"username"`
	Image    string        `json:"foto_profil"`
	Comments []CommentInfo `json:"comments"`
}

type CommentInfo struct {
	ID       uint   `json:"comment_id"`
	Komentar string `json:"komentar"`
	UserID   uint   `json:"user_id"`
	PostID   uint   `json:"post_id"`
	Username string `json:"username"`
	Image    string `json:"foto_profil"`
}

type UpdateRequest struct {
	ID      uint   `json:"id" form:"id" validate:"required"`
	Artikel string `json:"artikel" form:"artikel" validate:"required"`
	Gambar  string `json:"gambar" form:"gambar"`
}

type UpdateResponse struct {
	Artikel string `json:"artikel"`
	Gambar  string `json:"gambar"`
}
