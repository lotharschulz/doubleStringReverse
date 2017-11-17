# golang double string reverse

## install
```
GOPATH=$(pwd)
echo $GOPATH
export PATH=$PATH:$GOPATH
cd src/myStringUtil
go install
```

## run
```
cd .. # (in src directory)
go run main.go
```

## test
```
go get -t -v ./... # get dependencies
go test -v ./myStringUtil # (from src directory)
```
additional [go test command flags](https://golang.org/cmd/go/#hdr-Description_of_testing_flags)
