package track

import "time"

type activityTracker interface {
	Track(ts time.Time, slug string)
}

type Activity struct {
	tracker activityTracker
}

func NewActivity(at activityTracker) *Activity {
	return &Activity{
		tracker: at,
	}
}

func (a *Activity) Track(slug string) {
	a.tracker.Track(time.Now(), slug)
}
