package domain

import (
	"time"

	"github.com/google/uuid"
)

type Auction struct {
	ID          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

type ProductCondition int
type AuctionStatus int

const (
	Active AuctionStatus = iota
	Completed
)

const (
	New ProductCondition = iota + 1
	Used
	Refurbished
)

func CreateAuction(
	productName,
	category,
	description string,
	condition ProductCondition,
) Auction {
	return Auction{
		ID:          uuid.NewString(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}
}
