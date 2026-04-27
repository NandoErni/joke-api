package models

import "gorm.io/gorm"

type Joke struct {
	gorm.Model
	Text string
}
