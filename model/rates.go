package model

import (
	"time"

	"github.com/google/uuid"
)

// User is a JSON user
type Rates struct {
	ID        uuid.UUID `json:"id"`
	AskPrice  float64   `json:"ask_price" validate:"required"`
	BidPrice  float64   `json:"bid_price" validate:"required"`
	TimeStamp time.Time `json:"time_stamp"`
}

// ToDB converts User to DBUser
func (rates *Rates) ToDB() *DBRates {
	return &DBRates{
		ID:        rates.ID,
		AskPrice:  rates.AskPrice,
		BidPrice:  rates.BidPrice,
		TimeStamp: rates.TimeStamp,
	}
}

// DBUser is a Postgres user
type DBRates struct {
	tableName struct{}  `pg:"rates" gorm:"primaryKey"`
	ID        uuid.UUID `pg:"id,notnull,pk"`
	AskPrice  float64   `pg:"ask_price,notnull"`
	BidPrice  float64   `pg:"bid_price,notnull"`
	TimeStamp time.Time `pg:"time_stamp,notnull"`
}

// TableName overrides default table name for gorm
func (DBRates) TableName() string {
	return "users"
}

// ToWeb converts DBUser to User
func (dbUser *DBRates) ToWeb() *Rates {
	return &Rates{
		ID:        dbUser.ID,
		AskPrice:  dbUser.AskPrice,
		BidPrice:  dbUser.BidPrice,
		TimeStamp: dbUser.TimeStamp,
	}
}
