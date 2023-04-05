#!/bin/bash

echo 'starting in build.sh'

SCRIPTDIR="$( cd "$(dirname "$0")" && pwd  )"
ROOTDIR="$( cd "$SCRIPTDIR"/.. && pwd )"

usage=$(cat <<- EOF
use:
  bash deploy/build.sh [options]
args:
  -t      构建Tag(例: admin-1.0.0)
  -u      docker harbor 用户名
  -p      docker harbor 密码
EOF
)

# 参数
while getopts t:u:p: OPT
do
  case "$OPT" in
  t) build_tag=${OPTARG};;
  u) docker_user=${OPTARG};;
  p) docker_password=${OPTARG};;
  ?)
    echo "$usage"
    exit 1
  esac
done

# 对必填项做输入检查
if [[ -z $build_tag || -z $docker_user || -z $docker_password ]]; then
  echo "$usage"
  echo "tag: $build_tag, user: $docker_user, password: ${docker_password:0:9}*****"
  exit 1
fi

echo "参数解析成功！"
echo "tag: $build_tag, user: $docker_user, password: ${docker_password:0:9}*****"

# 服务
service=${build_tag%-*}
# 版本
version=${build_tag#*-}
# 镜像Tag
image_tag=wuyuw/uam-"$service":"$version"

if [[ $service != "job" && $service != "rpc" && $service != "admin" && $service != "api" ]]; then
  echo "service必须为: job|rpc|admin|api"
  exit 1
fi


echo "开始执行构建..."
echo "docker login..."
docker login --username "$docker_user" --password "$docker_password"  docker.io
echo "docker build..."
docker build --build-arg SERVICE="$service" -t "$image_tag" -f deploy/Dockerfile "$ROOTDIR"
echo "docker push..."
docker push "$image_tag"

echo "打包完成！"
echo "镜像: $image_tag"
