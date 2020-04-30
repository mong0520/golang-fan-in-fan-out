# Description

An implementation of Go Concurrency Patterns

Output
```
> go run main.go

[Worker 2] Consuming slow task... Task 1, it will take 5 seconds
[Worker 1] Consuming slow task... Task 4, it will take 5 seconds
[Worker 0] Consuming slow task... Task 2, it will take 5 seconds
[Worker 3] Consuming slow task... Task 5, it will take 5 seconds
[Worker 4] Consuming slow task... Task 3, it will take 5 seconds
-> Task 5 is completed, input = 5, output = 25
-> Task 4 is completed, input = 4, output = 16
-> Task 3 is completed, input = 3, output = 9
-> Task 1 is completed, input = 1, output = 1
-> Task 2 is completed, input = 2, output = 4
time elapsed 5.004699585s
```