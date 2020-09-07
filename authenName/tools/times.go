package tools

import "time"

func TimeNow() time.Time {
	t, _ := time.Parse("2 Jan 2006 15:04:05", time.Now().String())
	return t
}
