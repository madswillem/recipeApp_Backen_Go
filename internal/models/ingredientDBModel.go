package models

type IngredientDBSchema struct {
	ID               uint             `gorm:"primarykey"`
	Name             string           `gorm:"not null"`
	StandardUnit     string           `gorm:"not null"`
	NutritionalValue NutritionalValue `gorm:"polymorphic:Owner"`
	UserWithAllergies []*UserModel	  `gorm:"many2many:user_allergies"`
}
