package models

type UserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Salt     []byte `json:"salt"`
}

type User struct {
	ID       uint   `gorm:"column:id;primaryKey;not null"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Salt     string `json:"salt"`
}
