package dto

type VisitDTO struct { // inside part to your response
	StoreID   string   `json:"store_id" validate:"required"`
	ImageURLs []string `json:"image_url" validate:"required,dive,url"`
	VisitTime string   `json:"visit_time" validate:"required"`
}

type JobRequestDTO struct { // response 
	Count  int        `json:"count" validate:"gte=1"`
	Visits []VisitDTO `json:"visits" validate:"required,dive"`
}


