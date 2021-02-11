# Problem Set #4 Notes

- Arrays are numbered sequences of elements of single types.
- Arrays are not used much in Go, but they are a good building block for slices, and slices are used a lot in Go.

* A SLICE allows you to group together VALUES of the same TYPE.
* Composite literals are expressions that create new instances each time they are evaluated.
* Composite literals construct values for structs, arrays, slices, and maps.

- Composite literals: "x := type{values}" => "x := []int{4, 5, 7, 8, 42}

* Composite literals consist of declaring the type and then declaring the values of the type in-between curly braces.
* "Variadic parameter" = unlimited amount of values of this type.
* If you know you're going to have a slice of a certain size, go ahead and make it that size so you won't have to change it later and use more processing power than necessary.
* To do the above, use "make".
