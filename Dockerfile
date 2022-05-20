FROM golang:1.18.1

WORKDIR /go/src

RUN apt update && apt install build-essential protobuf-compiler librdkafka-dev -y
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

CMD ["tail","-f","/dev/null"]