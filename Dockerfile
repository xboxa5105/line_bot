FROM golang:1.12.5

WORKDIR /build/lineBot

# Fetch dependencies
COPY go.mod ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 go build

EXPOSE 8080
CMD ["./lineBot"]