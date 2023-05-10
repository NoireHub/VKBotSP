FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -v ./cmd/botVK

CMD [ "./botVK" ]