package entity

import(
	"gorm.io/gorm"
)

type Member struct{
	gorm.Model
	Phone string `valid:"required~Phone is required, stringlength(10|10)~Phone is not 10 digits"`
	Password string
	Url string `valid:"required~Url is required, url~Url is invalid"`
}