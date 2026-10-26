# Lab 7: Chrono Solver

**CS 411: Computing Paradigms - Spring 2027**  
**Week 8: March 1-7, 2027**  
**Due:** March 7, 2027  
**Value:** 100 points

**Starter repository:** [Open the Spring 2027 starter](https://github.com/cwillpinto/cs411-spring2027-lab07-chrono-solver)

## Purpose

Recover causal order from a distributed message trace. Compute Lamport timestamps, compute vector clocks, and identify which event pairs remain concurrent.

## Learning objectives

- Apply Lamport update rules
- Compute vector clocks by process
- Distinguish causal order from concurrency

## Before you begin

1. Review the Week 8 slide deck and assigned VoiceThread or reading.
2. Use the Lab 7 GitHub starter when its link appears in Blackboard. Make your own repository from that starter and clone it locally.
3. From the repository root, run `go test ./...` and `go run . data/example_trace.json`. The unchanged starter runs but fails the rule tests at the marked `TODO` methods; save this baseline before editing.

The starter includes `clock.go`, `solver.go`, a command-line runner, `data/example_trace.json`, and public tests. It requires Go 1.23 or newer. A trace lists processes and events in causal input order, with each send appearing before its matching receive.

## Part A Lamport clocks (30 points)

Calculate every event in `data/example_trace.json` by hand and record the result in `Lab7_Trace.pdf` or `Lab7_Trace.md`.

- Increment before each local or send event.
- On receive, take the maximum timestamp and add one.
- Explain why sorting Lamport values alone can overstate order.
- Show the local value and carried message value at each receive.

## Part B Vector clocks (40 points)

Track one component per process for the same trace.

- Update the local component on every event.
- Merge component-wise maxima on receive.
- Mark incomparable event pairs.
- State the component-wise comparison that proves one causal pair and one concurrent pair.

## Part C Solver and explanation (30 points)

Complete the `TODO` updates in `clock.go` and the vector comparison in `solver.go`. Keep the provided trace parser and test expectations intact.

- Run `go test ./...` from the repository root.
- Run `go run . data/example_trace.json > Lab7_Output.json` and compare the JSON values with the hand trace.
- Save the final test output and command used in `Lab7_Terminal.txt`.

## Verification

- `go test ./...` passes, including the local, receive, vector-relation, and full-trace checks.
- `go run . data/example_trace.json` emits six annotated events matching the hand trace.
- Every receive follows its send in both Lamport and vector order; incomparable vectors are reported as concurrent.

## What to submit

Submit your own repository URL in Blackboard and include these files in that repository or the Blackboard handoff as directed:

1. `clock.go` and `solver.go` with the completed rules.
2. `Lab7_Trace.pdf` or `Lab7_Trace.md` with the full hand trace and causal/concurrent explanation.
3. `Lab7_Output.json` produced by the command-line runner.
4. `Lab7_Terminal.txt` showing the final `go test ./...` command and result.

## Rubric

| Criterion | Points |
| --- | ---: |
| Lamport trace and receive calculations | 30 |
| Vector trace and component-wise comparisons | 25 |
| Causal and concurrent explanation | 15 |
| Working clock and solver rules | 20 |
| Output and test evidence | 10 |
| **Total** | **100** |

## AI usage and academic integrity

Add an **AI Usage** section to the required analysis file. Name each tool and list the prompts or interactions that affected the work. If you did not use AI, write `AI Usage: None.` You must be able to explain and defend every submitted design and line of code.

## Late work

Late work may be accepted for up to three calendar days with a 25 percent deduction per day, following the course syllabus.
