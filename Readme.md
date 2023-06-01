# Iban Check Service

Exposes a virtual resource `/v1/iban/$iban`. 

## Example Usage
```
docker build . -t iban-check
docker run --rm -d -p 8080:8080 iban-check
curl http://localhost:8080/v1/iban/GB82WEST12345698765432
```

## Testing
Only have unit tests for now of iban domain
```
go test -v ./...
```
