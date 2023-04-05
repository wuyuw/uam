FROM golang:1.17

ARG SERVICE
ENV SERVICE=${SERVICE}

USER root
# 修改时区
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

# 为我们的镜像设置必要的环境变量
ENV GOPROXY=https://goproxy.cn,direct \
  GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

# 移动到工作目录：/build
RUN mkdir /build
WORKDIR /build

COPY . .

RUN go build -o app ./services/${SERVICE}/

RUN mkdir /dist && cp /build/app /dist && rm -rf /build

WORKDIR /

EXPOSE 8888

ENTRYPOINT ["/dist/app"]
