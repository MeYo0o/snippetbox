package components

import "time"

func formateDate(t time.Time) string {
	return t.Format("Jan 02, 2006")
}
