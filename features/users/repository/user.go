package repository

import (
	"BE-Sosmed/features/users"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	FirstName     string
	LastName string
	Gender string
	Hp string
	Email string
	Password string
	Image string
	// Coupons  []repository.CouponModel `gorm:"foreignKey:UserID"`
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
result.Email = userData.Email
result.FirstName = userData.FirstName
result.LastName = userData.LastName
result.Gender = userData.Gender
result.Hp = userData.Hp
result.Password = userData.Password

	return *result, nil
}
