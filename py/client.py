import socket
import time
import json
import utils
import random

HOST = "localhost"
PORT = 1966

with open(utils.pathGen("..", "datas.json"), 'r') as load_f:
    load_datas = json.load(load_f)

with open(utils.pathGen("..", "control.json"), 'r') as load_f:
    load_settings = json.load(load_f)

settings = []
for i in range(load_settings["total"]):
    settings.append(
        (load_settings[str(i)]["times"], load_settings[str(i)]["gap"]))

speed_list = []
networkdata_list = []
networkSend = 0
networkRec = 0
for i, pair in enumerate(settings):

    num, gap = pair
    start = time.time()
    for j in range(num):
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
            resNum = random.randint(0, int(load_datas["sum"])-1)
            s.connect((HOST, PORT))
            sendMsg = ('gemini://gemini.circumlunar.space/%s/\r\n' %
                       resNum)
            networkSend += len(sendMsg)
            s.sendall(sendMsg.encode())
            recMsg = s.recv(1024)
            networkRec += len(recMsg)
            if recMsg.decode()[3:-2] != load_datas[str(resNum)]:
                print(j)
                print(recMsg)
                raise Exception("Invalid res from req %s" % resNum)
        time.sleep(gap)
    end = time.time()
    total = end-start
    # print("client total send bytes: %s" % networkSend)
    # print("client total recev bytes: %s" % networkRec)
    # print("(network_data_size): %s" % (networkSend+networkRec))
    # print("client total request time: %s" % num)
    # print("client total time using: %s" % total)
    networkdata_list.append(networkSend+networkRec)
    speed_list.append([total, num])
    print("client no. %s test done" % i)
    if i != len(settings)-1:
        time.sleep(1.5)
    else:
        print("client ended")

print("netdata: ", networkdata_list)
print("speed: ", speed_list)
