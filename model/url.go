package model

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	Original string `gorm:"size:512;not null" json:"original"`
	Short    string `gorm:"size:64;unique;not null" json:"short"`
	Lifespan int    `gorm:"not null" json:"lifespan"`
}
