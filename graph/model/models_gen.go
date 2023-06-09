// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Arms struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type InputArms struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type InputLegs struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type InputOwnedPartPart struct {
	Part PartType `json:"Part"`
	ID   string   `json:"id"`
}

type InputPreviousOwnerData struct {
	User      *string  `json:"user,omitempty"`
	SellPrice *float64 `json:"sellPrice,omitempty"`
}

type Legs struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type ModifyPart struct {
	ID             string                    `json:"id"`
	Part           PartType                  `json:"part"`
	Name           string                    `json:"name"`
	Description    string                    `json:"description"`
	Price          float64                   `json:"price"`
	ImageLocation  *string                   `json:"imageLocation,omitempty"`
	Owner          string                    `json:"owner"`
	PreviousPrices []*InputPreviousOwnerData `json:"previousPrices,omitempty"`
	Side           Side                      `json:"side"`
}

type ModifyTeddyBear struct {
	ID         string     `json:"id"`
	TotalPrice *float64   `json:"totalPrice,omitempty"`
	Head       *string    `json:"head,omitempty"`
	Body       *string    `json:"body,omitempty"`
	Arms       *InputArms `json:"arms,omitempty"`
	Legs       *InputLegs `json:"legs,omitempty"`
}

type ModifyUser struct {
	ID             string                `json:"id"`
	Username       string                `json:"username"`
	Password       string                `json:"password"`
	OwnedBears     []string              `json:"ownedBears,omitempty"`
	OwnedParts     []*InputOwnedPartPart `json:"ownedParts,omitempty"`
	AccountBalance *float64              `json:"accountBalance,omitempty"`
	LastPassword   *string               `json:"lastPassword,omitempty"`
	Email          string                `json:"email"`
}

type NewPart struct {
	Part           PartType                  `json:"part"`
	Name           string                    `json:"name"`
	Description    string                    `json:"description"`
	Price          float64                   `json:"price"`
	ImageLocation  *string                   `json:"imageLocation,omitempty"`
	Owner          string                    `json:"owner"`
	PreviousPrices []*InputPreviousOwnerData `json:"previousPrices,omitempty"`
	Side           Side                      `json:"side"`
}

type NewTeddyBear struct {
	TotalPrice *float64   `json:"totalPrice,omitempty"`
	Head       *string    `json:"head,omitempty"`
	Body       *string    `json:"body,omitempty"`
	Arms       *InputArms `json:"arms,omitempty"`
	Legs       *InputLegs `json:"legs,omitempty"`
}

type NewUser struct {
	Username       string                `json:"username"`
	Password       string                `json:"password"`
	OwnedBears     []string              `json:"ownedBears,omitempty"`
	OwnedParts     []*InputOwnedPartPart `json:"ownedParts,omitempty"`
	AccountBalance *float64              `json:"accountBalance,omitempty"`
	LastPassword   *string               `json:"lastPassword,omitempty"`
	Email          string                `json:"email"`
}

type OwnedPartPart struct {
	Part PartType `json:"Part"`
	ID   string   `json:"id"`
}

type Part struct {
	ID             string               `json:"id"`
	Part           PartType             `json:"part"`
	Name           string               `json:"name"`
	Description    string               `json:"description"`
	Price          float64              `json:"price"`
	ImageLocation  *string              `json:"imageLocation,omitempty"`
	Owner          string               `json:"owner"`
	PreviousPrices []*PreviousOwnerData `json:"previousPrices,omitempty"`
	Side           Side                 `json:"side"`
}

type PreviousOwnerData struct {
	User      *string  `json:"user,omitempty"`
	SellPrice *float64 `json:"sellPrice,omitempty"`
}

type TeddyBear struct {
	ID         string  `json:"id"`
	TotalPrice float64 `json:"totalPrice"`
	Head       *string `json:"head,omitempty"`
	Body       *string `json:"body,omitempty"`
	Arms       *Arms   `json:"arms,omitempty"`
	Legs       *Legs   `json:"legs,omitempty"`
}

type User struct {
	ID             string           `json:"id"`
	Username       string           `json:"username"`
	Password       string           `json:"password"`
	OwnedBears     []string         `json:"ownedBears,omitempty"`
	OwnedParts     []*OwnedPartPart `json:"ownedParts,omitempty"`
	AccountBalance *float64         `json:"accountBalance,omitempty"`
	LastPassword   *string          `json:"lastPassword,omitempty"`
	Email          string           `json:"email"`
}

type PartType string

const (
	PartTypeLeg  PartType = "Leg"
	PartTypeArm  PartType = "Arm"
	PartTypeBody PartType = "Body"
	PartTypeHead PartType = "Head"
)

var AllPartType = []PartType{
	PartTypeLeg,
	PartTypeArm,
	PartTypeBody,
	PartTypeHead,
}

func (e PartType) IsValid() bool {
	switch e {
	case PartTypeLeg, PartTypeArm, PartTypeBody, PartTypeHead:
		return true
	}
	return false
}

func (e PartType) String() string {
	return string(e)
}

func (e *PartType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PartType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PartType", str)
	}
	return nil
}

func (e PartType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Side string

const (
	SideLeft  Side = "left"
	SideRight Side = "right"
	SideBoth  Side = "both"
)

var AllSide = []Side{
	SideLeft,
	SideRight,
	SideBoth,
}

func (e Side) IsValid() bool {
	switch e {
	case SideLeft, SideRight, SideBoth:
		return true
	}
	return false
}

func (e Side) String() string {
	return string(e)
}

func (e *Side) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Side(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Side", str)
	}
	return nil
}

func (e Side) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
