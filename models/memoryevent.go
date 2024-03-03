package models

import "time"

type ExtentRemembered int64

const (
	All ExtentRemembered = iota
	OnlyWord
	OnlyArticle
	None
)

func (e ExtentRemembered) String() string {
	switch e {
	case All:
		return "all"
	case OnlyWord:
		return "only-word"
	case OnlyArticle:
		return "only-article"
	case None:
		return "none"
	}
	return "unknown"
}

type MemoryEvent struct {
	Id               string
	NounId           string
	TimeStamp        time.Time
	ExtentRemembered ExtentRemembered
}
