package domain

import "time"

type PVZ struct {
	ID               string    `db:"id" json:"id"`
	RegistrationDate time.Time `db:"registration_date" json:"registrationDate"`
	City             string    `db:"city" json:"city"`
}

type Reception struct {
	ID       string    `db:"id" json:"id"`
	PVZID    string    `db:"pvz_id" json:"pvzId"`
	DateTime time.Time `db:"date_time" json:"dateTime"`
	Status   string    `db:"status" json:"status"`
}

type Product struct {
	ID          string    `db:"id" json:"id"`
	ReceptionID string    `db:"reception_id" json:"receptionId"`
	DateTime    time.Time `db:"date_time" json:"dateTime"`
	Type        string    `db:"type" json:"type"`
}
