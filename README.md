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
