# Part 1

An introduction to building things in Go!

**What we'll cover:**
1) Running a go app
2) Packages & Modules
3) The Go toolchain 

### A go application
When we write Go, we are aiming to achieve one of the following two things: 
1) Building a library that others will use in their Go application.
2) Creating an executable that we can run. 

In this workshop, we'll be focusing on the latter, creating a web application to serve weather based
attire recommendations! 

Creating a go application, the go compiler will look for the `main` package to find the code it needs to run, 
and in this package, the `main()` function is where execution starts.

Take a look at this [example](https://play.golang.org/p/HmnNoBf0p1z):
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```
This is the hello world application of Go, and the anatomy of what source files in Go look like.

Let's step through it:
```go
package main
``` 
This is declaring what package the file belongs to, we'll cover packages more later in this part - but for now, just note that it's the *main* package, the one the compiler looks for when building an executable.

```go
import "fmt"

// You will often see this in block form, when there is more than a single dependancy
// import (
//    "fmt"
//    "anotherdepenancy"
// )

```
Following the package, the import block is declared. This is where we include all the dependencies we need in this file. In this simple example, we're including the [fmt](https://www.google.com) package, which is part of the standard library - it provides formatted I/O functionality.

```go
func main() {...}
```
This is the main function, the entry point to the go application. Although it serves a special purpose, it is defined exactly like any other function in Go. From the signature, we can see it has the following properties:
1) It's name, *main*, which is private (more on this later).
2) It has no parameters.
3) It has no return values.

The main function, is where our go apps start and ends, and is often used in services to perform setup, load configuration, connect and verify external dependancies and manage the lifecycle of an application. 
For now, let's just print some nice words; later on we'll revisit the main package later to tie together our application. 

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world")
}

```

### Building & Running
Once we have our main func, we need to run it! In Go, there's a few options to get us going.

To quickly run our code, we can do `go run main.go` - which will build, and then run the executable. This doesn't leave behind any binary we can ship or share, but it enables us to develop, and test our code quickly, giving a tight feedback loop.

When we're happy with our code, we can build the executable. We do this with `go build`. 
Go's compiler enables compilation to any platform, as long as specified by the environment variable `GOENV`. 

What about when we download a go source project, and want to install it? We'd have to build it and then move the binary into a path to run - or we could use `go install` which simplifies this into a single step. 
Go install builds the executable, and then moves it into the bin to run. 

### Packages & Modules
How do we organise code in Go, where do group files and functionality together?

In Go, the smallest atomic unit of code is the package - a package is simply a collection of .go files in the same directory. 
Anything inside the package has access to everything else within the package. Only exposed elements outside of the package are available publicly - more on this later!

Building on this, a Go module is a collection of packages grouped together. In a Go module you'll find the `go.mod` file, which defines a few things:
1) The module identifier (which will be the import path)
2) The module dependencies
3) The go version 

A couple rules with packages:
1) Cyclic dependencies will be detected and cause an error during compilation.
2) Importing but not using a package will cause an error during compilation.


### Diving into syntax
While we're here, it'll be useful to dive into a little of Go's syntax that we'll use when building our API later!

**functions**
While the `main()` function is quite simple, functions can get a little more interesting so let's take a look at some examples:
```go
// This is a function with a single parameter - 'name' - which is of type string
func DoStuff(name string) {}

// This is a function with two parameters - 'first' and 'last'. Because they are the same type we can leave the type until the final param 
func DoStuff(first, last string) {}

// This is a function that has no parameters - but returns an integer (the above examples have no return value)
func DoStuff() int {}

// This is a function that returns 2 values - and integer and a bool
func DoStuff() (int, bool) {}

// This is a function that returns 2 values - which are names. This means that they are already defined inside the function
func DoStuff() (name string, age int) {}
``` 

**variables**
There are lots of ways for us to create variables. 
Although Go is a statically typed language, we don't always have to specify the variable type; instead it can automatically infer the type of variable based on the value specified.
Here are some examples: 
```go
// This declares a variable x of type int which has no value (Go automatically initializes it with the zero value)
var x int

// This declares a variable x of type int, and assigns the value 1 to it
var x int = 1

// This declares a variable x of type int, and assigns the value 1 to it using type inference
var x = 1

// This declares two variables, x and y of types int - with values 1 and 2 respectively  - using the var block
var (
    x = 1
    y = 2
)

// This creates a variable x of type int with value 1 using shorthand notation
x := 1

// This creates two variables x, and y of types int - with values 1 and 2 respectively
x, y := 1, 2
```

In Go, creating a variable and not doing anything with it will cause a compilation error! This is because Go believes in writing
lean code and keeping its footprint as small as possible.
As a note, this is also the case with importing packages that aren't used!
```go
func bla() {
    x := 1
}
// Won't compile!
```
There is however, an exception to this rule - that allows us to suppress this error with the blank identifier `_`. 
The blank identifier tells the Go compiler that we are deliberately assigning this value to nothing, and it should not complain.
Let's take a look:
```go
func bla(){
    x := 1

    _ = x // Assign the value of x to the blank identifier
}
// Will compile!
``` 

**control blocks**
We can perform conditional logic with the `if` statement as you'd expect(note the missing parenthesis):
```go
if x == true {
    // Do stuff
}

if x == false {
    // Do stuff
} else {
    // Do other stuff
}


if x == 1 {
    // Do stuff
} else if x < 1 {
    // Do other stuff
} else {
    // Do other other stuff
}
```

**For loops**
We can also iterate over arrays, and maps with the `for` keyword:
```go
for {
    // Infinite loop (shorthand)
}

for x > 1 {
    // Continually loop until satisfied (similar to that of a while loop in other languages)
}

for idx, val := range myArray {
    // Iterate over an array (or slice, map) - giving the index, and copy of the value in the respective variables
}
```

### Go tooling
There are some great tools avaiable as part of the Go toolchain that make for a great developer experience when writing Go. 
Their usages vary from formatting, linting and vetting, serving documentation, and profiling / benchmarking.

Let's take a look at a few of the common ones that make day to day programming in Go easier and safer.

**go fmt**
Go fmt is a tool that formats our go code to comply with the standards defined for the language. There's a few reasons that this is [desirable](https://blog.golang.org/gofmt):
1) Don't waste time or effort worrying about formatting
2) Easier to read Go code when it's all styled the same
3) Easier to maintain Go code (and others' Go code)
4) Never debate about style again

Let's try running it with `go fmt ./...` 

**go lint**
Go lint is concerned with style (in the same way as go fmt), however where it differs is that it prints out style issues, rather than modifying the source.
While go fmt can (and should) be run automatically and often, go lint is not perfect but instead aims to offer a helping hand.

To run the linter, it should be installed following the instructions found [here](https://github.com/golang/lint). 

**go vet**
While the above are concerned with styling and syntax, Go vet aims to focus and report on suspicious constructs in code that could give rise to issues and bugs.
While vet does not guarantee correctness, it goes further than the compiler to detect problems ahead of runtime.

It can be invoked with `go vet ./...`

### References
1) [Hello World application (in playground)](https://play.golang.org/p/AH8TPSfGkFd)
2) [How To Write Go Code](https://golang.org/doc/code.html)
3) [Go Packages and Modules](https://golangbot.com/go-packages/)