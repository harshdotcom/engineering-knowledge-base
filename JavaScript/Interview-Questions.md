# JavaScript

## 1. Explain the concept of "hoisting" in JavaScript

Hoisting is JavaScript’s behavior of moving declarations to the top of their scope before code execution.

| Keyword | Hoisted | Initialized       | Can use before declaration |
| ------- | ------- | ----------------- | -------------------------- |
| `var`   | Yes     | Yes (`undefined`) | ✅ Yes                      |
| `let`   | Yes     | No (TDZ)          | ❌ No                       |
| `const` | Yes     | No (TDZ)          | ❌ No                       |


## 2. Difference between `var`, `let`, and `const`

### 1. Hoisting

- `var` is hoisted and initialized as `undefined`.
- `let` and `const` are hoisted but remain in the **Temporal Dead Zone (TDZ)** until they are declared.

---

### 2. Scope

- `var` is **function-scoped**.
- `let` and `const` are **block-scoped**.

---

### 3. Reassignment

- `var` can be reassigned.
- `let` can be reassigned.
- `const` cannot be reassigned.

---

### 4. Initialization

- `var` and `let` can be declared without initialization.
- `const` must be initialized at the time of declaration.

---

### Summary Table

| Feature        | var               | let               | const               |
|----------------|-------------------|-------------------|---------------------|
| Hoisting       | Yes (initialized as `undefined`) | Yes (TDZ applies) | Yes (TDZ applies)   |
| Scope          | Function scope    | Block scope       | Block scope         |
| Reassignment   | Allowed           | Allowed           | Not allowed         |
| Initialization | Optional          | Optional          | Required            |

## 3. What is the difference between == and === in JavaScript?

== (loose equality) compares values after performing type conversion.

=== (strict equality) compares both value and data type without type conversion.

## 4. What is the event loop in JavaScript runtimes?

The **event loop** is a mechanism in the JavaScript runtime environment that handles the execution of synchronous and asynchronous operations without blocking the main thread.

---

### How the Event Loop Works

1. The JavaScript engine starts executing the script and places all **synchronous operations** on the **call stack**.

2. When an **asynchronous operation** is encountered (such as `setTimeout()` or an HTTP request), it is offloaded to the appropriate **Web API** (in browsers) or **Node.js API** (in Node.js) to run in the background.

3. Once the asynchronous operation completes, its callback function is placed into one of the following queues:
   - **Microtask Queue**
   - **Macrotask Queue** (also called the Task Queue or Callback Queue)

---

### Queue Types

#### Microtask Queue (Higher Priority)
Includes:
- Promise callbacks: `then`, `catch`, `finally`
- `MutationObserver` callbacks
- `queueMicrotask()`

#### Macrotask Queue (Lower Priority)
Includes:
- `setTimeout()`, `setInterval()`
- HTTP/network callbacks
- UI events (click, scroll, etc.)

---

### Execution Order

4. The event loop continuously checks the **call stack**.

5. When the call stack becomes empty:
   - First, the **microtask queue** is processed.
   - Each microtask is pushed to the call stack and executed.
   - This continues until the microtask queue is empty.

6. After the microtask queue is empty:
   - The event loop processes **one macrotask** from the macrotask queue and pushes it to the call stack.

7. After executing that macrotask:
   - The event loop again checks the **microtask queue**.
   - If microtasks exist, they are executed before the next macrotask.

8. This cycle continues indefinitely, allowing JavaScript to efficiently handle asynchronous operations without blocking the main thread.

---

## 5. Explain event delegation in JavaScript

**Event delegation** is a technique in JavaScript where a single event listener is attached to a parent element instead of attaching separate listeners to multiple child elements.  
When an event occurs on a child element, it bubbles up through the DOM, and the parent element’s event listener handles the event by checking the event target.

---

### How Event Delegation Works

1. An event listener is added to a parent element.
2. A user interacts with one of the child elements.
3. The event bubbles up to the parent element.
4. The parent element determines which child triggered the event using `event.target`.

---

### Benefits of Event Delegation

#### 1. Improved Performance
- Only one event listener is attached instead of many.
- Reduces memory usage.
- Efficient for large lists or tables.

#### 2. Simplified Event Handling
- Event handling logic is written once in the parent.
- Code becomes cleaner and easier to maintain.

#### 3. Dynamic Element Support
- Works automatically for elements added or removed dynamically.
- No need to reattach event listeners when the DOM changes.

---

### Limitations

- The target element must be correctly identified using `event.target`.
- Not all events can be delegated because some do not bubble.

#### Non-bubbling events include:
- `focus`
- `blur`
- `scroll`
- `mouseenter`
- `mouseleave`
- `resize`

---

### Example

```html
<ul id="list">
  <li>Item 1</li>
  <li>Item 2</li>
  <li>Item 3</li>
</ul>




