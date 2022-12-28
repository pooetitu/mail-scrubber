FROM golang as build

WORKDIR /usr/app/scrubber
ENV GOPATH /usr/app

COPY go.mod go.sum ./
COPY main.go .
COPY server/ server/

RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/scrubber

FROM ubuntu

RUN mkdir /home/app
WORKDIR /home/app
COPY --from=build /usr/local/bin/scrubber /home/app/scrubber

ENV GIN_MODE=release
ENV PORT=8181

EXPOSE 8181

CMD ["/home/app/scrubber"]