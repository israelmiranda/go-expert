package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	fmt.Println("booking....")
	select {
	case <-ctx.Done():
		fmt.Println("Booking cancelled. Timeout reached.")
	case <-time.After(4 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
