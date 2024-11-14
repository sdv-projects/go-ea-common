package unitofwork_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gomock "go.uber.org/mock/gomock"
)

func TestDatabase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UnitOfWork Suite")
}

var (
	ctl *gomock.Controller
	ctx context.Context = context.Background()
)

var _ = BeforeSuite(func() {
	ctl = gomock.NewController(GinkgoT())

	DeferCleanup(func() {
		ctl.Finish()
	})
})
