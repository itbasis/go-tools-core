//go:build windows

package os_test

import (
	itbasisCoreOs "github.com/itbasis/go-tools-core/os"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"CleanPath", func() {
		defer ginkgo.GinkgoRecover()

		ginkgo.DescribeTable(
			"only windows", func(path string, cleanPaths []string, wantPath string) {
				itbasisCoreOs.SkipIfMinGW()

				gomega.Expect(itbasisCoreOs.CleanPath(path, cleanPaths...)).To(gomega.Equal(wantPath))
			},
			ginkgo.Entry(nil, "a;b;c", []string{}, "a;b;c"),
			ginkgo.Entry(nil, "a;b;c", []string{"a"}, "b;c"),
			ginkgo.Entry(nil, "a;b;a;c", []string{"a"}, "b;c"),
			ginkgo.Entry(nil, "a;b;c", []string{"b"}, "a;c"),
			ginkgo.Entry(nil, "a;b;c", []string{"c"}, "a;b"),
			ginkgo.Entry(nil, "a;b;c", []string{"a", "c"}, "b"),
		)

		ginkgo.DescribeTable(
			"only MinGW", func(path string, cleanPaths []string, wantPath string) {
				itbasisCoreOs.SkipIfNotMinGW()

				gomega.Expect(itbasisCoreOs.CleanPath(path, cleanPaths...)).To(gomega.Equal(wantPath))
			},
			ginkgo.Entry(nil, "a;b;c", []string{}, "a:b:c"),
			ginkgo.Entry(nil, "a;b;c", []string{"a"}, "b:c"),
			ginkgo.Entry(nil, "a;b;a;c", []string{"a"}, "b:c"),
			ginkgo.Entry(nil, "a;b;c", []string{"b"}, "a:c"),
			ginkgo.Entry(nil, "a;b;c", []string{"c"}, "a:b"),
			ginkgo.Entry(nil, "a;b;c", []string{"a", "c"}, "b"),
		)
	},
)

var _ = ginkgo.Describe(
	"AddBeforePath", func() {
		defer ginkgo.GinkgoRecover()

		ginkgo.DescribeTable(
			"only windows", func(path string, addPaths []string, wantPath string) {
				itbasisCoreOs.SkipIfMinGW()

				gomega.Expect(itbasisCoreOs.AddBeforePath(path, addPaths...)).To(gomega.Equal(wantPath))
			},
			ginkgo.Entry(nil, "a;b;c", []string{}, "a;b;c"),
			ginkgo.Entry(nil, "a;b;c", []string{"d"}, "d;a;b;c"),
			ginkgo.Entry(nil, "a;b;c", []string{"d", "e"}, "d;e;a;b;c"),
		)

		ginkgo.DescribeTable(
			"only MinGW", func(path string, addPaths []string, wantPath string) {
				itbasisCoreOs.SkipIfNotMinGW()

				gomega.Expect(itbasisCoreOs.AddBeforePath(path, addPaths...)).To(gomega.Equal(wantPath))
			},
			ginkgo.Entry(nil, "a;b;c", []string{}, "a:b:c"),
			ginkgo.Entry(nil, "a;b;c", []string{"d"}, "d:a:b:c"),
			ginkgo.Entry(nil, "a;b;c", []string{"d", "e"}, "d:e:a:b:c"),
		)
	},
)
