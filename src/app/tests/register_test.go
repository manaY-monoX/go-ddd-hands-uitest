package tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/playwright-community/playwright-go"
)

var _ = Describe("登録画面の品質検証\n", Ordered, func() {
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
		// 登録ページに移動する
		_, err := env.Page.Goto(testConfig.RegisterURL,
			// ページの移動時に待機する条件を指定するオプション
			// ページの移動が完了し、DOM(Document Object Model)がロードされるまで待つ
			playwright.PageGotoOptions{
				WaitUntil: playwright.WaitUntilStateDomcontentloaded,
			})
		Expect(err).NotTo(HaveOccurred(), "URLに移動できませんでした")
	})

	When("適切に値を入力して結果画面に遷移する\n", func() {
		It("登録結果画面に遷移することを検証する", func() {
			// 商品名を入力する
			err := env.Page.Locator("#productName").Fill("砂消しゴム")
			Expect(err).NotTo(HaveOccurred(), "商品名の入力に失敗しました")
			// 単価を入力する
			err = env.Page.Locator("#productPrice").Fill("110")
			Expect(err).NotTo(HaveOccurred(), "単価の入力に失敗しました")
			// カテゴリを選択する
			values := []string{"b1524011-b6af-417e-8bf2-f449dd58b5c0"} // カテゴリを選択する
			_, err = env.Page.Locator("#category").SelectOption(playwright.SelectOptionValues{
				Values: &values,
			})
			Expect(err).NotTo(HaveOccurred(), "カテゴリの選択に失敗しました")
			// [登録]ボタンをクリックする
			err = env.Page.Locator("button[type='submit']").Click()
			Expect(err).NotTo(HaveOccurred(), "フォームの送信に失敗しました")
			// 登録結果画面に遷移したことを検証する
			Expect(env.Page.Locator("p.mt-3").TextContent()).To(Equal("商品: 砂消しゴムを登録しました。"),
				"確認メッセージが表示される")
		})
	})

	When("登録済商品を登録する\n", func() {
		It("登録結果画面にエラーメッセージが表示されることを検証する", func() {
			// 商品を入力する
			err := env.Page.Locator("#productName").Fill("砂消しゴム")
			Expect(err).NotTo(HaveOccurred(), "商品名の入力に失敗しました")
			// 単価を入力する
			err = env.Page.Locator("#productPrice").Fill("110")
			Expect(err).NotTo(HaveOccurred(), "単価の入力に失敗しました")
			// カテゴリを選択する
			values := []string{"b1524011-b6af-417e-8bf2-f449dd58b5c0"} // カテゴリを選択する
			_, err = env.Page.Locator("#category").SelectOption(playwright.SelectOptionValues{
				Values: &values,
			})
			Expect(err).NotTo(HaveOccurred(), "カテゴリの選択に失敗しました")
			// [登録]ボタンをクリックする
			err = env.Page.Locator("button[type='submit']").Click()
			Expect(err).NotTo(HaveOccurred(), "フォームの送信に失敗しました")
			// 登録結果画面に遷移したことを検証する
			Expect(env.Page.Locator("p.mt-3").TextContent()).To(Equal("商品:砂消しゴムは、既に登録済です。"),
				"エラーメッセージが表示される")
		})
	})
})
