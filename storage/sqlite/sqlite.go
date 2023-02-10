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
func (s *Storage) PickRandom(ctx context.Context, q int) (*storage.Dish, error) {
	proteinSelect := `SELECT name FROM proteins ORDER BY RANDOM() LIMIT ?`
	carbSelect := `SELECT name FROM carbs ORDER BY RANDOM() LIMIT ?`
	fatSelect := `SELECT name FROM fats ORDER BY RANDOM() LIMIT ?`
	fiberSelect := `SELECT name FROM fibers ORDER BY RANDOM() LIMIT ?`

	var protein, carb, fat, fiber []string

	// func pickIngridient (ing string)
	err := s.db.QueryRowContext(ctx, proteinSelect, q).Scan(&protein)
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

//
// Работающий код, переписать с него запросы
//
// proteinSelect := `SELECT name FROM proteins ORDER BY RANDOM() LIMIT ?`

// var protein []string
// rows, err := db.Query(proteinSelect, 2)
// if err != nil {
// 	fmt.Println("Не могу извлечь строки из базы")
// }

// // Loop through rows, using Scan to assign column data to struct fields.
// for rows.Next() {
// 	var ingridient string
// 	if err := rows.Scan(&ingridient); err != nil {
// 		fmt.Println("Не могу считать следующую строку")
// 	}
// 	protein = append(protein, ingridient)
// }
// if err = rows.Err(); err != nil {
// 	fmt.Println("Ошибка строк")
// }

// fmt.Println(protein)
