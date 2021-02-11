# Problem Set #6 Notes

- Functions
  - We define functions with PARAMETERS (if any)
  - We call our functions by passing in ARGUMENTS (if any)
- "Decoupling" code means to "break up" code into modular components that can be run on their own. Functions help with this.
- The "Defer" keyword
  - This will defer the execution of a function until wherever the function is being called comes to an end/"returns" something.
  - A good application of this is to automatically close files:
    - Right when you open a file, you can write a function to close a file, but "defer" it until func main() is finished/exits.
    - Allows memory to be optimized and not have too many tabs open and taking up space.
- Methods
  - When you have a receiver, it's going to attach a function to a type.
  - As a result, any VALUE of that TYPE then has access to the METHOD.
- Interfaces
  - They allow us to define behavior and utilize polymorphism.
  - A VALUE can be of more than one TYPE.
  - "If you have these METHODS, you are also of my TYPE."
- Anonymous Functions
  - These are nameless functions that otherwise work just like normal functions.
- Func Expressions
  - Recall that expressions always come down to a value.
  - Func Expressions basically just assign the output of functions to variables.
- Returning a Function
  - A function is a type, too!
  - Useful for when you want to build a func, get it a certain way, and return something else.
- Callbacks
  - Callbacks are passing functions as arguments.
- Closure
  - "We want to enclose the scope of a variable in some code so that it is contained within a certain area."
  - Variables declared in the outer scope are accessible in inner scopes.
  - Limiting scope is a good general thing to do.
- Recursion
  - Anything that can be done with recursion can also be done with a loop.
  - "Recursion" is whenever a function calls itself.