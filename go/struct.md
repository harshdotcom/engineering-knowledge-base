# Structs in Go

## 🔹 What is a Struct?

A **struct** is a user-defined data type in Go that groups multiple related variables (fields) of different data types into a single unit.

Structs are useful for:
- Representing real-world entities (User, Product, Car, etc.)
- Organizing related data together  
- Creating complex data types  
- Improving code readability  

---

## 🔹 Declaring a Struct

```go
type Person struct {
    Name string
    Age  int
}


### A struct is a collection of fields.

``` go
type Vertex struct {
	X int
	Y int
}

// Initialization of struct with . 

v := Vertex{1, 2}
v.X = 4
fmt.Println(v.X)


// Struct with pointer

v := Vertex{1, 2}
p := &v
p.X = 1e9
fmt.Println(v)

```


| Concept            | Meaning             |
| ------------------ | ------------------- |
| `type X struct {}` | Define struct type  |
| `x.Field`          | Access field        |
| `&x`               | Pointer to struct   |
| `*X`               | Struct pointer type |
| Anonymous struct   | Struct without name |


### Structs in Functions (Pass by Value)

``` go
func updateAge(p Person) {
    p.Age = 100
}

func main() {
    p := Person{Name: "Harsh", Age: 22}
    updateAge(p)
    fmt.Println(p.Age) // Output: 22
}
```

### 🔹 Struct Pointer (Pass by Reference)

``` go
func updateAge(p *Person) {
    p.Age = 100
}

func main() {
    p := Person{Name: "Harsh", Age: 22}
    updateAge(&p)
    fmt.Println(p.Age) // Output: 100
}

```

