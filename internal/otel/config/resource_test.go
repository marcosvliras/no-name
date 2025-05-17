package config

import "testing"

func TestNewResource(t *testing.T) {
	// Test the newResource function
	err := newResource()
	if err != nil {
		t.Fatalf("newResource() failed: %v", err)
	}

	// Check if the Resource is not nil
	if Resource == nil {
		t.Fatal("Resource is nil")
	}
}
