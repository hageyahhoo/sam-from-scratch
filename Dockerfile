FROM golang:1.15-alpine AS builder

WORKDIR /go/src/app
COPY ./src/samScratchLambdaService.go .
COPY ./src/samScratchTableClient.go .

RUN go mod init
RUN go get -v "github.com/aws/aws-sdk-go/aws@v1.15.77"
RUN go get -v "github.com/aws/aws-sdk-go/aws/session@v1.15.77"
RUN go get -v "github.com/aws/aws-sdk-go/service/dynamodb@v1.15.77"
RUN go get -v "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute@v1.15.77"
RUN go install -v

FROM alpine AS app
COPY --from=builder /go/bin/app /bin/app

EXPOSE 8080
CMD ["app"]
