package models 

import (
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model 
	Question string `json:"question" gorm:"text: not null;default: null"`
	Answer string `json:"answer" gorm:"text: not null;default: null"` 
}