package main

import (
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

	err := PrettyPrint(project)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
