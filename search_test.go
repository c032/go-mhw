package mhw_test

import (
	"strings"
	"testing"

	"github.com/c032/go-mhw"
)

func TestMonsters(t *testing.T) {
	monsters, err := mhw.Monsters()
	if err != nil {
		t.Fatalf("Monsters() returned an error: %s", err)
	}

	if got := len(monsters); got == 0 {
		t.Errorf("len(Monsters()) = %d; want non-zero", got)
	} else {
		for i, monster := range monsters {
			if got := monster.Name(); strings.TrimSpace(got) == "" {
				t.Errorf("Monsters()[%d].Name() = %#v; want non-empty string", i, got)
			}
			if got := monster.RawURL(); strings.TrimSpace(got) == "" {
				t.Errorf("Monsters()[%d].RawURL() = %#v; want non-empty string", i, got)
			}
		}
	}
}
