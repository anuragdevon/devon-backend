# Build Phase------------------------------------------------
FROM golang:1.13.8-alpine3.11 as build 

# Setup working dir
RUN mkdir -p /devon-backend
WORKDIR /devon-backend

# Copy and add files to workdir
ADD contact /devon-backend/contact/
COPY main.go /devon-backend/
COPY go.sum /devon-backend/
COPY go.mod /devon-backend/
COPY .env /devon-backend/

# Build init
RUN CGO_ENABLED=0 go build -o /devon-backend/

# Image Serve Phase-------------------------------------------
FROM alpine:3.14

COPY --from=build /devon-backend /devon-backend

# Expose Ports
EXPOSE 8080

# Start Server
CMD cd devon-backend && ./main | tee devon-backend.log
