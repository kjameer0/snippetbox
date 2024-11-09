# Chapter 2

Handlers are like controllers in that they execute application logic and write responses. They are the actors in the program that get back to the user.

A router(servemux) will map routes to handlers. One app will have one servemux for all routes.

A server can listen for requests and send responses.

You can pass strings to the first parameter of `ListenAndServe()` that can be looked up for their port.

You can run a go module by its module url(go run snippetbox.khalidjameer.com)

## Questions

- what exact type is a http.Request?
