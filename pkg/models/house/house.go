package house

import (
	"time"

	"gorm.io/gorm"
)

type House struct {
	PublicHouse
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type PublicHouse struct {
	Name       string `json:"name,omitempty"`
	Founder    string `json:"founder,omitempty"`
	Animal     string `json:"animal,omitempty"`
	Element    string `json:"element,omitempty"`
	CommonRoom string `json:"common_room,omitempty"`
}
