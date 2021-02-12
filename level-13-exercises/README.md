# Testing/Benchmarking Notes

- Testing is good because it checks to make sure you always get your expected outcomes.
- Go Tests MUST:
  - Be in a file that end with "\_test.go".
  - Be put in the same package as the file being tested.
  - Be in a function with a signature "func TestNameHere(testing.T).
- Run tests by cd-ing into the package with the test and the file to be tested and run "go test".
- TABLE TESTS check tests with multiple inputs at a time.
- "gofmt" formats Go code.
- "go vet" reports suspicious constructs.
- "golint" reports poor coding style.
- 70%-80% test coverage is usually good.
  - Find out how much coverage your tests have by running "go test -cover" in the package where your tests reside.
- "go test -bench" can be used to figure out how efficient different programs are.
  - This is useful because you can compare the efficiency of two programs that do the same thing, but in different ways.

* Remember to BET:
  - Benchmark
  - Example
  - Test
