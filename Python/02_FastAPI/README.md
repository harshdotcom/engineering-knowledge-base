# FastAPI — `main.py` Explained (Beginner Friendly)

This file is a tiny FastAPI application. It shows the **three most common things**
you do with a web API:

1. Return some data (GET)
2. Read a value from the URL (GET with a path parameter)
3. Send data *into* the API (PUT with a request body)

Think of an **API** as a waiter in a restaurant: you (the browser or app) ask for
something, and the API brings it back to you.

---

## The code, line by line

### 1. Imports

```python
from fastapi import FastAPI
from pydantic import BaseModel
```

- `FastAPI` — the framework that turns your Python functions into web endpoints.
- `BaseModel` (from **Pydantic**) — helps you describe the *shape* of data and
  automatically checks it is valid (e.g. that `price` is really a number).

### 2. Create the app

```python
app = FastAPI()
```

`app` is your application. Every route (URL) you create gets attached to it.
FastAPI also uses this to auto-generate documentation for you (more on that below).

### 3. Define a data model

```python
class Item(BaseModel):
    name: str
    price: float
    is_offer: bool | None = None
```

This describes what an "Item" looks like when someone sends one to your API:

| Field      | Type              | Meaning                                             |
|------------|-------------------|-----------------------------------------------------|
| `name`     | `str` (text)      | Required — the item's name                           |
| `price`    | `float` (number)  | Required — the item's price                          |
| `is_offer` | `bool` or nothing | Optional — defaults to `None` if not provided        |

> `bool | None = None` means: *"this can be True/False, or it can be missing.
> If missing, use `None`."*

If someone sends the wrong type (e.g. text instead of a number for `price`),
FastAPI rejects it automatically with a helpful error. You don't write that code.

### 4. The root endpoint (GET `/`)

```python
@app.get("/")
def read_root():
    return {"Hello": "World"}
```

- `@app.get("/")` is a **decorator**. It says: *"when someone visits `/`,
  run the function below."*
- Visiting `http://127.0.0.1:8000/` returns the JSON: `{"Hello": "World"}`.

### 5. Read an item (GET `/items/{item_id}`)

```python
@app.get("/items/{item_id}")
def read_item(item_id: int, q: str | None = None):
    return {"item_id": item_id, "q": q}
```

- `{item_id}` in the URL is a **path parameter** — a value taken from the URL itself.
- `item_id: int` tells FastAPI to convert it to an integer (and reject non-numbers).
- `q: str | None = None` is a **query parameter** — an optional extra you add with `?`.

Examples:

| You visit                         | You get back                          |
|-----------------------------------|---------------------------------------|
| `/items/5`                        | `{"item_id": 5, "q": null}`           |
| `/items/5?q=hello`                | `{"item_id": 5, "q": "hello"}`        |
| `/items/abc`                      | Error — `abc` is not an integer       |

### 6. Update an item (PUT `/items/{item_id}`)

```python
@app.put("/items/{item_id}")
def update_item(item_id: int, item: Item):
    return {"item_name": item.price, "item_id": item_id}
```

- `PUT` is used to **send/update data**. The `item: Item` argument means the request
  must include a JSON body matching the `Item` model above.
- FastAPI reads the JSON body, validates it, and gives you an `item` object.

Example request body:

```json
{
  "name": "Notebook",
  "price": 12.5,
  "is_offer": true
}
```

Response:

```json
{
  "item_name": 12.5,
  "item_id": 3
}
```

> ⚠️ **Small bug to notice (great learning moment):** the key is called
> `"item_name"` but it returns `item.price` (the price, not the name).
> It should probably be `item.name`. This is a common copy-paste slip —
> the code runs fine, but the *meaning* is wrong.

---

## How to run it

From inside the `02_FastAPI` folder:

```bash
# 1. Activate your virtual environment (already created here as "venv")
venv\Scripts\activate          # Windows PowerShell

# 2. Install FastAPI + a server (only needed once)
pip install "fastapi[standard]"

# 3. Start the development server
fastapi dev main.py
```

Then open your browser at:

- **App:** http://127.0.0.1:8000/
- **Interactive docs (Swagger UI):** http://127.0.0.1:8000/docs
- **Alternative docs (ReDoc):** http://127.0.0.1:8000/redoc

The `/docs` page is the best part of FastAPI: it lets you **try every endpoint
from the browser** without writing any extra code.

---

## Key ideas to remember

| Concept            | What it does                                            |
|--------------------|--------------------------------------------------------|
| `@app.get / .put`  | Connects a URL + HTTP method to a Python function      |
| Path parameter     | A value taken from the URL path (`/items/{item_id}`)   |
| Query parameter    | Optional value after `?` (`?q=hello`)                   |
| Request body       | JSON data sent *into* the API, validated by a model    |
| Pydantic `BaseModel` | Describes and validates data shapes automatically    |
| Type hints         | FastAPI uses them to convert, validate, and document   |

---

## Quick glossary

- **Endpoint / Route** — a URL your API responds to (e.g. `/items/5`).
- **HTTP method** — the *kind* of action: `GET` (read), `PUT` (update),
  `POST` (create), `DELETE` (remove).
- **JSON** — the text format APIs use to send data: `{"key": "value"}`.
- **Decorator** — the `@app.get(...)` line that "labels" a function as a route.
