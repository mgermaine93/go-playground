# Error Handling Notes

- The creators of Go felt that "coupling exceptions to a control structure", e.g., the try-catch-finally idiom, resulted in convoluted code.
- Go's multi-value function returns make it easy to report an error without overloading the return value.
- Go's error-handling is "pleasant", but quite different than other languages.
- Type "error" is built into the language. It is an interface.
  - Any type that has the method "error" attached to it will implicitly implement it and will also be of type "error".
- ALWAYS check your errors...
  - ...almost always.
  - Check them by doing "if err != nil {fmt.Println(err)}"
- "defer f.Close()" is helpful for closing files that are done running.
  - Like closing tabs on your computer.

* "Defer" statements allow us to think about closing each file right after opening it, guaranteeing that, regardless of the number of return statements in the function, the files WILL be closed.
  - A deferred function's arguments are evaluated when the defer statement is evaluated.
  - Deferred function calls are executed in L-I-F-O order after the surrounding function returns.
  - Deferred functions may read and assign to the returning function's named return values.
* "Recover" is only useful inside deferred functions.
  - "Recover" is a built-in function that regains control of a panicking goroutine.
  - During normal execution, a call to "recover" will return "nil" and have no other effect.
  - If the current Goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
* "errors.New("**\_\_**")" allows for custom error messages.
