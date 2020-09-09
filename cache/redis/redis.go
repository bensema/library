package redis

import "time"

type Error string

func (err Error) Error() string { return string(err) }

// SlowLog represents a redis SlowLog
type SlowLog struct {
	// ID is a unique progressive identifier for every slow log entry.
	ID int64

	// Time is the unix timestamp at which the logged command was processed.
	Time time.Time

	// ExecutationTime is the amount of time needed for the command execution.
	ExecutionTime time.Duration

	// Args is the command name and arguments
	Args []string

	// ClientAddr is the client IP address (4.0 only).
	ClientAddr string

	// ClientName is the name set via the CLIENT SETNAME command (4.0 only).
	ClientName string
}
