
FROM golang:latest
COPY . /go/src/github.com/TommyStarK/hyperledger-fabric-kubernetes/chaincode-as-external-service
WORKDIR /go/src/github.com/TommyStarK/hyperledger-fabric-kubernetes/chaincode-as-external-service
RUN CGO_ENABLED=0 GOOS=linux go build -v -mod=vendor -o chaincode

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/TommyStarK/hyperledger-fabric-kubernetes/chaincode-as-external-service/chaincode .
EXPOSE 9999
CMD ["./chaincode"]  
