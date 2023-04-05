#!/usr/bin/env bash

# 1.在部署服务器上检出代码
# 2.确认待上线分支已push到远程仓库，且pipline执行结束
# 3.在部署服务器代码根目录下执行本脚本更新对应服务版本

SCRIPTDIR="$( cd "$( dirname "$0"  )" && pwd  )"
ROOTDIR="$( cd "$SCRIPTDIR"/.. && pwd )"
CODEDIR="$( cd "$ROOTDIR"/.. && pwd )"
HOMEDIR="$( cd "$CODEDIR"/.. && pwd )"

usage=$(cat <<- EOF
use:
  bash deploy/apply.sh [options]
args:
  -t      Tag(xx-x.x.x)
  -f      docker-compose.yaml配置文件路径
EOF
)

# 参数
while getopts t:f: OPT; do
  case "$OPT" in
  t) apply_tag=${OPTARG}
    ;;
  f) compose_path=${OPTARG}
  ;;
  \?)
    echo "${OPT}"
    echo "$usage"
    exit 1
  esac
done

#对必填项做输入检查
if [[ -z $apply_tag || -z $compose_path ]]; then
  echo "$usage"
  exit 1
fi

echo "参数解析成功！"
echo "tag: $apply_tag  docker-compose path: $compose_path"
# 服务
service=${apply_tag%-*}
# 版本
version=${apply_tag#*-}
# 镜像Tag
image=wuyuw/uam-"$service"

case "$service" in
job) 
  config_name="uam-job.yaml"
  log_dir="uam-job"
  ;;
rpc) 
  config_name="uamrpc.yaml"
  log_dir="uam-rpc"
  ;;
admin) 
  config_name="uam-admin-api.yaml"
  log_dir="uam-admin"
  ;;
api) 
  config_name="uam-api.yaml"
  log_dir="uam-api"
  ;;
*)
  echo "tag模式必须为: /^(admin|api|rpc|job)-\d+\.\d+\.\d+$/"
  exit 1
esac

# 文件不存在则创建
compose_dir=$(dirname "$compose_path")
if [[ ! -d $compose_dir ]]; then
  mkdir -p "$compose_dir"
fi
if [[ ! -f "$compose_path" ]]; then
  cp "${SCRIPTDIR}"/docker-compose-template.yaml "$compose_path"
fi


# 更新image版本
sed -i "s#$image:.*#$image:$version#g" "$compose_path"
# 更新config vol
sed -i "s#/.*/uam/etc/$config_name#$ROOTDIR/etc/$config_name#g" "$compose_path"
# 更新logs vol
sed -i "s#/.*/var/logs/$log_dir#$HOMEDIR/var/logs/$log_dir#g" "$compose_path"

echo "$compose_path 配置完成"
cat "$compose_path"
