import ctypes
import os

# gcc -o libpycall.so -shared -fPIC pycall.c


def ExeAndPrintCPUusage(f, num):
    ll = ctypes.cdll.LoadLibrary
    path, _ = os.path.split(os.path.abspath(__file__))
    lib = ll(os.path.join(path, "libpycall.so"))

    cputime1 = lib.getThreadCpuTimeNs()
    f(num)
    cputime2 = lib.getThreadCpuTimeNs()

    return cputime2-cputime1
    # return "CPU time = %d ns" % (cputime2-cputime1)
