//go:build windows

package os

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"GetPathSeparator", func() {
		defer ginkgo.GinkgoRecover()

		ginkgo.It(
			"only windows", func() {
				SkipIfMinGW()

				gomega.Expect(GetPathSeparator()).To(gomega.Equal(_windowsPathSeparator))
			},
		)

		ginkgo.It(
			"only MinGW", func() {
				SkipIfNotMinGW()

				gomega.Expect(GetPathSeparator()).To(gomega.Equal(_unixPathSeparator))
			},
		)

	},
)

var _ = ginkgo.Describe(
	"GetPathListSeparator", func() {
		defer ginkgo.GinkgoRecover()

		ginkgo.It(
			"only windows", func() {
				SkipIfMinGW()

				gomega.Expect(GetPathListSeparator()).To(gomega.Equal(_windowsPathListSeparator))
			},
		)

		ginkgo.It(
			"only MinGW", func() {
				SkipIfNotMinGW()

				gomega.Expect(GetPathListSeparator()).To(gomega.Equal(_unixPathListSeparator))
			},
		)
	},
)
