from client.main_socket_client import *
from flask import Flask
from flask import render_template

# set the project root directory as the static folder, you can set others.
app = Flask(__name__)
inputs = []
outputs = []

@app.route('/')
def ping():
    return "<p>Hello, World!</p>"

@app.route('/py-client/<name>')
def root(name):
    input = name
    output = socketClient(input)

    inputs.append(input)
    outputs.append(output)
    l = len(inputs)
    return render_template('index.html', nameClients=inputs, dataServers=outputs, dataLengths=l)

if __name__ == "__main__":
    app.run(debug=True, host = "127.0.0.1", port = 5001)

# pip3 install flask