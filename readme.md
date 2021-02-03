### Run all test function on package
```go test -v```

### Run all test in folder
```go test -v ./...```

### Run only single test function on package
```go test -v -run=TestHello```

### Run only single subtest function on package
```go test -v -run=TestHello/subtest1```

### Run all subtest with specified name function on package
```go test -v -run=/subtest1```

### Run unit test with benchmark
```go test -v -bench=```

### Run benchmark without unit test
```go test -v -run=NotMatchUnitTest -bench=BenchmarkTest```

### Run all benchmark in folder
```go test -v -run=NotMatchUnitTest -bench=. ./...```
