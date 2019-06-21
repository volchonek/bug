# step 1 build executable binary
# main container image as builder on goland
FROM golang:alpine as build-go

# variables for use in Dockerfile
LABEL maintainer="Ilay Volkov" 

# add all files from current directoty to the container directoty
ADD . /app

# install all the necessary packages for the operation of our go application
RUN cd /app && go build -o goapp . 

# step 2
# main for goland container from images alpine
FROM alpine

# sets the working directory
WORKDIR /app

COPY --from=build-go /app/goapp /app/

# sets the port to access outside, i.e. example when deploying a container from the image "docker run -p xxxx: 1989 recomendent"
EXPOSE 1989

# the command that will be executed when the container is expanded, the application start
ENTRYPOINT ["/app/goapp"]

