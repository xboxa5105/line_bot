FROM golang:1.12.5

WORKDIR /build/lineBot

# Fetch dependencies
# COPY go.mod ./

# Build
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 go build

EXPOSE 8080
CMD ["./lineBot"]