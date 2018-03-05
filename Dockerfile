# this base image is used so that subdirectory Dockerfiles can access the entire buck project in order to build with buck
FROM vmagro/buck:latest

# install golang so that buck can build our go code
RUN apt-get install -y curl unzip
RUN curl https://dl.google.com/go/go1.10.linux-amd64.tar.gz -o /tmp/go1.10.linux-amd64.tar.gz && tar -C /usr/local -xzf /tmp/go1.10.linux-amd64.tar.gz && rm -rf /tmp/go1.10.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

# install protobuf
RUN curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip && unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/local/bin/
RUN mv protoc3/include/* /usr/local/include/
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN mv /root/go/bin/protoc-gen-go /usr/local/bin

RUN mkdir /app
WORKDIR /app
COPY . /app

ENTRYPOINT ["/bin/bash"]
