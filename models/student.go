package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"nome" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]*$"`
}
