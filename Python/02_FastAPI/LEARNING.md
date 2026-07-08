# FastAPI Flashcards — Learn from Ground Zero 🧠🚀

A self-study Q&A deck. Cover the **answer**, read the **question**, answer in your
head, then check. Cards are grouped into **chunks of 4** and ordered
**Easy → Medium → Hard** so you can learn from absolute zero to advanced.

> Legend: 🟢 Easy · 🟡 Medium · 🔴 Hard
> Tip: Master all 🟢 chunks first, then move to 🟡, then 🔴.

---

# 🟢 PART 1 — EASY (Fundamentals)

## Chunk 1 — What & Why

**Q1. What is FastAPI?**
> A modern, high-performance Python web framework for building APIs, using standard
> Python type hints.

**Q2. Name three key benefits of FastAPI.**
> Fast performance, automatic interactive docs, and data validation from type hints.

**Q3. What is an API in one sentence?**
> A contract that lets two programs talk: one sends a request, the other returns a response.

**Q4. What language feature does FastAPI rely on heavily?**
> Python **type hints** (annotations like `item_id: int`).

## Chunk 2 — Installation & First Run

**Q5. How do you install FastAPI with everything needed?**
> `pip install "fastapi[standard]"`

**Q6. What server actually runs a FastAPI app?**
> An ASGI server, usually **Uvicorn**.

**Q7. How do you start the dev server for `main.py`?**
> `fastapi dev main.py`

**Q8. What is the default local URL after starting?**
> `http://127.0.0.1:8000`

## Chunk 3 — The App Object

**Q9. How do you create a FastAPI application instance?**
> ```python
> from fastapi import FastAPI
> app = FastAPI()
> ```

**Q10. What is `app` used for?**
> It holds your routes and configuration; every endpoint attaches to it.

**Q11. What is a "route" or "endpoint"?**
> A URL + HTTP method your API responds to (e.g. `GET /items`).

**Q12. What is a decorator in FastAPI routing?**
> The `@app.get("/")` line above a function that connects a URL to that function.

## Chunk 4 — Your First Endpoints

**Q13. Write a root endpoint returning `{"hello": "world"}`.**
> ```python
> @app.get("/")
> def read_root():
>     return {"hello": "world"}
> ```

**Q14. What format does FastAPI return a Python dict as?**
> JSON (automatically converted).

**Q15. What does the `/docs` URL show?**
> The interactive Swagger UI documentation.

**Q16. What does `/redoc` show?**
> Alternative auto-generated documentation (ReDoc).

## Chunk 5 — HTTP Methods

**Q17. Which decorator handles a GET request?**
> `@app.get(...)`

**Q18. Which HTTP method creates new data?**
> `POST` → `@app.post(...)`

**Q19. Which HTTP method updates/replaces data?**
> `PUT` → `@app.put(...)`

**Q20. Which HTTP method removes data?**
> `DELETE` → `@app.delete(...)`

## Chunk 6 — Path Parameters

**Q21. What is a path parameter?**
> A value captured from inside the URL path, e.g. `/items/{item_id}`.

**Q22. Write an endpoint reading an integer `item_id` from the path.**
> ```python
> @app.get("/items/{item_id}")
> def read(item_id: int): return {"id": item_id}
> ```

**Q23. What happens if you visit `/items/abc` when `item_id: int`?**
> FastAPI returns a `422` validation error — `abc` is not an integer.

**Q24. Can you have multiple path parameters?**
> Yes, e.g. `/users/{user_id}/items/{item_id}`.

## Chunk 7 — Query Parameters

**Q25. What is a query parameter?**
> A value passed after `?` in the URL, e.g. `/items?q=hello`.

**Q26. How do you make a query parameter optional?**
> Give it a default: `q: str | None = None`.

**Q27. How do you make a query parameter required?**
> Give it no default value: `q: str`.

**Q28. How does FastAPI decide path vs query parameter?**
> If the name is in the route path `{}` → path; otherwise → query.

## Chunk 8 — Request Body Basics

**Q29. What library defines request body models?**
> Pydantic (via `BaseModel`).

**Q30. Define an `Item` model with `name` and `price`.**
> ```python
> from pydantic import BaseModel
> class Item(BaseModel):
>     name: str
>     price: float
> ```

**Q31. How do you receive a body in an endpoint?**
> Add a typed argument: `def create(item: Item): ...`

**Q32. How do you read a field from the body?**
> Dot notation: `item.name`, `item.price`.

## Chunk 9 — Data Types & Validation Basics

**Q33. What does `is_offer: bool | None = None` mean?**
> An optional boolean field that defaults to `None` if omitted.

**Q34. What error code means "your input is invalid"?**
> `422 Unprocessable Entity`.

**Q35. Does FastAPI validate types automatically?**
> Yes — based on your type hints and Pydantic models.

**Q36. What Python type is used for whole numbers vs decimals?**
> `int` for whole numbers, `float` for decimals.

## Chunk 10 — Responses & Status

**Q37. What status code means success by default?**
> `200 OK`.

**Q38. How do you set a custom status code on an endpoint?**
> `@app.post("/items", status_code=201)`

**Q39. What does `201` mean?**
> Created — commonly used after a successful POST.

**Q40. What does `404` mean?**
> Not Found.

## Chunk 11 — Simple Errors

**Q41. How do you raise a 404 error?**
> ```python
> from fastapi import HTTPException
> raise HTTPException(status_code=404, detail="Not found")
> ```

**Q42. What does the `detail` argument do?**
> Provides the error message returned to the client.

**Q43. What status code means "server crashed"?**
> `500 Internal Server Error`.

**Q44. What status code means "you must log in"?**
> `401 Unauthorized`.

## Chunk 12 — Docs & Basics Wrap-up

**Q45. True/False: FastAPI generates docs automatically.**
> True — from your type hints and models.

**Q46. What is JSON?**
> A text data format for APIs: `{"key": "value"}`.

**Q47. What underlies FastAPI's speed?**
> Starlette (web) + Pydantic (data) + async ASGI.

**Q48. What's the difference between FastAPI and Uvicorn?**
> FastAPI = framework you code with; Uvicorn = server that runs it.

---

# 🟡 PART 2 — MEDIUM (Building Real APIs)

## Chunk 13 — Pydantic Deeper

**Q49. What is `Field` used for in Pydantic?**
> Adding validation/metadata: `price: float = Field(gt=0)`.

**Q50. How do you set a maximum length on a string field?**
> `name: str = Field(max_length=50)`.

**Q51. How do you give a field a description for the docs?**
> `Field(description="The item price")`.

**Q52. What Pydantic version does modern FastAPI use?**
> Pydantic **v2** (with `model_config`, `model_dump()`, etc.).

## Chunk 14 — Nested & Complex Models

**Q53. Can a Pydantic model contain another model?**
> Yes — nest them: `owner: User` inside `Item`.

**Q54. How do you define a list field?**
> `tags: list[str] = []`.

**Q55. How do you define a dict field?**
> `meta: dict[str, str] = {}`.

**Q56. How do you validate an email field easily?**
> Use `EmailStr` from `pydantic` (needs `email-validator`).

## Chunk 15 — Response Models

**Q57. What is `response_model` for?**
> Defining/filtering the output shape of an endpoint.

**Q58. How do you hide a password field in responses?**
> Use a separate output model without that field as `response_model`.

**Q59. What does `response_model_exclude_unset=True` do?**
> Omits fields the client didn't explicitly set.

**Q60. Can you return a Pydantic model directly from an endpoint?**
> Yes — FastAPI serializes it to JSON automatically.

## Chunk 16 — Validation Metadata

**Q61. How do you validate a query param with constraints?**
> ```python
> from fastapi import Query
> q: str = Query(min_length=3)
> ```

**Q62. How do you add metadata to a path param?**
> `item_id: int = Path(ge=1)`.

**Q63. What does `ge` and `le` mean in constraints?**
> Greater-than-or-equal and less-than-or-equal.

**Q64. How do you set a default value with `Query`?**
> `q: str = Query(default="all")`.

## Chunk 17 — Forms & Files

**Q65. How do you receive form fields (not JSON)?**
> ```python
> from fastapi import Form
> def login(username: str = Form()): ...
> ```

**Q66. What package is needed for form/file support?**
> `python-multipart`.

**Q67. How do you receive an uploaded file?**
> `from fastapi import UploadFile` → `file: UploadFile`.

**Q68. How do you read an uploaded file's contents?**
> `contents = await file.read()`.

## Chunk 18 — Headers & Cookies

**Q69. How do you read a request header?**
> ```python
> from fastapi import Header
> def h(user_agent: str = Header()): ...
> ```

**Q70. How do you read a cookie?**
> `from fastapi import Cookie` → `session: str = Cookie()`.

**Q71. How do you set a cookie in a response?**
> `response.set_cookie(key="x", value="y")`.

**Q72. How do you set a custom response header?**
> `response.headers["X-Custom"] = "value"`.

## Chunk 19 — Dependency Injection

**Q73. What is dependency injection in FastAPI?**
> Reusing shared logic (DB, auth, params) via `Depends`.

**Q74. Write a simple dependency for pagination.**
> ```python
> def paginate(skip: int = 0, limit: int = 10):
>     return {"skip": skip, "limit": limit}
> ```

**Q75. How do you use a dependency in an endpoint?**
> `def read(p: dict = Depends(paginate)): ...`

**Q76. Can dependencies depend on other dependencies?**
> Yes — sub-dependencies are supported and resolved automatically.

## Chunk 20 — Routers & Structure

**Q77. What is `APIRouter` for?**
> Splitting routes across multiple files/modules.

**Q78. How do you include a router in the app?**
> `app.include_router(router)`.

**Q79. How do you add a URL prefix to a router?**
> `APIRouter(prefix="/users")`.

**Q80. How do you group endpoints in the docs?**
> Use `tags=["users"]`.

## Chunk 21 — Middleware & CORS

**Q81. What is middleware?**
> Code that runs on every request/response (logging, timing, etc.).

**Q82. How do you add custom HTTP middleware?**
> `@app.middleware("http")` on an `async def(request, call_next)`.

**Q83. What is CORS?**
> Cross-Origin Resource Sharing — lets browsers call your API from other domains.

**Q84. How do you enable CORS?**
> `app.add_middleware(CORSMiddleware, allow_origins=[...])`.

## Chunk 22 — Background Tasks & Events

**Q85. How do you run work after returning a response?**
> ```python
> from fastapi import BackgroundTasks
> def send(bg: BackgroundTasks): bg.add_task(fn)
> ```

**Q86. Give a use case for background tasks.**
> Sending emails, writing logs, cleanup — anything not needed for the response.

**Q87. What replaced `@app.on_event("startup")`?**
> The `lifespan` context manager.

**Q88. What is a startup event used for?**
> Setup like opening DB connections or loading models.

## Chunk 23 — Async

**Q89. When should you use `async def`?**
> When calling `await`-able (non-blocking) I/O like async DB/HTTP.

**Q90. What happens if you use plain `def` in FastAPI?**
> It runs in a threadpool — safe for blocking code.

**Q91. What keyword pauses on an async operation?**
> `await`.

**Q92. What library makes async HTTP requests?**
> `httpx` (async-capable).

## Chunk 24 — Settings, Config & Testing

**Q93. How do you manage configuration/settings cleanly?**
> `pydantic-settings` → a `BaseSettings` class reading env vars.

**Q94. What tool tests FastAPI apps?**
> `TestClient` (from `fastapi.testclient`) or `httpx.AsyncClient`.

**Q95. Write a basic test for `GET /`.**
> ```python
> def test_root():
>     r = client.get("/")
>     assert r.status_code == 200
> ```

**Q96. How do you override a dependency in tests?**
> `app.dependency_overrides[real] = fake`.

---

# 🔴 PART 3 — HARD (Advanced & Production)

## Chunk 25 — Security & Auth

**Q97. What scheme reads a Bearer token?**
> `OAuth2PasswordBearer(tokenUrl="token")`.

**Q98. What is JWT?**
> JSON Web Token — a signed, self-contained token for stateless auth.

**Q99. Which library commonly encodes/decodes JWTs?**
> `python-jose` or `pyjwt`.

**Q100. How do you protect an endpoint with a token dependency?**
> `def me(token: str = Depends(oauth2_scheme)): ...`

## Chunk 26 — Scopes & Permissions

**Q101. What are OAuth2 scopes?**
> Named permissions attached to a token (e.g. `items:read`).

**Q102. What class declares required scopes?**
> `Security(dep, scopes=["items:read"])`.

**Q103. How do you hash passwords securely?**
> With `passlib` (e.g. bcrypt) — never store plaintext.

**Q104. Why use short-lived access + refresh tokens?**
> Limits damage if a token leaks; refresh renews access without re-login.

## Chunk 27 — Databases

**Q105. What ORM pairs naturally with FastAPI?**
> SQLAlchemy, or **SQLModel** (by FastAPI's author).

**Q106. How is a DB session usually provided?**
> As a dependency yielding a session with `try/finally` cleanup.

**Q107. Write a DB session dependency skeleton.**
> ```python
> def get_db():
>     db = SessionLocal()
>     try: yield db
>     finally: db.close()
> ```

**Q108. What supports fully async DB access?**
> SQLAlchemy async engine + `asyncpg`/`aiomysql`, or async SQLModel.

## Chunk 28 — Migrations & Modeling

**Q109. What tool handles DB schema migrations?**
> Alembic.

**Q110. Why separate DB models from Pydantic schemas?**
> DB models manage storage; schemas manage API input/output shape.

**Q111. What does `orm_mode`/`from_attributes` enable?**
> Reading data directly from ORM objects into Pydantic models.

**Q112. How do you prevent N+1 query problems?**
> Eager-load relationships (e.g. `selectinload`/`joinedload`).

## Chunk 29 — WebSockets & Realtime

**Q113. How do you declare a WebSocket endpoint?**
> `@app.websocket("/ws")` on `async def(ws: WebSocket)`.

**Q114. How do you accept a WebSocket connection?**
> `await websocket.accept()`.

**Q115. How do you send/receive text over a WebSocket?**
> `await ws.send_text(...)` / `await ws.receive_text()`.

**Q116. Use case for WebSockets?**
> Chat, live notifications, dashboards — persistent two-way communication.

## Chunk 30 — Custom & Streaming Responses

**Q117. How do you return raw HTML?**
> `from fastapi.responses import HTMLResponse`.

**Q118. How do you stream a large response?**
> `StreamingResponse(generator, media_type=...)`.

**Q119. How do you return a file download?**
> `FileResponse("path/to/file")`.

**Q120. How do you return a redirect?**
> `RedirectResponse(url="/new", status_code=307)`.

## Chunk 31 — Error Handling & OpenAPI

**Q121. How do you add a global exception handler?**
> `@app.exception_handler(MyError)` returning a custom response.

**Q122. How do you customize the OpenAPI schema?**
> Override `app.openapi` or pass `title`, `version`, `description` to `FastAPI()`.

**Q123. How do you document possible responses per endpoint?**
> `responses={404: {"description": "Not found"}}` in the decorator.

**Q124. How do you hide an endpoint from the docs?**
> `include_in_schema=False`.

## Chunk 32 — Performance & Deployment

**Q125. How do you run FastAPI in production?**
> Uvicorn workers (often via Gunicorn) behind a reverse proxy like Nginx.

**Q126. Why use multiple worker processes?**
> To use multiple CPU cores and handle more concurrent requests.

**Q127. What's a common way to containerize a FastAPI app?**
> A Docker image with `uvicorn`/`gunicorn` as the entrypoint.

**Q128. Name two ways to boost performance.**
> Use async I/O properly and add caching (e.g. Redis) for hot data.

## Chunk 33 — Advanced Patterns

**Q129. How do you version an API?**
> Route prefixes (`/v1`, `/v2`) or separate routers/sub-apps.

**Q130. How do you mount a sub-application?**
> `app.mount("/sub", other_app)`.

**Q131. How do you add rate limiting?**
> A library like `slowapi`, or a gateway/reverse proxy limit.

**Q132. How do you cache an expensive dependency per request?**
> Dependencies are cached per-request by default (disable with `use_cache=False`).

**Q133. How do you validate data returned to clients?**
> Via `response_model` — FastAPI validates/filters output against it.

**Q134. What's a clean way to structure a large FastAPI project?**
> Split into `routers/`, `models/`, `schemas/`, `services/`, `core/` (config) modules.

---

# 🟢 PART 4 — EASY (More Fundamentals)

## Chunk 34 — Routing Details

**Q135. Does route order matter in FastAPI?**
> Yes — the first matching route wins, so declare specific paths before generic ones.

**Q136. How do you avoid `/users/me` being caught by `/users/{id}`?**
> Declare `/users/me` **before** `/users/{id}`.

**Q137. How do you accept a path param that contains slashes?**
> Use the path convertor: `/files/{file_path:path}`.

**Q138. What's FastAPI's default trailing-slash behavior?**
> It redirects between `/items` and `/items/` (HTTP 307) automatically.

## Chunk 35 — More HTTP Methods

**Q139. Which method does a partial update?**
> `PATCH` → `@app.patch(...)`.

**Q140. What are HEAD and OPTIONS used for?**
> HEAD = headers only (no body); OPTIONS = capability/CORS preflight checks.

**Q141. What does `204` mean?**
> No Content — success with no response body.

**Q142. How do you return a 204 correctly?**
> Set `status_code=204` and return `None`.

## Chunk 36 — Typing Extras

**Q143. How do you restrict a path param to fixed choices?**
> Use an `Enum` subclassing `str`: `class Color(str, Enum): ...`.

**Q144. How do you accept a repeated query param as a list?**
> `q: list[str] = Query()` → `?q=a&q=b`.

**Q145. What query values parse to `True` for a `bool`?**
> `true`, `1`, `on`, `yes` (case-insensitive).

**Q146. Is `Optional[str]` the same as `str | None`?**
> Yes — identical meaning; `| None` is the modern style.

## Chunk 37 — Body Extras

**Q147. What happens with two Pydantic body params?**
> FastAPI nests them by parameter name in the JSON body.

**Q148. How do you accept a single value in the body?**
> `importance: int = Body()`.

**Q149. How do you force a single model to be nested under its key?**
> `item: Item = Body(embed=True)`.

**Q150. How do you make a body optional?**
> Give the argument a default: `item: Item | None = None`.

## Chunk 38 — Docs Metadata

**Q151. How do you add an example to a model in the docs?**
> `model_config = {"json_schema_extra": {"examples": [{...}]}}`.

**Q152. How do you add a summary/description to an endpoint?**
> Pass `summary=`/`description=` to the decorator, or use the function docstring.

**Q153. How do you mark an endpoint deprecated?**
> `@app.get(..., deprecated=True)`.

**Q154. How do you describe tag groups in the docs?**
> Pass `openapi_tags=[{"name": "...", "description": "..."}]` to `FastAPI()`.

## Chunk 39 — Special Types

**Q155. How does FastAPI handle `datetime` fields?**
> It parses ISO 8601 strings and serializes back to ISO strings.

**Q156. What type validates a UUID?**
> `uuid.UUID`.

**Q157. Which type should you use for money?**
> `Decimal` — avoids floating-point rounding errors.

**Q158. What type hides a secret value in logs/repr?**
> `SecretStr` (from Pydantic).

## Chunk 40 — App Configuration

**Q159. How do you set the API title and version?**
> `FastAPI(title="My API", version="1.0.0")`.

**Q160. How do you disable the docs entirely?**
> `FastAPI(docs_url=None, redoc_url=None)`.

**Q161. How do you change the docs URL?**
> `FastAPI(docs_url="/documentation")`.

**Q162. What is `root_path` for?**
> Telling FastAPI it's served under a prefix by a proxy (correct doc URLs).

## Chunk 41 — Running & Dev

**Q163. How do you enable auto-reload?**
> `fastapi dev` reloads automatically; with uvicorn use `--reload`.

**Q164. How do you make the server reachable on the network?**
> `uvicorn main:app --host 0.0.0.0 --port 80`.

**Q165. Difference between `fastapi dev` and `fastapi run`?**
> `dev` = reload + debugging; `run` = production-style (no reload).

**Q166. In `uvicorn main:app`, what are `main` and `app`?**
> `main` = the module/file; `app` = the FastAPI instance variable.

## Chunk 42 — Pydantic Field Basics

**Q167. How do you give a field a mutable default like `[]`?**
> `Field(default_factory=list)` — never use a bare `[]` default.

**Q168. How do you map a field to a different JSON key?**
> `Field(alias="itemName")`.

**Q169. How do you allow both the field name and its alias as input?**
> `model_config = {"populate_by_name": True}`.

**Q170. How do you make a model immutable?**
> `model_config = {"frozen": True}`.

## Chunk 43 — Returning Data

**Q171. How do you declare an endpoint returns a list of items?**
> `response_model=list[Item]`.

**Q172. How do you return a dictionary of models?**
> `response_model=dict[str, Item]`.

**Q173. How do you return an empty successful response?**
> Return `None` with `status_code=204`.

**Q174. How do you set the success status code for a route?**
> `@app.post(..., status_code=status.HTTP_201_CREATED)`.

## Chunk 44 — Errors Deeper

**Q175. How do you add headers to an HTTPException?**
> `HTTPException(status_code=401, detail="...", headers={"WWW-Authenticate": "Bearer"})`.

**Q176. When is 400 vs 422 appropriate?**
> 400 = general bad request; 422 = request body/params failed validation.

**Q177. What does 409 mean?**
> Conflict — e.g. creating a duplicate resource.

**Q178. Difference between 401 and 403?**
> 401 = not authenticated; 403 = authenticated but not allowed.

## Chunk 45 — Request Object

**Q179. How do you access the raw request?**
> `from fastapi import Request` → `def f(request: Request)`.

**Q180. How do you get the client IP?**
> `request.client.host`.

**Q181. How do you read raw query params?**
> `request.query_params`.

**Q182. How do you get the request method and URL?**
> `request.method` and `request.url`.

## Chunk 46 — Response Object

**Q183. How do you set cookies/headers on a normal endpoint?**
> Inject `response: Response` and modify it.

**Q184. How do you change the status code dynamically?**
> `response.status_code = 201`.

**Q185. How do you delete a cookie?**
> `response.delete_cookie("name")`.

**Q186. How do you set a default response class app-wide?**
> `FastAPI(default_response_class=ORJSONResponse)`.

## Chunk 47 — JSON Handling

**Q187. What does `jsonable_encoder` do?**
> Converts models/datetimes/etc. into JSON-serializable Python types.

**Q188. Why return a `JSONResponse` directly?**
> Full control over content, status code, and headers.

**Q189. How do you get faster JSON serialization?**
> Use `ORJSONResponse` (backed by `orjson`).

**Q190. When is manual JSON control needed?**
> Custom envelopes, non-standard status/header combos, or third-party payloads.

## Chunk 48 — Enums & Constants

**Q191. How do you use readable status code constants?**
> `from fastapi import status` → `status.HTTP_201_CREATED`.

**Q192. How is a `str` Enum serialized in a response?**
> As its `.value` string.

**Q193. How is an `IntEnum` serialized?**
> As its integer value.

**Q194. Why prefer status constants over numbers?**
> Readability and fewer magic-number mistakes.

## Chunk 49 — Foundations Recap

**Q195. What is Starlette?**
> The lightweight ASGI toolkit FastAPI is built on (routing, middleware, etc.).

**Q196. What is Pydantic's job in FastAPI?**
> Data validation and serialization.

**Q197. Difference between ASGI and WSGI?**
> ASGI supports async and WebSockets; WSGI is synchronous only.

**Q198. Where can you see the raw OpenAPI schema?**
> At `/openapi.json`.

---

# 🟡 PART 5 — MEDIUM (Deeper Skills)

## Chunk 50 — Pydantic Validators

**Q199. What does `@field_validator` do?**
> Validates or transforms a single field.

**Q200. What does `@model_validator` do?**
> Validates across multiple fields (e.g. `mode="after"`).

**Q201. How do you signal invalid data inside a validator?**
> Raise `ValueError` — FastAPI turns it into a 422.

**Q202. How do you reuse validation logic across models?**
> Shared `Annotated` types or shared validator functions.

## Chunk 51 — Pydantic Advanced

**Q203. What is `@computed_field`?**
> A derived, read-only property included in serialized output.

**Q204. What does `model_config` `extra` control?**
> Handling of unexpected fields: `"ignore"`, `"forbid"`, or `"allow"`.

**Q205. How do you constrain a value inline?**
> `Annotated[int, Field(gt=0, le=100)]`.

**Q206. What is a discriminated (tagged) union?**
> A union resolved by a field: `Field(discriminator="type")`.

## Chunk 52 — Annotated Style

**Q207. What's the modern way to declare a validated query param?**
> `q: Annotated[str, Query(min_length=3)]`.

**Q208. Why is `Annotated` recommended?**
> Reusable, type-checker friendly, and keeps defaults clear.

**Q209. How do you make a reusable dependency type?**
> `CommonDep = Annotated[dict, Depends(get_common)]`.

**Q210. How do you give an `Annotated` param a default?**
> `q: Annotated[str, Query()] = "all"`.

## Chunk 53 — Dependency Patterns

**Q211. What does a `yield` dependency enable?**
> Setup before and teardown after the request (like context managers).

**Q212. How do you use a class as a dependency?**
> `Depends(MyDependency)` — FastAPI calls it and injects the instance.

**Q213. How do you make a configurable (parameterized) dependency?**
> A class storing config; pass an instance to `Depends`.

**Q214. How do you apply a dependency to the whole app?**
> `FastAPI(dependencies=[Depends(verify_key)])`.

## Chunk 54 — Dependency Scope

**Q215. How do you apply dependencies to one router?**
> `APIRouter(dependencies=[Depends(...)])`.

**Q216. Are dependencies re-run for each use in a request?**
> No — cached per request by default.

**Q217. How do you disable dependency caching?**
> `Depends(fn, use_cache=False)`.

**Q218. Can a dependency return nothing?**
> Yes — useful for validation/side effects only.

## Chunk 55 — Current-User Pattern

**Q219. How is `get_current_user` typically built?**
> A dependency that decodes the token and loads the user.

**Q220. How do you build `get_current_active_user`?**
> A dependency that depends on `get_current_user` and checks `is_active`.

**Q221. How do you inject the user into a route?**
> `user: User = Depends(get_current_user)`.

**Q222. What header should a 401 include for Bearer auth?**
> `WWW-Authenticate: Bearer`.

## Chunk 56 — Files & Forms Advanced

**Q223. How do you accept multiple file uploads?**
> `files: list[UploadFile]`.

**Q224. How do you accept a form field and a file together?**
> Mix `Form()` params with an `UploadFile` param.

**Q225. `UploadFile` vs `bytes` for uploads?**
> `UploadFile` streams (spooled to disk); `bytes` loads it all in memory.

**Q226. How do you read an upload's name and type?**
> `.filename` and `.content_type`.

## Chunk 57 — Static & Templates

**Q227. How do you serve static files?**
> `app.mount("/static", StaticFiles(directory="static"))`.

**Q228. How do you set up Jinja2 templates?**
> `templates = Jinja2Templates(directory="templates")`.

**Q229. How do you render a template?**
> `templates.TemplateResponse("page.html", {"request": request, ...})`.

**Q230. Why must `request` be in the template context?**
> `Jinja2Templates` requires it to build URLs and context.

## Chunk 58 — Middleware Deeper

**Q231. What does `GZipMiddleware` do?**
> Compresses responses above a size threshold.

**Q232. What does `TrustedHostMiddleware` do?**
> Blocks requests with unexpected `Host` headers.

**Q233. What does `HTTPSRedirectMiddleware` do?**
> Redirects HTTP requests to HTTPS.

**Q234. In what order does middleware run?**
> Last added is outermost (runs first on the way in).

## Chunk 59 — CORS Deeper

**Q235. What must change to allow credentials in CORS?**
> `allow_credentials=True` and you cannot use `"*"` for origins.

**Q236. How do you restrict allowed methods/headers?**
> `allow_methods=[...]`, `allow_headers=[...]`.

**Q237. What is a CORS preflight request?**
> A browser `OPTIONS` request sent before certain cross-origin calls.

**Q238. How do you let JS read a custom response header?**
> `expose_headers=["X-Custom"]`.

## Chunk 60 — Testing Deeper

**Q239. How do you make a reusable test client fixture?**
> A pytest fixture that yields `TestClient(app)`.

**Q240. How do you test async endpoints directly?**
> Use `httpx.AsyncClient` with `ASGITransport(app=app)`.

**Q241. How do you use a test database?**
> Override `get_db` with a dependency returning a test session.

**Q242. How do you bypass auth in tests?**
> Override the `get_current_user` dependency.

## Chunk 61 — Testing More

**Q243. How do you test a file upload?**
> `client.post(url, files={"file": ("name.txt", data)})`.

**Q244. How do you send custom headers in a test?**
> `client.get(url, headers={"Authorization": "Bearer ..."})`.

**Q245. How do you test a WebSocket?**
> `with client.websocket_connect("/ws") as ws: ...`.

**Q246. How do you run one test with many inputs?**
> `@pytest.mark.parametrize(...)`.

## Chunk 62 — Settings & Config

**Q247. What does `BaseSettings` provide?**
> Reads config from environment variables and `.env` files automatically.

**Q248. Why cache the settings object?**
> To avoid re-reading env on every request — `@lru_cache` on `get_settings`.

**Q249. How do you group nested config?**
> Nested `BaseSettings`/models with an env prefix.

**Q250. How do you load Docker secrets?**
> Configure `secrets_dir` in the settings config.

## Chunk 63 — Logging

**Q251. How do you configure logging in a FastAPI app?**
> `logging.config.dictConfig({...})` at startup.

**Q252. How do you correlate logs for one request?**
> Middleware attaches a request/correlation ID to log records.

**Q253. What are structured logs?**
> Machine-readable (usually JSON) logs — e.g. via `structlog`.

**Q254. How do you control Uvicorn access logs?**
> Enable/disable via Uvicorn config or `--no-access-log`.

## Chunk 64 — Pagination & Querying

**Q255. How does offset/limit pagination work?**
> `skip` and `limit` query params slice the result set.

**Q256. What is cursor pagination and why use it?**
> Paging by an opaque cursor — stable under inserts/deletes.

**Q257. How do you add filtering?**
> Optional query params mapped to `WHERE` conditions.

**Q258. How do you safely add sorting?**
> Validate a `sort` param against an allowlist of fields.

## Chunk 65 — API Keys & Basic Auth

**Q259. How do you read an API key from a header?**
> `APIKeyHeader(name="X-API-Key")` as a dependency.

**Q260. How do you read an API key from the query string?**
> `APIKeyQuery(name="api_key")`.

**Q261. How do you implement HTTP Basic auth?**
> `HTTPBasic()` dependency returning username/password.

**Q262. How do you compare secrets safely?**
> `secrets.compare_digest(...)` to prevent timing attacks.

## Chunk 66 — OpenAPI Customization

**Q263. How do you declare API servers in OpenAPI?**
> `FastAPI(servers=[{"url": "https://api.example.com"}])`.

**Q264. How do you add rich tag metadata?**
> `openapi_tags` with `name`, `description`, `externalDocs`.

**Q265. Why set a custom `operation_id`?**
> Cleaner method names in generated client SDKs.

**Q266. How do security schemes appear in the docs?**
> Auto-added from your security dependencies (OAuth2, API key, etc.).

## Chunk 67 — Response Tuning

**Q267. How do you drop specific fields from a response?**
> `response_model_exclude={"password"}`.

**Q268. How do you whitelist response fields?**
> `response_model_include={"id", "name"}`.

**Q269. How do you output using field aliases?**
> `response_model_by_alias=True`.

**Q270. How do you document multiple possible response shapes?**
> A `Union` response model or `responses={...}` per status code.

## Chunk 68 — Serialization

**Q271. How do you add a custom encoder to `jsonable_encoder`?**
> `jsonable_encoder(obj, custom_encoder={MyType: str})`.

**Q272. What format are datetimes serialized in?**
> ISO 8601 by default; customizable via a serializer.

**Q273. How do you globally drop `None` fields from responses?**
> `response_model_exclude_none=True`.

**Q274. How do you serialize directly from ORM objects?**
> `model_config = {"from_attributes": True}`.

## Chunk 69 — Concurrency Basics

**Q275. How do you run blocking code inside an async route?**
> `await run_in_threadpool(fn)` or `asyncio.to_thread(fn)`.

**Q276. What must you avoid in an `async def` route?**
> Blocking (sync) I/O — it stalls the event loop.

**Q277. How do you run several async calls concurrently?**
> `await asyncio.gather(...)`.

**Q278. When are `BackgroundTasks` not enough?**
> For heavy, long, retryable, or durable work — use a task queue.

## Chunk 70 — HTTP Clients

**Q279. How should you manage an outbound `httpx` client?**
> Create one in `lifespan`, store on `app.state`, reuse it.

**Q280. Why always set timeouts on HTTP calls?**
> To prevent hanging requests from exhausting resources.

**Q281. How do you add retries to outbound calls?**
> httpx transport retries or a library like `tenacity`.

**Q282. Why reuse a single `AsyncClient`?**
> Connection pooling — faster than creating one per request.

## Chunk 71 — Streaming & SSE

**Q283. How do you implement Server-Sent Events?**
> Stream `text/event-stream` via `StreamingResponse`.

**Q284. How do you stream a large file download?**
> Yield chunks from a generator into `StreamingResponse`.

**Q285. How do you support partial media downloads?**
> Handle Range requests and return `206 Partial Content`.

**Q286. What does a streaming generator yield?**
> Chunks of `bytes` (or strings), one at a time.

## Chunk 72 — Events & Lifespan

**Q287. What belongs in a `lifespan` handler?**
> Opening resources (DB/HTTP pools) before `yield`, closing after.

**Q288. Where do you store shared resources?**
> On `app.state` (e.g. `app.state.http = client`).

**Q289. How do you access app state inside a route?**
> `request.app.state`.

**Q290. Where does graceful shutdown logic go?**
> After the `yield` in the `lifespan` context manager.

## Chunk 73 — Sub-Apps & Mounting

**Q291. How do you mount a second FastAPI app?**
> `app.mount("/v2", app_v2)`.

**Q292. Do mounted apps get their own docs?**
> Yes — each has independent `/docs`.

**Q293. How do you mount a Flask/Django (WSGI) app?**
> Wrap it with `WSGIMiddleware` and `app.mount(...)`.

**Q294. Are mounted apps isolated?**
> Yes — separate routing/middleware scopes.

## Chunk 74 — Validation UX

**Q295. How do you customize the 422 error format?**
> Override the `RequestValidationError` exception handler.

**Q296. How do you make validation errors friendlier?**
> Reshape `exc.errors()` into a simpler structure.

**Q297. What is `problem+json`?**
> RFC 7807 — a standard error response format.

**Q298. What happens if a response fails its `response_model`?**
> A server error (500) is raised — it's treated as a bug.

---

# 🔴 PART 6 — HARD (Production & Architecture)

## Chunk 75 — JWT Deep

**Q299. What are the three parts of a JWT?**
> Header, payload, and signature (base64url-encoded, dot-separated).

**Q300. Name common JWT claims.**
> `sub`, `exp`, `iat`, `nbf`, plus custom claims.

**Q301. How do you verify a JWT?**
> Decode with the secret/public key, pin allowed algorithms, check `exp`.

**Q302. How do you prevent the algorithm-confusion attack?**
> Explicitly whitelist algorithms and never accept `"none"`.

## Chunk 76 — Token Lifecycle

**Q303. Why separate access and refresh tokens?**
> Short-lived access limits exposure; refresh renews without re-login.

**Q304. What is refresh token rotation?**
> Issuing a new refresh token each use and invalidating the old one.

**Q305. How do you revoke JWTs (which are stateless)?**
> Keep a denylist (by `jti`) in a store like Redis.

**Q306. How do you implement logout with tokens?**
> Revoke the refresh token / clear the auth cookie.

## Chunk 77 — Auth Architecture

**Q307. Where should tokens be stored client-side?**
> An httpOnly, secure cookie, or the `Authorization` header.

**Q308. What is RBAC?**
> Role-Based Access Control — roles grant permissions, checked in a dependency.

**Q309. Scopes vs roles?**
> Scopes = fine-grained token permissions; roles = groupings of users.

**Q310. What is the OAuth2 password flow?**
> Client posts username/password to a token endpoint and receives a token.

## Chunk 78 — Password Security

**Q311. Which algorithms hash passwords well?**
> bcrypt or argon2 (via `passlib`).

**Q312. What role does a salt play?**
> Makes identical passwords hash differently; bcrypt adds it automatically.

**Q313. How do you verify a password?**
> `pwd_context.verify(plain, hashed)`.

**Q314. Where do you enforce password strength?**
> In a Pydantic validator on the input model.

## Chunk 79 — SQLAlchemy 2.0

**Q315. How do you define models in SQLAlchemy 2.0?**
> Subclass a `DeclarativeBase` with typed `Mapped[...]` columns.

**Q316. How is a DB session provided per request?**
> A `yield` dependency creating and closing a session.

**Q317. What's the 2.0 query style?**
> `session.execute(select(Model).where(...))`.

**Q318. When do you commit vs rollback?**
> Commit on success; rollback in the error path.

## Chunk 80 — Async DB

**Q319. How do you create an async engine?**
> `create_async_engine(url)`.

**Q320. What session type is used for async?**
> `AsyncSession`, with `await session.execute(...)`.

**Q321. What Postgres driver enables async?**
> `asyncpg`.

**Q322. What's a common async DB pitfall?**
> Mixing sync and async sessions, or blocking calls in async code.

## Chunk 81 — Data Architecture

**Q323. What is the repository pattern?**
> Abstracting data access behind an interface, decoupling routes from the ORM.

**Q324. What is a service layer?**
> Business logic kept separate from HTTP routing.

**Q325. What is a unit of work?**
> Grouping operations into a single transaction boundary.

**Q326. DTO/schema vs entity/model?**
> Schema shapes API I/O; model represents DB storage.

## Chunk 82 — Transactions & Integrity

**Q327. How do you make multiple writes atomic?**
> Wrap them in one transaction (commit/rollback together).

**Q328. What is optimistic locking?**
> A version column detects concurrent modification on write.

**Q329. What is pessimistic locking?**
> Locking rows with `SELECT ... FOR UPDATE`.

**Q330. How do you handle a unique-constraint violation?**
> Catch `IntegrityError` and return `409 Conflict`.

## Chunk 83 — Caching

**Q331. What's commonly used as a cache layer?**
> Redis, for hot/read-heavy data.

**Q332. What is the cache-aside pattern?**
> Check cache → on miss load from DB → populate cache.

**Q333. How do you keep a cache fresh?**
> TTLs plus explicit invalidation on writes.

**Q334. What library caches FastAPI responses?**
> `fastapi-cache2` (decorator-based caching).

## Chunk 84 — Background Processing

**Q335. What is Celery?**
> A distributed task queue with a broker and worker processes.

**Q336. Which async-friendly queues exist?**
> `arq` and `dramatiq`.

**Q337. When do you outgrow `BackgroundTasks`?**
> For long, heavy, retryable, or cross-process durable jobs.

**Q338. How do you schedule recurring jobs?**
> APScheduler, Celery beat, or external cron.

## Chunk 85 — Messaging

**Q339. What are common message brokers?**
> RabbitMQ and Kafka.

**Q340. What is pub/sub?**
> Publishers emit events; subscribers consume them, decoupled.

**Q341. What is event-driven architecture?**
> Services react asynchronously to domain events.

**Q342. Why is idempotency important in messaging?**
> Messages can be delivered more than once; idempotency keys make retries safe.

## Chunk 86 — WebSockets Advanced

**Q343. How do you broadcast to many WebSocket clients?**
> Maintain a set of active connections and send to each.

**Q344. How do you implement chat "rooms"?**
> Group connections by a room key.

**Q345. How do you authenticate a WebSocket?**
> Validate a token from the query string or the first message.

**Q346. How do you detect a client disconnect?**
> Catch `WebSocketDisconnect`.

## Chunk 87 — Realtime Scaling

**Q347. How do you scale WebSockets across workers?**
> Use a backplane like Redis pub/sub to share messages.

**Q348. Why send heartbeats/pings?**
> To keep connections alive and detect dead ones.

**Q349. What is backpressure in streaming?**
> Preventing send queues from growing unbounded when clients are slow.

**Q350. SSE vs WebSockets — when to pick which?**
> SSE for simple one-way streams; WebSockets for two-way interaction.

## Chunk 88 — Observability

**Q351. What is a liveness health check?**
> A `/health` endpoint proving the process is up.

**Q352. What does a readiness check add?**
> Verifies dependencies (DB, cache) are reachable before serving traffic.

**Q353. How do you expose Prometheus metrics?**
> Use `prometheus-fastapi-instrumentator` to publish `/metrics`.

**Q354. What does distributed tracing give you?**
> End-to-end request spans across services (OpenTelemetry).

## Chunk 89 — Error Handling Architecture

**Q355. Why define domain exceptions?**
> To map business errors cleanly to HTTP responses.

**Q356. How do you handle a custom exception globally?**
> `@app.exception_handler(DomainError)` returning a response.

**Q357. Why add a catch-all 500 handler?**
> To log unexpected errors and hide internal details from clients.

**Q358. Why standardize an error envelope?**
> Consistent error shape makes clients easier to build.

## Chunk 90 — Deployment

**Q359. How do you run FastAPI under Gunicorn?**
> `gunicorn -k uvicorn.workers.UvicornWorker main:app`.

**Q360. How many workers should you run?**
> Start near `2 * cores + 1`, then tune for CPU and memory.

**Q361. How do you speed up Uvicorn?**
> Install `uvloop` and `httptools`.

**Q362. Why put a reverse proxy in front?**
> TLS termination, static files, load balancing, and buffering.

## Chunk 91 — Docker

**Q363. What base image is recommended?**
> A slim official Python image.

**Q364. What is a multistage Docker build?**
> Build dependencies in one stage, copy artifacts into a slim runtime stage.

**Q365. Why run the container as non-root?**
> Reduces blast radius if the app is compromised.

**Q366. How do you add a container healthcheck?**
> A `HEALTHCHECK` instruction hitting the `/health` endpoint.

## Chunk 92 — Kubernetes / Cloud

**Q367. How do probes map to FastAPI?**
> Liveness → `/health`; readiness → dependency-checking endpoint.

**Q368. How should config and secrets be supplied?**
> Env vars (12-factor); secrets via a secret manager, never in the image.

**Q369. How do you shut down gracefully in k8s?**
> Handle SIGTERM and finish in-flight requests before exiting.

**Q370. What enables horizontal scaling?**
> A stateless app with state kept in external stores (DB/Redis).

## Chunk 93 — Performance

**Q371. How do you profile a FastAPI app?**
> `py-spy` (sampling) or `cProfile`.

**Q372. How do you detect event-loop blocking?**
> Measure event-loop lag / slow async responses.

**Q373. Why is Pydantic v2 fast?**
> Its validation core (`pydantic-core`) is written in Rust.

**Q374. How do you reduce serialization cost?**
> Use `orjson` responses and lean response models.

## Chunk 94 — Security Hardening

**Q375. How do you prevent SQL injection?**
> Use the ORM / parameterized queries — never string-format SQL.

**Q376. Which security headers should you set?**
> HSTS, `X-Content-Type-Options`, and a Content-Security-Policy.

**Q377. How do you limit request body size?**
> Enforce a max size (proxy limit or middleware) and reject oversized bodies.

**Q378. How do you make file uploads safe?**
> Validate type/size and store outside the web root.

## Chunk 95 — Rate Limiting & Abuse

**Q379. How do you add rate limiting?**
> A library like `slowapi`, or limits at the gateway/proxy.

**Q380. Name two rate-limiting algorithms.**
> Token bucket and sliding window.

**Q381. Should you limit per user or per IP?**
> Prefer per identity (user/API key); fall back to IP for anonymous traffic.

**Q382. What status code signals rate limiting?**
> `429 Too Many Requests`.

## Chunk 96 — HTTP Caching Semantics

**Q383. What is an ETag?**
> A validator identifying a resource version for conditional requests.

**Q384. What does `Cache-Control` do?**
> Instructs clients/proxies how and how long to cache.

**Q385. When do you return `304 Not Modified`?**
> When the client's cached version is still current.

**Q386. Which HTTP methods are idempotent?**
> GET, PUT, DELETE (and HEAD) — POST is not.

## Chunk 97 — Versioning & Lifecycle

**Q387. What are common API versioning strategies?**
> URL path, request header, or media-type versioning.

**Q388. How do you deprecate an endpoint responsibly?**
> Mark it deprecated, document a sunset date, and communicate it.

**Q389. What keeps changes backward compatible?**
> Additive-only changes — don't remove/rename existing fields.

**Q390. How do you generate client SDKs?**
> From the OpenAPI schema (e.g. `openapi-generator`).

## Chunk 98 — Advanced Integrations

**Q391. How do you add GraphQL to FastAPI?**
> Integrate Strawberry (or similar) as a mounted GraphQL app.

**Q392. Why run gRPC alongside FastAPI?**
> Efficient internal service-to-service communication.

**Q393. How do microservices communicate?**
> Via REST, gRPC, or asynchronous messaging.

**Q394. What does an API gateway handle?**
> Edge routing, auth, rate limiting, and aggregation.

## Chunk 99 — Testing & Quality at Scale

**Q395. What is contract testing?**
> Verifying an API conforms to its agreed schema/contract.

**Q396. Which tools do load testing?**
> Locust and k6.

**Q397. How do you swap externals in integration tests?**
> `app.dependency_overrides` to inject fakes/stubs.

**Q398. Why use factory fixtures for test data?**
> To build valid, varied test objects quickly and consistently.

## Chunk 100 — Architecture Wrap-Up

**Q399. What do clean-architecture layers look like in FastAPI?**
> Routes → services → repositories → models, each depending only inward.

**Q400. What's the golden rule of FastAPI?**
> Let type hints do the work — they drive validation, docs, and editor support.

---

## 📌 Master Cheat-Sheet

| Task               | Syntax                                              |
|--------------------|----------------------------------------------------|
| Create app         | `app = FastAPI()`                                   |
| Route              | `@app.get/post/put/delete("/path")`                |
| Path param         | `/items/{id}` + `id: int`                          |
| Query param        | `q: str | None = None`                             |
| Body               | `item: Item` (a `BaseModel`)                        |
| Validate field     | `Field(gt=0, max_length=50)`                        |
| Validate query     | `Query(min_length=3)`                              |
| Output shape       | `response_model=Item`                              |
| Custom status      | `status_code=201`                                 |
| Raise error        | `raise HTTPException(404, "msg")`                  |
| Dependency         | `x = Depends(fn)`                                  |
| Auth token         | `Depends(oauth2_scheme)`                          |
| DB session         | `Depends(get_db)`                                 |
| Router             | `APIRouter()` + `app.include_router(r)`           |
| Middleware         | `@app.middleware("http")`                         |
| CORS               | `add_middleware(CORSMiddleware, ...)`             |
| Background task    | `BackgroundTasks.add_task(fn)`                    |
| WebSocket          | `@app.websocket("/ws")`                           |
| Stream/File        | `StreamingResponse` / `FileResponse`             |
| Test               | `TestClient(app)`                                 |
| Run dev            | `fastapi dev main.py`                             |
| Run prod           | `gunicorn -k uvicorn.workers.UvicornWorker`      |

---

### How to use this deck
1. Study one **chunk (4 cards)** at a time.
2. Finish all 🟢 chunks before moving to 🟡, then 🔴.
3. Mark cards you miss; revisit them next session.
4. After each chunk, try writing tiny code to *prove* you understood it.

Happy learning! 🚀
