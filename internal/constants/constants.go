package constants

type DomainStatus int64

const (
	CatchAll DomainStatus = iota
	NotCatchAll
)

const EventsThreshHold = 1000

func (d DomainStatus) String() string {
	switch d {
	case CatchAll:
		return "catch-all"
	case NotCatchAll:
		return "not catch-all"
	}
	return "unknown"
}
