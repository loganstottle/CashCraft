package main

import (
	"context"
	"log"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func TestFinance(apiKey string) {

	c := polygon.New(apiKey)

	params := models.ListTickersParams{}.
		WithMarket(models.AssetClass("stocks")).
		WithActive(true).
		WithOrder(models.Order("asc")).
		WithLimit(100).
		WithSort(models.Sort("ticker"))

	iter := c.ListTickers(context.Background(), params)

	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
