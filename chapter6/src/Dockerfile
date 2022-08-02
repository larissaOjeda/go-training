FROM golang:latest as builder
RUN cd ..
RUN mkdir test
WORKDIR test
COPY . ./
ARG version=dev
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -ldflags "-X main.version=$version" -o test ./cmd/test/main.go 

FROM scratch
COPY --from=builder /go/test/test .
CMD ["./test"]
