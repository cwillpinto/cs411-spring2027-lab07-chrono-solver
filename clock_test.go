package main

import (
	"reflect"
	"testing"
)

func TestLocalAndReceiveRules(t *testing.T) {
	a, b := NewClock(2, 0), NewClock(2, 1)
	first := a.Local()
	if first.Lamport != 1 || !reflect.DeepEqual(first.Vector, []int{1, 0}) {
		t.Fatalf("first local event = %+v, want L=1 V=[1 0]", first)
	}
	second := a.Local()
	if second.Lamport != 2 || !reflect.DeepEqual(second.Vector, []int{2, 0}) {
		t.Fatalf("second local/send = %+v, want L=2 V=[2 0]", second)
	}
	b.Local()
	received := b.Receive(second)
	if received.Lamport != 3 || !reflect.DeepEqual(received.Vector, []int{2, 2}) {
		t.Fatalf("receive = %+v, want L=3 V=[2 2]", received)
	}
	if !reflect.DeepEqual(first.Vector, []int{1, 0}) {
		t.Fatal("an earlier stamp changed after another event")
	}
}

func TestReceiveUsesLargerLocalTime(t *testing.T) {
	c := NewClock(2, 1)
	for i := 0; i < 8; i++ {
		c.Local()
	}
	got := c.Receive(Stamp{Lamport: 5, Vector: []int{2, 0}})
	if got.Lamport != 9 || !reflect.DeepEqual(got.Vector, []int{2, 9}) {
		t.Fatalf("receive with larger local time = %+v, want L=9 V=[2 9]", got)
	}
}
