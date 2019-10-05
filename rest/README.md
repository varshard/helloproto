# REST HTTP gateway for gRPC

Following https://github.com/grpc-ecosystem/grpc-gateway.
Use a HTTP server as a reverse proxy that will proxy HTTP request to gRPC.

## Prerequisite

* [google.api.http annotation](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto)

## Generate proto stub

```bash
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:./addressbook \
./addressbook.proto
```

## Generate proto stub gateway

```bash
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:./addressbook \
./addressbook.proto
```
