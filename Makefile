all: clean  yaccdata build test

test:
	go test -v
build:
	go build
clean:
	rm -f ./goyaccidl.y.go

yaccdata:
	goyacc  -o ./goyaccidl.y.go  -p "YaccIdl"  ./goyaccidl.y
