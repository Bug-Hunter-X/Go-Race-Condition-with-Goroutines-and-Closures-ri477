# Go Race Condition with Goroutines and Closures

This repository demonstrates a common race condition in Go programs using goroutines and closures. The program aims to increment a shared counter using multiple goroutines. However, without proper synchronization, the final count might be less than expected due to race conditions.

## Problem

The `count` variable is accessed and modified by multiple goroutines concurrently without any synchronization mechanism. This leads to race conditions and unpredictable results.  The `defer wg.Done()` ensures that the goroutines are properly coordinated with the waitgroup, however the increment of count itself is not protected and the expected value of `10` isn't always what's received.

## Solution

The solution involves using a mutex to protect the `count` variable from concurrent access.  This ensures that only one goroutine can modify the variable at a time, preventing the race condition.