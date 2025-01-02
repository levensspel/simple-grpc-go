# Get Started
This project is mostly initiated based on the official quickstart docs https://grpc.io/docs/languages/go/quickstart/

For more detailed steps to getting started, you can follow the quickstart docs in given link above, since this source code may continue to update for learning purpose.

This app uses `protoc` version 3. You need to ensure your Golang version is any one of the two latest major releases of Go as stated in the quickstart docs in the link above.

### Setup the Plugin for Protocol Compiler

#### 1. Install the protocol compiler plugins for Go using the following commands
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

#### 2. Update your PATH so that the protoc compiler can find the plugins
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

#### 3. Ensure the dependencies installation
```bash
go mod tidy
```

### Apply Proto Buffer Changes
If you want to apply some changes to the proto buffer definition (the `.proto` file), you will need to run the following commands to re-generate the gRPC code (the translated proto buffer into GoLang, eg. `yourdomain_grpc.pb.go` and `yourdomain.pb.go`):
```bash
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
proto/user/user.proto
```
*You may need to adjust the `proto/user/user.proto` name to your actual source path

### Run the App
```bash
go run ./server/main.go
```

```bash
go run ./client/main.go
```

