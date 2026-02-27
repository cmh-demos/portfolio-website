# Backend service base image
# Default template now builds a Go binary. Python example remains but is
# no longer the primary path.

# build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
# copy only module definition first; go.sum will be generated in the builder
COPY go.mod ./
# fetch dependencies and produce go.sum inside the image
RUN go mod tidy

# now copy application source and build
COPY main.go ./
RUN go build -o server .

# final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server ./
EXPOSE 8080
CMD ["./server"]

# To run Python version instead, switch base image to python:3.11-slim and
# uncomment lines below.
# COPY requirements.txt ./
# RUN pip install --no-cache-dir -r requirements.txt
# COPY . /app
# CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8080"]
