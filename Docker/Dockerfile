
# Use an official Go runtime as a parent image
FROM golang:1.20.4-buster

# Install system dependencies and Git
RUN apt-get update && apt-get install -y wget git net-tools tcpdump iperf


RUN echo 4
# Clone the go-gtp repository
RUN git clone https://github.com/AbdallahRustom/go-gtp.git /go-gtp

WORKDIR /go-gtp

RUN go mod tidy
RUN go build -o ./examples/gw-tester/mme/mme ./examples/gw-tester/mme
RUN go build -o ./examples/gw-tester/enb/enb ./examples/gw-tester/enb


WORKDIR /go-gtp
COPY init.sh /go-gtp/init.sh
CMD /go-gtp/init.sh
