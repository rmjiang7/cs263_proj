# CS263 Project

by Brandon Huynh & Richard Jiang

## Project Vision

### Comparing the gccgo and gc compilers for GoLang

We plan to investigate and evaluate the effect of the using the two default GoLang compilers: gc and gccgo.
In particular, we will document the variety of the optimization options available using gccgo and indicate how they modify the runtime execution for some key features of the language.
We will empirically evaluate the tradeoff of compilation time and runtime aiming to understand when and where to use the gccgo compiler in place of the standard gc compiler.
One particular feature that we will explore more in depth is the ability to link C code, which Gccgo is reportedly better at, offering faster performance when calling C functions.
We plan to evaluate these claims and speculate on how they might affect the performance of applications that rely heavily on C libraries, such as Deep Learning and Graphics.

### Links and Resources

https://stackoverflow.com/questions/25811445/what-are-the-primary-differences-between-gc-and-gccgo

https://www.reddit.com/r/golang/comments/41dw90/gccgo_vs_go_calling_overhead_with_c_few_questions/

https://golang.org/doc/install/gccgo

## Documentation

### Installation

Follow [this setup guide](https://golang.org/doc/code.html).
Then:

```bash
cd $GOPATH/src/github.com/user
git clone https://github.com/rmjiang7/cs263_proj.git
cd cs263_proj
go install
```

### Profiling Code

**TODO:** Current code doesn't seem to run long enough to generate and profiling data...

How to generate profiling files:

```bash
cs263_proj -cpuprofile cpu.prof -memprofile mem.prof
```

How to view profiling files:

```bash
go tool pprof cpu.prof
```

[Link to more commands and examples](https://blog.golang.org/profiling-go-programs)

**TODO:** Figure out how to use webserver for visualizations? :

```go
import _ "net/http/pprof"
```
