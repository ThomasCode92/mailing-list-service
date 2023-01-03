# MailingList Microservice

This application, a Mailing list service, is written in [Go](https://go.dev/) and it utilizes some of Go's advanced features (like Goroutines, WaitGroups, ...)!<br />
It is a project of the _[Go Programming: The Complete Developer's Guide](https://www.udemy.com/course/go-programming-golang-the-complete-developers-guide/)_ course from [ZTM](https://zerotomastery.io/).

## Setup

This project requires a `gcc` compiler installed and the `protobuf` code generation tools.

### Install protobuf compiler

Install the `protoc` tool using the instructions available at [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/).<br />
Alternatively you can download a pre-built binary from [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases) and placing the extracted binary somewhere in your `$PATH`.

### Install Go protobuf codegen tools

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`<br />
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

### Generate Go code from .proto files

```
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  Proto/mail.proto
```
