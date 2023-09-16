#!/bin/bash
sudo docker rm -f $(sudo docker ps -qa)
sudo docker run -d -it --name regional -p 50052:50052 --network mi-red --expose 50052 lab1:latest go run regional-test/grpc_regional.go
sudo docker run -d -it --name central --network mi-red --expose 50052 lab1:latest go run central-test2/grpc_central.go