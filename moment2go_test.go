package moment2go

import (
	"testing"
	"time"
)

func TestConvertMomentFormat(t *testing.T) {
	momentFormat := "YYYY-MM-DD hh:mm:ss"
	expectedGoFormat := "2006-01-02 03:04:05"

	goFormat := ConvertMomentFormat(momentFormat)

	if goFormat != expectedGoFormat {
		t.Errorf("Unexpected Go format. Expected: %s, Got: %s", expectedGoFormat, goFormat)
	}
}

func TestConvertMomentToGoLayout(t *testing.T) {
	momentLayout := "YYYY-MM-DD hh:mm:ss"
	expectedGoLayout := "2006-01-02 03:04:05"

	goLayout := ConvertMomentToGoLayout(momentLayout)

	if goLayout != expectedGoLayout {
		t.Errorf("Unexpected Go layout. Expected: %s, Got: %s", expectedGoLayout, goLayout)
	}
}

func TestConvertMomentToGoLayoutWithLocation(t *testing.T) {
	momentLayout := "YYYY-MM-DD hh:mm:ss"
	location, err := time.LoadLocation("UTC")
	if err != nil {
		t.Fatalf("Failed to load time zone location: %v", err)
	}
	expectedGoLayout := "2006-01-02 03:04:05 00:00"

	goLayout := ConvertMomentToGoLayoutWithLocation(momentLayout, location)

	if goLayout != expectedGoLayout {
		t.Errorf("Unexpected Go layout with location. Expected: %s, Got: %s", expectedGoLayout, goLayout)
	}
}
