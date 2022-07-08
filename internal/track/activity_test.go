package track_test

import (
	"habitual/internal/track"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type mockActivityTracker struct {
	ts   time.Time
	slug string
}

func (m *mockActivityTracker) Track(ts time.Time, slug string) {
	m.ts = ts
	m.slug = slug
}

type ActivityTrackerTestSuite struct {
	suite.Suite
}

func TestTracksIncludesTimestampAndSlug(t *testing.T) {
	mt := &mockActivityTracker{}
	a := track.NewActivity(mt)

	a.Track("test")

	min := time.Now().Add(-(time.Hour * 24))
	if !mt.ts.After(min) {
		assert.Fail(t, "time should be the result of a time.Now()")
	}
}
