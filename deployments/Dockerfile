FROM golang:1.21.4-alpine3.18
WORKDIR /app
ADD . .
RUN go install ./
COPY . .
ENTRYPOINT tilt-my-dev