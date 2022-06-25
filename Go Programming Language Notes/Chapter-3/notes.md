# Chapter 3 - Basic Data Types

### 3.1 Integers

Go provides both signed and unsigned integer arithmetic. There are four distinct sizes of signed integers - 8, 16, 32,
and 64 bits - represented by the types int8, int16, int32, and int64 and corresponding unsigned versions
uint8,uint16,uint32, and uint64.

There are also two types called just int and uint that are the natural or most efficient size for signed and unsigned
integers on a particular platform; int is by far the most widely used numeric type. Both these types have the same size,
either 32 or 64 bits, but one must not make assumptions about which; different compilers may make different choices even
on identical hardware.

The type rune is a synonym for int32.

Regardless of their size, int, uint, and uintptr are different types from their explicitly sized siblings. This int is
not the same type as int32, even if the natural size of integers is 32 bits

Signed numbers are represented in 2's-complement form, in which the high-order bit is reserved for the sign of the
number and the range of values of an n-bit number is from -2^(n-1) to 2^(n-1) - 1. Unsigned integers use the full range
of bits for non-negative values and thus have the range 0 to 2^n - 1. For instance, the range of int8 is -128 to 127,
whereas the range of uint8 is 0 to 255

The remainder operator % applies only to integers.In Go, the sign of the remainder is always the same as the sign of the
divided, so -5%3 and -5%-3 are both -2

Go also provides the following bitwise binary operators, the first four of which treat their operands as

`& bitwise AND`

`| bitwise OR`

`^ bitwise XOR`

`&^ bit clear(AND NOT)`

`<< left shift`

`>> right shift`

### 3.2 Floating-Point Numbers

Go Provides two sizes of floating -point numbers, float 32 and float 64. A float32 provides approximately six decimal
digits of precision, whereas a float64 provides 15 digits.

The math package has functions for creating and detecting the special values defined by IEEE 754: the positive and
negative infinities, which represent numbers of excessifve magnitude and the result of division by zero; NaN("not a
number""), the result of such mathematically dubious operations as 0/0 or Sqrt(-1).

``` go
var z float64
fmt.Println(z,-z,1/z,-1/z,z/z)  // "0 -0 +Inf -Inf NaN"
```

### 3.5 Strings

A string is an immutable sequence of bytes. Strings may contain arbitrary data, including bytes with value 0, but
usually they contain human-readable text. Text strings are conventionally interpreted as UTF-8-encoded sequences of
Unicode code points.

The built-in len functions returns the number of bytes(not runes) in a string and index operation s[i] retrieves the
i-th byte of string s, where 0 <= i < len(s)

```go
s := "hello, world"
fmt.Println(len(s))     // "12"
fmt.Println(s[0], s[7]) // "104 119"  ('h' and 'w')
```

The i-th byte of a string is not necessarily the i-th character of a string, because the UTF-8 encoding of a non-ASCII
code point requires two or more bytes.

The substring operation s[i:j] yields new string consisting of the bytes of the original string starting at index i and
continuing up to, but not including, the byte at index j. The result contains j0i bytes.

```go
fmt.Println(s[0:5]) // "hello"
fmt.Println(s[:5]) // "hello"
fmt.Println(s[7:]) // "world"
fmt.Println(s[:])  // "hello, world"
```

Strings may be compared with comparison operators like == and <; the comparison is done byte by byte, so the result is
the natural lexicographic ordering.

Strings values are immutable: the byte sequence contained in a string value can never be changed, though of course we
can assign a new value to a string variable. To append string to another, for instance, we can write

```go
s := "left foot"
t := s
s += ", right foot"
```

This does not modify the string that s originally held but causes s to hold the new string formed by the += statement;
meanwhile, t still contains the old string

```go
fmt.Println(s) // "left foot, right foot"
fmt.Println(t) // "left foot"
```

Since strings are immutable, constructions that try to modify a string's data in place are not allowed:

```go
s[0] = 'L' // compile error: cannot assign to s[0]
```

### 3.5.4 Strings and Byte Slices

Four standard packages are particularly important for manipulating strings: bytes, strings, strconv, and unicode. The
strings package provides many functions for searching, replacing comparing, trimming, splitting, and joining.

The bytes package has similar functions for manipulating slices of bytes, of type []byte, which share some properties
with strings. Because strings are immutable, building up strings incrementally can involve a lot of allocation and
copying. In such cases, it's more efficient to use the bytes.

The strconv package provides functions for converting boolean, integer, and floating-point values to and from their
string representations, and functions for quoting and unquoting strings

The unicode package provides functions like IsDigit, IsLetter, IsUpper, and IsLower for classifying runes. Each function
takes a single rune argument and returns boolean. Conversion functions like ToUpper and ToLower convert a rune into the
given

### 3.6 Constants

Constants are expressions whose value is known to the compiler and whose evaluation is guaranteed to occur at compile
time, not at run time. The underlying type of every constant is basic type: boolean, string or number.

A const declaration defines named values that look syntactically like variables but whose value is constant, which
prevents accidental (or nefarious) changes during program execution. For instance, a constant is more appropriate than a
variable for a mathematical constant like pi, since its value won't change:

```go
const pi = 3.14159
const (
    e  = 2.71828182845904523536028747135266249775724709369995957496696763
    pi = 3.14159265358979323846264338327950288419716939937510582097494459
)
```

When a sequence of constants is declared as a group, the right-hand side expression may be omitted for all but th first
of the group, implying that the previous expression and its type should be used again

```go
const ( 
    a=1
    b
    c=2
    d
)
fmt.Println(a, b, c, d) // "1 1 2 2"
```

A const declaration may use the constant generator iota, which is used to create a sequence of related values without
spelling out each one explicitly. In a const declaration, the value of iota begins at zero and increments by one for
each item in the sequence.

```go
type Weekday int
const (
      Sunday Weekday = iota
      Monday
      Tuesday
      Wednesday
      Thursday
      Friday
      Saturday
)
```
This declares Sunday to be 0, Monday to be 1, and so on.

As a more complex example of iota, this declaration names the powers of 1024:
```go
const (
    _ = 1 << (10 * iota)
    KiB // 1024
    MiB // 1048576
    GiB // 1073741824
    TiB // 1099511627776
    PiB // 1125899906842624
    EiB // 1152921504606846976
    ZiB // 1180591620717411303424
    YiB // 1208925819614629174706176
)
```