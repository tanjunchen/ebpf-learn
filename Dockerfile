FROM ubuntu:20.04

RUN apt update -y -q
RUN DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y -q curl build-essential ca-certificates
RUN curl -s https://storage.googleapis.com/golang/go1.20.1.linux-amd64.tar.gz| tar -v -C /usr/local -xz
ENV PATH $PATH:/usr/local/go/bin
RUN apt install -y wget gnupg2
RUN printf "deb http://apt.llvm.org/xenial/ llvm-toolchain-xenial-12 main" | tee /etc/apt/sources.list.d/llvm-toolchain-xenial-12.list
RUN wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key | apt-key add -
RUN apt -y update
RUN apt install -y llvm clang git

WORKDIR /ebpf
COPY . .
RUN make
RUN chmod a+x /ebpf
ENTRYPOINT ["./ebpf"]
CMD ["./ebpf"]

