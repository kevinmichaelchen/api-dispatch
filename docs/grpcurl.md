[grpcurl](https://github.com/fullstorydev/grpcurl) is pretty useful for testing:
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