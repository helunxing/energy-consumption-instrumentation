import CPUusage

def doWork():
    ans=0
    for i in range(100):
        ans+=i

CPUusage.ExeAndPrintCPUusage(doWork)
