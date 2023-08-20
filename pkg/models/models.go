package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentPortal struct {
	gorm.Model
	Name       string        `gorm:"type:varchar(20);not null"`
	Email      string        `gorm:"type:varchar(20);unique;not null"`
	Batch      string        `gorm:"type:varchar(10);not null"`
	Domain     string        `gorm:"type:varchar(10);not null"`
	Week       string        `gorm:"type:numeric(2);not null"`
	Status     string        `gorm:"type:status;not null;default:'active'"`
	Start_Date time.Time     `gorm:"type:date;not null"`
	End_Date   time.Time     `gorm:"type:date"`
	Duration   time.Duration `gorm:"type:interval"`
}
