# Go Collagen

[![Build Status](https://travis-ci.org/Szymongib/go-collagen.svg?branch=master)](https://travis-ci.org/Szymongib/go-collagen)

Collagen generates methods to operate on collection for a specific type.

## Usage

Collagen generated methods operates on collection type, which is type alias on a slice of structs.

For the following `Friend` struct:
```go
type Friend struct {
    Name string
    Age  int
}
```

The generated collection type will be:
```go
type Friends []Friend
```

To generate the collection type as well as all methods, run:
```
collagen --name Friend
```

or annotate the struct using `go generate`:
```go
//go:generate collagen --name Friend
type Friend struct {
    Name string
    Age  int
}
```
and run:
```
go generate
```

To specify the name of the collection type, use the `--plural` option:
```
collagen --name Friend --plural FriendCollection
```

## Collection methods

Collagen can generate the following methods:
- Contains
- Drop
- Exists
- Filter
- Map
- Take
 
To generate only subset of methods use `--functions` flag, with coma-separated names of the methods. 
For example, to generate only `Filter`, `Exists` and `Map` methods for the `Friend` struct, run:
```
collagen --name Friend --functions "Filter,Exists,Map"
```
