package models

type Store struct {
	ID       int64  `db:"id" json:"id"`
	storeName     string `db:"storeName" json:"storeName"`
	AreaCode int64  `db:"AreaCode" json:"AreaCode"`
}
