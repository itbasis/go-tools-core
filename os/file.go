package os

import (
	"io/fs"
	"log/slog"

	itbasisCoreLog "github.com/itbasis/go-tools-core/log"
)

func BeARegularFile(fsys fs.FS, path string) bool {
	fileInfo, err := fs.Stat(fsys, path)
	if err != nil {
		slog.Debug("fail get file info", itbasisCoreLog.SlogAttrError(err))

		return false
	}

	return fileInfo.Mode().IsRegular()
}
