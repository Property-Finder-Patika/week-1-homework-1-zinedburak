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

The math package has functions for creating and detecting the
special values defined by IEEE 754: the positive and negative infinities, which represent numbers of excessifve
magnitude and the result of division by zero; NaN("not a number""), the result of such mathematically dubious operations
as  0/0 or Sqrt(-1).

``` go
var z float64
fmt.Println(z,-z,1/z,-1/z,z/z)  // "0 -0 +Inf -Inf NaN"
```