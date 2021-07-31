package mhw_test

import (
	"strings"
	"testing"

	"github.com/c032/go-mhw"
)

func TestLargeMonsters(t *testing.T) {
	monsters, err := mhw.LargeMonsters()
	if err != nil {
		t.Fatalf("LargeMonsters() returned an error: %s", err)
	}

	if got := len(monsters); got == 0 {
		t.Errorf("len(LargeMonsters()) = %d; want non-zero", got)
	} else {
		for i, monster := range monsters {
			if got := monster.Name(); strings.TrimSpace(got) == "" {
				t.Errorf("LargeMonsters()[%d].Name() = %#v; want non-empty string", i, got)
			}
			if got := monster.RawURL(); strings.TrimSpace(got) == "" {
				t.Errorf("LargeMonsters()[%d].RawURL() = %#v; want non-empty string", i, got)
			}
		}
	}
}

func TestSmallMonsters(t *testing.T) {
	monsters, err := mhw.SmallMonsters()
	if err != nil {
		t.Fatalf("SmallMonsters() returned an error: %s", err)
	}

	if got := len(monsters); got == 0 {
		t.Errorf("len(SmallMonsters()) = %d; want non-zero", got)
	} else {
		for i, monster := range monsters {
			if got := monster.Name(); strings.TrimSpace(got) == "" {
				t.Errorf("SmallMonsters()[%d].Name() = %#v; want non-empty string", i, got)
			}
			if got := monster.RawURL(); strings.TrimSpace(got) == "" {
				t.Errorf("SmallMonsters()[%d].RawURL() = %#v; want non-empty string", i, got)
			}
		}
	}
}

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
