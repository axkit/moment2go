// The package moment2go provides functionality to convert Moment.js date and time format to Go date and time format.
package moment2go

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Map Moment.js format tokens to Go format tokens
// Important: the order of tokens is important, because some tokens are substrings of other tokens.
var formatTokens = []struct {
	momentFormat string
	goFormat     string
}{
	{"DDDD", "Monday"},
	{"DD", "02"},
	{"D", "2"},
	{"dddd", "Monday"},
	{"ddd", "Mon"},
	{"MMMM", "January"},
	{"MMM", "Jan"},
	{"MM", "01"},
	{"M", "1"},
	{"YYYY", "2006"},
	{"YY", "06"},
	{"HH", "15"},
	{"H", "3"},
	{"hh", "15"},
	{"h", "3"},
	{"mm", "04"},
	{"m", "4"},
	{"ss", "05"},
	{"s", "5"},
	{"A", "PM"},
	{"a", "pm"},
	{"ZZ", "-0700"},
	{"Z", "-07:00"},
}

// ConvertMomentFormat converts a Moment.js date and time format to a Go date and time format.
func ConvertMomentFormat(momentFormat string) string {

	goFormat := momentFormat
	for i := range formatTokens {
		goFormat = strings.ReplaceAll(goFormat, formatTokens[i].momentFormat, formatTokens[i].goFormat)
	}

	return goFormat
}

// ConvertMomentToGoLayout converts a Moment.js date and time layout to a Go date and time layout.
func ConvertMomentToGoLayout(momentLayout string) string {

	layoutParts := strings.Split(momentLayout, " ")

	var goLayout string
	if len(layoutParts) > 1 {
		// If there are separate date and time parts, convert them separately
		datePart := layoutParts[0]
		timePart := layoutParts[1]

		goDateFormat := ConvertMomentFormat(datePart)
		goTimeFormat := ConvertMomentFormat(timePart)

		goLayout = goDateFormat + " " + goTimeFormat
	} else {
		// If there is only one part, convert it as a whole
		goLayout = ConvertMomentFormat(momentLayout)
	}

	return goLayout
}

// ConvertMomentToGoLayoutWithLocation converts a Moment.js date and time layout to a Go date
// and time layout with a time zone offset.
func ConvertMomentToGoLayoutWithLocation(momentLayout string, location *time.Location) string {
	goLayout := ConvertMomentToGoLayout(momentLayout)

	// Add the time zone offset to the layout
	_, offset := time.Now().In(location).Zone()
	goLayout += fmt.Sprintf(" %02d:%02d", offset/3600, (offset%3600)/60)

	return goLayout
}

// Moment2GoConverter is a thread-safe converter for Moment.js date and time formats to Go date and time formats.
type Moment2GoConverter struct {
	mux sync.RWMutex
	m   map[string]string
}

// NewConverter creates a new Moment2GoConverter.
func NewConverter() *Moment2GoConverter {
	return &Moment2GoConverter{
		m: make(map[string]string),
	}
}

// Convert converts a Moment.js date and time format to a Go date and time format.
func (c *Moment2GoConverter) Convert(momentLayout string) string {
	c.mux.RLock()
	goLayout, ok := c.m[momentLayout]
	c.mux.RUnlock()
	if ok {
		return goLayout
	}

	goLayout = ConvertMomentToGoLayout(momentLayout)

	c.mux.Lock()
	c.m[momentLayout] = goLayout
	c.mux.Unlock()

	return goLayout
}

// Format formats a time.Time value using a Moment.js date and time format.
func (c *Moment2GoConverter) Format(momentLayout string, t time.Time) string {

	goLayout := c.Convert(momentLayout)
	return t.Format(goLayout)
}

// Parse parses a Moment.js date and time format and caches the corresponding Go date and time format.
func (c *Moment2GoConverter) Parse(momentLayout string) {
	goLayout := c.Convert(momentLayout)
	c.m[momentLayout] = goLayout
}
