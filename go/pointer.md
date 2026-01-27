#  Pointers in Go

## 🔹 What is a Pointer?

A **pointer** is a variable that stores the **memory address** of another variable instead of storing the actual value.

Pointers are useful for:
- Modifying values inside functions  
- Avoiding unnecessary copying of large data  
- Efficient memory usage  

---

## 🔹 Declaring a Pointer

```go
var p *int
```

This means:

p is a pointer to an int

It can store the address of an integer variable

Default value of p is nil

## 🔹 Getting Address of a Variable (& operator)

```go
i := 42
p := &i
```

Explanation:

i stores value 42

&i gives the memory address of i

p stores that address


## 🔹 Dereferencing a Pointer (* operator)

```go
mt.Println(*p) // Output: 42
```

## 🔹 Modifying Value Using Pointer

```go
i := 10
p := &i

*p = 20
fmt.Println(i) // Output: 20
```

## 🔹Pointer Initialization
Method 1: Using address of a variable
```go
x := 10
p := &x
```
Method 2: Using new
```go
p := new(int)
*p = 10
```
## 🔹Pointer Initialization
```go
func update(n *int) {
    *n = 100
}

func main() {
    x := 10
    update(&x)
    fmt.Println(x) // Output: 100
}
```
## 🔹Important Notes
* Go has no pointer arithmetic (unlike C/C++)

* You cannot add or subtract pointers

* Safer memory handling than C/C++

| Concept      | Meaning                 |
| ------------ | ----------------------- |
| `&x`         | Address of variable     |
| `*p`         | Value at the address    |
| `var p *int` | Declare pointer         |
| `new(int)`   | Allocate memory         |
| `nil`        | Pointer with no address |

## A pointer stores the address of a variable, and dereferencing retrieves the value at that address.