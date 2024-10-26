package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

func calculatePoints(receipt Receipt) int {
	points := 0

	// Rule 1: Points for alphanumeric characters in retailer name
	points += countAlphanumeric(receipt.Retailer)

	// Rule 2: 50 points if the total is a round dollar amount
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == math.Floor(total) {
		points += 50
	}

	// Rule 3: 25 points if total is multiple of 0.25
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points for description length multiple of 3
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day is odd
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time is between 2:00 pm and 4:00 pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 {
		points += 10
	}

	return points
}

// Helper function to count alphanumeric characters
func countAlphanumeric(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			count++
		}
	}
	return count
}
