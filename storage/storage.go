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
	ShowAllIngridients(ctx context.Context, ingridient string) ([]string, error)
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

func (d *Dishes) TakeDish() string {
	var result string
	for i := 0; i < len(d.Proteins); i++ {
		result += "🍲 " + d.Proteins[i] + " " + d.Carbs[i] + " " + d.Fats[i] + " " + d.Fibers[i] + "\n"
	}
	return result
}
