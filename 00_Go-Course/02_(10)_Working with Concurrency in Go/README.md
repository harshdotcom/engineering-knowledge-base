#### 1.  Don't Communicate by sharing memory, share memory by communicating

#### 2.  Don't over-enginner things by using shared memory and complicated, error-prone synchronization primitives; instead use message-padding between GoRoutine so varibles and data can be used in the appropriate sequence.

#### 3.  Keep your application's complexity to an absoulute minimum; it's easier to write, easier to understand, easier to maintain.

- We'll start by looking at the basics types in the sync package: mutexes (semaphores) and wait groups
- We'll go through three of the classic computer science problems:
- The Porducer/Consumer problem
- The Dining philosopher problem
- The Sleeping Barber Problem

#### 4. These are classic problems for a reason: they give students excellent exposure to the problem inherent in concurrent programming, and they force them to figure out an efficient method of solving such problems


#### 5. We'll also cover a more real-world scenario, where we build a subset of a larger (imaginary) problem where a user of a service wants to subscribe and buy one of series of available subscriptions.

#### 6. When they purchase a subscription, we will generate invoice, sent an email and generate pdf manual and send that to the user.

#### 7. We'll do these things concurrently


- Mutex = "Mutual exclusion" - allowed us to deal with race conditions
- Relative simple to use
- Dealing with shared resourses and concurrent/Parallel goroutines
- Lock/Unlock
- We can test for race conditions when running code, or testing it  
- Race conditions occur when multiple GoRoutines try to access the same data
- Can be difficult to spot when reading Code
- Go allows us to check for them when running a Program, or when testing out code with go test.
- Channels are a means of having Go Routines share data
- They can talk to each other
- This is Go' philosopy of having things share memory by communicating, rather than communicating by sharing memory.
- The producer/consumer problem
