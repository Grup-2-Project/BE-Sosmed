package handler

type AddRequest struct {
	Artikel string `json:"artikel" form:"artikel" validate:"required"`
	Gambar  string `json:"gambar" form:"gambar"`
}

type AddResponse struct {
	ID      uint   `json:"id"`
	Artikel string `json:"artikel"`
	Gambar  string `json:"gambar"`
}

type GetResponse struct {
	ID      uint   `json:"id"`
	Artikel string `json:"artikel"`
	Gambar  string `json:"gambar"`
	UserID  uint   `json:"user_id"`
}

type UpdateRequest struct {
	ID      uint   `json:"id" form:"id" validate:"required"`
	Artikel string `json:"artikel" form:"artikel" validate:"required"`
	Gambar  string `json:"gambar" form:"gambar"`
}

type UpdateResponse struct {
	ID      uint   `json:"id"`
	Artikel string `json:"artikel"`
	Gambar  string `json:"gambar"`
	UserID  uint   `json:"user_id"`
}
