package test

import (
	"github.com/raidnav/algorist/math"
	"testing"
)

func TestPrimeSieve(t *testing.T) {
	if !math.IsPrime(1000_000_007) {
		t.Fail()
	}
	if !math.IsPrime(7) {
		t.Fail()
	}
	if math.IsPrime(4) {
		t.Fail()
	}
}
