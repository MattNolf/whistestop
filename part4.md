# Part 4

Now it's time to put it all together and expose our work over http!

**What we'll cover:**
1) HTTP servers

### Serving traffic over http
To enable our service to serve traffic, and expose our API - we need to build a transport layer to accept requests, and handle them appropriately by doing the right thing.

In Go, we have lots of options to help us achieve this, but we'll be leaning on the standard library to get the job done. The standard library in Go is rich, and this gives us a good oppertunity to take a look.

Head over to the public godoc reference for the http package [here](https://golang.org/pkg/net/http/).

The http package provides lots of functionality, for both clients and servers, and for now we'll be focusing on building and serving traffic.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", someHandler)

    http.ListenAndServe(":8080", nil)
}
```

In this example, we are registering the path `/` as a route, and telling the router to invoke the handler `someHandler` with the request. 

Next, we call the function `ListenAndServe` to listen to traffic at address `:8080`. The final argument, is a handler interface for a server, which is often left `nil`.

To add an additional path to our server, we just need to attach the handler function to the path, in the exact same way we did with the root 
- the multiplexer will then match the request against the pattern that matches most closely.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", someHandler)
    http.HandleFunc("/another", someOtherHandler)

    http.ListenAndServe(":8080", nil)
}
```

Something we can observe here is that we are providing a function as an argument in the function `HandleFunc`. Let's take a look at the function signature to see what it expects:
```go
HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

As we can see, it takes 2 parameters the second being a function. Functions are first class citizens in Go, and can be treated as values, it's perfectly acceptable (and often idiomatic) to provide an anonymous function here like so:
```go
// Provide an anonymous func
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
})

// Assign a function to the variable myFunc
myFunc = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
})
http.HandleFunc("/", myFunc)
```

There are two parameters to the function - firstly `w` which is the response writer. 
This argument is used to construct a response to the original request, such as setting the body and response code.
Secondly, `r` is the original request - where we can extract values from the request, such as query parameters, and the body (if there is one).

We'll need to use both in our case! The request will give us the query param for the location - and then we can use the writer to give our attire recommendations!   

Next, let's add the routes we need for the specification we have been given.


### Returning the Attire
Now we need to return the `attire` struct we have received from our code. To do this, we can make use of the `encoding/json` package to 
marshall our response.

We can achieve this with the following: 
```go
json.NewEncoder(w).Encode(attire)
```
This creates a new json encoder, and give it the response writer to encode into; and then immediately calls the `Encode` function with our
struct! Nice!

Without setting any http status code, the default is set to 200 OKAY. In the case of any errors we encounter along the way, we can easily set the header code
with `w.WriteHeader(http.StatusInternalServerError)`

### Starting our Server
So it looks like we now have everything we need:
1) A client to talk to our weather API
2) A business layer to determine the best attire
3) A transport layer to serve traffic and call our service

All we need to do now, is head back to our main function, and tie it all together!

### Further Reading
[Go HTTP Server](https://gowebexamples.com/http-server/)
[Writing Web Applications](https://golang.org/doc/articles/wiki/)
[Advice for Maintainable Go Programs](https://dave.cheney.net/practical-go/presentations/qcon-china.html)
[Gorilla Mux HTTP package](https://github.com/gorilla/mux)