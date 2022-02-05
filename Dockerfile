FROM golang:1.17-buster

# Setup basic dvel
RUN apt-get update

# Setup working dir
RUN mkdir -p /backend
WORKDIR /backend

# Copy and add files to workdir
ADD contact /backend/contact/
COPY main.go /backend/
COPY go.sum /backend/
COPY go.mod /backend/
COPY .env /backend/

# Expose Ports
EXPOSE 8080

CMD go run /backend/main.go | tee backen.log
# DOCKER_BUILDKIT=1 docker build .
# docker build -t auth:1 . 
# DOCKER_BUILDKIT=1 docker build -t auth:1 .
