# [grpcurl](https://github.com/fullstorydev/grpcurl)
`grpcurl` is pretty useful for testing:

## Make request
We can use a [heredoc](https://linuxize.com/post/bash-heredoc/) to pass multiple
lines of input to the `grpcurl` command. We use the `-plaintext` flag to disable
TLS. The `-d @` flag means we're piping data from stdin.
```bash
(
cat << EOF
{
  "location": {
    "latitude": 40.73010864595388,
    "longitude": -73.95094555260256
  }
}
EOF
) | grpcurl -plaintext -d @ localhost:8080 coop.drivers.dispatch.v1beta1.DispatchService/Dispatch
```

## Check App Health
```bash
grpcurl -plaintext -d '{"service":"foobar"}' localhost:8080 grpc.health.v1.Health/Check
```