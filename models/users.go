package models

import (
	"time"
)

// Model
type Users struct {
	ID        int64     `gorm:"primary_key;auto_incriment" json:"id"`
	FirstName string    `gorm:"size:50" json:"fitst_name"`
	LastName  string    `gorm:"size:50" json:"last_name"`
	Sex       string    `gorm:"size:10" json:"sex"`
	Age       string    `gorm:"size:20" json:"age"`
	Email     string    `gorm:"size:200" json:"email"`
	Phone     string    `gorm:"size:50" json:"phone"`
	Username  string    `gorm:"size:200" json:"username"`
	Password  string    `gorm:"size:50" json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (users *Users) TableName() string {
	return "users"
}

func (users *Users) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = users.ID
	resp["first_name"] = users.FirstName
	resp["last_name"] = users.LastName
	resp["sex"] = users.Sex
	resp["age"] = users.Age
	resp["email"] = users.Email
	resp["phone"] = users.Phone
	resp["username"] = users.Username
	resp["password"] = users.Password

	return resp
}
