# Part 2

In this part, we will build a package to encapsulate and fetch weather data from a third party API.

**What we'll cover:**
1) Types
2) Exporting
3) Errors

### Fetching weather data
In order to give accurate attire predictions, we need to first find out the actual weather in the area requested. To do this, we'll make use of a third party weather API to get the data we need.

Let's build a simple client that enables us to interact, and query the third party API. 
We'll create a new package to do this, to enable us to encapsulate and abstract the implementation away from others who need to get weather, so they don't have to care about the any of the details!

As the third party API supports REST, we'll make use of the `net/http` standard library to make the http requests, and then decode the expected response into a struct with the built-in `encoding/json` package.

First, let's define what we expect the response to be, into a struct. Structs are collections of related data in a single unit, useful for forming records. 
You can compare them to objects from other languages. Following the documented response in the introduction, we can define our response struct like follows:

```go
package weather

type Response struct {
		Main          string 
		Description   string 
		Location      string 
		TempMin       float32
		TempMax       float32
		WindSpeed     int    
		WindDirection string 
}
```

This is how we define types in go - the `type` keyword followed by a name and the definition (in this case a struct).

Inside the struct, we define all the fields, and their types. 
In addition, we can add additional tags to define properties, such as json key names when encoding / decoding.

```go
package weather

type Response struct {
		Main          string  `json:"main"`
		Description   string  `json:"description"`
		Location      string  `json:"location"`
		TempMin       float32 `json:"temp_min"`
		TempMax       float32 `json:"temp_max"`
		WindSpeed     int     `json:"wind_speed"`
		WindDirection string  `json:"wind_direction"`
}
```
Once we've defined our type, we can easily instantiate an instance of it: 
```go
// Create a new Response, and assign to the variable response
myResponse := Response {
	Main: "Some-Main"
	...
}

// Create a new response and assign the reference to responseRef
myResponseRef := &Response {
	Main: "Some-Main"
}

// Create a response variable that is unassigned 
var empty Response
```

While we can let consumers of our package create instances of the type like above (literally), it is common to define constructors to control the instantiation. We can define a function, called `New()` that will behave like a constructor:

```go
package weather

func New(main string) Response {
	return Response {
		Main: main,
	}
}
```

### Public / Private

Before we go any further - we should draw attention to an important part of naming things in Go. 
You may (or may not) have noticed the use inconsistent use of capitalization for naming throughout the session. 
Some things have been uppercase such as the `Response` struct, and it's constructor `New()`,  while the `main()` function, and it's helpers were lowercase. 
When do we decide what's the right casing?

In Go, the use of casing is what defines whether our type/func is package internal or external. 
If the type is uppercase, such as `Response`, or `New()`, then it is exposed, and consumers of the package can make use of it with `weather.Response{}` and `weather.New()`, however with lowercase values, is it not possible to use outside of the package.

When building packages in Go, we should think carefully about what consumers of our package *need* to be exposed, and what is internal implementation that shouldn't be exposed.

This makes it very simple in Go to understand if something is package private, or exposed - just by checking if it's uppercase!

An additional note, while function overloading is not possible, functions with the same name but in lower/uppercase are completely fine, and not uncommon:

```go
package weather

func GetWeather() Weather{
	return getWeather()
}

func getWeather() Weather {
	// This compiles just fine
	return Weather{}
}
```

The same applies to fields inside structs. Fields with uppercase identifiers are exposed, and can be referenced outside the package they are declared in, while those that are lower cannot:

```go
package weather

type Response struct {
	Main      string  `json:"main"`
	feelsLike string
}

...
package external

resp := weather.Response{}

resp.Main = "main" // Allowed
resp.feelsLike = "feels like" // Won't compile
```

When decoding json values into structs and vice versa, only public fields can be accessed by the `encoding/json` package to populate - hence the example above doesn't have a json tag, as it won't have an impact. 

### Errors
Nobody likes errors, but in Go, it's extremely common to deal with them! 
In Go, the error is just another built-in type: `error`. 
Errors are treated like this in Go with the thinking that: errors should always be considered, and are just another normal part of a flow - we should expect errors when they may occur, and have mechanisms in place to handle them when they do.

To give an example, let's take a look at a function that will return an error if the input is invalid:

```go
func CalculateTip(people int, price int) (int, error)
	if people < 1 {
		return 0, errors.New("no people")
	}
	return ((price/100) * 0.05)/people , nil
}
```
'Throwing' an error is just like returning any other value, we simply return the error value. The standard `errors` package gives us an easy way to create a new error.

The caller of this function then needs to handle this error in their flow! The standard approach for this is to try to handle it, or return an error up the stack.
(errors are commonly named `err` in Go)
```go
tip, err := CalculateTip(1, 100)
if err != nil {
	return err
}
```

What exactly is an error in Go? We know it's a just a type, but what kind? If we take a look, it is actually an interface (more on these later), requiring the function `Error()` returning a string. 
```go 
type error interface {
    Error() string
}
```
This means, if we wanted to, we could even define our own errors - to provide more context for [example](https://blog.golang.org/error-handling-and-go).
```go
type SyntaxError struct {
    msg    string // description of error
    Offset int64  // error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string { return e.msg }
```

### References
1) [Structs By Example](https://gobyexample.com/structs)
2) [Error Handling in Go](https://blog.golang.org/error-handling-and-go)
3) [Errors Package](https://pkg.go.dev/github.com/pkg/errors?tab=doc)