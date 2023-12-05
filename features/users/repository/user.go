package repository

import (
	cr "BE-Sosmed/features/comments/repository"
	pr "BE-Sosmed/features/postings/repository"
	"BE-Sosmed/features/users"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	FirstName string
	LastName  string
	Gender    string
	Hp        string
	Email     string `gorm:"unique"`
	Password  string
	Image     string
	Username  string            `gorm:"unique"`
	Postings  []pr.PostingModel `gorm:"foreignKey:UserID"`
	Comments  []cr.CommentModel `gorm:"foreignKey:UserID"`
}

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.Repository {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) InsertUser(newUser users.User) (users.User, error) {
	var inputDB = new(UserModel)
	inputDB.FirstName = newUser.FirstName
	inputDB.LastName = newUser.LastName
	inputDB.Gender = newUser.Gender
	inputDB.Hp = newUser.Hp
	inputDB.Email = newUser.Email
	inputDB.Password = newUser.Password
	inputDB.Username = newUser.Username
	inputDB.Image = "default"

	if err := uq.db.Create(&inputDB).Error; err != nil {
		return users.User{}, err
	}

	newUser.ID = inputDB.ID

	return newUser, nil
}

func (uq *userQuery) Login(email string) (users.User, error) {
	var userData = new(UserModel)

	if err := uq.db.Where("email = ?", email).First(userData).Error; err != nil {
		return users.User{}, err
	}

	var result = new(users.User)
	result.ID = userData.ID
	result.FirstName = userData.FirstName
	result.Password = userData.Password
	result.Username = userData.Username

	return *result, nil
}

func (uq *userQuery) ReadUserById(UserID uint) (users.User, error) {
	var userData UserModel

	if err := uq.db.Where("id = ?", UserID).First(&userData).Error; err != nil {
		return users.User{}, err
	}

	var result = new(users.User)
	result.ID = userData.ID
	result.FirstName = userData.FirstName
	result.Username = userData.Username

	return *result, nil
}

func (uq *userQuery) UpdateUser(UserID uint, updatedUser users.User) (users.User, error) {
	var userData UserModel
	if err := uq.db.First(&userData, UserID).Error; err != nil {
		return users.User{}, err
	}

	userData.FirstName = updatedUser.FirstName
	userData.LastName = updatedUser.LastName
	userData.Username = updatedUser.Username
	userData.Hp = updatedUser.Hp
	userData.Image = updatedUser.Image
	userData.Password = updatedUser.Password

	if err := uq.db.Save(&userData).Error; err != nil {
		return users.User{}, err
	}

	var result = new(users.User)
	result.ID = userData.ID
	result.FirstName = userData.FirstName
	result.Username = userData.Username

	return *result, nil
}

func (uq *userQuery) DeleteUser(UserID uint) error {
	if err := uq.db.Delete(&UserModel{}, UserID).Error; err != nil {
		return err
	}

	return nil
}
