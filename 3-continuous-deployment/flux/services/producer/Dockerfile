FROM golang:1.17-alpine AS build

WORKDIR /gowork
ADD . .
RUN go get
RUN go build main.go


FROM alpine:3.8 AS run 
WORKDIR /the/workdir/path
COPY --from=build /gowork/main ./main

CMD ./main