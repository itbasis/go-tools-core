//go:build !windows

package os

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"GetPathSeparator", func() {
		gomega.Expect(GetPathSeparator()).To(gomega.Equal(_unixPathSeparator))
	},
)

var _ = ginkgo.Describe(
	"GetPathListSeparator", func() {
		gomega.Expect(GetPathListSeparator()).To(gomega.Equal(_unixPathListSeparator))
	},
)
