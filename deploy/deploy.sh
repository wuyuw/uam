#!/bin/bash

usage=$(cat <<- EOF
use:
  bash deploy/apply.sh [options]
args:
  -f      docker-compose.yaml配置文件路径
EOF
)

# 参数
while getopts f: OPT; do
  case "$OPT" in
  f) compose_path=${OPTARG}
  ;;
  \?)
    echo "${OPT}"
    echo "$usage"
    exit 1
  esac
done

#对必填项做输入检查
if [[ -z $compose_path ]]; then
  echo "$usage"
  exit 1
fi


cat "$compose_path"
docker-compose -f "$compose_path" -p uam up -d
echo "部署成功！"

