#!/bin/bash

# 设置应用名称和版本号
APP_NAME="donkey"
APP_VERSION="1.0.1"

# 根据当前机器架构选择 Dockerfile
ARCH=$(uname -m)
if [ "$ARCH" == "x86_64" ]; then
  DOCKERFILE="Dockerfile.amd64"
elif [ "$ARCH" == "arm64" ]; then
  DOCKERFILE="Dockerfile.arm64v8"
else
  echo "Unsupported architecture: $ARCH"
  exit 1
fi

# 构建 Docker 镜像
docker build -t $APP_NAME:$APP_VERSION -f $DOCKERFILE .

#停止并删除历史容器
docker stop $APP_NAME && docker rm $APP_NAME

# 运行 Docker 镜像
docker run -d -p 8080:8080 --name $APP_NAME $APP_NAME:$APP_VERSION 

# 推送 Docker 镜像
# docker push $APP_NAME:$APP_VERSION