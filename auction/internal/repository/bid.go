package repository

import (
	"context"
	"time"

	"github.com/israelmiranda/go-expert/auction/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidMongo struct {
	ID        string  `bson:"_id"`
	UserID    string  `bson:"user_id"`
	AuctionID string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

func fromBid(bid domain.Bid) BidMongo {
	return BidMongo{
		ID:        bid.ID,
		UserID:    bid.UserID,
		AuctionID: bid.AuctionID,
		Amount:    bid.Amount,
		Timestamp: bid.Timestamp.Unix(),
	}
}

func (b BidMongo) toBid() domain.Bid {
	return domain.Bid{
		ID:        b.ID,
		UserID:    b.UserID,
		AuctionID: b.AuctionID,
		Amount:    b.Amount,
		Timestamp: time.Unix(b.Timestamp, 0),
	}
}

type BidRepository struct {
	collection *mongo.Collection
}

func NewBidRepository(database *mongo.Database) BidRepository {
	return BidRepository{
		collection: database.Collection("bids"),
	}
}

func (r BidRepository) FindByAuctionId(ctx context.Context, auctionID string) ([]domain.Bid, error) {
	filter := bson.M{"auction_id": auctionID}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return []domain.Bid{}, err
	}
	defer cursor.Close(ctx)

	var bidsMongo []BidMongo
	if err := cursor.All(ctx, &bidsMongo); err != nil {
		return []domain.Bid{}, err
	}

	var bids []domain.Bid
	for _, bidMongo := range bidsMongo {
		bids = append(bids, bidMongo.toBid())
	}

	return bids, nil
}

func (r BidRepository) FindWinningByAuctionId(ctx context.Context, auctionID string) (domain.Bid, error) {
	return domain.Bid{}, nil
}

func (r BidRepository) Create(ctx context.Context, bid domain.Bid) error {
	_, err := r.collection.InsertOne(ctx, fromBid(bid))
	if err != nil {
		return err
	}

	return nil
}
