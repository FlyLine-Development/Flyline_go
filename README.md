# Flyline-go 

A Go client library for the [Flyline API](https://flyline.io/docs/).

## Table of Contents

- [flyline-go](#Flyline-go)
    * [Install](#install)
    * [Developing](#developing)
    * [License](#license)

## Install

```console
$ go get github.com/flyline-development/flyline-go
```



## Documentation

The module supports all Flyline API endpoints.

GoDoc: [![GoDoc](https://godoc.org/github.com/flyline-development/flyline_go?status.svg)](https://godoc.org/github.com/flyline-development/flyline-go/client)

## Getting Started

### Calling Endpoints

To call an endpoint you must set Token using `SetToken()` object.

```go
import (
    "net/http"
    "os"

    "github.com/flyline-development/flyline_go"
)


flyline := flyline.SetToken("test_####")
response := flyline.GetSeatTypes()

```

Each endpoint returns an object which contains the parsed JSON from the HTTP response.

## Developing

1. Download this repo into your Go source directory
2. Run `make setup` pull down all dependencies etc

## Support

Open an [issue](https://github.com/flyline-development/flyline_go/issues/new)!

## License
