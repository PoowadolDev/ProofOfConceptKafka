package main

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string
	Amout int
}

type Order struct {
	gorm.Model
	UserUid string
	Status  string
}
