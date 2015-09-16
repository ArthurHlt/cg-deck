// +build acceptance

package util

import (
	// . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/codeskyblue/go-sh"
	"github.com/sclevine/agouti"
)

type User struct {
	testEnvVars AcceptanceTestEnvVars
	session     *sh.Session
}

type UserId int

const (
	TestUser01Id UserId = iota
)

type UserAssetsMap map[UserId]AssetId

// userAssets is a subset of the assets that maps an individual user to its directory of assets.
var userAssets = map[UserId]AssetId {
}


func StartUserSessionWith(testEnvVars AcceptanceTestEnvVars, userId UserId) User {
	// Create multiple CLI sessions on the same computer by setting CF_HOME. https://github.com/cloudfoundry/cli/issues/330#issuecomment-70450866
	user := User{testEnvVars, sh.NewSession().SetEnv("CF_HOME", assets[TestUser01Id])}
	user.LoginToCLI()
	return user
}

func (u User) LoginTo(page *agouti.Page) {
	Expect(page.Navigate(u.testEnvVars.Hostname)).To(Succeed())
	delayForRendering()
	Eventually(Expect(page.Find("#login-btn").Click()).To(Succeed()))
	Eventually(Expect(page).To(HaveURL(u.testEnvVars.LoginURL + "login")))
	Expect(page.FindByName("username").Fill(u.testEnvVars.Username)).To(Succeed())
	Expect(page.FindByName("password").Fill(u.testEnvVars.Password)).To(Succeed())
	Expect(page.FindByButton("Sign in").Click()).To(Succeed())
	Expect(page).To(HaveURL(u.testEnvVars.Hostname + "/#/dashboard"))
}

func (u User) LogoutOf(page *agouti.Page) {
	Expect(page.Find("#logout-btn").Click()).To(Succeed())
	Eventually(Expect(page).To(HaveURL(u.testEnvVars.LoginURL + "login")))
}

func (u User) LoginToCLI() {
	// Make sure we have signed out
	u.LogoutOfCLI()
	u.session.Command("cf", "api", u.testEnvVars.APIURL).Run()
	// TODO. Will figure out logic for doing multiple accounts once we get them.
	u.session.Command("cf", "auth", u.testEnvVars.Username, u.testEnvVars.Password).Run()
	u.session.Command("cf", "target").Run()
}

func (u User) LogoutOfCLI() {
	u.session.Command("cf", "logout").Run()
}
