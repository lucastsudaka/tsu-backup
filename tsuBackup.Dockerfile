# Choose whatever you want, version >= 1.16
FROM mariadb:jammy

WORKDIR /tsu-backup

RUN apt update && apt upgrade -y

RUN apt install npm make nano mariadb-backup wget -y

ARG GOLANG_VERSION=1.22.1
RUN wget https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz && \
    rm go${GOLANG_VERSION}.linux-amd64.tar.gz
ENV PATH="${PATH}:/usr/local/go/bin"

RUN cd /tsu-backup
#RUN go install golang.org/dl/go1.22.0@latest

COPY go.mod go.sum  ./

RUN go mod download

CMD ["./air", "-c", ".air.toml"]
