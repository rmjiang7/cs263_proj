package main

import "flag"
import "fmt"
import "log"
import "os"
import "runtime"
import "runtime/pprof"
import "time"
import _ "net/http/pprof" // For setting up visualization server (?)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func ExpensiveFunc() {
	time.Sleep(1000 * time.Millisecond)
}


func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

	
	fmt.Println("Begin ExpensiveFunc.")
	start := time.Now()

	ExpensiveFunc()

	elapsed := time.Since(start)
	fmt.Println("End ExpensiveFunc.")
	log.Printf("ExpensiveFunc took %s", elapsed)
	

    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
        f.Close()
    }
}