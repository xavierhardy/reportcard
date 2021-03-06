[![Go Report Card](https://goreportcard.com/badge/xavierhardy/reportcard)](https://goreportcard.com/report/xavierhardy/reportcard) [![Build Status](https://travis-ci.org/xavierhardy/reportcard.svg?branch=master)](https://travis-ci.org/xavierhardy/reportcard) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/xavierhardy/reportcard/blob/master/LICENSE)

# Report Card

A fork of [Go Report Card](https://github.com/gojp/goreportcard). A web application that generates a report on the quality of an open source project. It uses several measures, including `gofmt`, `go vet`, `go lint` and `gocyclo` for Go projects. The objective is support several programming languages and formats. To get a report on your own project, run `reportcard-cli`.

### Installation

Assuming you already have a recent version of Go installed, pull down the code with `go get`:

```
go get github.com/xavierhardy/reportcard
```

Go into the source directory and pull down the project dependencies:

```
cd $GOPATH/src/github.com/xavierhardy/reportcard
make install
```

Now run

```
make start
```

and you should see

```
Running on 127.0.0.1:8000...
```

Navigate to that URL in your browser and check that you can see the front page.

### Command Line Interface

There is also a CLI available for grading applications on your local machine.

Example usage:
```
go get github.com/xavierhardy/reportcard/cmd/reportcard-cli
cd $GOPATH/src/github.com/xavierhardy/reportcard
reportcard-cli
```

```
Grade: A+ (99.9%)
Files: 362
Issues: 2
gofmt: 100%
go_vet: 99%
gocyclo: 99%
golint: 100%
ineffassign: 100%
license: 100%
misspell: 100%
readme: 100%
```

Verbose output is also available:
```
reportcard-cli -v
```

```
Grade: A+ (99.9%)
Files: 332
Issues: 2
gofmt: 100%
go_vet: 99%
go_vet  vendor/github.com/prometheus/client_golang/prometheus/desc.go:25
        error: cannot find package "github.com/prometheus/client_model/go" in any of: (vet)

gocyclo: 99%
gocyclo download/download.go:22
        warning: cyclomatic complexity 17 of function download() is high (> 15) (gocyclo)

golint: 100%
ineffassign: 100%
license: 100%
misspell: 100%
readme: 100%
```

### Contributing

Report Card is an open source project run by volunteers, and contributions are welcome! Check out the [Issues](https://github.com/xavierhardy/reportcard/issues) page to see if your idea for a contribution has already been mentioned, and feel free to raise an issue or submit a pull request.

### License

The code is licensed under the permissive Apache v2.0 licence. This means you can do what you like with the software, as long as you include the required notices. [Read this](https://tldrlegal.com/license/apache-license-2.0-(apache-2.0)) for a summary.

### Notes

We don't support it on Windows since we have no way to test it on Windows.
