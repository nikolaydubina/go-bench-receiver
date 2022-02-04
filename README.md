## What is more efficient value or pointer method receivers?

### A) Why would you think `struct` is more efficient?

You are making single call to fetch struct and its fields.

### B) Why would you think `pointer` is more efficient?

You are not copying unnecessary data, only pointer is passed to function.

### Result

Overall, difference is negligible.
Difference is in order of few ns, which is dominated by logic of function (typically hundreds of ns).


- small structs >>> structs few ns faster
    - Subject to codebase, but this is one of common situations. The effect is unnoticeable and dominated by business logic.
- large structs are deeply nested >>> pointers is few hundred ns faster
    - You have to have very large and nested structs to see difference. This is only case when pointe is noticeably faster. Difference can be big. This is effect B).
- small structs and deeply nested >>> structs is few ns faster
    - You have to have very deep nesting to see difference to observe this effect. This is effect A).


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
