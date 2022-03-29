package model

import "time"

type Users struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Status    int8
	Dob       time.Time
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID        uint `gorm:"primaryKey"`
	NameProduct string
}
