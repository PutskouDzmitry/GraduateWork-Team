FROM golang:latest as builder
RUN mkdir /app
WORKDIR /app
COPY . .
ARG version=dev
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags "-X main.version=$version" -o kn-be-sd -v ./server/cmd/cmd.go

FROM alpine
COPY --from=builder /app/kn-be-sd /
CMD ["./kn-be-sd"]