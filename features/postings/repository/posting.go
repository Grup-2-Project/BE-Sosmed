package repository

import (
	"gorm.io/gorm"
	"BE-Sosmed/features/postings"
)

type PostingModel struct {
	gorm.Model
	Artikel string
	Gambar      string
	UserID      uint
}

type postingQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) postings.Repository {
	return &postingQuery{
		db: db,
	}
}

func (pq *postingQuery) InsertPosting(userID uint, newPosting postings.Posting) (postings.Posting, error) {
	var inputData = new(PostingModel)
	inputData.UserID = userID
	inputData.Artikel = newPosting.Artikel
	inputData.Gambar = newPosting.Gambar

	if err := pq.db.Create(&inputData).Error; err != nil {
		return postings.Posting{}, err
	}

	newPosting.ID = inputData.ID

	return newPosting, nil
}