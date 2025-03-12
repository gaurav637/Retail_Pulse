package models
import "time"

type Image struct {
	ID          int64  `db:"id" json:"id"`
	Description string `db:"description" json:"description"`
	Width       int    `db:"width" json:"width"`
	Height      int    `db:"height" json:"height"`
	Parimeter   int    `db:"parameter" json:"parameter"`
	JobID       int64  `db:"job_id" json:"job_id"`
	StoreID     int64  `db:"store_id" json:"store_id"`
	URL         string `db:"url" json:"url"`
	VisitTime  time.Time `db: "visit_time" json:"visit_time"`
}
