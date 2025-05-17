package models

import (
	"github.com/todzuko/snippetbox-lets-go/internal/assert"
	"testing"
)

func TestUserModelExists(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}
	tests := []struct {
		name     string
		userID   int
		expected bool
	}{
		{
			name:     "Valid ID",
			userID:   1,
			expected: true,
		},
		{
			name:     "Zero ID",
			userID:   0,
			expected: false,
		},
		{
			name:     "Non-existent ID",
			userID:   2,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{DB: db}

			exists, err := m.Exists(tt.userID)

			assert.NilError(t, err)
			assert.Equal(t, exists, tt.expected)
		})
	}
}
