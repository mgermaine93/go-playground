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

*
