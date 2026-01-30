package main

import (
	"context"
	"fmt"
	"log"

	"github.com/israelmiranda/go-expert/auction/internal/configuration/database/mongodb"
	"github.com/israelmiranda/go-expert/auction/internal/domain"
	"github.com/israelmiranda/go-expert/auction/internal/repository"
	"github.com/israelmiranda/go-expert/auction/internal/usecase"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	databaseConnection, err := mongodb.NewMongoConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("-----Users-----")

	userRepository := repository.NewUserRepository(databaseConnection)
	// createUserUseCase := usecase.NewCreateUserUseCase(userRepository)
	// createUserUseCase.Create(ctx, usecase.UserInput{Name: "John Doe"})
	// createUserUseCase.Create(ctx, usecase.UserInput{Name: "Jane Doe"})

	findUserUseCase := usecase.NewFindUserUseCase(userRepository)

	users, _ := findUserUseCase.FindAll(ctx)
	for _, user := range users {
		fmt.Println(user)
	}

	fmt.Println("-----user by id-----")
	user, _ := findUserUseCase.FindById(ctx, users[0].ID)
	fmt.Println(user)

	fmt.Println("-----Auctions-----")

	auctionRepository := repository.NewAuctionRepository(databaseConnection)
	// createAuctionUseCase := usecase.NewCreateAuctionUseCase(auctionRepository)
	// createAuctionUseCase.Create(ctx, usecase.AuctionInput{
	// 	ProductName: "Product 1",
	// 	Category:    "Category 1",
	// 	Description: "Description 1",
	// 	Condition:   domain.New,
	// })
	// createAuctionUseCase.Create(ctx, usecase.AuctionInput{
	// 	ProductName: "Product 2",
	// 	Category:    "Category 1",
	// 	Description: "Description 2",
	// 	Condition:   domain.Used,
	// })
	// createAuctionUseCase.Create(ctx, usecase.AuctionInput{
	// 	ProductName: "Product 3",
	// 	Category:    "Category 1",
	// 	Description: "Description 3",
	// 	Condition:   domain.Refurbished,
	// })

	findAuctionUseCase := usecase.NewFindAuctionUseCase(auctionRepository)

	auctions, _ := findAuctionUseCase.FindAllBy(ctx, repository.AuctionParams{
		Category: "Category 1",
		Status:   domain.Active,
	})
	for _, auction := range auctions {
		fmt.Println(auction)
	}

	fmt.Println("-----Bids-----")

	bidRepository := repository.NewBidRepository(databaseConnection)
	// createBidUseCase := usecase.NewCreateBidUseCase(bidRepository)
	// createBidUseCase.Create(ctx, usecase.BidInput{
	// 	UserID:    users[0].ID,
	// 	AuctionID: auctions[0].ID,
	// 	Amount:    10.5,
	// })
	// createBidUseCase.Create(ctx, usecase.BidInput{
	// 	UserID:    users[1].ID,
	// 	AuctionID: auctions[0].ID,
	// 	Amount:    30.5,
	// })

	findBidUseCase := usecase.NewFindBidUseCase(bidRepository)

	bids, _ := findBidUseCase.FindByAuctionId(ctx, auctions[0].ID)
	for _, bid := range bids {
		fmt.Println(bid)
	}
}
