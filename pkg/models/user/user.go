package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	PublicUser
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type PublicUser struct {
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	BloodStatus string `json:"blood_status,omitempty"`
	House       string `json:"house,omitempty"`
	Species     string `json:"species,omitempty"`
	Occupation  string `json:"occupation,omitempty"`
}
