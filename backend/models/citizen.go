package models

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type gender string

// Enum values for Status
const (
	Male   gender = "Male"
	Female gender = "Female"
	Trans  gender = "Trans"
)

type StateEnum string

const (
	AndhraPradesh    StateEnum = "Andhra Pradesh"
	ArunachalPradesh StateEnum = "Arunachal Pradesh"
	Assam            StateEnum = "Assam"
	Bihar            StateEnum = "Bihar"
	Chhattisgarh     StateEnum = "Chhattisgarh"
	Goa              StateEnum = "Goa"
	Gujarat          StateEnum = "Gujarat"
	Haryana          StateEnum = "Haryana"
	HimachalPradesh  StateEnum = "Himachal Pradesh"
	Jharkhand        StateEnum = "Jharkhand"
	Karnataka        StateEnum = "Karnataka"
	Kerala           StateEnum = "Kerala"
	MadhyaPradesh    StateEnum = "Madhya Pradesh"
	Maharashtra      StateEnum = "Maharashtra"
	Manipur          StateEnum = "Manipur"
	Meghalaya        StateEnum = "Meghalaya"
	Mizoram          StateEnum = "Mizoram"
	Nagaland         StateEnum = "Nagaland"
	Odisha           StateEnum = "Odisha"
	Punjab           StateEnum = "Punjab"
	Rajasthan        StateEnum = "Rajasthan"
	Sikkim           StateEnum = "Sikkim"
	TamilNadu        StateEnum = "Tamil Nadu"
	Telangana        StateEnum = "Telangana"
	Tripura          StateEnum = "Tripura"
	UttarPradesh     StateEnum = "Uttar Pradesh"
	Uttarakhand      StateEnum = "Uttarakhand"
	WestBengal       StateEnum = "West Bengal"
)

type Citizen struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name,omitempty" validate:"required"`
	LastName  string             `bson:"last_name,omitempty" validate:"required"`
	DOB       string             `bson:"dob,omitempty" validate:"required"`
	Gender    gender             `bson:"gender,omitempty" validate:"oneof=Male Female Trans"`
	Address   string             `bson:"address,omitempty" validate:"required"`
	City      string             `bson:"city,omitempty" validate:"required"`
	PinCode   string             `bson:"pin_code,omitempty" validate:"required,len=6"`
	State     StateEnum          `bson:"state,omitempty" validate:"required,oneof=Andhra Pradesh Arunachal Pradesh Assam Bihar Chhattisgarh Goa Gujarat Haryana Himachal Pradesh Jharkhand Karnataka Kerala Madhya Pradesh Maharashtra Manipur Meghalaya Mizoram Nagaland Odisha Punjab Rajasthan Sikkim Tamil Nadu Telangana Tripura Uttar Pradesh Uttarakhand West Bengal"`
}

func (c *Citizen) ValidateStruct(db *gorm.DB) error {
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		return err
	}
	return nil
}
