package domain

import "time"

type Category struct {
	ID          int       `json:"id" bson:"id"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Image       string    `json:"image" bson:"image"`
	Version     int    `json:"version" bson:"version"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Categories []Category

func NewCategory() *Category {
	now := time.Now()

	return &Category{
		Version: 0,
		CreatedAt: now,
		UpdatedAt: now,
	}
}