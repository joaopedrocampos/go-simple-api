[![codecov](https://codecov.io/gh/joaopedrocampos/go-simple-api/branch/main/graph/badge.svg?token=485T4O7H9W)](https://codecov.io/gh/joaopedrocampos/go-simple-api)

# Go Simple API

The goal of this API is to start learning Golang

## Start application

You can start the application by running the following command on application root folder:

```
$ go run cmd/server/main.go
```

## Build

If you want to compile and generate application executable you can run:

```
$ go build cmd/server/main.go
```

And run it with:

```
$ ./main
```

## Tests

Tests can be ran with the command:

```
$ go test ./...
```

If you want the tests to run verbosely just add the `-v` flag to the command

Sequence diagrams will be generated automatically inside the `.sequence` folders and to see it you just need to open these files on any web browser.

## Coverage

To generate coverage report just run the test command with the `-race`, `-covermode=atomic` and `-coverprofile=coverage.out` flags

```
$ go test ./... -race -covermode=atomic -coverprofile=coverage.out
```

With this the `coverage.out` report gonna be generated and you can render it on your browser automatically with

```
$ go tool cover -html=coverage.out
```
