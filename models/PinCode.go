package models

type ResetPassword struct {
	ID       uint   `gorm:"column:id;primary_key;not null"`
	Pin      string `gorm:"size:255;not null" json:"pin"`
	Usermail string `gorm:"size:255;not null" json:"UserMail"`
}

type PinCodeInput struct {
	Pin      string `json:"pin"`
	Usermail string `json:"UserMail" binding:"required"`
	Password string `json:"Password"`
}
