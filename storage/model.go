package storage

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"      json:"id"`
	Name        string    `gorm:"type:varchar(255);uniqueIndex" json:"name"`
	Country     string    `gorm:"type:varchar(100)"             json:"country"`
	Price       float64   `gorm:"type:decimal(10,2)"            json:"price"`
	ReleaseDate time.Time `gorm:"type:date"                     json:"release_date"`
}
