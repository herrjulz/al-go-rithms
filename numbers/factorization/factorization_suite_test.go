package factorization_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFactorization(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Factorization Suite")
}
