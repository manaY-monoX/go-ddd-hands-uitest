package tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/playwright-community/playwright-go"
)

var _ = Describe("Cookieにセッションが登録されていることを検証\n", Ordered, func() {
	var env *UITestEnvironment

	BeforeAll(func() {
		var err error
		env = SetupTestEnvironment()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterAll(func() {
		TeardownTestEnvironment(env)
	})

	When("Cookieの値を取得して値を検証する\n", func() {
		BeforeEach(func() {
			// 登録ページに移動する
			_, err := env.Page.Goto(testConfig.RegisterURL,
				// ページの移動時に待機する条件を指定するオプション
				// ページの移動が完了し、DOM(Document Object Model)がロードされるまで待つ
				playwright.PageGotoOptions{
					WaitUntil: playwright.WaitUntilStateDomcontentloaded,
				})
			Expect(err).NotTo(HaveOccurred(), "URLに移動できませんでした")
		})
		It("Cookieの値を評価する", func() {
			// Cookieを取得する
			cookies, err := env.Context.Cookies()
			Expect(err).NotTo(HaveOccurred())
			var appSessionCookie *playwright.Cookie
			for _, cookie := range cookies {
				if cookie.Name == "app-session" {
					appSessionCookie = &cookie
					break
				}
			}
			Expect(appSessionCookie).NotTo(BeNil(), "Cookie「app-session」が存在する必要がある")
			Expect(appSessionCookie.Domain).To(Equal("front_exercise"), "Cookie「app-session」には正しいドメインが必要")
			Expect(appSessionCookie.Path).To(Equal("/"), "Cookie「app-session」には正しいパスが必要")
			Expect(appSessionCookie.HttpOnly).To(BeTrue(), "Cookie 'app-session' は HttpOnly である")
		})
	})

})
