import socket
import time

#from selectors import EVENT_READ, EVENT_WRITE

from loop import Loop

host = "127.0.0.1"  # The server's hostname or IP address
port = 5566  # The port used by the server

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
s.bind((host, port))
s.listen(10240)
s.setblocking(False)

loop = Loop()

def handler(conn):
    while True:
        msg = yield from loop.recv(conn, 1024)
        if not msg:
            conn.close()
            break
        local_time = time.ctime()
        v = msg.decode('utf-8')
        d = "%s-%s" %(v, local_time)
        yield from loop.send(conn, d.encode())

def main():
    while True:
        conn, addr = yield from loop.accept(s)
        conn.setblocking(False)
        loop.create_task((handler(conn), None))

loop.create_task((main(), None))
loop.run()