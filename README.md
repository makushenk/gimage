# Clean Architecture on Go

### About
This is an attempt to apply principles described in the Robert Martin's [article](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)


### Tests


To run tests:
```
go test ./... -coverprofile=test.coverage
```

To see test coverage:
```
go tool cover -html=test.coverage 
```
