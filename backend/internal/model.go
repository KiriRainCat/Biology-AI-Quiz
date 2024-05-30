package internal

import "time"

type record struct {
	Name       string `json:"name,omitempty" gorm:"primaryKey"`
	Passphrase string `json:"passphrase,omitempty"`

	Accuracy  float64   `json:"accuracy,omitempty"`
	TimeTaken float64   `json:"time_taken,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type question struct {
	Question string   `json:"question,omitempty"`
	Choices  []string `json:"choices,omitempty"`
	Answer   int      `json:"answer"`
}
