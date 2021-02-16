# gomock-mgo-example

An example of using [gomock](https://github.com/golang/gomock) for mocking the 
[mgo](https://pkg.go.dev/gopkg.in/mgo.v2) interfaces for unit testing.

For more details, take a look at the `repository/users_test.go` and the `Makefile`.

The mocks were generated using the `mockgen` tool, provided by `gomock`.
```sh
# Generating the mocks
$ make mocks
# Running the tests
$ make test
```
