package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SourceLang struct {
	En string `json:"en" bson:"en"`
	Bn string `json:"bn" bson:"bn"`
}

type UserInfo struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}

func TestPretty(t *testing.T) {
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
	assert.NoError(t, err, "PrettyPrint should not return an error")
}

func TestPrettyPrintWithLocation(t *testing.T) {
	location := struct {
		Id          string     `bson:"_id,omitempty" json:"id,omitempty"`
		Title       SourceLang `bson:"title" json:"title"`
		Address     SourceLang `bson:"address" json:"address"`
		Description SourceLang `bson:"description" json:"description"`
		MapLink     string     `bson:"map_link" json:"map_link"`
		CreatedAt   time.Time  `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at,omitempty" bson:"updated_at"`
		CreatedBy   UserInfo   `bson:"created_by" json:"created_by"`
		UpdatedBy   UserInfo   `bson:"updated_by" json:"updated_by"`
	}{
		Id: "673dcdb47484940c3f9fd08c",
		Title: SourceLang{
			En: "Uttara, EC",
			Bn: "উত্তরা,EC",
		},
		Address: SourceLang{
			En: "8th Floor, Millennium Tower, House 2, Road 7, Sector 3, Uttara",
			Bn: "৮ম ফ্লোর, মিলেনিয়াম টাওয়ার, হাউজ ২, রোড ৭, সেক্টর ৩, উত্তরা।",
		},
		Description: SourceLang{
			En: "Uttara Center",
			Bn: "উত্তরা কেন্দ্র",
		},
		MapLink:   "https://www.google.com/maps/dir/",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: UserInfo{Username: "creator_user", Email: "creator@example.com"},
		UpdatedBy: UserInfo{Username: "updater_user", Email: "updater@example.com"},
	}

	err := PrettyPrint(location)
	assert.NoError(t, err, "PrettyPrint should not return an error")
}
