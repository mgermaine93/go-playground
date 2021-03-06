# Concurrency Notes

- 2006 -- Intel releases the first dual-core CPU.
- 2007 -- Google begins creating Go specifically to take advantage of multiple-core processes.
- 2012 -- V1 of Go is released.
- This has significant implications for concurrent code in Go.
- Concurrency is natively built right into Go -- "eeeeeease of programming".
- "Makes concurrency easy".
- Concurrency vs. Parallelism
  - If you write Go code on a single-core CPU machine, it's not going to run in parallel.
  - Multiple threads of Go code will not run. It will run sequentially, one line after another.
  - CONCURRENCY is a DESIGN PATTERN, or a way we WRITE CODE.
  - One of the primary factors of whether or not your code runs in parallel is if you have more than one CPU.
  - Concurrent programming in many environments is made difficult by the subtleties required to implement correct access to shared variables.
- So what's the actual difference between the two?
  - CONCURRENCY is the COMPOSITION of independently executing processes. CONCURRENCY is about DEALING WITH lots of things at once.
  - PARALLELISM is the simultaneous EXECUTION of (possibly related) computations. PARALLELISM is about DOING lots of things at once.
- WaitGroups
  - When designing code, you need to think about all of the things happening at once rather than sequential stuff.
  - A WAITGROUP waits for collection of GOROUTINES to finish.
    - .Add() sets the number Goroutines to wait for.
    - .Done() is called when each Goroutine is finished.
    - .Wait() can be used to block stuff until all Goroutines have finished.
  - Writing concurrent code just involves (among other things) putting "go" in front of a function or method call, e.g. "go foo()" rather than just "foo()".
  - For WaitGroups, remember that if the RECEIVER is of type POINTER, then we can only use a value of type POINTER.
- Goroutines
  - A Goroutine is a function executing concurrently with other Goroutines (functions) in the same address space.
  - Each one is lightweight, costing little more than allocations of stack space.
- Method Sets
  - "The METHOD SET of a TYPE determines the INTERFACES that the TYPE implements..."
- Race Conditions
  - Race conditions are not exactly good code.
  - They will give unpredictable results.
  - Mutexes (?) are a solution.
  - But what actually are they...? Still need to figure this out.
- Mutex
  - A "mutex" is a mutual exclusion lock.
  - They allow us to "lock down" code so that only one goroutine can access that locked chunk of code at a time.
- Package Atomic
  - Also prevents race conditions.
  - The functions here require great care to be used correctly.
  - Except for special, low-level applications, synchronization is better done with channels or the facilities of the "sync" package.
  - "Share memory by communicating, don't communicate by sharing memory".

* "Go run your_program_here" runs a program.
* "Go build your_program_here" drops a program into a repo, from which you can run the program.
* "rm \_\_\_ your_program_here" removes the program.
* "Go install your_program_here" installs the program into your Go Path.
