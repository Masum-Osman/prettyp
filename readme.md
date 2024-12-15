# PrettyPrint JSON Utility

This project provides a utility for pretty-printing JSON with color-coded output. The utility is designed to print keys in blue and values in red, making it easier to read and debug JSON data.

## Project Structure

```
prettyp/
├── go.mod
├── go.sum
├── prettyp.go
├── prettyp_test.go
├── main.go
└── readme.md
```

## Installation

1. Clone the repository:

```sh
git clone https://github.com/Masum-Osman/prettyp.git
cd prettyp
```

2. Initialize the module (if not already done):

```sh
go mod tidy
```

## Usage

### Install 
```
go get github.com/Masum-Osman/prettyp/prettyp
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
    "github.com/Masum-Osman/prettyp"
    "fmt"
)

func main() {
    project := struct {
        Id           int64                  `json:"project_id"`
        Title        string                 `json:"title"`
        Name         string                 `json:"name"`
        Budget       float64                `json:"budget"`
        Active       bool                   `json:"active"`
        Tags         []string               `json:"tags"`
        Milestones   []int                  `json:"milestones"`
        Metadata     map[string]interface{} `json:"metadata"`
        NestedStruct struct {
            Description string `json:"description"`
            Completed   bool   `json:"completed"`
        } `json:"nested_struct"`
    }{
        Id:         123,
        Title:      "Test Project",
        Name:       "Test Name",
        Budget:     1000000.50,
        Active:     true,
        Tags:       []string{"Go", "JSON", "PrettyPrint"},
        Milestones: []int{1, 2, 3, 4},
        Metadata: map[string]interface{}{
            "created_by": "User123",
            "priority":   "High",
        },
        NestedStruct: struct {
            Description string `json:"description"`
            Completed   bool   `json:"completed"`
        }{
            Description: "Nested Struct Description",
            Completed:   false,
        },
    }

    err := prettyp.PrettyPrint(project)
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

## Testing

To run the tests for the PrettyPrint function, use the following command:

```sh
go test ./prettyp
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
