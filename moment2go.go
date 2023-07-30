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
	{"DDD", "_2"},
	{"DD", "02"},
	{"D", "2"},
	{"MMMM", "January"},
	{"MMM", "Jan"},
	{"MM", "01"},
	{"M", "1"},
	{"YYYY", "2006"},
	{"YY", "06"},
	{"hh", "03"},
	{"H", "15"},
	{"mm", "04"},
	{"ss", "05"},
	{"A", "PM"},
	{"a", "pm"},
	{"ZZ", "-0700"},
	{"Z", "-07:00"},
}

func ConvertMomentFormat(momentFormat string) string {

	// Replace Moment.js format tokens with Go format tokens
	goFormat := momentFormat
	for i := range formatTokens {
		goFormat = strings.ReplaceAll(goFormat, formatTokens[i].momentFormat, formatTokens[i].goFormat)
	}

	return goFormat
}

func ConvertMomentToGoLayout(momentLayout string) string {
	// Split the Moment.js layout into date and time parts
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

func ConvertMomentToGoLayoutWithLocation(momentLayout string, location *time.Location) string {
	goLayout := ConvertMomentToGoLayout(momentLayout)

	// Add the time zone offset to the layout
	_, offset := time.Now().In(location).Zone()
	goLayout += fmt.Sprintf(" %02d:%02d", offset/3600, (offset%3600)/60)

	return goLayout
}

type Moment2GoConverter struct {
	mux sync.RWMutex
	m   map[string]string
}

func NewMoment2GoConverter() *Moment2GoConverter {
	return &Moment2GoConverter{
		m: make(map[string]string),
	}
}

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

func (c *Moment2GoConverter) Format(momentLayout string, t time.Time) string {
	goLayout := c.Convert(momentLayout)
	return t.Format(goLayout)
}

func (c *Moment2GoConverter) Parse(momentLayout string) {
	goLayout := c.Convert(momentLayout)
	c.m[momentLayout] = goLayout
}
