import grpc
import networkstuff_pb2_grpc as pb2_grpc
import networkstuff_pb2 as pb2


if __name__ == "__main__":
    channel = grpc.insecure_channel("localhost:50051")
    stub = pb2_grpc.RouterServiceStub(channel)
    router_request = pb2.RouterRequest(id=2)
    result = stub.GetRouter(router_request)
    print(f"{result}")
