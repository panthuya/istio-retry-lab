from flask import Flask
import grpc
import echo_pb2
import echo_pb2_grpc

app = Flask(__name__)

@app.route("/")
def call_backend():
    channel = grpc.insecure_channel("backend:6565")
    stub = echo_pb2_grpc.EchoServiceStub(channel)
    resp = stub.Echo(echo_pb2.EchoRequest(message="Hello from frontend"))
    return resp.message

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=80)
