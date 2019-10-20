package boing

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
	t.Run("Tone", func(t *testing.T) {
		Tone()
	})
	t.Run("Chicken", func(t *testing.T) {
		Chicken()
		Play(Chickens)
		Play(Chicken2)
		Play(Chicken3)
		Play(Chicken4)
		Play(ChickensInBarn)
		Play(Rooster)
	})
}
