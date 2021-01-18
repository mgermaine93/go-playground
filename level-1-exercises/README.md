# Problem Set #1 Notes:

- A STATEMENT is the smallest standalone element of a program that expresses some action to be carried out. They usually take up a single line in a program.
- An EXPRESSION is a combination of one or more explicit values, constants, variables, operators, and functions that a programming language interprets and computes to produce another value.
- The ZERO VALUE is the default value the Go compiler assigns to the variables when no value is provided.
- [golang.org](https://golang.org/) is the official website of the Go programming language, and only has documentation for the standard library.
- [godoc.org](https://godoc.org/) houses documentation for the standard Go library AS WELL AS for third-party packages.
- The short declaration operator (":=") can only be used at the block level, i.e. between curly braces/inside functions. Use it whenever possible.
  - Allows you to declare a variable and assign a value to the variable at the same time.

* "=" allows you to "re-assign" a previously declared variable to another value.
* The blank operator ("\_") serves as an anonymous placeholder that tells the compiler that you are not going to use a value returned by a function.
* Variadic parameters have type "...interface{}", and it means the function can take as many values of that type as you want to pass in.
* Every value in Go is also of type "empty interface", which is expressed as "interface{}".
* Packages are a way to organize/group similar code together.
  - Imports bring packages into programs.
  - Reading the index of packages can be informative.
* Think of the compiler as a "police officer". It ensures that program code is well-written.
* Go is a static programming language, so you can't "re-declare" variables of one type to be a variable of a different type.
* A VARIABLE is DECLARED to hold a VALUE of a certain TYPE.
  _ In Go, you can create your own type (e.g., "type hotdog int", "var b hotdog").
  But in static programming languages, you cannot store a the value of a variable of one type into another variable of another type.
  _ "Conversion" is to take a value of a certain type and convert it to another type.
  _ The same idea as "casting" in another languages, but Go calls it "conversion".
  _ e.g.:
  `var a int`
  `type hotdog int`
  `var b hotdog`
  `a = b // this does not run successfully`
  `a = int(b) // this runs successfully, because the underlying type of "hotdog" is "int".`
