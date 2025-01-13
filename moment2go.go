// The package moment2go provides functionality to convert Moment.js date and time format to Go date and time format.
package moment2go

import (
	"fmt"
	"regexp"
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

// compileRegexp compiles a regular expression to match Moment.js format tokens.
// It shall not panic, because the pattern is hardcoded.
func compileRegexp() *regexp.Regexp {

	var tokenPatterns []string
	for _, token := range formatTokens {
		tokenPatterns = append(tokenPatterns, regexp.QuoteMeta(token.momentFormat))
	}
	pattern := strings.Join(tokenPatterns, "|")

	return regexp.MustCompile(pattern)
}

// ConvertMomentFormat converts a Moment.js date and time format to a Go date and time format.
func ConvertMomentFormat(momentFormat string) string {

	re := compileRegexp()

	// Replace tokens using regex with context-aware replacement.
	return re.ReplaceAllStringFunc(momentFormat, func(match string) string {
		for _, token := range formatTokens {
			if match == token.momentFormat {
				return token.goFormat
			}
		}
		return match // Fallback (should not occur).
	})
}

// ConvertMomentToGoLayoutWithLocation converts a Moment.js date and time layout to a Go date
// and time layout with a time zone offset.
func ConvertMomentToGoLayoutWithLocation(momentLayout string, location *time.Location) string {
	goLayout := ConvertMomentFormat(momentLayout)

	// Add the time zone offset to the layout.
	_, offset := time.Now().In(location).Zone()
	goLayout += fmt.Sprintf(" %02d:%02d", offset/3600, (offset%3600)/60)

	return goLayout
}

// Moment2Go is a thread-safe converter for Moment.js date and time formats to Go date and time formats.
type Moment2Go struct {
	re  *regexp.Regexp
	mux sync.RWMutex
	m   map[string]string
}

// New creates a new Moment2GoConverter.
func New() *Moment2Go {
	return &Moment2Go{
		re: compileRegexp(),
		m:  make(map[string]string),
	}
}

// Convert converts a Moment.js date and time format to a Go date and time format.
func (c *Moment2Go) Convert(momentLayout string) string {
	c.mux.RLock()
	goLayout, ok := c.m[momentLayout]
	c.mux.RUnlock()
	if ok {
		return goLayout
	}

	goLayout = ConvertMomentFormat(momentLayout)

	c.mux.Lock()
	c.m[momentLayout] = goLayout
	c.mux.Unlock()

	return goLayout
}

// Format formats a time.Time value using a Moment.js date and time format.
func (c *Moment2Go) Format(momentLayout string, t time.Time) string {

	goLayout := c.Convert(momentLayout)
	return t.Format(goLayout)
}
