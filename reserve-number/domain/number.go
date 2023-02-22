package domain

import "time"

type Number struct {
	UUID      string    `json:"uuid" form:"uuid" gorm:"primary_key"`
	Number    int64     `json:"number" form:"number" binding:"required"`
	CreatedAt time.Time `gorm:"-" bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"-" bson:"updated_at" json:"updated_at,omitempty"`
	DeletedAt time.Time `gorm:"-" bson:"deleted_at" json:"deleted_at,omitempty"`
}
