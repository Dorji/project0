package main

import (
	"testing"
)

func TestNewObject(t *testing.T) {
	tests := []struct {
		name     string
		input    Family
		wantName string
		wantErr  bool
	}{
		{
			name:     "Base case",
			input:    &Base{},
			wantName: "Parent",
			wantErr:  false,
		},
		{
			name:     "Child case",
			input:    &Child{},
			wantName: "Child",
			wantErr:  false,
		},
		{
			name:     "Nil case",
			input:    nil,
			wantName: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewObject(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			switch v := got.(type) {
			case *Base:
				if v.name != tt.wantName {
					t.Errorf("Base.name = %v, want %v", v.name, tt.wantName)
				}
			case *Child:
				if v.name != tt.wantName {
					t.Errorf("Child.name = %v, want %v", v.name, tt.wantName)
				}
				if v.lastName != "Inherited" {
					t.Errorf("Child.lastName = %v, want Inherited", v.lastName)
				}
			default:
				t.Errorf("Unexpected type %T", v)
			}
		})
	}
}
