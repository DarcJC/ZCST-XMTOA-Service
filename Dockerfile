FROM golang:1.17.7-alpine

WORKDIR $GOPATH/src/Onboarding
COPY . $GOPATH/src/Onboarding
ENV GOPROXY=https://goproxy.cn,direct
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init -d .,handler
RUN go build -o build/web . && go build -o build/runner ./task/binary/runner.go

EXPOSE 8000
ENTRYPOINT ["./build/web"]
