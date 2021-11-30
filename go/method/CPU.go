package method

/*
#include <pthread.h>
#include <time.h>
#include <stdio.h>

long long getThreadCpuTimeNs() {
    struct timespec t;
    if (clock_gettime(CLOCK_THREAD_CPUTIME_ID, &t)) {
        perror("clock_gettime");
        return 0;
    }
    return t.tv_sec * 1000000000LL + t.tv_nsec;
}
*/
import "C"
import (
	"fmt"
	"strings"

	"github.com/klauspost/cpuid/v2"
)

func ExeAndPrintCPUusage(doWork func(int), num int) int64 {
	cputime1 := C.getThreadCpuTimeNs()
	doWork(num)
	cputime2 := C.getThreadCpuTimeNs()
	fmt.Println(cputime1)
	fmt.Println(cputime2)
	delta := cputime2 - cputime1
	fmt.Println(delta)
	return int64(delta)
}

// Show the CPU speed
func PrintCPU() {
	// Print basic CPU information:
	fmt.Println("Name:", cpuid.CPU.BrandName)
	fmt.Println("PhysicalCores:", cpuid.CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", cpuid.CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", cpuid.CPU.LogicalCores)
	fmt.Println("Family", cpuid.CPU.Family, "Model:", cpuid.CPU.Model, "Vendor ID:", cpuid.CPU.VendorID)
	fmt.Println("Features:", strings.Join(cpuid.CPU.FeatureSet(), ","))
	fmt.Println("Cacheline bytes:", cpuid.CPU.CacheLine)
	fmt.Println("L1 Data Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", cpuid.CPU.Cache.L1I, "bytes")
	fmt.Println("L2 Cache:", cpuid.CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", cpuid.CPU.Cache.L3, "bytes")
	fmt.Println("Frequency", cpuid.CPU.Hz, "hz")

	// Test if we have these specific features:
	if cpuid.CPU.Supports(cpuid.SSE, cpuid.SSE2) {
		fmt.Println("We have Streaming SIMD 2 Extensions")
	}
}
func main() {
	PrintCPU()
}
