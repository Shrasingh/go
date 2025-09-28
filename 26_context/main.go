package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
  main.go - Context demo

  This file demonstrates:
  - What `context` is (briefly)
  - Why you should use it
  - How to use common context patterns:
      - context.Background()
      - context.WithCancel()
      - context.WithTimeout()
      - context.WithValue()
  - How contexts are passed to goroutines and how goroutines should listen for cancellation.

  Best practices (short):
  - Pass context as the first parameter to functions that do I/O or long work: func Foo(ctx context.Context, ...)
  - Don't store large values in a context; use it for request-scoped values only (IDs, auth tokens, small metadata).
  - Always cancel contexts you create with WithCancel/WithTimeout/WithDeadline to release resources.
*/

// ctxKey is a private type for context keys to avoid collisions with other packages.
type ctxKey string

// doWork simulates a chunked task that checks ctx.Done() between steps.
// It demonstrates cooperative cancellation: the goroutine stops when the context is canceled.
func doWork(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Example of reading a request-scoped value from context.
	if v := ctx.Value(ctxKey("reqID")); v != nil {
		fmt.Printf("[worker %d] got reqID from context: %v\n", id, v)
	}

	// Simulate work: 10 steps, but stop early if ctx is canceled.
	for step := 1; step <= 10; step++ {
		select {
		case <-ctx.Done():
			// ctx.Err() returns the cancellation reason (context.Canceled or context.DeadlineExceeded).
			fmt.Printf("[worker %d] stopping at step %d: reason=%v\n", id, step, ctx.Err())
			return
		default:
			// Do a piece of work (simulate with Sleep).
			fmt.Printf("[worker %d] working step %d/10\n", id, step)
			time.Sleep(300 * time.Millisecond) // simulate work chunk
		}
	}

	fmt.Printf("[worker %d] completed all steps successfully\n", id)
}

func main() {
	fmt.Println("=== Go context: what / why / how (runnable demo) ===")
	fmt.Println()

	// -----------------------
	// 1) Root context
	// -----------------------
	// What: context.Background() returns a non-cancellable, empty root context.
	// Why: use it at the top level (main, initialization, tests) as the base for derived contexts.
	// How: derive other contexts from this (WithCancel, WithTimeout, WithValue, etc).
	root := context.Background()

	// -----------------------
	// 2) WithCancel example
	// -----------------------
	// What: WithCancel returns (ctx, cancel). Calling cancel() signals cancellation to all derived contexts.
	// Why: use it when you want to manually stop work (e.g., client closed connection, user clicked stop).
	// How: start goroutines with the ctx; call cancel() from the controlling goroutine to stop them.
	fmt.Println("-> Example: context.WithCancel")
	{
		ctx, cancel := context.WithCancel(root)
		var wg sync.WaitGroup

		wg.Add(1)
		go doWork(ctx, 1, &wg)

		// Let the worker run for ~1 second then cancel.
		time.Sleep(1 * time.Second)
		fmt.Println("[main] calling cancel() to stop worker 1")
		cancel() // signals the worker to stop

		wg.Wait()
		fmt.Println()
	}

	// -----------------------
	// 3) WithTimeout example
	// -----------------------
	// What: WithTimeout returns a context that is automatically canceled after the timeout.
	// Why: protect against slow remote calls (don't block forever).
	// How: pass ctx to I/O operations (DB calls, HTTP requests). If the operation exceeds the timeout, it will be canceled.
	fmt.Println("-> Example: context.WithTimeout")
	{
		// Create a context that will be cancelled automatically after 900ms.
		ctx, cancel := context.WithTimeout(root, 900*time.Millisecond)
		defer cancel() // good practice to call cancel to free resources even when the timeout triggers

		var wg sync.WaitGroup
		wg.Add(1)
		go doWork(ctx, 2, &wg)

		// Wait longer than the timeout so we can see deadline exceeded.
		time.Sleep(2 * time.Second)
		wg.Wait()
		fmt.Println()
	}

	// -----------------------
	// 4) WithValue example
	// -----------------------
	// What: WithValue attaches a small, request-scoped value to the context.
	// Why: carry metadata (request ID, auth token) through call chains without changing many function signatures.
	// How: define a private key type, store the value with WithValue, and retrieve it in lower-level functions.
	fmt.Println("-> Example: context.WithValue (request-scoped metadata)")
	{
		// Create a context that carries a request ID.
		reqCtx := context.WithValue(root, ctxKey("reqID"), "request-abc-123")

		var wg sync.WaitGroup
		wg.Add(1)
		go doWork(reqCtx, 3, &wg)

		wg.Wait()
		fmt.Println()
	}

	// -----------------------
	// 5) Summary / best-practices (printed)
	// -----------------------
	fmt.Println("Summary / best practices:")
	fmt.Println("- Pass context as the first parameter: func Foo(ctx context.Context, ...).")
	fmt.Println("- Use WithCancel to manually stop work; always call cancel() to release resources.")
	fmt.Println("- Use WithTimeout/WithDeadline to automatically bound work duration.")
	fmt.Println("- Use WithValue sparingly and only for request-scoped small values (IDs, small metadata).")
	fmt.Println("- Have goroutines check <-ctx.Done() and return quickly when canceled.")
	fmt.Println()

	// -----------------------
	// 6) How this maps to MongoDB (example snippet, not executed)
	// -----------------------
	// In real DB code you typically use a context with timeout for connect/ping and queries, e.g.:
	//    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//    defer cancel()
	//    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	//    // then use ctx for operations like client.Ping(ctx, nil)
	//
	// That way if the network is slow the context cancels and your program doesn't hang forever.
	//
	// (If you want the exact MongoDB snippet in your project, I can provide it — but keep DB imports commented out here
	// so this demo remains runnable without extra dependencies.)

	fmt.Println("Demo finished.")
}
