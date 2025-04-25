package tests

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/playwright-community/playwright-go"
)

var _ = Describe("画面遷移の品質検証\n", Ordered, func() {
	var env *UITestEnvironment

	BeforeAll(func() {
		var err error
		env = SetupTestEnvironment()
		Expect(err).NotTo(HaveOccurred())
	})

	AfterAll(func() {
		TeardownTestEnvironment(env)
	})

	BeforeEach(func() {
		// トップページに移動する
		_, err := env.Page.Goto(testConfig.TopURL,
			// ページの移動時に待機する条件を指定するオプション
			// ページの移動が完了し、DOM(Document Object Model)がロードされるまで待つ
			playwright.PageGotoOptions{
				WaitUntil: playwright.WaitUntilStateDomcontentloaded,
			})
		Expect(err).NotTo(HaveOccurred(), "URLに移動できませんでした")
	})

	When("トップページへのナビゲーション\n", func() {
		It("ブラウザのタイトルが「トップ」になる", func() {
			// ページタイトルを取得する
			title, err := env.Page.Title()
			Expect(err).NotTo(HaveOccurred(), "ページタイトルを取得できませんでした")
			Expect(title).To(Equal(testConfig.TopPageTitle))
		})
	})

	When("検索ページへのナビゲーション\n", func() {
		It("ブラウザのタイトルが「商品検索」になる", func() {
			// 「キーワード検索」リンクをクリックする
			searchLink := env.Page.Locator(testConfig.SearchLink)
			err := searchLink.Click()
			Expect(err).NotTo(HaveOccurred(), "検索リンクをクリックできませんでした")
			// ページタイトルを取得して検証する
			title, err := env.Page.Title()
			Expect(err).NotTo(HaveOccurred(), "ページタイトルを取得できませんでした")
			Expect(title).To(Equal(testConfig.SearchPageTitle))
			// URLを取得して検証する
			url := env.Page.URL()
			Expect(url).To(Equal(testConfig.SearchURL),
				fmt.Sprintf("URLは、%sである必要があります。", testConfig.SearchURL))
		})
	})

	When("登録ページへのナビゲーション\n", func() {
		It("ブラウザのタイトルが「商品登録」になる", func() {
			// 「登録」リンクをクリックする
			searchLink := env.Page.Locator(testConfig.RegisterLink)
			err := searchLink.Click()
			Expect(err).NotTo(HaveOccurred(), "登録リンクをクリックできませんでした")
			// ページタイトルを検証する
			title, err := env.Page.Title()
			Expect(err).NotTo(HaveOccurred(), "ページタイトルを取得できませんでした")
			Expect(title).To(Equal(testConfig.RegisterPageTitle))
			// URLを取得して検証する
			url := env.Page.URL()
			Expect(url).To(Equal(testConfig.RegisterURL),
				fmt.Sprintf("URLは、%sである必要があります。", testConfig.RegisterURL))
		})
	})
})
