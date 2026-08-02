# engineering-knowledge-base
A structured, evolving knowledge base documenting practical engineering learnings, with a primary focus on Go (Golang), backend systems, and system design. This repository captures concepts, trade-offs, applied examples, and references derived from hands-on work and continuous study.

```
engineering-knowledge-base
├─ 00_Go-Course
│  ├─ 01_(11.5)_Master the Go Programming Language
│  │  ├─ 01
│  │  │  ├─ 02_doctor
│  │  │  │  └─ doctor.go
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 02_variable
│  │  │  └─ main.go
│  │  ├─ 03_Scope
│  │  │  ├─ go.mod
│  │  │  ├─ main.go
│  │  │  └─ packageone
│  │  │     └─ package.go
│  │  ├─ 04_ConsoleApp
│  │  │  ├─ go.mod
│  │  │  ├─ go.sum
│  │  │  └─ main.go
│  │  ├─ 05_Hammer-Bitcoin
│  │  │  ├─ game
│  │  │  │  └─ bitcoin-miner.go
│  │  │  ├─ go.mod
│  │  │  ├─ go.sum
│  │  │  └─ main.go
│  │  ├─ 06_String-Interpolation
│  │  │  ├─ go.mod
│  │  │  ├─ go.sum
│  │  │  └─ main.go
│  │  ├─ 07_Built-In-Types-Working
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 08_Channels
│  │  │  ├─ go.mod
│  │  │  ├─ go.sum
│  │  │  └─ main.go
│  │  ├─ 09_Interface
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 10_Expression
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 11_Composition
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 12_Exported
│  │  │  ├─ go.mod
│  │  │  ├─ main.go
│  │  │  └─ staff
│  │  │     └─ staff.go
│  │  ├─ 13_Three-Part-Loop
│  │  │  ├─ go.mod
│  │  │  ├─ main.go
│  │  │  └─ mylogger
│  │  │     └─ mylogger.go
│  │  ├─ 14_NestedLoop-Debugger
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 15_If
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 16_Select
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 17_Operators
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 18_RPS
│  │  │  ├─ go.mod
│  │  │  ├─ index.html
│  │  │  └─ main.go
│  │  └─ README.md
│  ├─ 02_(10)_Working with Concurrency in Go
│  │  ├─ 01_First-Example
│  │  │  ├─ go.mod
│  │  │  ├─ main.go
│  │  │  └─ main_test.go
│  │  ├─ 02_Mutex
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 03_More-Mutex
│  │  │  ├─ go.mod
│  │  │  ├─ main.go
│  │  │  └─ main_test.go
│  │  ├─ 03_Producer-Cosumer
│  │  │  ├─ go.mod
│  │  │  ├─ go.sum
│  │  │  └─ main.go
│  │  ├─ 04_Dinning-Philospher_Problem
│  │  │  ├─ go.mod
│  │  │  ├─ go.sum
│  │  │  └─ main.go
│  │  ├─ 05_Channels
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 06_Select
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 07_Buffered_Channels
│  │  │  ├─ go.mod
│  │  │  └─ main.go
│  │  ├─ 08_Sleeping-Barber
│  │  │  ├─ barbershop.go
│  │  │  ├─ go.mod
│  │  │  ├─ go.sum
│  │  │  └─ main.go
│  │  ├─ 09_FInal_Project
│  │  │  ├─ .gocache
│  │  │  │  ├─ 00
│  │  │  │  ├─ 01
│  │  │  │  ├─ 02
│  │  │  │  ├─ 03
│  │  │  │  ├─ 04
│  │  │  │  ├─ 05
│  │  │  │  ├─ 06
│  │  │  │  ├─ 07
│  │  │  │  ├─ 08
│  │  │  │  ├─ 09
│  │  │  │  ├─ 0a
│  │  │  │  ├─ 0b
│  │  │  │  ├─ 0c
│  │  │  │  ├─ 0d
│  │  │  │  ├─ 0e
│  │  │  │  ├─ 0f
│  │  │  │  ├─ 10
│  │  │  │  ├─ 11
│  │  │  │  ├─ 12
│  │  │  │  ├─ 13
│  │  │  │  ├─ 14
│  │  │  │  ├─ 15
│  │  │  │  ├─ 16
│  │  │  │  ├─ 17
│  │  │  │  ├─ 18
│  │  │  │  ├─ 19
│  │  │  │  ├─ 1a
│  │  │  │  ├─ 1b
│  │  │  │  ├─ 1c
│  │  │  │  ├─ 1d
│  │  │  │  ├─ 1e
│  │  │  │  ├─ 1f
│  │  │  │  ├─ 20
│  │  │  │  ├─ 21
│  │  │  │  ├─ 22
│  │  │  │  ├─ 23
│  │  │  │  ├─ 24
│  │  │  │  ├─ 25
│  │  │  │  ├─ 26
│  │  │  │  ├─ 27
│  │  │  │  ├─ 28
│  │  │  │  ├─ 29
│  │  │  │  ├─ 2a
│  │  │  │  ├─ 2b
│  │  │  │  ├─ 2c
│  │  │  │  ├─ 2d
│  │  │  │  ├─ 2e
│  │  │  │  ├─ 2f
│  │  │  │  ├─ 30
│  │  │  │  ├─ 31
│  │  │  │  ├─ 32
│  │  │  │  ├─ 33
│  │  │  │  ├─ 34
│  │  │  │  ├─ 35
│  │  │  │  ├─ 36
│  │  │  │  ├─ 37
│  │  │  │  ├─ 38
│  │  │  │  ├─ 39
│  │  │  │  ├─ 3a
│  │  │  │  ├─ 3b
│  │  │  │  ├─ 3c
│  │  │  │  ├─ 3d
│  │  │  │  ├─ 3e
│  │  │  │  ├─ 3f
│  │  │  │  ├─ 40
│  │  │  │  ├─ 41
│  │  │  │  ├─ 42
│  │  │  │  ├─ 43
│  │  │  │  ├─ 44
│  │  │  │  ├─ 45
│  │  │  │  ├─ 46
│  │  │  │  ├─ 47
│  │  │  │  ├─ 48
│  │  │  │  ├─ 49
│  │  │  │  ├─ 4a
│  │  │  │  ├─ 4b
│  │  │  │  ├─ 4c
│  │  │  │  ├─ 4d
│  │  │  │  ├─ 4e
│  │  │  │  ├─ 4f
│  │  │  │  ├─ 50
│  │  │  │  ├─ 51
│  │  │  │  ├─ 52
│  │  │  │  ├─ 53
│  │  │  │  ├─ 54
│  │  │  │  ├─ 55
│  │  │  │  ├─ 56
│  │  │  │  ├─ 57
│  │  │  │  ├─ 58
│  │  │  │  ├─ 59
│  │  │  │  ├─ 5a
│  │  │  │  ├─ 5b
│  │  │  │  ├─ 5c
│  │  │  │  ├─ 5d
│  │  │  │  ├─ 5e
│  │  │  │  ├─ 5f
│  │  │  │  ├─ 60
│  │  │  │  ├─ 61
│  │  │  │  ├─ 62
│  │  │  │  ├─ 63
│  │  │  │  ├─ 64
│  │  │  │  ├─ 65
│  │  │  │  ├─ 66
│  │  │  │  ├─ 67
│  │  │  │  ├─ 68
│  │  │  │  ├─ 69
│  │  │  │  ├─ 6a
│  │  │  │  ├─ 6b
│  │  │  │  ├─ 6c
│  │  │  │  ├─ 6d
│  │  │  │  ├─ 6e
│  │  │  │  ├─ 6f
│  │  │  │  ├─ 70
│  │  │  │  ├─ 71
│  │  │  │  ├─ 72
│  │  │  │  ├─ 73
│  │  │  │  ├─ 74
│  │  │  │  ├─ 75
│  │  │  │  ├─ 76
│  │  │  │  ├─ 77
│  │  │  │  ├─ 78
│  │  │  │  ├─ 79
│  │  │  │  ├─ 7a
│  │  │  │  ├─ 7b
│  │  │  │  ├─ 7c
│  │  │  │  ├─ 7d
│  │  │  │  ├─ 7e
│  │  │  │  ├─ 7f
│  │  │  │  ├─ 80
│  │  │  │  ├─ 81
│  │  │  │  ├─ 82
│  │  │  │  ├─ 83
│  │  │  │  ├─ 84
│  │  │  │  ├─ 85
│  │  │  │  ├─ 86
│  │  │  │  ├─ 87
│  │  │  │  ├─ 88
│  │  │  │  ├─ 89
│  │  │  │  ├─ 8a
│  │  │  │  ├─ 8b
│  │  │  │  ├─ 8c
│  │  │  │  ├─ 8d
│  │  │  │  ├─ 8e
│  │  │  │  ├─ 8f
│  │  │  │  ├─ 90
│  │  │  │  ├─ 91
│  │  │  │  ├─ 92
│  │  │  │  ├─ 93
│  │  │  │  ├─ 94
│  │  │  │  ├─ 95
│  │  │  │  ├─ 96
│  │  │  │  ├─ 97
│  │  │  │  ├─ 98
│  │  │  │  ├─ 99
│  │  │  │  ├─ 9a
│  │  │  │  ├─ 9b
│  │  │  │  ├─ 9c
│  │  │  │  ├─ 9d
│  │  │  │  ├─ 9e
│  │  │  │  ├─ 9f
│  │  │  │  ├─ a0
│  │  │  │  ├─ a1
│  │  │  │  ├─ a2
│  │  │  │  ├─ a3
│  │  │  │  ├─ a4
│  │  │  │  ├─ a5
│  │  │  │  ├─ a6
│  │  │  │  ├─ a7
│  │  │  │  ├─ a8
│  │  │  │  ├─ a9
│  │  │  │  ├─ aa
│  │  │  │  ├─ ab
│  │  │  │  ├─ ac
│  │  │  │  ├─ ad
│  │  │  │  ├─ ae
│  │  │  │  ├─ af
│  │  │  │  ├─ b0
│  │  │  │  ├─ b1
│  │  │  │  ├─ b2
│  │  │  │  ├─ b3
│  │  │  │  ├─ b4
│  │  │  │  ├─ b5
│  │  │  │  ├─ b6
│  │  │  │  ├─ b7
│  │  │  │  ├─ b8
│  │  │  │  ├─ b9
│  │  │  │  ├─ ba
│  │  │  │  ├─ bb
│  │  │  │  ├─ bc
│  │  │  │  ├─ bd
│  │  │  │  ├─ be
│  │  │  │  ├─ bf
│  │  │  │  ├─ c0
│  │  │  │  ├─ c1
│  │  │  │  ├─ c2
│  │  │  │  ├─ c3
│  │  │  │  ├─ c4
│  │  │  │  ├─ c5
│  │  │  │  ├─ c6
│  │  │  │  ├─ c7
│  │  │  │  ├─ c8
│  │  │  │  ├─ c9
│  │  │  │  ├─ ca
│  │  │  │  ├─ cb
│  │  │  │  ├─ cc
│  │  │  │  ├─ cd
│  │  │  │  ├─ ce
│  │  │  │  ├─ cf
│  │  │  │  ├─ d0
│  │  │  │  ├─ d1
│  │  │  │  ├─ d2
│  │  │  │  ├─ d3
│  │  │  │  ├─ d4
│  │  │  │  ├─ d5
│  │  │  │  ├─ d6
│  │  │  │  ├─ d7
│  │  │  │  ├─ d8
│  │  │  │  ├─ d9
│  │  │  │  ├─ da
│  │  │  │  ├─ db
│  │  │  │  ├─ dc
│  │  │  │  ├─ dd
│  │  │  │  ├─ de
│  │  │  │  ├─ df
│  │  │  │  ├─ e0
│  │  │  │  ├─ e1
│  │  │  │  ├─ e2
│  │  │  │  ├─ e3
│  │  │  │  ├─ e4
│  │  │  │  ├─ e5
│  │  │  │  ├─ e6
│  │  │  │  ├─ e7
│  │  │  │  ├─ e8
│  │  │  │  ├─ e9
│  │  │  │  ├─ ea
│  │  │  │  ├─ eb
│  │  │  │  ├─ ec
│  │  │  │  ├─ ed
│  │  │  │  ├─ ee
│  │  │  │  ├─ ef
│  │  │  │  ├─ f0
│  │  │  │  ├─ f1
│  │  │  │  ├─ f2
│  │  │  │  ├─ f3
│  │  │  │  ├─ f4
│  │  │  │  ├─ f5
│  │  │  │  ├─ f6
│  │  │  │  ├─ f7
│  │  │  │  ├─ f8
│  │  │  │  ├─ f9
│  │  │  │  ├─ fa
│  │  │  │  ├─ fb
│  │  │  │  ├─ fc
│  │  │  │  ├─ fd
│  │  │  │  ├─ fe
│  │  │  │  ├─ ff
│  │  │  │  └─ README
│  │  │  ├─ cmd
│  │  │  │  └─ web
│  │  │  │     ├─ config.go
│  │  │  │     ├─ handlers.go
│  │  │  │     ├─ main.go
│  │  │  │     ├─ middleware.go
│  │  │  │     ├─ render.go
│  │  │  │     ├─ routes.go
│  │  │  │     └─ templates
│  │  │  │        ├─ alerts.partial.gohtml
│  │  │  │        ├─ base.layout.gohtml
│  │  │  │        ├─ footer.partial.gohtml
│  │  │  │        ├─ header.partial.gohtml
│  │  │  │        ├─ home.page.gohtml
│  │  │  │        ├─ login.page.gohtml
│  │  │  │        ├─ mail.html.gohtml
│  │  │  │        ├─ mail.plain.gohtml
│  │  │  │        ├─ navbar.partial.gohtml
│  │  │  │        └─ register.page.gohtml
│  │  │  ├─ docker-compose.yml
│  │  │  ├─ go.mod
│  │  │  ├─ go.sum
│  │  │  └─ Makefile
│  │  └─ README.md
│  ├─ 03_(12)_Working with Microservices in Go
│  │  ├─ image-1.png
│  │  ├─ image.png
│  │  └─ README.md
│  ├─ 04_(11)_Working with Design Patterns in Go
│  ├─ 05_(13)_Introduction to Testing in Go
│  ├─ 06_(40)_Introduction to AI and Machine Learning with Go (Golang)
│  │  └─ 01_AI-Search
│  │     ├─ main.go
│  │     ├─ maze-100-steps.txt
│  │     ├─ maze.txt
│  │     └─ maze2.txt
│  ├─ 07_(30)_Building Modern Web Applications with Go
│  ├─ 08_(20)_Building Web Applications with Go – Intermediate
│  ├─ 09_(15)_Working with WebSockets in Go
│  ├─ 10_(5)_Building a Module in Go
│  ├─ image.png
│  └─ README.md
├─ AI&ML
│  ├─ 00_Python
│  │  └─ main.py
│  └─ neural-network.md
├─ applied-examples
│  └─ README.md
├─ architecture
│  ├─ cqrs.md
│  ├─ event-driven-architecture.md
│  ├─ monolith-vs-microservices.md
│  └─ saga.md
├─ backend
│  ├─ api-design.md
│  ├─ caching.md
│  ├─ database-design.md
│  ├─ messaging.md
│  └─ observability.md
├─ DSA
│  └─ Graph.md
├─ engineering-practices
│  ├─ code-reviews.md
│  ├─ debugging.md
│  ├─ decision-making.md
│  ├─ incident-response.md
│  └─ mentoring.md
├─ go
│  ├─ Arrays.md
│  ├─ channels.md
│  ├─ concurrency.md
│  ├─ context.md
│  ├─ error-handling.md
│  ├─ LLD
│  │  ├─ 01_FizzBuzz
│  │  │  └─ main.go
│  │  ├─ 02_Reverse_String
│  │  │  └─ main.go
│  │  ├─ 03_Check-Palindrome
│  │  │  └─ main.go
│  │  ├─ 04_Generate_first_n_Fibonacci_numbers
│  │  │  └─ main.go
│  │  ├─ 05_Factorial_of_n
│  │  │  └─ main.go
│  │  ├─ 06_Count_Vowels
│  │  │  └─ main.go
│  │  ├─ 07_CheckPrime
│  │  │  └─ main.go
│  │  ├─ 08_Find_Max_in_Slice
│  │  │  └─ main.go
│  │  ├─ 09_Sum_of_Slice
│  │  │  └─ main.go
│  │  ├─ 10_Remove_Duplicates
│  │  │  └─ main.go
│  │  ├─ 11_Word_Count
│  │  │  └─ main.go
│  │  ├─ 12_Two_Sum
│  │  │  └─ main.go
│  │  ├─ 13_Caesar_Cipher
│  │  │  └─ main.go
│  │  ├─ 14_Stack
│  │  │  └─ main.go
│  │  └─ 15_Queue
│  │     └─ main.go
│  ├─ memory.md
│  ├─ performance.md
│  ├─ pointer.md
│  ├─ Practice
│  │  ├─ 01_Hello-World
│  │  │  └─ main.go
│  │  ├─ 02_Time
│  │  │  └─ main.go
│  │  ├─ 03_Math
│  │  │  └─ main.go
│  │  ├─ 04_Flow-Control
│  │  │  └─ main.go
│  │  └─ 05_struct-slice-map
│  │     ├─ cmd
│  │     │  └─ main.go
│  │     └─ main.go
│  ├─ Slices.md
│  └─ struct.md
├─ go.mod
├─ go.sum
├─ JavaScript
│  └─ Interview-Questions.md
├─ LICENSE
├─ microservices
│  ├─ 03_Saga-Pattern
│  │  ├─ 03_Saga-Pattern.md
│  │  ├─ image-1.png
│  │  ├─ image-2.png
│  │  ├─ image-3.png
│  │  ├─ image-4.png
│  │  ├─ image-5.png
│  │  ├─ image-6.png
│  │  └─ image.png
│  ├─ 04_MonolithVsMicroservices
│  │  ├─ image-1.png
│  │  ├─ image-2.png
│  │  ├─ image-3.png
│  │  ├─ image-4.png
│  │  ├─ image-5.png
│  │  ├─ image.png
│  │  └─ README.MD
│  └─ 05_DataBase-Per-Service
│     ├─ image-1.png
│     ├─ image-2.png
│     ├─ image-3.png
│     ├─ image-4.png
│     ├─ image.png
│     └─ README.md
├─ MLOps
│  └─ 01_Python
│     ├─ 01
│     │  └─ 01_Python_Basics
│     ├─ 01_Helloworld.py
│     ├─ 02_Variable.py
│     └─ 03_Input.py
├─ Python
│  ├─ 00_Syntex
│  │  ├─ main.py
│  │  └─ Readme.md
│  ├─ 01_Day
│  │  ├─ main.py
│  │  └─ main2.py
│  ├─ 02_FastAPI
│  │  ├─ .fastapicloud
│  │  │  ├─ cloud.json
│  │  │  └─ README.md
│  │  ├─ LEARNING.md
│  │  ├─ main.py
│  │  └─ README.md
│  └─ 03_LLM
│     ├─ main.py
│     └─ requirements.txt
├─ README.md
├─ references
│  ├─ checklists.md
│  ├─ commands.md
│  └─ diagrams.md
├─ resources
│  ├─ articles.md
│  ├─ books.md
│  ├─ courses.md
│  └─ talks.md
├─ system-design
│  ├─ 01_User-Crud.md
│  ├─ capacity-planning.md
│  ├─ consistency.md
│  ├─ load-balancing.md
│  ├─ rate-limiter.md
│  └─ sharding.md
├─ templates
│  ├─ applied-example.md
│  ├─ concept-note.md
│  └─ system-design-note.md
└─ tooling
   ├─ docker.md
   ├─ github-actions.md
   └─ observability-tools.md

```