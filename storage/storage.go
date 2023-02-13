package storage

import (
	"context"
	// "crypto/sha1"
	"errors"
	// "fmt"
	// "io"
)

type Storage interface {
	PickRandomDish(ctx context.Context, quantity int) (*Dishes, error)
	PickRandomIngridient(ctx context.Context, ingridient string, quantity int) ([]string, error)
}

var ErrNoIngridients = errors.New("no ingredients available")

type Dishes struct {
	Proteins []string
	Carbs    []string
	Fats     []string
	Fibers   []string
}

type Plate struct {
	Plate []string
}

// func takeDish(d Dishes) {
// 	for i := 0; i < len(d.Proteins); i++ {
// 		fmt.Println(d.Proteins[i], d.Carbs[i], d.Fats[i], d.Fibers[i])
// 	}
// }