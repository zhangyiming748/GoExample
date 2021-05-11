go test -bench=.
go test -bench=BenchmarkMySquare
go test -bench=BenchmarkMyCube

go test -bench=. -benchmem -v -cover
go test -bench=BenchmarkMySquare -benchmem
go test -bench=BenchmarkMyCube -benchmem

#-bench regexp 执行相应的 benchmarks,例如 -bench=.
#-cover 开启测试覆盖率
#-run regexp 只运行 regexp 匹配的函数,例如 -run=Array 那么就执行包含有 Array 开头的函数;
#-v 显示测试的详细命令.