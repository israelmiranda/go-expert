run tests
$ go test

run tests verbose
$ go test -v

coverage %
$ go test -coverprofile=coverage.out

show coverage details
$ go tool cover -html=coverage.out

benchmark
$ go test -bench=.

$ go test -bench=. -run=ˆ#

$ go test -bench=. -run=ˆ# -count=10

$ go test -bench=. -run=ˆ# -count=10 -benchtime=3s

$ go test -bench=. -run=ˆ# -benchmem

fuzzing
$ go test -fuzz=. -run=ˆ#