package models

import "time"

type Job struct {
	ID               string     `db:"id" json:"id"`
	Created_at        time.Time `db:"created_at" json:"created_at"`
	Updated_at        time.Time `db:"updated_at" json:"updated_at"`
	Processing_status  string   `db:"processing_status" json:"processing_status"`
}
