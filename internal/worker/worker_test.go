package worker

import (
	"testing"

	"dc-simulator/internal/model"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "Nombres positifs classiques",
			numbers:  []int{1, 2, 3},
			expected: 14, // 1² + 2² + 3² = 1 + 4 + 9 = 14
		},
		{
			name:     "Liste vide",
			numbers:  []int{},
			expected: 0,
		},
		{
			name:     "Nombres négatifs",
			numbers:  []int{-2, 4},
			expected: 20, // (-2)² + 4² = 4 + 16 = 20
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jobs := make(chan model.Job, 1)
			results := make(chan model.Result, 1)

			jobs <- model.Job{ID: 1, Numbers: tt.numbers}
			
			close(jobs)

			Process(99, jobs, results)
			close(results)

			res, ok := <-results
			if !ok {
				t.Fatalf("Aucun résultat retourné par le worker")
			}

			if res.Sum != tt.expected {
				t.Errorf("Résultat incorrect: attendu %d, obtenu %d", tt.expected, res.Sum)
			}
		})
	}
}