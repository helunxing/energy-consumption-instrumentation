import socket
import CPUusage
import time
import os
import utils
import json

IO_time_list = []
CPU_time_list = []


def server(num):
    HOST = "localhost"
    PORT = 1966

    start = time.time()
    f = open(utils.pathGen("..", "datas.json"), 'r')
    load_datas = json.load(f)
    f.close()
    end = time.time()
    # print("reading datas file using time (IO_time): %s" %
    #       (end-start)*pow(0.1, 6))
    IO_time_list.append((end-start)*pow(10, 6))

    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        s.bind((HOST, PORT))
        s.listen()
        for _ in range(num):
            conn, addr = s.accept()
            with conn:
                data = conn.recv(1050)
                if not data:
                    break

                tmp = data.decode()
                end = tmp.find('\r\n')
                url = tmp[:end]
                paths = url.split('/')
                if not paths[-1]:
                    paths.pop()
                key = paths[-1]
                # print('request %s from %s' % (key, addr))
                meta = "context"
                conn.sendall(('20 %s\r\n' % (load_datas[key])).encode())
        s.close()


with open(utils.pathGen("..", "control.json"), 'r') as load_f:
    load_settings = json.load(load_f)

settings = []
for i in range(load_settings["total"]):
    settings.append(
        (load_settings[str(i)]["times"], load_settings[str(i)]["gap"]))


for i, pair in enumerate(settings):
    num, _ = pair
    cpu_time = CPUusage.ExeAndPrintCPUusage(server, num)*0.001
    # print("CPU using time (CPU_time): %s" % cpu_time)
    CPU_time_list.append(cpu_time)
    print("python server no. %s test done" % i)
    if i != len(settings)-1:
        # print("sleep for 1 second, wait the port")
        time.sleep(1)
    else:
        print("server end")

print("IO time: ", IO_time_list)
print("CPU time: ", CPU_time_list)
