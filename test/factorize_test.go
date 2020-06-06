package test

import (
	"github.com/raidnav/algorist/math"
	"testing"
)

func TestGCD(t *testing.T) {
	if 10 != math.GCD(10, 20) {
		t.Fail()
	}
}

func TestLCM(t *testing.T) {
	if 20 != math.LCM(10, 20) {
		t.Fail()
	}
}
