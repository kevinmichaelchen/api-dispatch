We use [Buf](https://buf.build/) for all things proto.

## Install Buf
Generally it's good to have it installed on your local machine:
```bash
go install \
  github.com/bufbuild/buf/cmd/buf@latest \
  google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
```

Because gRPC is not the most debuggable protocol, we spin up a 
[connect-go](https://connect.build/docs/go/getting-started/) server on port 8081.
This lets you use curl (instead of [grpcurl](./grpcurl.md)). 

## Get nearest drivers
```bash
curl \
    --header "Content-Type: application/json" \
    --data '{"limit": 5, "pickup_location": {"latitude": 40.73010864595388, "longitude": -73.95094555260256}}' \
    localhost:8081/coop.drivers.dispatch.v1beta1.DispatchService/GetNearestDrivers | jq
```

## Get list of drivers
```bash
curl \
    --header "Content-Type: application/json" \
    --data '{"page_size": 1}' \
    localhost:8081/coop.drivers.dispatch.v1beta1.DispatchService/ListDrivers | jq
```