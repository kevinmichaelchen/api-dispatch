```bash
# Get nearest drivers to a specified pickup (passenger) location
go run cmd/dispatch/*.go nearest drivers --latitude 40.73010864595388 --longitude -73.95094555260256

# Get nearest trips to a specified driver location
go run cmd/dispatch/*.go nearest trips --latitude 40.73010864595388 --longitude -73.95094555260256
```