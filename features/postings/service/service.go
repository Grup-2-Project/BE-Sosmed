package service

import (
	"BE-Sosmed/features/postings"
	"BE-Sosmed/helper/jwt"
	"errors"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type PostingService struct {
	m postings.Repository
}

func New(model postings.Repository) postings.Service {
	return &PostingService{
		m: model,
	}
}

func (ps *PostingService) TambahPosting(token *golangjwt.Token, newPosting postings.Posting) (postings.Posting, error) {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return postings.Posting{}, err
	}

	result, err := ps.m.InsertPosting(userID, newPosting)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return postings.Posting{}, errors.New("posting sudah ada pada sistem")
		}
		return postings.Posting{}, errors.New("terjadi kesalahan server")
	}

	return result, nil
}

func (ps *PostingService) SemuaPosting() ([]postings.Posting, error) {
	posts, err := ps.m.GetAllPost()

	if err != nil {
		return nil, err
	}

	return posts, nil
}


func (ps *PostingService) UpdatePosting(token *golangjwt.Token, updatePosting postings.Posting) (postings.Posting, error) {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return postings.Posting{}, err
	}


	updatedPost, err := ps.m.UpdatePost(userID, updatePosting)
	if err != nil {
		return postings.Posting{}, err
	}

	return updatedPost, nil
}



func (ps *PostingService) DeletePosting(token *golangjwt.Token, postID uint) (error) {
    userID, err := jwt.ExtractToken(token)
    if err != nil {
        return err
    }

    err = ps.m.DeletePost(userID, postID)
    if err != nil {
        return err
    }

    return nil
}