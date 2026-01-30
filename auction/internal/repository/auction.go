package repository

import (
	"context"
	"time"

	"github.com/israelmiranda/go-expert/auction/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionParams struct {
	ProductName string
	Category    string
	Status      domain.AuctionStatus
}

type AuctionMongo struct {
	ID          string                  `bson:"_id"`
	ProductName string                  `bson:"product_name"`
	Category    string                  `bson:"category"`
	Description string                  `bson:"description"`
	Condition   domain.ProductCondition `bson:"condition"`
	Status      domain.AuctionStatus    `bson:"status"`
	Timestamp   int64                   `bson:"timestamp"`
}

func fromAuction(auction domain.Auction) AuctionMongo {
	return AuctionMongo{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp.Unix(),
	}
}

func (a AuctionMongo) toAuction() domain.Auction {
	return domain.Auction{
		ID:          a.ID,
		ProductName: a.ProductName,
		Category:    a.Category,
		Description: a.Description,
		Condition:   a.Condition,
		Status:      a.Status,
		Timestamp:   time.Unix(a.Timestamp, 0),
	}
}

type AuctionRepository struct {
	collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) AuctionRepository {
	return AuctionRepository{
		collection: database.Collection("auctions"),
	}
}

func (r AuctionRepository) FindAllBy(ctx context.Context, params AuctionParams) ([]domain.Auction, error) {
	filter := bson.M{}

	if params.Category != "" {
		filter["category"] = params.Category
	}

	if params.ProductName != "" {
		filter["product_name"] = params.ProductName
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return []domain.Auction{}, err
	}
	defer cursor.Close(ctx)

	var auctionsMongo []AuctionMongo
	if err := cursor.All(ctx, &auctionsMongo); err != nil {
		return []domain.Auction{}, err
	}

	var auctions []domain.Auction
	for _, auctionMongo := range auctionsMongo {
		auctions = append(auctions, auctionMongo.toAuction())
	}

	return auctions, nil
}

func (r AuctionRepository) Create(ctx context.Context, auction domain.Auction) error {
	_, err := r.collection.InsertOne(ctx, fromAuction(auction))
	if err != nil {
		return err
	}

	return nil
}
