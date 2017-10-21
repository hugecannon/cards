# Cards

A golang playing cards library.

Sets of cards/hands/deck are backed by a single int64. The primary aim being to investigate how to implement card games with only bitwise operations.

## Getting Started

Simply go-get the package to use as a library

```
$ go get github.com/hugecannon/cards
```

### Prerequisites

Developed and tested on `go1.8.1`


## Running the tests

```
$ go test ./...
```


## Built With

* [Golang 1.8.1](https://golang.org/dl/)


## Authors

* **Hugh Cannon** -  [hugecannon](https://github.com/hugecannon)

See also the list of [contributors](https://github.com/hugecannon/cards/contributors) who participated in this project.

## TODO

- Complete godoc and improve with examples
- Complete poker unit tests
- Improve bitwise-yness of some of the poker package functions. (Not happy about Take/Peek/Count usage)
- Implement poker server
    - Dockerize poker server

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
