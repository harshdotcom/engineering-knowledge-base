# Monolith vs Microservices

---

## TL;DR

> **Monolith** = One giant app that does everything together. Simple to start, hard to scale.
> **Microservices** = Many small apps, each doing one thing well. Complex to manage, but scales like a beast.
> **When to switch?** When your monolith becomes painful — slow deploys, one bug crashing everything, teams blocking each other.
> **How to switch?** Gradually, using the **Strangler Fig Pattern** — peel off one service at a time, never rewrite everything at once.

---

## The Pizza Shop Analogy (Yes, Really)

Imagine you run a pizza shop. At first, **you do everything** — take orders, make pizza, handle delivery, manage billing. That's a **Monolith**.

Now your shop blows up. You have 500 orders a day. You can't do it all. So you **hire specialists** — one guy just takes orders, another just makes pizza, a delivery team handles only delivery. Each team works independently. That's **Microservices**.

---

## What is a Monolith?

A monolith is a single application where **all features live in one codebase and deploy together**.

```
[ One Big App ]
    ├── User Registration
    ├── Search & Filters
    ├── Cart & Wishlist
    ├── Checkout & Payment
    ├── Order Tracking
    ├── Push Notifications
    ├── Reviews & Ratings
    └── Loyalty Programs
```

### Why Monoliths Are Great (at first)

- **Simple to build** — one repo, one deploy, one database
- **Easy to debug** — everything is in one place
- **Fast to get started** — no distributed system headaches
- **Great for small teams** — no coordination overhead

### Why Monoliths Become Painful (as you grow)

| Problem | What Happens |
|---|---|
| One bug, whole app crashes | Payment bug? Users can't even log in |
| Slow deployments | Change 1 line? Redeploy the entire app |
| Hard to scale | Need more search power? You must scale the WHOLE app |
| Team conflicts | 10 devs editing the same codebase = merge hell |
| Tech lock-in | Started with PHP? Stuck with PHP forever |

---

## What are Microservices?

Each feature becomes its own **independent service** with its own:
- Codebase
- Database
- Deployment pipeline
- Team ownership

```
[ API Gateway ]
      |
  ┌───┼───┬──────┬────────┬──────────┐
  │   │   │      │        │          │
User Search Cart Payment Notifications Orders
 Svc  Svc  Svc   Svc      Svc         Svc
  │   │   │      │        │          │
 DB  DB  DB     DB       DB          DB
```

### Why Microservices Are Great

- **Independent deploys** — fix payment without touching search
- **Independent scaling** — Black Friday? Scale only the checkout service
- **Team autonomy** — each team owns their service end-to-end
- **Tech freedom** — Search service in Go, ML service in Python, no problem
- **Fault isolation** — Notifications crash? Users can still buy stuff

### Why Microservices Are Hard

| Problem | What Happens |
|---|---|
| Network calls everywhere | Services talk over HTTP/gRPC — latency, failures |
| Distributed tracing | A bug hops across 5 services — good luck debugging |
| Data consistency | No shared DB means keeping data in sync is hard |
| Operational overhead | 20 services = 20 deployments, logs, alerts to manage |
| Testing complexity | Integration tests become massive |
| Service discovery | How does Service A even find Service B? |

---

## Real World: How the Giants Did It

### Amazon

Amazon started as a **monolith** in 1995. By early 2000s, it was a nightmare — slow builds, massive team conflicts, one team's bad deploy would take down the whole site.

Jeff Bezos issued the famous **"API Mandate"**:
> *Every team must expose their data via APIs. No backdoors. No exceptions.*

This forced them to think in services. They **gradually** migrated over years — not in one big rewrite.

Today Amazon runs **thousands of microservices**.

### Netflix

Netflix hit a massive **database corruption** in 2008 that took them down for 3 days. That was the wakeup call.

They migrated to microservices over **7 years** (2008–2016). Key moves:
- Moved to AWS cloud
- Split into hundreds of services
- Built tools like **Eureka** (service discovery) and **Hystrix** (fault tolerance) — which they later open-sourced

Today Netflix runs **700+ microservices**.

---

## The Strangler Fig Pattern — How to Actually Migrate

Named after the Strangler Fig tree — it wraps around an old tree, slowly takes over, and eventually the old tree dies. The new one stands alone.

```
Phase 1:           Phase 2:           Phase 3:
                                      
[ Monolith ]    [ Monolith ]      [  Empty  ]
    │           ┌────┴────┐             │
    │           │   New   │           [ New  ]
    │           │ Service │           [Services]
    │           └─────────┘
```

### Steps to Migrate

1. **Don't rewrite everything** — that's how projects die
2. Pick the **most painful or most independent** piece first (e.g., notifications, user auth)
3. Build the new service **alongside** the monolith
4. Route traffic to the new service gradually (using an API Gateway or feature flags)
5. Once stable, **delete the code from the monolith**
6. Repeat for the next service

### What to Extract First (Good Candidates)

- Services that **scale differently** (e.g., search needs more CPU)
- Services with **independent teams**
- Services that **change frequently** (lots of deploys)
- Services that **crash other things** when they fail

### What to Keep in Monolith Longer (Bad Candidates)

- Tightly coupled features that share the same database transactions
- Features that are rarely changed
- Features with no clear ownership

---

## When to Use What

### Stick with Monolith When:
- You're a **startup or small team** (< 10 engineers)
- The product is **still finding product-market fit**
- You need to **move fast** and the overhead isn't worth it
- The app is **simple** and won't grow massively

### Move to Microservices When:
- **Teams are blocking each other** constantly
- **Deployments take forever** and are risky
- You need to **scale specific parts** of the app independently
- The codebase is so large that **no one understands the whole thing**
- You have **different scaling/tech requirements** per feature

---

## Tools You'll Need for Microservices

| Category | Tools |
|---|---|
| Container runtime | Docker |
| Orchestration | Kubernetes, ECS |
| API Gateway | Kong, AWS API Gateway, Nginx |
| Service discovery | Consul, Eureka, Kubernetes DNS |
| Message queues | Kafka, RabbitMQ, SQS |
| Distributed tracing | Jaeger, Zipkin, Datadog |
| Monitoring | Prometheus + Grafana, Datadog |
| CI/CD | GitHub Actions, ArgoCD, Jenkins |
| Service mesh | Istio, Linkerd |

---

## Common Mistakes to Avoid

**Starting with microservices** — don't. Start with a modular monolith, migrate later when pain is real.

**Making services too small** — nano-services create more overhead than value. Each service should own a full business domain.

**Shared database** — if two services share a DB, they're not independent. You just have a distributed monolith.

**Synchronous calls everywhere** — long chains of HTTP calls = cascading failures. Use async messaging (Kafka, queues) for non-critical paths.

**No observability** — without distributed tracing and centralized logs, debugging microservices is impossible.

**Big bang migration** — rewriting everything at once almost always fails. Strangle gradually.

---

## Quick Reference Cheat Sheet

```
MONOLITH                        MICROSERVICES
─────────────────────────────   ─────────────────────────────
One codebase                    Many codebases
One DB                          DB per service
One deploy                      Independent deploys
Simple to start                 Complex to operate
Hard to scale parts             Scale parts independently
Teams conflict                  Teams autonomous
One language/framework          Polyglot (any lang per service)
One failure = full outage       Failures are isolated
Easy to debug                   Needs distributed tracing
Great for startups              Great for scale
```

---

## Key Takeaway

> There is no "better" architecture. A **monolith** is not bad — premature microservices are.
> Start simple. Migrate when the pain of NOT migrating is greater than the pain of migrating.
> And when you migrate, **use the Strangler Fig Pattern** — slice off pieces, never blow up and rebuild.
