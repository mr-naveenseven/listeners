# Listeners

A simple Go project demonstrating **concurrent message listeners** with **graceful shutdown**.

---

## 🧩 What It Does

This project spawns multiple goroutines (“listeners”) that poll a shared message at different intervals using `time.Ticker`.  
Each listener reads the message safely via mutex protection and can be stopped gracefully using channels and `sync.WaitGroup`.

---

## 🚀 Quick Start

### Navigate to the project directory
    cd listeners

### Run directly
    go run .

### Or build and run
    go build -o listeners main.go
    ./listeners

## Expected Output
Message data from Pro plan listener: initial message\
Message data from Pro plan listener: initial message\
Message data from Base plan listener: initial message\
Pro plan listener stopped!\
Base plan listener stopped!

## 🗂️ Project Structure

| File | Description |
|------|--------------|
| **main.go** | Creates message, starts two listeners (1s and 5s intervals), runs for 6 seconds, then stops them gracefully. |
| **listener.go** | Defines the `Listener` struct with `Start()` and `Stop()` methods for managing goroutines. |
| **message.go** | Defines the thread-safe `Message` struct with mutex-protected read access. |

---

## 🧠 Key Concepts

- **Goroutine Management** — Uses `sync.WaitGroup` to track and wait for all listener goroutines to complete.  
- **Graceful Shutdown** — Each listener stops cleanly via a `Done` channel signal.  
- **Thread Safety** — Shared `Message` reads are protected by `sync.Mutex`.  
- **Idempotent Stop** — Uses `sync.Once` to safely call `Stop()` multiple times without panics.

---

## 👨‍💻 Developer

**GitHub:** [mr-naveenseven](https://github.com/mr-naveenseven)
