import socket

HOST = "localhost"
PORT = 1966

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.bind((HOST, PORT))
    s.listen()
    while True:
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
            print('request %s from %s'%(key, addr))
            meta="context"
            conn.sendall(('20 %s\r\n'%(meta)).encode())
