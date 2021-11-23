import socket
import CPUusage
import time
import os

def server():
    HOST = "localhost"
    PORT = 1966

    # f = open("")
    # long running
    # do something other

    start = time.time()
    f = open(os.path.join("..", "README.md"))
    f.close()
    # long running
    # do something other
    end = time.time()
    print(end-start)

    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.bind((HOST, PORT))
        s.listen()
        cont = True
        # while cont:
        for i in range(5):
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
                print('request %s from %s' % (key, addr))
                meta = "context"
                conn.sendall(('20 %s\r\n' % (meta)).encode())


print(CPUusage.ExeAndPrintCPUusage(server))
