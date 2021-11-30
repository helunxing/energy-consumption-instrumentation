import random
import json
import utils


def ranstr(num):
    H = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz0123456789'

    res = ''
    for _ in range(num):
        res += random.choice(H)

    return res


dataSum = 100
data_dict = {}
data_dict['sum'] = str(dataSum)
dataMaxLen = 1000

for i in range(dataSum):
    data_dict[i] = ranstr(random.randint(0, dataMaxLen))


with open(utils.pathGen("..", "datas.json"), "w") as f:
    json.dump(data_dict, f)
