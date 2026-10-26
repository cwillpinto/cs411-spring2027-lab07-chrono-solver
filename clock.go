package main

// Stamp is the state attached to an event or carried by a message.
type Stamp struct {
	Lamport int   `json:"lamport"`
	Vector  []int `json:"vector"`
}

// Clock owns one process's logical time. Process IDs are zero-based here;
// the trace file uses names such as P1 and P2.
type Clock struct {
	id      int
	lamport int
	vector  []int
}

func NewClock(processCount, id int) *Clock {
	return &Clock{id: id, vector: make([]int, processCount)}
}

func (c *Clock) snapshot() Stamp {
	return Stamp{Lamport: c.lamport, Vector: append([]int(nil), c.vector...)}
}

// Local records a local or send event. Increment both logical clocks before
// returning a snapshot. The returned vector must not alias c.vector.
func (c *Clock) Local() Stamp {
	// TODO: implement the local-event update rule.
	return c.snapshot()
}

// Receive records a message receipt. Merge the carried vector component-wise,
// then increment this process's component. The Lamport rule is max + 1.
func (c *Clock) Receive(message Stamp) Stamp {
	// TODO: implement both receive update rules.
	return c.snapshot()
}
