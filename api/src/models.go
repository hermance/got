package main

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	name string
}

type User struct {
	gorm.Model
	email    string
	name     string
	password string
}

type UserConfig struct {
	gorm.Model
}

type WorkDay struct {
	gorm.Model
	date     string
	day_type string // on_site / remote / off
}
