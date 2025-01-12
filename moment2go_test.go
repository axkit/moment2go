package moment2go_test

import (
	"sync"
	"testing"
	"time"

	"github.com/axkit/moment2go"
	"github.com/stretchr/testify/assert"
)

func TestConvertMomentFormat(t *testing.T) {
	tests := []struct {
		name         string
		momentFormat string
		expected     string
	}{
		{"Full date format", "YYYY-MM-DD", "2006-01-02"},
		{"Time with AM/PM", "hh:mm A", "15:04 PM"},
		{"Day and Month", "dddd, MMMM", "Monday, January"},
		{"Complex format", "YYYY-MM-DDTHH:mm:ssZZ", "2006-01-02T15:04:05-0700"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := moment2go.ConvertMomentFormat(test.momentFormat)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestConvertMomentToGoLayout(t *testing.T) {
	tests := []struct {
		name         string
		momentLayout string
		expected     string
	}{
		{"Date only", "YYYY-MM-DD", "2006-01-02"},
		{"Date and time", "YYYY-MM-DD HH:mm:ss", "2006-01-02 15:04:05"},
		{"Time only", "HH:mm:ss", "15:04:05"},
		{"Custom format", "ddd, MMM DD, YYYY", "Mon, Jan 02, 2006"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := moment2go.ConvertMomentToGoLayout(test.momentLayout)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestConvertMomentToGoLayoutWithLocation(t *testing.T) {
	tests := []struct {
		name         string
		momentLayout string
		location     *time.Location
	}{
		{"UTC location", "YYYY-MM-DD", time.UTC},
		{"Local location", "YYYY-MM-DD", time.Now().Location()},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := moment2go.ConvertMomentToGoLayoutWithLocation(test.momentLayout, test.location)
			assert.Contains(t, result, "2006-01-02")
		})
	}
}

func TestMoment2GoConverter(t *testing.T) {
	converter := moment2go.NewConverter()
	tests := []struct {
		name         string
		momentLayout string
		timeInput    time.Time
		expected     string
	}{
		{"Simple format", "YYYY-MM-DD", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), "2023-01-01"},
		{"With time", "YYYY-MM-DD HH:mm", time.Date(2023, 1, 1, 15, 30, 0, 0, time.UTC), "2023-01-01 15:30"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := converter.Format(test.momentLayout, test.timeInput)
			assert.Equal(t, test.expected, result)
		})
	}

	t.Run("Thread-safe conversion", func(t *testing.T) {
		var wg sync.WaitGroup
		momentLayout := "YYYY-MM-DD"
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				converter.Convert(momentLayout)
			}()
		}
		wg.Wait()
	})
}
