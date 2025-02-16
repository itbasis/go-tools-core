package env

import (
	"strings"

	"golang.org/x/tools/godoc/vfs"
)

const (
	KeyPath   = "PATH"
	KeyGoBin  = "GOBIN"
	KeyGoPath = string(vfs.RootTypeGoPath)
	KeyGoRoot = string(vfs.RootTypeGoRoot)
)

type (
	Map       = map[string]string
	List      = []string
	ListOrMap interface {
		List | Map
	}
)

func MapToSlices(envMap Map) List {
	var result = make(List, 0, len(envMap))

	for k, v := range envMap {
		result = append(result, k+"="+v)
	}

	return result
}

func SlicesToMap(env List) Map {
	var result = make(Map, len(env))

	for _, v := range env {
		envSplit := strings.SplitN(v, "=", 2) //nolint:mnd // _
		result[envSplit[0]] = envSplit[1]
	}

	return result
}

func MergeEnvs[E ListOrMap](source E, additions ...Map) Map {
	if len(source) == 0 && len(additions) == 1 {
		return additions[0]
	}

	var result Map
	switch e := any(source).(type) {
	case List:
		result = SlicesToMap(e)

	case Map:
		result = e
	}

	for _, addition := range additions {
		for k, v := range addition {
			result[k] = v
		}
	}

	return result
}
