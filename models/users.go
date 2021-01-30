package models

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ErrNotFound = errors.New("models: resource not found")

	ErrInvalidId = errors.New("Invalid ID")
)

func NewUserService(connectInfo string) (*UserService, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(connectInfo), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}
	return &UserService{
		db: db,
	}, nil
}

func (us *UserService) DestructiveReset() error {

	if err := us.db.Migrator().DropTable(&User{}); err != nil {
		return err
	}
	return us.AutoMigrate()
}

func (us *UserService) AutoMigrate() error {

	if err := us.db.AutoMigrate(&User{}); err != nil {
		return err
	}
	return nil
}

type UserService struct {
	db *gorm.DB
}

func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id=?", id)
	err := first(db, &user)
	return &user, err
}

func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email=?", email)
	err := first(db, &user)
	return &user, err
}

func first(db *gorm.DB, user *User) error {
	err := db.First(user).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	} else {
		return err
	}
}

func (us *UserService) Create(user *User) error {
	if err := us.db.Create(user).Error; err != nil {
		fmt.Println("Debug:error in the create statement")
		return err
	}
	return nil
}

func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidId
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
