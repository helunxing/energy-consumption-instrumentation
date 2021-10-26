#include <pthread.h>
#include <time.h>
#include <stdio.h>
#include <stdlib.h>

long long getThreadCpuTimeNs()
{
    struct timespec t;
    if (clock_gettime(CLOCK_THREAD_CPUTIME_ID, &t))
    {
        perror("clock_gettime");
        return 0;
    }
    return t.tv_sec * 1000000000LL + t.tv_nsec;
}