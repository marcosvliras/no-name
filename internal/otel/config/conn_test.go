package config

import "testing"

func TestInitConn(t *testing.T) {
	// Test the initConn function
	err := initConn()
	if err != nil {
		t.Fatalf("initConn() failed: %v", err)
	}

	// Check if the Conn is not nil
	if Conn == nil {
		t.Fatal("Conn is nil")
	}
}
