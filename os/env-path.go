package os

import (
	"context"
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

// CleanPath Removes the sdkm directory from paths
//
//nolint:nonamedreturns // To optimize result logging
func CleanPath(path string, cleanPaths ...string) (result string) {
	var logger = slog.Default()

	defer slog.Debug("clean path: " + result)

	if logger.Enabled(context.Background(), slog.LevelDebug) {
		logger.Debug("path: " + path)
		logger.Debug(fmt.Sprintf("cleaning env path: %s", cleanPaths))
	}

	if len(cleanPaths) == 0 {
		result = FixPath(path)

		return result
	}

	var splitPaths = SplitPathList(path)

	splitPaths = slices.DeleteFunc(
		splitPaths, func(s string) bool {
			return slices.Contains(cleanPaths, s)
		},
	)

	result = strings.Join(splitPaths, _pathListSeparator)

	return result
}

func AddBeforePath(path string, paths ...string) string {
	path = FixPath(path)

	if len(paths) == 0 {
		return path
	}

	return strings.Join(paths, _pathListSeparator) + _pathListSeparator + path
}
