package helper

import "time"

func TimeNow() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
}
