[![Go Reference](https://pkg.go.dev/badge/github.com/eighty4/ebauche.svg)](https://pkg.go.dev/github.com/eighty4/ebauche)

# Èbauche

```
\ āˈbōsh \
plural -s
: an incomplete watch movement consisting of plates,
  bridges, wheels, and barrels to be finished and fitted
  with jewels, escapement, mainspring, hands, and dial
```
Time-related fn utils

```
go get github.com/eighty4/ebauche
```

## IntervalIndefinitely

Executes fn indefinitely on interval

```go
ebauche.IntervalIndefinitely(time.Second * 1, func() {
	fmt.Println("i print every second")
})
```

## Interval

Executes predicate indefinitely on interval until it returns false

```go
limit := 84
count := 0
fmt.Println("i will print", limit, "times")
ebauche.Interval(time.Millisecond * 1, func() bool {
    fmt.Println("print", count)
    count++
    return count < limit
})
```
