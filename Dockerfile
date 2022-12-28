FROM golang as build

WORKDIR /usr/app/scrubber
ENV GOPATH /usr/app

COPY go.mod go.sum ./
COPY main.go .
COPY server/ server/

RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/scrubber

FROM alpine
COPY --from=build /usr/local/bin/scrubber /usr/local/bin/scrubber

RUN mkdir /home/app
WORKDIR /home/app

ENV GIN_MODE=release
ENV PORT=8181

EXPOSE 8181

ENTRYPOINT ["/usr/local/bin/scrubber"]