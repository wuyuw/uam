package setup

import (
	"fmt"
	"log"
	"strings"
	"uam/services/admin/internal/svc"
	"uam/services/admin/internal/sysadmin"

	"github.com/pkg/errors"
)

// 构建路由搜索树
func MustBuildRouterTree(svcCtx *svc.ServiceContext) {
	err := BuildRouterTree(svcCtx)
	if err != nil {
		log.Fatalf("路由搜索树构建失败: %s", err.Error())
	}
}

func BuildRouterTree(svcCtx *svc.ServiceContext) error {
	var err error
	for _, api := range sysadmin.ApiPermList {
		err = svcCtx.RouterTree.Add(strings.ToUpper(api.Method), api.Path)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Method: %s Path: %s", api.Method, api.Path))
		}
	}
	return nil
}
