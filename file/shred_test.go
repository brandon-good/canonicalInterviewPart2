package file_test

import (
	"bytes"
	"os"
	"shred/file"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("When attempting to shred a file with Shred", func() {
	It("should return an error if file does not exist", func() {
		err := file.Shred("nonExistentFile")
		Expect(err).ToNot(BeNil())
	})

	It("should not return an error if the file has no read permissions", func() {
		// create the test file
		err := os.WriteFile("test_data/WPerms.txt", []byte("No reading!"), 0300)
		Expect(err).To(BeNil())

		err = file.Shred("test_data/WPerms.txt")
		Expect(err).To(BeNil())

		_, err = os.Stat("test_data/WPerms.txt")
		Expect(os.IsNotExist(err)).To(BeTrue())
	})

	It("should return an error if the file has no write permissions", func() {
		// create the test file
		err := os.WriteFile("test_data/RPerms.txt", []byte("No writing!"), 0400)
		Expect(err).To(BeNil())

		err = file.Shred("test_data/RPerms.txt")
		Expect(err).ToNot(BeNil())

		// rm the test file
		err = os.Remove("test_data/RPerms.txt")
		Expect(err).To(BeNil())
	})

	It("should successfully be removed after calling Shred", func() {
		// create the test file
		err := os.WriteFile("test_data/RWPerms.txt", []byte("These are some random bytes!"), 0600)
		Expect(err).To(BeNil())

		err = file.Shred("test_data/RWPerms.txt")
		Expect(err).To(BeNil())

		_, err = os.Stat("test_data/RWPerms.txt")
		Expect(os.IsNotExist(err)).To(BeTrue())
	})
})

var _ = Describe("When calling Mangle", func() {
	BeforeEach(func() {
		// create the test file bigger than the default buff size
		err := os.WriteFile("test_data/RWPerms.txt", []byte(strings.Repeat("Canonical", file.BUFF_SIZE)), 0600)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		// rm the test file
		err := os.Remove("test_data/RWPerms.txt")
		Expect(err).To(BeNil())
	})

	It("should have different contents before and after the call to mangle", func() {
		contentsBefore, err := os.ReadFile("test_data/RWPerms.txt")
		Expect(err).To(BeNil())

		err = file.Mangle("test_data/RWPerms.txt")
		Expect(err).To(BeNil())

		contentsAfter, err := os.ReadFile("test_data/RWPerms.txt")
		Expect(err).To(BeNil())

		Expect(bytes.Equal(contentsBefore, contentsAfter)).ToNot(BeTrue())
	})

})
