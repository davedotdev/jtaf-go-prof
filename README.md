# Instructions

Sometimes when working on Terraform providers, you want to profile CPU and memory to figure out what's actually going on. Things are as never as they seem.

In `main.go` there is some profiling code, which generates two files per invocation:

-	a file with `_cpuprof` as a postfix containing CPU pprof data
- 	a file with `_memprof` as a postfix containing memory pprof data

It's possible to open those files and explore them in a simple web tool with:

```bash
go tool pprof -http=":8000" <specific_file>_cpuprof
```

Simples.
