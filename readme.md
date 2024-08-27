# Go Intermediate

## Magesh Kuppan
- tkmagesh77@gmail.com

## Schedule
|       | Duration |
|-------|----------|
| Session-01 | 1:50 hrs |
| Tea Break | 20 mins |
| Session-02 | 1:50 mins |

## Software Requirements
- Go tools (https://go.dev/dl)
- Any Editor
- Git Client

## Methodology
- No powerpoints
- Code & Discuss
- No dedicated Q & A time

## Repository
- https://github.com/tkmagesh/nutanix-gointermediate-aug-2024-2

## Prerequisites
- Data Types, Variables, Constants, iota
- Programming Constructs (if else, for, switch case)
- Functions
    - Higher Order Functions
    - Deferred Functions
- Errors
- Panic & Recovery
- Structs & Methods
    - Struct Composition
- Interfaces
- Modules & Packages

## Agenda
- Concurrency
- Concurrency Patterns
- IO
- Testing
- Database Programming
- HTTP Services

## Concurrency Programming
- Concurrency is NOT parallelism
- Application with more than one execution path
- Typically achieved using OS Threads
- OS Threads are costly
    - ~2MB 
    - Creating & Destroying (Thread Pools)
    - Context switch
## Go Concurrency
- Builtin scheduler
- Concurrent operations are represented as goroutines
- Goroutines are cheap (~4KB)
- Context switches are fast

- Concurrency support is built in the language
    - go keyword, channel data type, channel operator ( <- ), range & select-case constructs

### WaitGroup
- semaphore based counter
- can block the execution of the current function until the counter becomes 0