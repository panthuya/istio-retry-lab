FROM python:3.11-alpine
WORKDIR /app
RUN pip install grpcio grpcio-tools flask
COPY app.py ./
COPY proto/ ./proto/
RUN python -m grpc_tools.protoc -I./proto --python_out=. --grpc_python_out=. ./proto/echo.proto
EXPOSE 80
CMD ["python", "app.py"]
