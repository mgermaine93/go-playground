# Problem Set #2 Notes

- Iota = a special pre-declared identifier that helps with enumerated constants
  int64, float32, etc., are really only necessary if your application will be at a massive scale and you need to specify the memory alloted to each variable. \* Otherwise, just use "int" for whole numbers and "float64" for real (decimal) numbers.
- From Go: "Go's iota identifier is used in const declarations to simplify definitions of incrementing numbers. Because it can be used in expressions, it provides a generality beyond that of simple enumerations."

* Most popular coding scheme today is UTF-8. Go source code is always UTF-8.

- "rune" is an alias for int32.
- "byte" is an alias for uint8.
- If type "int" is used, then the compiler will choose whether int32 or int64 is used, i.e., implementation-specific sizes.

* A string is a sequence of bytes that represent Unicode code points, called runes.
