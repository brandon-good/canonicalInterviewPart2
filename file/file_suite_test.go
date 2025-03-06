package file_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestFile(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: GinkgoWriter})
	RegisterFailHandler(Fail)

	RunSpecs(t, "File Suite")
}
