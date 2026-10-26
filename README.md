# CS 411 Lab 7 Chrono Solver starter

Starter template: https://github.com/cwillpinto/cs411-spring2027-lab07-chrono-solver

This is the intentionally incomplete Go starter for the Spring 2027 Lab 7 handout. The provided trace, command-line runner, and tests establish the contract. Complete the `TODO` rules in `clock.go` and `solver.go`; do not change the expected results in the tests.

The full prompt is `Lab_07_Chrono_Solver.md`. `Lab7_Trace_Template.md` is a blank analysis scaffold; rename your finished copy to `Lab7_Trace.md`.

## Start

Use Go 1.23 or newer. From this repository root:

```text
go test ./...
go run . data/example_trace.json > Lab7_Output.json
```

The unchanged starter compiles and writes JSON, but its clock values are intentionally wrong and the tests fail. After implementing the rules, save the final `go test ./...` output in `Lab7_Terminal.txt` and submit the evidence listed in the handout. The supplied JSON trace is ordered so each send appears before its matching receive.

Lamport timestamps preserve the direction of a causal edge but cannot prove that two events are causally related. Use vector comparison for the required concurrent pair.

## Work in your own GitHub repository

The instructor starter is an upstream source. Keep your completed work in a repository under your own GitHub account, then submit the repository URL and the evidence listed in the assignment handout through Blackboard.

If Blackboard links a template repository, choose **Use this template**, create the new repository under your account, and clone that new repository. Check `git remote -v` before pushing; `origin` must be your repository.

For a direct clone, first create an empty repository under your account. Replace the placeholders below with the starter link supplied in Blackboard and your own empty repository URL:

```sh
git clone https://github.com/cwillpinto/cs411-spring2027-lab07-chrono-solver
cd <cloned-directory>
git remote rename origin upstream
git remote add origin <your-empty-repository-url>
git remote -v
git push -u origin HEAD
```

Commit and push your completed work to your own `origin`. Do not push to the instructor's upstream repository. Use the exact submission files and verification commands in the assignment handout.
