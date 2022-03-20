# Memory Cache

A zero-dependency cache library for storing data in memory with **generics**.

## Installation

```bash
go get -u github.com/rodrigo-brito/memory-cache
```

## Examples of usage

```go
type User struct {
	Name string
}

cache := cache.New[User](time.Hour)
cache.Set("user-1", User{Name: "Rodrigo Brito"}, time.Minute)

user, ok := cache.Get("user-1")
if ok {
    fmt.Println(user.Name) // output: "Rodrigo Brito"
}
```
