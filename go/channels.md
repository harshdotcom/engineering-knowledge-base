# 📡 Channels

## 📖 Introduction

**Channels** are a communication mechanism used in concurrent programming to allow different execution units (such as goroutines or threads) to safely exchange data.

Instead of sharing memory directly, channels promote **communication through data flow**, which helps reduce race conditions and makes concurrent code easier to understand and maintain.

> 🧠 "Do not communicate by sharing memory; instead, share memory by communicating."

---

## 🎯 Purpose of Channels

Channels are mainly used to:

- Transfer data between concurrent tasks  
- Synchronize execution  
- Avoid explicit locking (mutexes)  
- Implement worker pools and pipelines  
- Manage fan-in and fan-out patterns  

---

## 🧩 Basic Concept

A channel has:
- A **type** (the kind of data it carries)
- A **direction** (send, receive, or both)

Example (Go):

```go
ch := make(chan int)
