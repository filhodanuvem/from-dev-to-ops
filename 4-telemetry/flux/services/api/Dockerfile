FROM golang:1.17-alpine AS build

WORKDIR /gowork
ADD . .
RUN go build main.go


FROM alpine:3.8 AS run 
COPY --from=build /gowork/main ./main

CMD ./main