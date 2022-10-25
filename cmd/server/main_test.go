package main

import (
	"os"
	"testing"
)

func TestDefaultPortHandler(t *testing.T) {

	port := portHandler()

	if port != ":8080" {
		t.Errorf("Error, default port should be :8080")
	} else {
		t.Logf("Success, default port is :8080")
	}
}

func TestEnvPorHandler(t *testing.T) {
	os.Setenv("PORT", "8888")

	port := portHandler()

	if port != ":8888" {
		t.Errorf("Error, port should be :8888")
	} else {
		t.Logf("Success, port changed by env to :8888")
	}

	os.Unsetenv("PORT")
}
