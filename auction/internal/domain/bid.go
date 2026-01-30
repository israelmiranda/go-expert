package domain

import (
	"time"

	"github.com/google/uuid"
)

type Bid struct {
	ID        string
	UserID    string
	AuctionID string
	Amount    float64
	Timestamp time.Time
}

func CreateBid(
	userID,
	auctionID string,
	amount float64,
) Bid {
	return Bid{
		ID:        uuid.NewString(),
		UserID:    userID,
		AuctionID: auctionID,
		Amount:    amount,
		Timestamp: time.Now(),
	}
}
