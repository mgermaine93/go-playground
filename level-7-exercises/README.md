# Problem Set #7 Notes

- Pointers
  - They "point" to some location in memory where a value is stored.
  - Use the operator "&" immediately before a variable to see the ADDRESS in memory where that variable's value is held.
    - "AMPERSAND equals ADDRESS"
  - The "\*" gives you the VALUE stored at the ADDRESS, when you have the address.
    - "\*" equals "VALUE STORED AT ADDRESS"
  - EVERYTHING IN GO IS PASSED BY VALUE.
  - INTS and POINTERS to ints are two different TYPES.
- When to Use Pointers?
  - Use them whenever you have a large chunk of data and you don't want to pass around the whole chunk of data.
    - Just store the data and memory and use a pointer to reference it.
  - Also use them if you need to CHANGE something at a certain location.
    - "De-reference the address and assign some new value there."
- Method Sets
  - Method sets are methods that are attached to a type.
  - Is the set of methods for a given type? Then THAT IS IT'S METHOD SET.
