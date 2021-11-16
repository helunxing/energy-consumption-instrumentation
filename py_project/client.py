import socket

HOST = "localhost"
PORT = 1966

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    s.sendall(b'gemini://gemini.circumlunar.space/docs/\r\n')
    data = s.recv(1024)
    print(data.decode())
