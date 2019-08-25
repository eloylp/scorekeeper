package scorekeeper_test

import (
	"github.com/eloylp/scorekeeper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInitialScore(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	assert.Equal(t, 0, scorer.Get("Sandy"))
}

func TestAddPoints(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	assert.Equal(t, 10, scorer.Get("Sandy"))
}

func TestAddPointsAgain(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	scorer.Add("Sandy", 10)
	assert.Equal(t, 20, scorer.Get("Sandy"))
}

func TestSubsPoints(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Add("Sandy", 10)
	scorer.Subs("Sandy", 5)
	assert.Equal(t, 5, scorer.Get("Sandy"))
}

func TestNonNegativePoints(t *testing.T) {
	scorer := scorekeeper.NewScorer()
	scorer.Subs("Sandy", 5)
	assert.Equal(t, 0, scorer.Get("Sandy"))
}
