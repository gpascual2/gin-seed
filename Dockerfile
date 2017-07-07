FROM golang:1.8.3-alpine
ENV API_PORT=80
EXPOSE $API_PORT
WORKDIR /go/src/bitbucket.org/gpascual2/gin-seed/api
ADD ./api/. /go/src/bitbucket.org/gpascual2/gin-seed/api
RUN go install && go build
CMD ["./api","-e","prod"]

