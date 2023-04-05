package sysadmin

import (
	"errors"
	"net/http"
	"path"
	"uam/tools/search"
)

type RouterTree struct {
	Trees map[string]*search.Tree
}

var emptyHandler = struct{}{}

func NewRouterTree() *RouterTree {
	return &RouterTree{
		Trees: make(map[string]*search.Tree),
	}
}

func (rt *RouterTree) Add(method, reqPath string) error {
	if !validMethod(method) {
		return errors.New("无效的请求方法")
	}

	if len(reqPath) == 0 || reqPath[0] != '/' {
		return errors.New("无效的Path")
	}

	cleanPath := path.Clean(reqPath)
	tree, ok := rt.Trees[method]
	if ok {
		return tree.Add(cleanPath, emptyHandler)
	}

	tree = search.NewTree()
	rt.Trees[method] = tree
	return tree.Add(cleanPath, emptyHandler)
}

func (rt *RouterTree) Search(r *http.Request) (*ApiPerm, error) {
	reqPath := path.Clean(r.URL.Path)
	if tree, ok := rt.Trees[r.Method]; ok {
		if result, ok := tree.Search(reqPath); ok {
			return &ApiPerm{
				Method: r.Method,
				Path:   result.GetRoute(),
			}, nil
		}
	}
	return nil, errors.New("not found")
}

func validMethod(method string) bool {
	return method == http.MethodDelete || method == http.MethodGet ||
		method == http.MethodHead || method == http.MethodOptions ||
		method == http.MethodPatch || method == http.MethodPost ||
		method == http.MethodPut
}
