package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/klevtcov/makemenu_go/storage"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

// New creates new SQLite storage.
func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	return &Storage{db: db}, nil
}

// PickRandom picks random page from storage.
func (s *Storage) PickRandom(ctx context.Context) (*storage.Dish, error) {
	proteinSelect := `SELECT name FROM proteins ORDER BY RANDOM() LIMIT 1`
	carbSelect := `SELECT name FROM carbs ORDER BY RANDOM() LIMIT 1`
	fatSelect := `SELECT name FROM fats ORDER BY RANDOM() LIMIT 1`
	fiberSelect := `SELECT name FROM fibers ORDER BY RANDOM() LIMIT 1`

	var protein, carb, fat, fiber string

	func pickIngridient (ing string) 
	err := s.db.QueryRowContext(ctx, proteinSelect).Scan(&protein)
	if err == sql.ErrNoRows {
		return nil, storage.ErrNoIngridients
	}
	if err != nil {
		return nil, fmt.Errorf("can't pick random ingridient: %w", err)
	}

	return &storage.Dish{
		Proteins: protein,
		Carbs:    carb,
		Fats:     fat,
		Fibers:   fiber,
	}, nil
}

// func (s *Storage) Init(ctx context.Context) error {
// 	q := `CREATE TABLE IF NOT EXISTS pages (url TEXT, user_name TEXT)`
// 	_, err := s.db.ExecContext(ctx, q)
// 	if err != nil {
// 		return fmt.Errorf("can't create table: %w", err)
// 	}
// 	return nil
// }
