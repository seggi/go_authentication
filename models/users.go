package models

import (
	"time"
)

// Model
type User struct {
	ID        int64     `gorm:"primary_key;auto_incriment" json:"id"`
	FirstName string    `json:"fitst_name"`
	LastName  string    `json:"last_name"`
	Sex       string    `gorm:"size:10" json:"sex"`
	Age       string    `json:"age"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Returns the table name of the Model
func (users *User) TableName() string {
	return "user"
}

// Request Binding for User Login
type UserLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"passsword" binding:"required"`
}

// Request Binding for User Register
type UserRegister struct {
	Phone     string `form:"phone" json:"phone"`
	Email     string `form:"email" json:"email" binding:"required"`
	Username  string `form:"username" json:"username"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
	Sex       string `form:"sex"`
}

func (user *User) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = user.ID
	resp["first_name"] = user.FirstName
	resp["last_name"] = user.LastName
	resp["sex"] = user.Sex
	resp["age"] = user.Age
	resp["email"] = user.Email
	resp["phone"] = user.Phone
	resp["username"] = user.Username
	resp["password"] = user.Password
	resp["is_active"] = user.IsActive

	return resp
}
