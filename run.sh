#!/bin/bash

# build container
sudo docker build ~/go_projects/src/btest -t golang:btest
# deploy container
sudo docker run -it --name=btest golang:btest