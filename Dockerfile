FROM node:16-alpine AS static

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY ./ ./
WORKDIR /app/templates
RUN export NODE_OPTIONS=--openssl-legacy-provider
RUN npm install
RUN npm run build --openssl-legacy-provider

FROM golang:1.17.6-alpine
WORKDIR /app
COPY --from=static ./app ./
RUN go mod download
RUN CGO_ENABLED=0 go test ./... -v -cover

WORKDIR /app/cmd
RUN go build -o global_log

WORKDIR /app
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

CMD ["/app/cmd/sf_log"]
