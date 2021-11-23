import socket
import os
import time
import json
import utils

# with open(os.path.join("..", "README.md")) as f:
#     pass


HOST = "localhost"
PORT = 1966

with open(utils.pathGen("..", "datas.json"), 'r') as load_f:
    load_dict = json.load(load_f)
    print(load_dict['1'])
    

settings = [(5, 0.001)]
for pair in settings:
    num, gap = pair
    start = time.time()
    for _ in range(num):
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            s.connect((HOST, PORT))
            s.sendall(b'gemini://gemini.circumlunar.space/docs/\r\n')
            data = s.recv(1024)
            print(data.decode())
        time.sleep(gap)
    end = time.time()
    total = end-start
    print(total)
