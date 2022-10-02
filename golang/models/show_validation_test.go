package tests

import (
	"testing"

	pb "github.com/organicveggie/live-shows/models"
)

func TestShowValidation(t *testing.T) {
	tests := []struct {
		name      string
		show      pb.Show
		wantError bool
	}{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.show.Validate(); (err != null) != test.wantError {
				t.Error("Validation mismatch. Expected error == %b, got error == %b", test.wantError, err != null)
			}
		})
	}
}
