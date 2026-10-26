# Lab 7 hand trace

Rename your completed copy to `Lab7_Trace.md` before submitting.

| Event ID | Process | Kind | Message | Local state before | Carried stamp | Lamport after | Vector after |
| --- | --- | --- | --- | --- | --- | ---: | --- |
| a | P1 | local | — | | — | | |
| b | P2 | local | — | | — | | |
| send-m | P1 | send | m | | — | | |
| receive-m | P2 | receive | m | | | | |
| send-n | P2 | send | n | | — | | |
| receive-n | P1 | receive | n | | | | |

## Causal pair

Name two events, give their vectors, and identify the path or component-wise comparison that establishes the order.

## Concurrent pair

Name two events, give their vectors, and show why neither vector dominates the other.

## Why Lamport sorting is insufficient

Explain what a Lamport timestamp can and cannot prove about the pair above.

## AI Usage

Name any tools and prompts or interactions that affected this work, or write `AI Usage: None.`
