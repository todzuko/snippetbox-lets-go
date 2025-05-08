package main

import (
	"github.com/todzuko/snippetbox-lets-go/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2025, 5, 7, 10, 15, 0, 0, time.UTC),
			want: "7 May 2025, 10:15",
		},
		{
			name: "empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2025, 5, 7, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "7 May 2025, 09:15",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}
