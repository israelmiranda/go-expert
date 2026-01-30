package usecase

import (
	"context"
	"time"

	"github.com/israelmiranda/go-expert/auction/internal/domain"
)

type BidOutput struct {
	ID        string
	UserID    string
	AuctionID string
	Amount    float64
	Timestamp time.Time
}

func fromBids(bids []domain.Bid) []BidOutput {
	var outputs []BidOutput
	for _, bid := range bids {
		outputs = append(outputs, fromBid(bid))
	}
	return outputs
}

func fromBid(bid domain.Bid) BidOutput {
	return BidOutput{
		ID:        bid.ID,
		UserID:    bid.UserID,
		AuctionID: bid.AuctionID,
		Amount:    bid.Amount,
		Timestamp: bid.Timestamp,
	}
}

type FindBid interface {
	FindByAuctionId(ctx context.Context, auctionID string) ([]domain.Bid, error)
	FindWinningByAuctionId(ctx context.Context, auctionID string) (domain.Bid, error)
}

type FindBidUseCase struct {
	repository FindBid
}

func NewFindBidUseCase(repository FindBid) FindBidUseCase {
	return FindBidUseCase{repository}
}

func (u FindBidUseCase) FindByAuctionId(ctx context.Context, auctionID string) ([]BidOutput, error) {
	bids, err := u.repository.FindByAuctionId(ctx, auctionID)
	if err != nil {
		return []BidOutput{}, err
	}

	return fromBids(bids), nil
}

func (u FindBidUseCase) FindWinningByAuctionId(ctx context.Context, auctionID string) (BidOutput, error) {
	bid, err := u.repository.FindWinningByAuctionId(ctx, auctionID)
	if err != nil {
		return BidOutput{}, err
	}

	return fromBid(bid), nil
}
