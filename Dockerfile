FROM golang:1.23

WORKDIR /usr/src/app

ENV AWS_CONFIG_FILE=/root/.aws/config
ENV AWS_SSO_SESSION=z-dev-sso

ARG AWS_ACCESS_KEY
ARG AWS_SECRET_KEY

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

CMD ["app"]

