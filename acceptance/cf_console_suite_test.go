// +build acceptance

package acceptance

import (
	"github.com/18F/cf-deck/controllers"
	"github.com/gocraft/web"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"

	"fmt"
	"net/http/httptest"
	"os"
	"path"
	"testing"
	"time"
)

// delayForRendering is to allow for test platforms to catch up and render.
func delayForRendering() {
	time.Sleep(helpers.TimeoutConstant)
}

// Helper function to handle all the weird work of creating a test server.
func startServer() (*httptest.Server, acceptanceTestEnvVars) {
	// Load the environment variables to conduct the tests.
	testEnvVars := acceptanceTestEnvVars{}
	testEnvVars.loadTestEnvVars()

	var err error
	// Attempt to initial routers
	app, settings, err := controllers.InitApp(testEnvVars.EnvVars)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Since we are running in a separate folder from the main package, we need to change the location of the static folder.
	app.Middleware(web.StaticMiddleware("../static", web.StaticOption{IndexFile: "index.html"}))

	// Create the httptest server
	server := httptest.NewUnstartedServer(app)
	server.Start()

	// Change config values to fit the URL of the httptest server that is created on a random port.
	testEnvVars.Hostname = server.URL
	settings.OAuthConfig.RedirectURL = server.URL + "/oauth2callback"

	return server, testEnvVars
}

func createPage() *agouti.Page {
	// Create a fresh page to navigate.
	page, err := agoutiDriver.NewPage()
	Expect(err).NotTo(HaveOccurred())
	// PhantomJS makes the window really small. For now, these tests will be for desktop sizes.
	page.Size(1024, 768)
	Expect(err).NotTo(HaveOccurred())
	page.ClearCookies()
	return page
}



func TestCfConsole(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CfConsole Suite")
}

var agoutiDriver *agouti.WebDriver

var _ = BeforeSuite(func() {
	// MAKE SURE YOU INSTALL PhantomJS. `brew install phantomjs`
	agoutiDriver = agouti.PhantomJS()

	Expect(agoutiDriver.Start()).To(Succeed())
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
})
