//go:build !windows

package os_test

import (
	"github.com/itbasis/go-tools-core/os"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.DescribeTable(
	"FixPath", func(source, want string) {
		defer ginkgo.GinkgoRecover()

		gomega.Expect(os.FixPath(source)).To(gomega.Equal(want))
	},
	ginkgo.Entry(nil, `/usr`, `/usr`),
	ginkgo.Entry(nil, `/usr:/home`, `/usr:/home`),
)

var _ = ginkgo.DescribeTable(
	"SplitPathList", func(source string, want []string) {
		defer ginkgo.GinkgoRecover()

		gomega.Expect(os.SplitPathList(source)).To(gomega.ConsistOf(want))
	},
	ginkgo.Entry(nil, `/usr`, []string{`/usr`}),
	ginkgo.Entry(nil, `/usr:/home`, []string{`/usr`, `/home`}),
)
