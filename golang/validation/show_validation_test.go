package validation

import (
	"testing"

	pb "github.com/organicveggie/live-shows/protos"
)

func TestLocationValidation(t *testing.T) {
	tests := []struct {
		name    string
		loc     *pb.Location
		wantErr bool
	}{
		{
			name:    "Empty fails",
			loc:     &pb.Location{},
			wantErr: true,
		},
		{
			name: "Missing name fails",
			loc: &pb.Location{
				City:    "Madison",
				State:   "WI",
				Country: "United States",
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.loc.Validate(); (err != nil) != test.wantErr {
				t.Errorf("Validation mismatch. Expected error == %t, got error == %t", test.wantErr, err != nil)
			}
		})
	}
}

func TestShowValidation(t *testing.T) {
	tests := []struct {
		name    string
		show    *pb.Show
		wantErr bool
	}{
		{
			name:    "Empty fails",
			show:    &pb.Show{},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.show.Validate(); (err != nil) != test.wantErr {
				t.Errorf("Validation mismatch. Expected error == %t, got error == %t", test.wantErr, err != nil)
			}
		})
	}
}
