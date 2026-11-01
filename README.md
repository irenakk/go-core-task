# go-core-task

My solutions to the Go-core module tasks. The task text is in [task.md](task.md) (in Russian). Each task is a small `main` package with unit tests, in a folder with the task number. Every task was done in its own branch first and then merged into `main`.

| Task | What it does |
|---|---|
| 1 | Prints the types of a few variables, joins them into one string, then hashes it as a rune slice with SHA-256 and the salt `go-2024` inserted in the middle |
| 2 | Slice helpers: random slice, keep even numbers, append, copy, remove by index |
| 3 | `StringIntMap` type with `Add`, `Remove`, `Copy`, `Exists`, `Get` |
| 4 | Strings from the first slice that are missing in the second |
| 5 | Intersection of two `[]int`, returns a `bool` and the common values |
| 6 | Random number generator on an unbuffered channel |
| 7 | Merges N channels into one |
| 8 | WaitGroup replacement built on a semaphore (buffered channel) instead of `sync.WaitGroup` |
| 9 | Pipeline: reads `uint8` from one channel, writes the cube as `float64` to another |

## How to run

Run all tests:

```bash
go test ./...
```

Or one task:

```bash
go test ./7/
go run ./7
```

Task 3 is only the `StringIntMap` type and its tests, so there's nothing to `go run` there.
