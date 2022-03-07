FROM golang:1.15-alpine as builder

RUN apk add --no-cache git
RUN apk add --update make
RUN mkdir -p /app

# Move to working directory /build
WORKDIR /app

# Copy the code into the container
COPY . .

# Copy and download dependency using go mod
RUN go mod download

RUN make build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /bin/app /bin/app

# Command to run when starting the container
CMD ["/bin/app"]