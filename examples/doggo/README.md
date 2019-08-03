# Doggos

Simple application that manages collection of Doggos.

## Run the example

### Generate collection code

To run the application, generate collection methods with `collagen`. You can use it directly by running:
```
collagen --name Doggo
```
or by using `go generate`:
```
go generate
```

You should see that `Doggos` has been generated.
```
Doggos generated with methods.
```

### Run the tests

Now you can run the tests:
```
go test ./...
```

### Run example code

To run the example application, type:
```
go run .
```

