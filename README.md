## What is more efficient value or pointer method receivers?

A) Why would you think struct is more efficient?

You are making single call to fetch struct and its fields.
Especially given that interce is a struct-wrapper around pointer with type.

B) Why would you think pointer is more efficient?

You are not copying unnecessary data, only pointer is passed to function.
Especially given that pointer is always same size.

## Result

Overall delta is negligible.

- small structs >>> structs few ns faster
    - Subject to codebase, but this is one of common situations. The effect is unnoticeable and dominated by business logic.
- deeply nested large structs >>> pointers is few hundred ns faster
    - You have to have very large and nested structs to observe this effect. This is only case when pointer is noticeably faster. Difference can be big. This is effect B)
- deeply nested small structs >>> structs is few ns faster
    - You have to have very deep nesting to observe this effect. This is effect A)

```
$ GOMAXPROCS=1 go test -timeout=1h -bench=. -benchtime=10s -benchmem ./...
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/go-bench-receiver
Benchmark_BasicServiceStruct                      	30910656	       373.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_BasicServicePointer                     	32115787	       374.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServiceLightStruct_UpStruct          	31909561	       374.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServiceLightStruct_UpPointer         	32234484	       374.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServiceLightPointer_UpStruct         	32067454	       374.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServiceLightPointer_UpPointer        	31839183	       375.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServiceLightStruct_UpStructMany_10   	30359691	       394.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServiceLightPointer_UpPointerMany_10 	30688312	       393.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServiceLightStruct_UpStructMany_50   	16043655	       748.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServiceLightPointer_UpPointerMany_50 	15883161	       750.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServiceStruct_UpStruct               	31512231	       377.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServiceStruct_UpPointer              	31865308	       377.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServicePointer_UpStruct              	31695372	       375.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServicePointer_UpPointer             	32216876	       374.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServiceStruct_UpStructMany_10        	23444461	       513.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_DepServicePointer_UpPointerMany_10      	30453177	       392.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_EmptyServiceStruct                      	 2979396	      4039.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_EmptyServicePointer                     	 2984328	      4034.0 ns/op	       0 B/op	       0 allocs/op
```

## Why is the result this way?

Looks like sturct/poitner does not matter much to Go runtime.
It must be using some lookup table based on type, regardless if it pointer or struct type, for methods of type.
Similarly, interface must be resolved to concrete method too.
Confirming this with Go source code is area of further research.
