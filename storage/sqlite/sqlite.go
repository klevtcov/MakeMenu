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

// Остановка базы, закрытие соединения
func Shutdown(s *Storage) error {
	if err := s.db.Close(); err != nil {
		return fmt.Errorf("can't close connection to database: %w", err)
	}
	return nil
}

// Выбор из базы заданного количества случайных ингридиентов
func (s *Storage) PickRandomIngridient(ctx context.Context, ingridient string, quantity int) ([]string, error) {
	q := fmt.Sprintf("SELECT name FROM %s ORDER BY RANDOM() LIMIT ?", ingridient)
	var result []string

	rows, err := s.db.Query(q, quantity)
	if err != nil {
		fmt.Println("Не могу извлечь строки из базы", err)
		return nil, err
	}

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var selectedIngridient string
		if err := rows.Scan(&selectedIngridient); err != nil {
			fmt.Println("Не могу считать следующую строку", err)
			return nil, err
		}
		result = append(result, selectedIngridient)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Ошибка строк", err)
		return nil, err
	}
	return result, nil
}

// Сборка блюд из случайно выбранных ингридиентов в заданном количестве
func (s *Storage) PickRandomDish(ctx context.Context, quantity int) (*storage.Dishes, error) {

	protein, err := s.PickRandomIngridient(ctx, "proteins", quantity)
	if err != nil {
		fmt.Println("Не могу выбрать протеины из базы", err)
		return nil, err
	}

	carb, err := s.PickRandomIngridient(ctx, "carbs", quantity)
	if err != nil {
		fmt.Println("Не могу выбрать углеводы из базы", err)
		return nil, err
	}

	fat, err := s.PickRandomIngridient(ctx, "fats", quantity)
	if err != nil {
		fmt.Println("Не могу выбрать жиры из базы", err)
		return nil, err
	}

	fiber, err := s.PickRandomIngridient(ctx, "fibers", quantity)
	if err != nil {
		fmt.Println("Не могу выбрать клетчатку из базы", err)
		return nil, err
	}

	return &storage.Dishes{
		Proteins: protein,
		Carbs:    carb,
		Fats:     fat,
		Fibers:   fiber,
	}, nil
}
