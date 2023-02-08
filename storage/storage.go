package storage

import (
	"context"
	// "crypto/sha1"
	"errors"
	// "fmt"
	// "io"
)

type Storage interface {
	PickRandom(ctx context.Context, Quantity int) (*Dish, error)
}

var ErrNoIngridients = errors.New("no ingredients available")

type Dish struct {
	Proteins []string
	Carbs    []string
	Fats     []string
	Fibers   []string
}
