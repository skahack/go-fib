# Fibonacci Generator

## Usage

```go
ch := fib.Fib(10)
<-ch // => 0
<-ch // => 1
<-ch // => 1
<-ch // => 2
<-ch // => 3
<-ch // => 3
```

```go

for v := range Fib(10) {
  fmt.Println(v)
}
```

## License

MIT
