package os

import (
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

const (
	_unixPathSeparator    = `/`
	_windowsPathSeparator = `\`

	_unixPathListSeparator    = `:`
	_windowsPathListSeparator = `;`
)

var (
	_pathSeparator     string
	_pathListSeparator string

	_reDriveReplace = regexp.MustCompile(`\b(\S:\\)`)
)

func init() {
	_ = GetPathSeparator()
	_ = GetPathListSeparator()
}

func GetPathSeparator() string {
	if _pathSeparator != "" {
		return _pathSeparator
	}

	if runtime.GOOS == "windows" && !IsMinGW() {
		_pathSeparator = _windowsPathSeparator
	} else {
		_pathSeparator = _unixPathSeparator
	}

	return _pathSeparator
}

func GetPathListSeparator() string {
	if _pathListSeparator != "" {
		return _pathListSeparator
	}

	if runtime.GOOS == "windows" && !IsMinGW() {
		_pathListSeparator = _windowsPathListSeparator
	} else {
		_pathListSeparator = _unixPathListSeparator
	}

	return _pathListSeparator
}

func FixPath(value string) string {
	if !IsMinGW() {
		return value
	}

	var result = _reDriveReplace.ReplaceAllStringFunc(
		value, func(s string) string {
			return _unixPathSeparator + strings.ToLower(s[0:1]) + _unixPathSeparator
		},
	)

	result = filepath.ToSlash(result)

	return strings.ReplaceAll(result, _windowsPathListSeparator, _unixPathListSeparator)
}

func SplitPathList(value string) []string {
	return strings.Split(FixPath(value), _pathListSeparator)
}
