package types

import "gorm.io/gorm"

type Organization struct {
	gorm.Model `json:"-"`
	Name       string
}

type User struct {
	gorm.Model `json:"-"`
	Email      string
	Name       string
}

type UserConfig struct {
	gorm.Model `json:"-"`
}

type WorkDay struct {
	gorm.Model `json:"-"`
	Data       string
	Day_type   string // on_site / remote / off
}
