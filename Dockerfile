FROM golang:1.10-alpine AS build-env
COPY . /go/src/github.com/tambykojak/whatsmyhourlyrate-api

RUN cd /go/src/github.com/tambykojak/whatsmyhourlyrate-api && go build -o api

# final stage
FROM alpine
ENV PORT=3000
WORKDIR /app
COPY --from=build-env /go/src/github.com/tambykojak/whatsmyhourlyrate-api/api /app/
EXPOSE 3000

ENTRYPOINT ./api 