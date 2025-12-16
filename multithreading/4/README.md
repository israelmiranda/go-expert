apache ab - single-threaded command-line utility for benchmarking and load testing HTTP web servers

$ brew install httpd
$ ab -V

$ ab -n 1000 -c 100 http://localhost:3000/

looking for race condition

$ go run -race main.go