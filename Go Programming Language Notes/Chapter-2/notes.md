# Chapter 2 - Program Structure

### 2.1 Names

A name begins with a letter or an underscore and may have any number of additional letters digits, and underscores. Case
maters heapsort and Heapsort are different names

Go has 25 keywords like if and switch that may be used only where the syntax permits; they can't be used as names

If an entity is declared within a function, it is local to that function. If declared outside of a function, however, it
is visible in all files of the package to which it belongs. The case of the first letter of a name determines its
visibility across package boundaries. If the name begins with an upper-case letter, it is exported, which means that it
is visible and accessible outside of its own package and may be referred by other parts of the program as with Printf in
the fmt package. **Package names themselves are always in lower case**
---

### 2.2 Declarations

A declaration names a program entity and specifies some or all of its properties. There are four major kinds of
declarations: var, const, type, and func.

A Go program is stored in one or more files whose names end in .go. Each file begins with a package declaration that
says what package the file is part of. The package declaration is followed by any import declarations, and then a
sequence of package-level declarations of types, variables constants and functions, in any order. For example, this
program declares a constant, a function and a couple of variables

``` Go
import "fmt"
const boilingF = 212.0
func main() {
    var f = boilingF
    var c = (f - 32) * 5 / 9
    fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
    // Output:
    // boiling point = 212°F or 100°C
}
```

The constant boilingF is a package-level declaration (as is main), whereas the variables f and c are local to the
function main. The name of each package-level entity is visible not only throughout the source file that contains its
declaration, but throughout all the files of the package ( **If these variables are declared with a small case letter
they are still not accessible from the files in the same package** )
. By contrast, local declarations are visible only within the function in which they are declared and perhaps only
within a small part of it
---

### 2.3 Variables

The general form of a declaration

``` Go
var name type = expression
```

---
It is possible to declare and optionally initialize a set of variables in a single declaration with a matching list of
expressions. Omitting the type allows declaration of multiple variables of different types

``` Go
var i, j, k int
var b, f, s = true, 2.3, "four"
```

---
A set of variables can also be initialized by calling a function that returns multiple values:

``` Go
var f, err = os.Open(name) // os.Open returns a file and an error
```

---

### Pointers

A variable is a piece of storage containing a value. Variables created by declarations are identified by a name, such as
x, but many variables are identified only by expressions like x[i] or x.f. All these expressions read the value of a
variable, except when they appear on the left-hand side of an assignment, in which case a new value is assigned to the
variable.

A pointer value is the address of a variable. A pointer is thus the location at which a value is stored. Not every value
has an address, but every variable does. With a pointer, we can read or update the value of a variable indirectly,
without using or even knowing the name of the variable, if indeed it has a name.

If a variable is declared var x int, the expression &x ("address of x") yields a pointer to an integer variable that is,
a value of type *int, which is pronounced "pointer to int". If this value is called p, we say "p points to x", or
equivalently "p contains the address of x." The variable to which p points is written *p. The expression *p yields the
value of that variable, an int but since *p denotes a variable, it may also appear on the left-hand side of an
assignment, in which case the assignment updates the variable

``` Go
x := 1
p := &x // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2         // equivalent to x = 2
fmt.Println(x) // "2"
```

### The new Function

Another way to create a variable is to use the built-in function new. The expression new(T) creates an unnamed variable
of type T, initializes ti to the zero value of T, and returns its address, which is a value of type *T

``` Go
p := new(int)   // p, of type *int, points to an unnamed int variable
fmt.Println(*p) // "0"
*p = 2          // sets the unnamed int to 2
fmt.Println(*p) // "2"
```

### 2.4 Assignments

### Tuple Assignment

Example of allowed tuple assignments

``` Go
x,y = y,x
a[i], a[j] = a[j], a[i]

// computing greatest common divisor
func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x 
}

// computing the n-th Fibonacci number iteratively
func fib(n int) int {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        x, y = y, x+y
    }
    return x 
}
i, j, k = 2, 3, 5
```

### 2.5 Type Declarations

A type declaration defines a new named type that has the same underlying type as an existing type. The named type
provides a way to separate different and perhaps incompatible uses of the underlying type so that they can't be mixed
unintentionally.

    type name underlying-type

Type declarations most often appear at package leve, where the named type is visible through-out the package, and if the
name is exported, it's accessible from other packages as well.

To illustrate type declarations, let's turn the different temperature scales into different types:

``` go
import "fmt"
type Celsius float64
type Fahrenheit float64
const (
    AbsoluteZeroC Celsius = -273.15
    FreezingC     Celsius = 0
    BoilingC      Celsius = 100
)
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// The underlying type of a named type determines its structure and representation, and also the of intrinsic 
// operationsit supports, which are the same as if the underlying type had been used directly

fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
boilingF := CToF(BoilingC)
fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F

// compile error: type mismatch because boillingF is Fahrenheit While FreezingC is Celsius
fmt.Printf("%g\n", boilingF-FreezingC) 


// Comparison operators can also be used to compare a value of a named type to another of the same type.
// or to a value of the underlying type. But two values of different named types cannot be compared directly

var c Celsius
var f Fahrenheit
fmt.Println(c == 0) // "true"
fmt.Println(f >= 0) // "true"
fmt.Println(c == f) // compile error: type mismatch
fmt.Println(c == Celsius(f)) // "true"!

// Note the last case carefully. In spite of its name, the type conversion Celsius(f) does not change the value of its
// argument, just its type. The test is true because c and f are both zero

// You can declare type methods just like they are objects 
// EXAMPLE:
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

c := FToC(212.0)
fmt.Println(c.String()) // "100°C"
fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
fmt.Printf("%s\n", c)   // "100°C"                    
fmt.Println(c)          // "100°C"                    
fmt.Printf("%g\n", c)   // "100"; does not call String
fmt.Println(float64(c)) // "100"; does not call String

```

This package defines two types, Celsius and Fahrenheit, for the two units of temperature. Even though both have the same
underlying type, float64, they are not the same type, so they cannot be compared or combined in arithmetic expressions.
Distinguishing the types makes it possible to avoid errors like inadvertently combining temperatures in the two
different scales; an explicit type conversion like Celsius(t) or Fahrenheit(t) is required to convert from a float64.
Celsius(t) and Fahrenheit(t) are conversions, not functions calls. They don't change the value or representation in any
way, but they make the change of meaning explicit. On the other hand, the functions CToF and FToC convert between two
scales; they do return different values.

