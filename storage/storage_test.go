package storage

import (
	"testing"
)

func TestDishes_TakeDish(t *testing.T) {
	type fields struct {
		Proteins []string
		Carbs    []string
		Fats     []string
		Fibers   []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "base test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//proteins, err := mocks.Storage.PickRandomIngridient(context.Context(), "protein", 1)
			//carbs, err
			//fats, err
			//fibers
			d := &Dishes{
				Proteins: tt.fields.Proteins,
				Carbs:    tt.fields.Carbs,
				Fats:     tt.fields.Fats,
				Fibers:   tt.fields.Fibers,
			}
			if got := d.TakeDish(); got != tt.want {
				t.Errorf("TakeDish() = %v, want %v", got, tt.want)
			}
		})
	}
}
