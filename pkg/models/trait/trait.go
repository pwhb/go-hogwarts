package trait

type Trait struct {
	PublicTrait
	ID uint `gorm:"primarykey" json:"id"`
}

type PublicTrait struct {
	Name  string `json:"name,omitempty"`
	House uint   `json:"house,omitempty"`
}
