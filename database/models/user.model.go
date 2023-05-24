package models

import "time"

type User struct {
	UserID       string    `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	Name         string    `json:"name,omitempty" valid:"-"`
	Username     string    `json:"username,omitempty" valid:"type(string),required~Username cannot be empty"`
	Email        string    `json:"email" valid:"email,required~Email cannot be empty"`
	Password     string    `json:"password,omitempty" valid:"type(string),required~Password cannot be empty"`
	Gender       string    `json:"gender" valid:"-"`
	Country      string    `json:"country" valid:"-"`
	DateOfBirth  string    `json:"date_of_birth,omitempty" schema:"date_of_birth" valid:"-"`
	MobileNumber string    `json:"mobile_number,omitempty" schema:"mobile_number" valid:"-"`
	Role         string    `json:"role,omitempty" gorm:"default: user" valid:"-"`
	Image        string    `json:"image,omitempty" valid:"-"`
	TokenVerify  string    `json:"token_verify" valid:"-"`
	IsVerified   bool      `json:"is_verified,omitempty" gorm:"default: false" valid:"-"`
	CreatedAt    time.Time `json:"created_at" valid:"-"`
	UpdatedAt    time.Time `json:"updated_at" valid:"-"`
}

type Users []User

func (User) TableName() string {
	return "users"
}
