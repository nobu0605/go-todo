package constants

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (s Status) String() string {
	switch s {
	case Todo:
		return "todo"
	case InProgress:
		return "inProgress"
	case Done:
		return "done"
	}
	return "unknown"
}

func GetStatuses() []Status {
	return []Status{Todo, InProgress,Done}
}