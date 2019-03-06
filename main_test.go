package main

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestSlack(t *testing.T) {
	godotenv.Load()
	if err := slack(); err != nil {
		panic(err)
	}
}
