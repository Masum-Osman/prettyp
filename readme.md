# PrettyPrint JSON Utility

This project provides a utility for pretty-printing JSON with color-coded output. The utility is designed to print keys in blue and values in red, making it easier to read and debug JSON data.


## Usage

### Install 
```
go get github.com/Masum-Osman/prettyp
```

### Import the Package

To use the PrettyPrint function in your Go project, import the package:

```go
import "github.com/Masum-Osman/prettyp"
```

### Example

Here’s an example of how to use the PrettyPrint function:

```go
package main

import (
    p "github.com/Masum-Osman/prettyp"
    "fmt"
)

func main() {
    cat := struct {
		Id            string    `json:"id"`
		Name          string    `json:"name"`
		Breed         string    `json:"breed"`
		Age           int       `json:"age"`
		IsIndoor      bool      `json:"is_indoor"`
		WeightKg      float64   `json:"weight_kg"`
		FavoriteFoods []string  `json:"favorite_foods"`
		AdoptedAt     time.Time `json:"adopted_at"`
		Owner         Owner     `json:"owner"`
	}{
		Id:            "673dcdb47484940c3f9fd08c",
		Name:          "Whiskers",
		Breed:         "British Shorthair",
		Age:           3,
		IsIndoor:      true,
		WeightKg:      4.2,
		FavoriteFoods: []string{"Tuna", "Salmon", "Chicken"},
		AdoptedAt:     time.Now(),
		Owner: Owner{
			Name:  "Owner",
			Email: "owner@example.com",
		},
	}

    err := p.PrettyPrint(project)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

### Running the Example

To run the example, use the following command:

```sh
go run main.go
```

![screenshot](doc/screenshot.png)

## Testing

To run the tests for the PrettyPrint function, use the following command:

```sh
go test
```

<!-- ## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. -->

## Contributing

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/fooBar`).
3. Commit your changes (`git commit -am 'Add some fooBar'`).
4. Push to the branch (`git push origin feature/fooBar`).
5. Create a new Pull Request.

## Acknowledgements

- Inspired by the need for better JSON readability in Go projects.

---
