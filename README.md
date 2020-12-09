# Clean Architecture on Go


### Tests


To run tests:
```
go test ./... -coverprofile=test.coverage
```

To see test coverage:
```
go tool cover -html=test.coverage 
```