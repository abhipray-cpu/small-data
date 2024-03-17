package models

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Role string

// Enum values for Role.
const (
	Admin    Role = "Admin"
	Employee Role = "Employee"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"user_name,omitempty" validate:"required"`
	Email    string             `bson:"user_mail,omitempty"  validate:"required,email"`
	Type     Role               `bson:"type,omitempty" validate:"oneof=Admin Employee"`
	Password string             `bson:"password,omitempty" validate:"required"`
}

// ValidateStruct validates the User struct using the validator package.
func (u *User) ValidateStruct(db *gorm.DB) error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}
