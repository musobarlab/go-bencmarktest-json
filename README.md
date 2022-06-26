## Benchmark JSON CamelCase vs SnakeCase

### Run Benchmark Test
```shell
$ make test_bench
```

### Result From My Machine
```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i5-4258U CPU @ 2.40GHz
```

Result

```
go test . -bench=.
goos: darwin
goarch: amd64
pkg: github.com/musobarlab/go-bencmarktest-json
cpu: Intel(R) Core(TM) i5-4258U CPU @ 2.40GHz
BenchmarkSerializeDataCamelToJSON-4              1848588               610.9 ns/op
BenchmarkSerializeDataSnakeToJSON-4              1821445               657.6 ns/op
BenchmarkDeserializeDataCamelFromJSON-4           449808              2612 ns/op
BenchmarkDeserializeDataSnakeFromJSON-4           444752              2684 ns/op
PASS
ok      github.com/musobarlab/go-bencmarktest-json      7.516s
```

JSON with `camelCase` is faster than `snakeCase` in encode and decode process.


```
Test Name                               Operation (Higher better)      NS per Operation (Lower better)

BenchmarkSerializeDataCamelToJSON-4              1848588               610.9 ns/op
BenchmarkSerializeDataSnakeToJSON-4              1821445               657.6 ns/op
BenchmarkDeserializeDataCamelFromJSON-4           449808              2612 ns/op
BenchmarkDeserializeDataSnakeFromJSON-4           444752              2684 ns/op
```

#
In this example, the data is structured and uses the same data type, only the json field is different

Structure with CamelCase
```go
// DataCamel type
type DataCamel struct {
	ID                string   `json:"id"`
	NameNameNome      string   `json:"nameNameNome"`
	BlaBleBloQuantity uint64   `json:"blaBleBloQuantity"`
	Images            []string `json:"images"`
}
```

Structure with Snake_Case
```go
// DataSnake type
type DataSnake struct {
	ID                string   `json:"id"`
	NameNameNome      string   `json:"name_name_nome"`
	BlaBleBloQuantity uint64   `json:"bla_ble_blo_quantity"`
	Images            []string `json:"images"`
}
```