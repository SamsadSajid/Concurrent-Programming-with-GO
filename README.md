# `async_service_call_v1.go` :

Make a netowrk call for each  `gihubUserName` and show the execution time. The network calling is not asynchronous. After one call finishes, the next call takes place. As these calls doesn't have any dependency, they should be executed parallely. `async_service_call_v2.go` solve this problem.

# `async_service_call_v2.go`
Does the same above work but the implementation is abit diffetent. This one use concurrency to make concurrent api call so that the network call can take place concurrently.

# Comparison of two approach
|Function Name|Number of Network call|Avg Time Taken|
|---|---|---|
|async_service_call_v1.go|2 Network call|2.135790661|
|async_service_call_v2.go|2 Network call|1.382750029|

# Result
V2 is 64.74% faster than V1

# Area of improvement for v2
As these network call doesn't depend on one another, we can parallely execute the network call which will concurrently execute.
Look the codes in `async_service_call_v3.go`

# Benchmark for V3
|Number of Core|Number of Network call|Avg Time Taken|
|---|---|---|
|2|2 Network call|1.260203267|
|4|2 Network call|1.135379339|

My device run with 4 core.
For 4 core, the time taken 1.135379339 which is 82.11% faster than v2