# Chapter 2

Handlers are like controllers in that they execute application logic and write responses. They are the actors in the program that get back to the user.

A router(servemux) will map routes to handlers. One app will have one servemux for all routes.

A server can listen for requests and send responses.

You can pass strings to the first parameter of `ListenAndServe()` that can be looked up for their port.

You can run a go module by its module url(go run snippetbox.khalidjameer.com)

Routes that end in a trailing slash only get their handlers run if there is an exact url match in the request url. So /snippet/view is only matched by a request to /snippet/view, not /snippet/view/. Trailing slash routes are called a subtree path pattern.

Avoid structural conflicts with routes as much as possible

`{$}` makes a route not match any subtree patterns

HEAD requests are used to get metadata like file size about a GET route.
The Allow header lets you specify what methods are allowed to hit a URL.
Always add methods to routes to avoid overlap
w.WriteHeader can be called once per response

## Interfaces

(https://www.alexedwards.net/blog/interfaces-explained)

The following quote is from page 41 of "Let's Go"
"because the http.ResponseWriter value in your
handlers has a Write() method, it satisfies the io.Writer interface... But at a practical level, it means that any functions where you see an io.Writer parameter, you can pass in your http.ResponseWriter value and whatever is being written will subsequently be sent as the body of the HTTP response."
io.Writer is an interface that implements one method that returns a string, Write, If another type has a Write method that returns a string, it satisfies or implements the io.Writer interface. Any function that takes an io.Writer interface type as a parameter can be used by any type that has a Write method.

The internal directory code can only be imported by folders inside parent folder. Other projects that download your code can't use the code in the internal folder.

I learned that malicious actors can move through my file system with relative paths if path names coming from users aren't cleaned.
## Questions

- what exact type is a http.Request?
