import socket

hostGoServer = "127.0.0.1"  # The server's hostname or IP address
portGoServer = 65432  # The port used by the server

def socketClient(input):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((hostGoServer, portGoServer))
        s.sendall(input.encode())
        data = s.recv(1024)
        # d = "msg 2"
        # s.sendall(d.encode())
        # data2 = s.recv(1024)
        # print(data, data2)
        return data
        # repr(data)
