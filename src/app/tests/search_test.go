package tests

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/playwright-community/playwright-go"
)

var _ = Describe("検索画面の品質検証\n", Ordered, func() {
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
		// 検索ページに移動する
		_, err := env.Page.Goto(testConfig.SearchURL,
			// ページの移動時に待機する条件を指定するオプション
			// ページの移動が完了し、DOM(Document Object Model)がロードされるまで待つ
			playwright.PageGotoOptions{
				WaitUntil: playwright.WaitUntilStateDomcontentloaded,
			})
		Expect(err).NotTo(HaveOccurred(), "URLに移動できませんでした")
	})

	When("存在しない商品キーワードで検索する\n", func() {
		It("存在しないことを表すエラーメッセージが表示される", func() {
			// キーワード「川」を入力するためのLocatorを作成する
			keywordInput := env.Page.Locator("#keyword")
			err := keywordInput.Fill("川")
			Expect(err).NotTo(HaveOccurred(), "キーワードフィールドを入力できませんでした")
			// [検索]ボタンをクリックする
			err = env.Page.Locator("button[type='submit']").Click()
			Expect(err).NotTo(HaveOccurred(), "[検索]ボタンをクリックできませんでした")
			// エラーメッセージを検証する
			errorMessage := env.Page.Locator("p.text-danger.mt-3")
			messageText, err := errorMessage.TextContent()
			Expect(err).NotTo(HaveOccurred(), "Failed to retrieve the error message")
			Expect(messageText).To(ContainSubstring("キーワード:'川'に該当する商品は見つかりませんでした。"),
				"エラー メッセージは商品が見つからなかったことを示すはず")
		})
	})
	When("存在する商品キーワードで検索する\n", func() {
		It("該当する商品の一覧が表示される", func() {
			// キーワード「山」を入力するためのLocatorを作成する
			keywordInput := env.Page.Locator("#keyword")
			err := keywordInput.Fill("ボールペン")
			Expect(err).NotTo(HaveOccurred(), "キーワードフィールドを入力できませんでした")
			// [検索]ボタンをクリックする
			err = env.Page.Locator("button[type='submit']").Click()
			Expect(err).NotTo(HaveOccurred(), "[検索]ボタンをクリックできませんでした")

			// テーブルの存在を確認する
			table := env.Page.Locator("table.table")
			Expect(table.Count()).To(Equal(1),
				"ページ上に 1 つの結果テーブルが存在する必要がある")
			// テーブルの行を検証
			rows := table.Locator("tbody tr")
			rowCount, err := rows.Count()
			Expect(err).NotTo(HaveOccurred(), "テーブル内の行をカウントできませんでした")
			Expect(rowCount).To(BeNumerically(">", 0),
				"結果テーブルには少なくとも 1 行が必要")
			// 特定の行のデータを検証する
			row := rows.Nth(0)
			Expect(row.Locator("td").Nth(0).TextContent()).To(Equal("1"),
				"最初の行の最初の列は「1」である必要がある")
			Expect(row.Locator("td").Nth(1).TextContent()).To(Equal("水性ボールペン(黒)"),
				"最初の行の2列目は「水性ボールペン(黒)」である必要がある")
			Expect(row.Locator("td").Nth(2).TextContent()).To(Equal("￥120"),
				"最初の行の3列目は「￥120」である必要がある")
			Expect(row.Locator("td").Nth(3).TextContent()).To(Equal("文房具"),
				"最初の行の4列目は「文房具」である必要がある")
			row = rows.Nth(1)
			Expect(row.Locator("td").Nth(0).TextContent()).To(Equal("2"),
				"最初の行の最初の列は「2」である必要がある")
			Expect(row.Locator("td").Nth(1).TextContent()).To(Equal("水性ボールペン(赤)"),
				"最初の行の2列目は「水性ボールペン(赤)」である必要がある")
			Expect(row.Locator("td").Nth(2).TextContent()).To(Equal("￥120"),
				"最初の行の3列目は「￥120」である必要がある")
			Expect(row.Locator("td").Nth(3).TextContent()).To(Equal("文房具"),
				"最初の行の4列目は「文房具」である必要がある")
			row = rows.Nth(2)
			Expect(row.Locator("td").Nth(0).TextContent()).To(Equal("3"),
				"最初の行の最初の列は「3」である必要がある")
			Expect(row.Locator("td").Nth(1).TextContent()).To(Equal("水性ボールペン(青)"),
				"最初の行の2列目は「水性ボールペン(青)」である必要がある")
			Expect(row.Locator("td").Nth(2).TextContent()).To(Equal("￥120"),
				"最初の行の3列目は「￥120」である必要がある")
			Expect(row.Locator("td").Nth(3).TextContent()).To(Equal("文房具"),
				"最初の行の4列目は「文房具」である必要がある")
			row = rows.Nth(3)
			Expect(row.Locator("td").Nth(0).TextContent()).To(Equal("4"),
				"最初の行の最初の列は「4」である必要がある")
			Expect(row.Locator("td").Nth(1).TextContent()).To(Equal("油性ボールペン(黒)"),
				"最初の行の2列目は「油性ボールペン(黒)」である必要がある")
			Expect(row.Locator("td").Nth(2).TextContent()).To(Equal("￥100"),
				"最初の行の3列目は「￥100」である必要がある")
			Expect(row.Locator("td").Nth(3).TextContent()).To(Equal("文房具"),
				"最初の行の4列目は「文房具」である必要がある")
			row = rows.Nth(4)
			Expect(row.Locator("td").Nth(0).TextContent()).To(Equal("5"),
				"最初の行の最初の列は「5」である必要がある")
			Expect(row.Locator("td").Nth(1).TextContent()).To(Equal("油性ボールペン(赤)"),
				"最初の行の2列目は「油性ボールペン(赤)」である必要がある")
			Expect(row.Locator("td").Nth(2).TextContent()).To(Equal("￥100"),
				"最初の行の3列目は「￥100」である必要がある")
			Expect(row.Locator("td").Nth(3).TextContent()).To(Equal("文房具"),
				"最初の行の4列目は「文房具」である必要がある")
			row = rows.Nth(5)
			Expect(row.Locator("td").Nth(0).TextContent()).To(Equal("6"),
				"最初の行の最初の列は「6」である必要がある")
			Expect(row.Locator("td").Nth(1).TextContent()).To(Equal("油性ボールペン(青)"),
				"最初の行の2列目は「油性ボールペン(青)」である必要がある")
			Expect(row.Locator("td").Nth(2).TextContent()).To(Equal("￥100"),
				"最初の行の3列目は「￥100」である必要がある")
			Expect(row.Locator("td").Nth(3).TextContent()).To(Equal("文房具"),
				"最初の行の4列目は「文房具」である必要がある")
		})
	})
})
