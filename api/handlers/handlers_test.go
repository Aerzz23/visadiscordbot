package handlers_test

import (
	"github.com/boltdb/bolt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/aerzz23/visadiscordbot/api/handlers"
)

var _ = Describe("Bothandlers", func() {
	var (
		testBotHandlers handlers.BotHandlers
	)

	BeforeEach(func() {
		testBotHandlers = handlers.New(&bolt.DB{})
	})

	Describe("BotHandlers has been initialized", func() {
		Context("When I call New()", func() {
			It("Should being assigned db", func() {
				Expect(testBotHandlers).ToNot(BeNil())
			})
		})

	})
})
