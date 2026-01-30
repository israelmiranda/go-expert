package usecase

import (
	"context"

	"github.com/israelmiranda/go-expert/auction/internal/domain"
)

type BidInput struct {
	UserID    string
	AuctionID string
	Amount    float64
}

func (b BidInput) toBid() domain.Bid {
	return domain.CreateBid(
		b.UserID,
		b.AuctionID,
		b.Amount,
	)
}

type CreateBid interface {
	Create(ctx context.Context, bid domain.Bid) error
}

type CreateBidUseCase struct {
	repository CreateBid
}

func NewCreateBidUseCase(repository CreateBid) CreateBidUseCase {
	return CreateBidUseCase{repository}
}

func (u CreateBidUseCase) Create(ctx context.Context, bidInput BidInput) error {
	return u.repository.Create(ctx, bidInput.toBid())
}
