# Memory Cache

A zero-dependency cache library for storing data in memory with **generics**.

## Requirements

- Golang 1.18+

## Installation

```bash
go get -u github.com/rodrigo-brito/memory-cache
```

## Examples of usage

```go
type User struct {
    Name string
}

// Create a new cache for type `User` with a clean-up interval of 1 hour
cache := cache.New[User](time.Hour)

// Store a new user with key "user-1" and TTL of 1 minute.
cache.Set("user-1", User{Name: "Rodrigo Brito"}, time.Minute)

// Retrieve the user with key "user-1"
user, ok := cache.Get("user-1")
if ok {
    fmt.Println(user.Name) // output: "Rodrigo Brito"
}
```

## License

This project is release under [MIT license](LICENSE).
