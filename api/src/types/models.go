package types

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name string
}

type User struct {
	gorm.Model
	Email    string
	Name     string
	Password string
}

type UserConfig struct {
	gorm.Model
}

type WorkDay struct {
	gorm.Model
	Data     string
	Day_type string // on_site / remote / off
}
