package tests

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/playwright-community/playwright-go"
)

// 画面遷移とUIのテスト
func TestUISuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "画面遷移とUIのテスト")
}

//lint:ignore U1000 ignore unused variable warning for testConfig
var testConfig = struct {
	TopURL            string
	TopPageTitle      string
	SearchLink        string
	SearchURL         string
	SearchPageTitle   string
	RegisterLink      string
	RegisterURL       string
	RegisterPageTitle string
}{
	TopURL:            "http://front_exercise:8081/exercise/top",
	TopPageTitle:      "トップページ",
	SearchLink:        "a[href='/exercise/search']",
	SearchURL:         "http://front_exercise:8081/exercise/search",
	SearchPageTitle:   "商品検索",
	RegisterLink:      "a[href='/exercise/register']",
	RegisterURL:       "http://front_exercise:8081/exercise/register",
	RegisterPageTitle: "商品登録",
}

// テストに必要なリソース
type UITestEnvironment struct {
	PW      *playwright.Playwright    // playwright
	Browser playwright.Browser        // ブラウザ
	Context playwright.BrowserContext // ブラウザコンテキスト
	Page    playwright.Page           // アクセスページ
}

// テストに必要な環境を準備する
func SetupTestEnvironment() *UITestEnvironment {
	var env UITestEnvironment
	// Playwrightを開始する
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Playwrightを開始できませんでした: %v", err)
	}
	env.PW = pw
	// Chromiumブラウザを起動する
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("Chromiumブラウザを起動できませんでした: %v", err)
	}
	env.Browser = browser
	// Contextを生成する
	context, err := browser.NewContext()
	if err != nil {
		log.Fatalf("コンテキストを作成できませんでした: %v", err)
	}
	env.Context = context
	// ページを生成する
	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("ページを作成できませんでした: %v", err)
	}
	env.Page = page
	return &env
}

// 終了処理をする
func TeardownTestEnvironment(env *UITestEnvironment) {
	if env.Context != nil {
		env.Context.Close()
	}
	if env.Browser != nil {
		env.Browser.Close()
	}
	if env.PW != nil {
		env.PW.Stop()
	}
}
