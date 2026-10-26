package main

import "fmt"

type Event struct {
	ID      string `json:"id"`
	Process string `json:"process"`
	Kind    string `json:"kind"` // local, send, or receive
	Message string `json:"message,omitempty"`
}

type Trace struct {
	Processes []string `json:"processes"`
	Events    []Event  `json:"events"`
}

type AnnotatedEvent struct {
	Event
	Stamp
}

type Relation string

const (
	Before     Relation = "before"
	After      Relation = "after"
	Equal      Relation = "equal"
	Concurrent Relation = "concurrent"
)

// Compare reports the component-wise relation between two vector clocks.
// Equal vectors are equal state; incomparable vectors are concurrent.
func Compare(left, right []int) (Relation, error) {
	if len(left) != len(right) {
		return "", fmt.Errorf("vector lengths differ: %d and %d", len(left), len(right))
	}
	// TODO: compare every component and return Before, After, Equal, or Concurrent.
	return Equal, nil
}

// Solve processes a topologically ordered trace. A send carries its event's
// stamp; a receive consumes the stamp of the matching earlier send.
func Solve(trace Trace) ([]AnnotatedEvent, error) {
	if len(trace.Processes) == 0 {
		return nil, fmt.Errorf("trace has no processes")
	}
	clocks := make(map[string]*Clock, len(trace.Processes))
	for index, name := range trace.Processes {
		if name == "" || clocks[name] != nil {
			return nil, fmt.Errorf("empty or duplicate process %q", name)
		}
		clocks[name] = NewClock(len(trace.Processes), index)
	}
	messages := make(map[string]Stamp)
	consumed := make(map[string]bool)
	seenEvents := make(map[string]bool)
	result := make([]AnnotatedEvent, 0, len(trace.Events))
	for _, event := range trace.Events {
		clock := clocks[event.Process]
		if event.ID == "" || seenEvents[event.ID] || clock == nil {
			return nil, fmt.Errorf("invalid event ID or process: %+v", event)
		}
		seenEvents[event.ID] = true
		var stamp Stamp
		switch event.Kind {
		case "local":
			if event.Message != "" {
				return nil, fmt.Errorf("local event %q has a message", event.ID)
			}
			stamp = clock.Local()
		case "send":
			if event.Message == "" || messages[event.Message].Vector != nil {
				return nil, fmt.Errorf("missing or duplicate send ID %q", event.Message)
			}
			stamp = clock.Local()
			messages[event.Message] = stamp
		case "receive":
			message, ok := messages[event.Message]
			if event.Message == "" || !ok || consumed[event.Message] {
				return nil, fmt.Errorf("receive %q has no unmatched earlier send", event.Message)
			}
			stamp = clock.Receive(message)
			consumed[event.Message] = true
		default:
			return nil, fmt.Errorf("event %q has unknown kind %q", event.ID, event.Kind)
		}
		result = append(result, AnnotatedEvent{Event: event, Stamp: stamp})
	}
	return result, nil
}
