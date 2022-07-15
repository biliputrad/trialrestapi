package Entity

import (
	"gorm.io/gorm"
)

type User struct {
	ID        int
	UserName  string `gorm:"size:20;notnull"`
	Password  string `gorm:"notnull"`
	FirstName string `gorm:"size 30;notnull"`
	LastName  string `gorm:"size 30;notnull"`
	Email     string `gorm:"notnull"`
	Phone     uint64 `gorm:"notnull"`
	DeletedAt gorm.DeletedAt
}
type UserRequest struct {
	UserName  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname"`
	Phone     uint64 `json:"phone" binding:"required,number"`
}
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type Users []User

type ForgetPass struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    uint64 `json:"phone" binding:"required,number"`
	Password string `json:"password" binding:"required,min=6"`
}
type DisplayUser struct {
	ID        int    `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     uint64 `json:"phone"`
}

type PDFUser struct {
	ID        string
	UserName  string
	Email     string
	FirstName string
	LastName  string
	Phone     string
}

type DisplayUsers []DisplayUser
