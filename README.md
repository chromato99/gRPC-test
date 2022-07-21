# gRPC-test

This is my implementation of a gRPC server and client using golang.

When a client requests a random number with max number, the server returns a random int number between 0 and max.

### Build and Run

```zsh
# Build server
go build -o server ./server

# Build client
go build -o client ./client

# Run server
./server/server

# Run client
./client/client
```
