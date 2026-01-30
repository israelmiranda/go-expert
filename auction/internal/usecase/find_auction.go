package usecase

import (
	"context"
	"time"

	"github.com/israelmiranda/go-expert/auction/internal/domain"
	"github.com/israelmiranda/go-expert/auction/internal/repository"
)

type AuctionOutput struct {
	ID          string
	ProductName string
	Category    string
	Description string
	Condition   domain.ProductCondition
	Status      domain.AuctionStatus
	Timestamp   time.Time
}

func fromAuctions(auctions []domain.Auction) []AuctionOutput {
	var outputs []AuctionOutput
	for _, auction := range auctions {
		outputs = append(outputs, fromAuction(auction))
	}
	return outputs
}

func fromAuction(auction domain.Auction) AuctionOutput {
	return AuctionOutput{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp,
	}
}

type FindAuction interface {
	FindAllBy(ctx context.Context, params repository.AuctionParams) ([]domain.Auction, error)
}

type FindAuctionUseCase struct {
	repository FindAuction
}

func NewFindAuctionUseCase(repository FindAuction) FindAuctionUseCase {
	return FindAuctionUseCase{repository}
}

func (u FindAuctionUseCase) FindAllBy(ctx context.Context, params repository.AuctionParams) ([]AuctionOutput, error) {
	auctions, err := u.repository.FindAllBy(ctx, params)
	if err != nil {
		return []AuctionOutput{}, err
	}

	return fromAuctions(auctions), nil
}
