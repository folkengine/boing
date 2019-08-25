package main

import "testing"

func TestPing(t *testing.T) {
	t.Run("Boing", func(t *testing.T) {
		Boing()
	})
	t.Run("Klaxon", func(t *testing.T) {
		Klaxon()
	})
	t.Run("Ping", func(t *testing.T) {
		Ping()
	})
}
