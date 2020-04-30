# Description

An implementation of Go Concurrency Patterns

Output
```
> go run main.go

[Worker 2] Consuming slow task... Task5, it will take 5 seconds
[Worker 1] Consuming slow task... Task1, it will take 5 seconds
[Worker 4] Consuming slow task... Task3, it will take 5 seconds
[Worker 3] Consuming slow task... Task4, it will take 5 seconds
[Worker 0] Consuming slow task... Task2, it will take 5 seconds
-> Task: Task4 is completed, input = 40, output = 1600
-> Task: Task1 is completed, input = 10, output = 100
-> Task: Task5 is completed, input = 50, output = 2500
-> Task: Task2 is completed, input = 20, output = 400
-> Task: Task3 is completed, input = 30, output = 900
time elapsed 5.004547962s
```