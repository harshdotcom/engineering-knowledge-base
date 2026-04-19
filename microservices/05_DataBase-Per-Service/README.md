# Database Per Service Pattern

## TL;DR

> Each microservice owns its own database. No service is allowed to directly touch another service's database. If services need to share data, they talk to each other via APIs or events.

---

## The Problem — Why Not Share One Database?

Imagine a company with 3 teams: **Orders**, **Users**, and **Inventory**. If all three services share one big database:

- The Orders team changes a table → breaks the Inventory service
- One slow query from Users → slows down the entire system
- You can't deploy or scale services independently

This is called **tight coupling** — everything depends on everything else. It's a nightmare.

```
BAD — Shared Database (Tightly Coupled)
=========================================

  [Orders Service]  ──┐
  [Users Service]   ──┼──► [One Big Database]
  [Inventory Service] ┘

Any change in the DB can break ALL services at once.
```

---

## The Solution — Database Per Service (Loosely Coupled)

Give each service its own private database. No service can peek into another service's DB directly.

```
GOOD — Database Per Service (Loosely Coupled)
===============================================

  [Orders Service]    ──► [Orders DB]
  [Users Service]     ──► [Users DB]
  [Inventory Service] ──► [Inventory DB]

Each service is fully independent. Changes in one don't break others.
```

### Why is this good?

| Benefit | Explanation |
|---|---|
| Independent deployments | You can update the Orders service without touching Users or Inventory |
| Independent scaling | Orders getting a lot of traffic? Scale only that service + its DB |
| Technology freedom | Orders can use PostgreSQL, Users can use MongoDB — whatever fits best |
| Fault isolation | One DB crashes? Only that service goes down, not the whole system |

---

## The Catch — Data Consistency is Hard

Here's the tricky part. What if placing an order requires:

1. Creating a record in **Orders DB**
2. Deducting stock from **Inventory DB**
3. Charging the user in **Payments DB**

In a single database, you'd wrap this in a **transaction** — either all steps succeed, or none do. Simple.

With 3 separate databases, you **cannot** do a single transaction across them. This is called a **distributed transaction**, and it's really complex.

```
The Distributed Transaction Problem
=====================================

  Step 1: Orders DB  ✅ (order created)
  Step 2: Inventory DB ✅ (stock deducted)
  Step 3: Payments DB ❌ (payment failed!)

  Now what? Orders DB and Inventory DB already changed.
  You need to UNDO those changes manually. Very hard.
```

---

## The Fix — Event-Driven Architecture

Instead of trying to do everything in one shot, services **publish events** and other services **react** to them.

Think of it like leaving sticky notes:

> "Hey, an order was placed! Anyone who cares, please react."

```
Event-Driven Flow
==================

  [Order Service]
      │
      │  publishes event: "OrderPlaced"
      ▼
  [Message Broker]  ◄── (e.g., Kafka, RabbitMQ)
      │
      ├──► [Inventory Service]  → deducts stock
      └──► [Payment Service]    → charges customer

Each service does its part independently.
If one fails, it can retry later from the event log.
```

### How does this solve the consistency problem?

- **Eventual consistency**: Data won't be consistent *instantly*, but it will become consistent *eventually*
- **Retry on failure**: If Payments fails, the event is still in the queue — it retries automatically
- **No tight coupling**: Order Service doesn't need to know that Inventory or Payments exist

---

## Saga Pattern (Bonus Concept)

When a business operation spans multiple services, you use a **Saga** — a sequence of local transactions, each publishing an event that triggers the next step.

```
Saga: Place Order
==================

  1. Orders Service    → creates order      → emits "OrderCreated"
        ↓
  2. Inventory Service → reserves stock     → emits "StockReserved"
        ↓
  3. Payment Service   → charges customer   → emits "PaymentDone"
        ↓
  4. Orders Service    → marks order "Confirmed"

If Step 3 fails:
  Payment Service → emits "PaymentFailed"
  Inventory Service → listens → releases reserved stock  (compensating transaction)
  Orders Service → marks order "Cancelled"
```

Each failure triggers a **compensating transaction** to undo previous steps — like an undo button for each step.

---

## Summary

```
Key Rules of Database Per Service
===================================

  Rule 1: Each service owns its own DB — no sharing
  Rule 2: Services communicate via APIs or events, not direct DB access
  Rule 3: Embrace eventual consistency — not everything needs to be instant
  Rule 4: Use event-driven patterns (like Saga) for multi-service operations
```

| Concept | Simple Version |
|---|---|
| Database per service | Every service has its own private database |
| Loose coupling | Services don't depend on each other's internals |
| Distributed transaction | A transaction that spans multiple databases — very hard |
| Event-driven architecture | Services talk by publishing/listening to events |
| Eventual consistency | Data will be correct *eventually*, not necessarily right now |
| Saga pattern | A chain of events and compensating actions to handle failures |

---

## Further Reading

- [03 — Saga Pattern](../03_Saga-Pattern/README.md)
- [04 — Monolith vs Microservices](../04_MonolithVsMicroservices/README.md)
