//go:build windows

package os_test

import (
	"github.com/itbasis/go-tools-core/os"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"FixPath", func() {
		defer ginkgo.GinkgoRecover()

		ginkgo.DescribeTable(
			"only windows", func(source, want string) {
				os.SkipIfMinGW()

				gomega.Expect(os.FixPath(source)).To(gomega.Equal(want))
			},
			ginkgo.Entry(nil, `C:\Windows\System32`, `C:\Windows\System32`),
			ginkgo.Entry(nil, `C:\Windows\System32;C:\Program Files\dotnet\`, `C:\Windows\System32;C:\Program Files\dotnet\`),
		)

		ginkgo.DescribeTable(
			"only MinGW", func(source, want string) {
				os.SkipIfNotMinGW()

				gomega.Expect(os.FixPath(source)).To(gomega.Equal(want))
			},
			ginkgo.Entry(nil, `C:\Windows\System32`, `/c/Windows/System32`),
			ginkgo.Entry(nil, `C:\Windows\System32;C:\Program Files\dotnet\`, `/c/Windows/System32:/c/Program Files/dotnet/`),
		)
	},
)

var _ = ginkgo.Describe(
	"SplitPathList", func() {
		defer ginkgo.GinkgoRecover()

		ginkgo.DescribeTable(
			"only windows", func(source string, want []string) {
				os.SkipIfMinGW()

				gomega.Expect(os.SplitPathList(source)).To(gomega.ConsistOf(want))
			},
			ginkgo.Entry(nil, `C:\Windows\System32`, []string{`C:\Windows\System32`}),
			ginkgo.Entry(nil, `C:\Windows\System32;C:\Program Files\dotnet\`, []string{`C:\Windows\System32`, `C:\Program Files\dotnet\`}),
		)

		ginkgo.DescribeTable(
			"only MinGW", func(source string, want []string) {
				os.SkipIfNotMinGW()

				gomega.Expect(os.SplitPathList(source)).To(gomega.ConsistOf(want))
			},
			ginkgo.Entry(nil, `C:\Windows\System32`, []string{`/c/Windows/System32`}),
			ginkgo.Entry(nil, `C:\Windows\System32;C:\Program Files\dotnet\`, []string{`/c/Windows/System32`, `/c/Program Files/dotnet/`}),
		)
	},
)
