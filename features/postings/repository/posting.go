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

func (pq *postingQuery) GetAllPost() ([]postings.Posting, error) {
	var posts []PostingModel

	if err := pq.db.Order("created_at desc").Find(&posts).Error; err != nil {
		return nil, err
	}

	var result []postings.Posting
	for _, post := range posts {
		result = append(result, postings.Posting{
			ID:     post.ID,
			Artikel: post.Artikel,
			Gambar:  post.Gambar,
			UserID:  post.UserID,
		})
	}

	return result, nil
}