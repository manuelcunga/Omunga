FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh && \
    chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

ARG BUILD_FLAGS
RUN go build ${BUILD_FLAGS} -o /app/main

CMD air start
