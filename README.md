# moment2go
[![Build Status](https://github.com/axkit/bitset/actions/workflows/go.yml/badge.svg)](https://github.com/axkit/moment2go/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/axkit/moment2go)](https://goreportcard.com/report/github.com/axkit/moment2go)
[![GoDoc](https://pkg.go.dev/badge/github.com/axkit/moment2go)](https://pkg.go.dev/github.com/axkit/moment2go)
[![Coverage Status](https://coveralls.io/repos/github/axkit/moment2go/badge.svg?branch=main)](https://coveralls.io/github/axkit/moment2go?branch=main)


The `moment2go` package provides functionality to convert Moment.js date and time formats into Go date and time layouts. It offers a seamless way to integrate familiar Moment.js-style date formatting into Go applications, while maintaining thread-safety and extensibility.

## Features

- Converts Moment.js format tokens to Go layout tokens.
- Thread-safe and efficient conversion for concurrent applications.
- Supports location-based date and time formatting.

## Installation

To use this package, install it with:

```bash
go get github.com/axkit/moment2go
```

## Usage

### Recommended Conversion Approach

Format a `time.Time` value using a Moment.js format:

```go
package main

import (
	"fmt"
	"time"
	"github.com/axkit/moment2go"
)

func main() {
	converter := moment2go.New()
	t := time.Date(2023, 1, 1, 15, 30, 0, 0, time.UTC)
	formattedTime := converter.Format("YYYY-MM-DD HH:mm", t)
	fmt.Println("Formatted Time:", formattedTime)
}
```

### Basic Conversion

Convert a Moment.js format string into a Go layout string:

```go
package main

import (
	"fmt"
	"github.com/axkit/moment2go"
)

func main() {
	momentFormat := "YYYY-MM-DD"
	goFormat := moment2go.ConvertMomentFormat(momentFormat)
	fmt.Println("Go Format:", goFormat)
}
```


### Location-Based Conversion

Convert a Moment.js format string and apply a location-specific time zone offset:

```go
package main

import (
	"fmt"
	"time"
	"github.com/axkit/moment2go"
)

func main() {
	location, _ := time.LoadLocation("America/New_York")
	momentLayout := "YYYY-MM-DD"
	goLayout := moment2go.ConvertMomentToGoLayoutWithLocation(momentLayout, location)
	fmt.Println("Go Layout with Location:", goLayout)
}
```

## API

### Functions

#### `ConvertMomentFormat(momentFormat string) string`
Converts a Moment.js date and time format to a Go date and time format.

#### `ConvertMomentToGoLayoutWithLocation(momentLayout string, location *time.Location) string`
Converts a Moment.js date and time layout to a Go layout with a time zone offset.

### Types

#### `Moment2Go`
Thread-safe converter for Moment.js formats.

- `New() *Moment2Go`: Creates a new `Moment2Go` instance.
- `Convert(momentLayout string) string`: Converts a Moment.js format string to a Go layout.
- `Format(momentLayout string, t time.Time) string`: Formats a `time.Time` value using a Moment.js layout.

## Testing

Run the tests with:

```bash
go test ./...
```

## Examples

Example test cases are included in the `moment2go_test.go` file, covering:

- Basic format conversions.
- Complex format handling.
- Thread-safety validations.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for bug fixes, feature requests, or improvements.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

