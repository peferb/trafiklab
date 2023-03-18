# Developer notes

_Just my own notes..._

## TODO
* Add support for "SL Route-planner v3.1" (in its own package to match api:s structure?)
* Make Api implement [Stringer interface](https://pkg.go.dev/fmt#Stringer)
* [Support debuging](https://medium.com/@cep21/go-1-7-httptrace-and-context-debug-patterns-608ae887224a) (?)
* Add func that takes a Writer as parameter? (for less GC...)
  * Does Writer interface allows HttpWriters to be passed?
* Option to get data as structs (?)
* Use [golangci-lint](https://golangci-lint.run/) (?)
* Use of /internal to hide public methods?
* Wrap errors...?
  * Avoid to creating lib specific errors of same reason of avoiding lib specific type safety?
  * Redundant wrapping them in such small lib?

## Links
* [Trafiklab login](https://developer.trafiklab.se/user/login)
* [Go modules publishing doc](https://go.dev/doc/modules/publishing)


* [Jack Lindamoods article "aspects of a good go library"](https://medium.com/@cep21/aspects-of-a-good-go-library-7082beabb403)
* [Is ExpVar intersting in some case?](https://sysdig.com/blog/golang-expvar-custom-metrics/)
* [Go Proverbs with Rob Pike @ Youtube](https://youtu.be/PAAkCSZUG1c)

## Stuff I keep forgetting

Run specific test by test func name:
```go
go test ./... -run TestRoutePlannerV3_1GisRouteResponseXml
```
