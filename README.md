Tratodo - is a simple todo application written in go and svelte. It was created for educational purposes and represents my first conscious practice of writing code in svelte and go.

<img width="1279" alt="Screenshot 2024-07-02 at 22 25 17" src="https://github.com/PillowPowa/tratodo/assets/97808006/eac029c7-3de0-479e-8995-86e0d5d1901a">

## Server side: 

The server side is a REST API and consists of a simple jwt authorization and a module with todo.
It also contains generated documentation with the swag tool.

In general, the server development was not difficult, go was extremely simple,
I just read the go tour, read what a [good structure](https://github.com/golang-standards/project-layout) should look like and that's it.

The server itself was written using `net/http` and a router `gorilla/mux`.
As a databse was used sqlite, cuz it is extremely easy to configure.

For migrations was used migrator, and it was isolated into a separate "app" (yes, migration tool without cli, weird, but interesting).

In general, I can note two things that surprised me:
1. Implementation of model updates (patch rest method) in go
2. Errors are values 
if err != nil { return err } 😄

As a TS maintainer, this seemed a bit strange to me, but I can't say that I didn't like it, it's pretty cool.
Also, it may seem strange to go guru that handlers return errors.
I get this idea from youtuber "Anthony GG" and it seemed convenient to me, although I read articles that this method is bad for middlrewares organization

In this project, I also tried to use yaml files as a configuration, which I didn't like, generally .env is the most convenient method. ;)

## Client side:

The client side, the implementation of the REST API, was written in svelte, why not in sveltekit?
idk, for some reason I thought that sveltekit was the same as next/remix for react.
I won't say that it would be difficult to transition, I just didn't do it.

In general, after react and angular, svelte seemed to me extremely yummy 😋, compilation (ahh 🤩), the bundle is extremely small, everything is fast thanks to vite (perfect 👌). 

Svelte hasn't become difficult for me as a react dev, the syntax is similar to JSX (the markup itself and the component approach). `svelte/store` is super cool, svelte reactivity is also great. (But I would still prefer react useEffect to [svelte's reactive guidelines](https://svelte.dev/examples/reactive-statements)). One thing I didn't like was the [event forwarding](https://svelte.dev/examples/event-forwarding), it's extremely "eoooww". Another thing that I didn't implement as I wanted was optimistic data mutation (it works, but it looks weird and feels like a bad thing to me).

Anyway, I should try the `svelte-kit` to feel the full power of svelte
