# README for [fgg](https://github.com/pldi-submission-448/fgg)

---
This implementation is forked from https://github.com/rhu1/fgg.

## Install
```
GO111MODULE=off go get github.com/pldi-submission-448/fgg
```

## Translating dispatcher.fgg
```shell script
go run . -fgg -dynamic -- dispatcher.fgg
go run . -fgg -dictpass -- dispatcher.fgg
go run . -fgg -dictpassnoassertion -- dispatcher.fgg
```

##Micro Benchmark 

We assume the GOPATH is `~/go`.

After running `go install .`, there should be an executable file 
`~/go/bin/fgg`.
Then use the shell scripts in `bencmark` directory.
For example, the command below executes the micro benchmarks
ten times for each configuration, and prints the execution time.

```shell script
./benchmark/microbenchmarks/exec_time_no_assertion.sh
```
