package dagshortestpath_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestShortestpath(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shortestpath Suite")
}
