# Create Builder Image
FROM --platform=linux/amd64 golang:1.18 as builder
LABEL maintainer="MKP Mobile Krenzzz <mkpproduction@gmail.com>"

ENV GIT_TERMINAL_PROMPT=1 GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set Working Directory
RUN mkdir -p /app
ADD . /app
WORKDIR /app
COPY . .

# Do Your Magic Here
#
#

# Build Go Binary File
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/main

# Create Second Image
FROM alpine:3.13.1

RUN touch .env
ENV TZ=Asia/Jakarta

RUN apk add --no-cache tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Copy Binary File from Builder Image
COPY --from=builder /app /

# Run Binary File
ENTRYPOINT ["/main"]