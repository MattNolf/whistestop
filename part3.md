# Part 3

Now it's time to implement some business logic, so we can codify our attire rules!

**What we'll cover:**
1) Interfaces
2) Testing

### An Attire Service
We should now build out our functionality to take some weather data, and apply our rules to determine what kind of attire should be worn.
To do this, we need to create a type that has access to the weather client we built in part 2.

Let's define a new struct for our service, and give it a weather client like so:
```go
type Service struct {
    client weather.Client
}
```

This will now mean any methods we call on our service, will have access to the weather client so that we can use GetWeather() method. 
This is great, but it builds a direct dependency on the client we built accessing the third party API. 
What if we decide in future we want to use a different weather API? Then we'd need to update our code to support that new implementation.  

What would be even nicer instead however, would be to instead specify the functionality we need from the client, rather than caring about specific concrete implementations about which API we're using.
To keep good isolation between our packages and limit the direct dependencies - we should define our parameters as "anything that can get me the weather given a location". 
To do that, we can use interfaces to define the functionality we need!

### Defining Interfaces

Interfaces in Go allow us to define behaviour in the form of function signatures, and we can use them as direct replacements for concrete values - so we'll often see them in function signatures, and struct definitions too. 
Let's take a look at defining an interface for our weather client. 
```go
type Forecaster interface {
    GetWeather(location string) (weather.Weather, error)
}
```
This means that any struct (implementation) that implements the `Forecaster` interface can give us the `GetWeather(...)` function!

The nice thing about Go interfaces, is that implementation of interfaces is implicit! 
What this means is the weather client we built in part 2 will automatically implement the interface without 
explicitly defining that behaviour (so we don't have to go back and change any of it).

As a result of this, it's good practice for consumers of the client to define an interface with just the functionality they need, 
keeping interfaces tight and limited to exactly what the client needs (and nothing more). This differs from other languages where
interfaces with methods not needed by a client end up being reused to avoid duplication.  

We can now use the interface in our struct like so:
```go
type Forecaster interface {
    GetWeather(location string) (weather.Weather, error)
}

type Service struct {
    forecaster Forecaster
}
```

### Service Method
Much like in object oriented languages, objects have methods attached to them, that can be invoked on them.
Go also supports such functionality, allowing us to define functions on structs, called methods. The syntax
for this involves prepending the struct to the declaration - like so:

```go
// A method with a service receiver 
func (s service) Recommend(location string) (Attire, error) {
    //...
}

// A method with a pointer to a service as a receiver - allowing us to modify s (the service) 
func (s *service) Recommend(location string) (Attire, error) {
    //...
}
```

This function is a method on the service struct, and could be invoked on a service value with dot notation:
`service.Recommend(someLocation)`. 
Note the receiver of the function - it's the service struct we defined earlier and (much like the param)
definition, we can specify a name for the argument. 
In this case we're specifying a value, however it's also possible to define methods on pointers too! Go automatically handles de-referencing so dot notation
behaves the same both with pointers and values.

Now let's go ahead and implement the `Recommend(location string)` function to satisfy the requirements documented in the introduction.

### Testing our rules
How can we be sure we accurately implemented the rules defined in the spec? The most effective way to check this is by testing it!

In Go, testing is easy, as it has great built in support to write, run and package tests - let's take a look!

If we create a function with that starts with *Test*, and it takes a single parameter - `testing.T` - then Go will automatically know that what we've defined is a test! 
This means it won't be included in any build, and can be triggered with the `go test` command (more on this later).

It was mentioned before that directories of .go files in Go must all be of the same package. 
Well, there is one exception to this rule, which is a test package - which follows the pattern `packagename_test`.

This is welcome, as it means that we are testing the package from an external point of view, and won't be testing the internals (as we're not in the same package, we can't access non-exposed functions/types/values).
This is known as black box testing.

```go
func Test_Recommend(t *testing.T) {
  // Make our assertions
}
```

The `testing.T` pointer gives us access to do various things during our test such as error, fail and skip tests just by invoking its methods.
It's automatically provided during invocation with the `go test` command. 

### Mocking Dependencies
Because we enabled our service struct to receive an interface instead of a concrete type, it means we can easily create a mock weather client
for use within our tests. This will enable us to provide certain weather data and verify the rules are satified!
Let's define our mock below:

```go
type weatherMock struct{}

func (wm weatherMock) GetWeather(location string) (weather.Weather, error) {
	if location == "Manchester" {
		return weather.Weather{
			Weather: "Rain",
			Temp:    21.0,
		}, nil
	}
	return weather.Weather{}, nil
}
```

### Defining Tests
Writing tests in Go can take a few different forms, as our testing argument `t` enables us to run child tests and define a hierarchy.
This is especially useful in cases where we can avoid setting up the same test cases multiple times and instead reuse parts of other tests.

Let's see what that looks like:
```go
func Test_Recommend(t *testing.T) {
	// Perform initial setup for tests
	t.Run("run test 1", func(t *testing.T) {
	    //	Perform test using initial setup
	})
	t.Run("run test 2", func(t *testing.T) {
	    //	Perform test using initial setup
	})
}
```
One thing worth mentioning here, is that the tests we define should always be independent; meaning they don't rely on any other tests to modify or set
state to make the necessary assertions.

If we find that all our tests are almost identical, with the only difference being the expected input / output - then this might be a good time to leverage test tables!
Test tables enable us to define all our inputs, and outputs - and then iterate over each case in the table and run the test. 
This makes our tests very succinct while also covering all the edge cases.

```go
t.Run("should return jacket in correct conditions", func(t *testing.T) {
	// Define our table
    table := getTable()

    for _, testCase := table {
        // Perform test with the values from our test case
    }
}
```

### Running our Tests
Once we have defined our tests, how do we actually run them?

The go toolchain makes it super easy to run our tests. Let's take a look at the different ways it supports:

```
go test ./...        # Run all tests.
go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
```

There are some additional flags that we can make use of:

The race flag will work to verify there are no race conditions in code that makes use of concurrency primitives:
```
go test ./... --race 
```

The count flag enables us to specify the number of times the tests run, to check for flaky tests for example:
```
go test ./... --count=100
```

### References
1) [How to Use Interfaces in Go](https://www.digitalocean.com/community/tutorials/how-to-use-interfaces-in-go)
2) [Using Go Interfaces](https://blog.chewxy.com/2018/03/18/golang-interfaces/)
3) [Writing Unit Tests](https://blog.alexellis.io/golang-writing-unit-tests/)