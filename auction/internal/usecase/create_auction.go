package usecase

import (
	"context"

	"github.com/israelmiranda/go-expert/auction/internal/domain"
)

type AuctionInput struct {
	ProductName string
	Category    string
	Description string
	Condition   domain.ProductCondition
}

func (a AuctionInput) toAuction() domain.Auction {
	return domain.CreateAuction(
		a.ProductName,
		a.Category,
		a.Description,
		a.Condition,
	)
}

type CreateAuction interface {
	Create(ctx context.Context, auction domain.Auction) error
}

type CreateAuctionUseCase struct {
	repository CreateAuction
}

func NewCreateAuctionUseCase(repository CreateAuction) CreateAuctionUseCase {
	return CreateAuctionUseCase{repository}
}

func (u CreateAuctionUseCase) Create(ctx context.Context, auctionInput AuctionInput) error {
	return u.repository.Create(ctx, auctionInput.toAuction())
}
