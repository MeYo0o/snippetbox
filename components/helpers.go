package components

import (
	"strings"
	"time"
)

func formateDate(t time.Time) string {
	return t.Format("Jan 02, 2006")
}

func renderContent(s string) string {
	// Accept both actual newline characters and literal "\n" escapes.
	s = strings.ReplaceAll(s, `\r\n`, "\n")
	s = strings.ReplaceAll(s, `\n`, "\n")
	return s
}
