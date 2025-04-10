package dto

import (
	"time"
)

type ProductDTO struct {
	ID       string    `json:"id"`
	Type     string    `json:"type"`
	DateTime time.Time `json:"dateTime"`
}

type ReceptionDTO struct {
	ID       string       `json:"id"`
	DateTime time.Time    `json:"dateTime"`
	Status   string       `json:"status"`
	Products []ProductDTO `json:"products"`
}

type PVZDTO struct {
	ID               string         `json:"id"`
	City             string         `json:"city"`
	RegistrationDate time.Time      `json:"registrationDate"`
	Receptions       []ReceptionDTO `json:"receptions"`
}
