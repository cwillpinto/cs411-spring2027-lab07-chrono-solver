package main

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestExampleTrace(t *testing.T) {
	data, err := os.ReadFile("data/example_trace.json")
	if err != nil {
		t.Fatal(err)
	}
	var trace Trace
	if err := json.Unmarshal(data, &trace); err != nil {
		t.Fatal(err)
	}
	got, err := Solve(trace)
	if err != nil {
		t.Fatal(err)
	}
	wantLamport := []int{1, 1, 2, 3, 4, 5}
	wantVector := [][]int{{1, 0}, {0, 1}, {2, 0}, {2, 2}, {2, 3}, {3, 3}}
	if len(got) != len(wantLamport) {
		t.Fatalf("got %d events", len(got))
	}
	for i := range got {
		if got[i].Lamport != wantLamport[i] || !reflect.DeepEqual(got[i].Vector, wantVector[i]) {
			t.Errorf("event %s = L%d V%v, want L%d V%v", got[i].ID, got[i].Lamport, got[i].Vector, wantLamport[i], wantVector[i])
		}
	}
}

func TestVectorRelations(t *testing.T) {
	cases := []struct {
		a, b []int
		want Relation
	}{
		{[]int{1, 0}, []int{2, 2}, Before},
		{[]int{2, 2}, []int{1, 0}, After},
		{[]int{1, 0}, []int{0, 1}, Concurrent},
		{[]int{2, 2}, []int{2, 2}, Equal},
	}
	for _, tc := range cases {
		got, err := Compare(tc.a, tc.b)
		if err != nil || got != tc.want {
			t.Errorf("Compare(%v,%v) = %q, %v; want %q", tc.a, tc.b, got, err, tc.want)
		}
	}
	if _, err := Compare([]int{1}, []int{1, 0}); err == nil {
		t.Error("mismatched vector lengths must fail")
	}
}

func TestBadTraceFailsClearly(t *testing.T) {
	_, err := Solve(Trace{Processes: []string{"P1"}, Events: []Event{{ID: "r", Process: "P1", Kind: "receive", Message: "absent"}}})
	if err == nil {
		t.Fatal("receive without a send must fail")
	}
}
