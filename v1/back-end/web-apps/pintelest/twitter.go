package pintelest

import (
	"encoding/json"
	"fmt"
	"freeCodeCamp/v1/back-end/web-apps/_help"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitter"
)

func init() {
	c.loadData(secretFile)
	c.Mysql.Db = "pintelest"
	mydb.initDB()
	store := sessions.NewFilesystemStore(os.TempDir(), []byte("goth-example"))
	store.MaxLength(math.MaxInt64)
	gothic.Store = store
	goth.UseProviders(
		twitter.New(c.Twitter.ConsumerKey, c.Twitter.ConsumerSecret, loginCallback),
	)
}

var c Conf

// Conf ...
type Conf struct {
	Twitter struct {
		ConsumerKey       string `json:"ConsumerKey"`
		ConsumerSecret    string `json:"ConsumerSecret"`
		AccessToken       string `json:"AccessToken"`
		AccessTokenSecret string `json:"AccessTokenSecret"`
	} `json:"twitter"`
	Mysql struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Db       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"mysql"`
	User goth.User
}

func (c *Conf) authTwitter(w http.ResponseWriter, r *http.Request) {
	fmt.Println(`---------- c.authTwitter ----------`)
	values := r.URL.Query()
	values.Add("provider", "twitter")
	r.URL.RawQuery = values.Encode()
	// try to get the user without re-authenticating
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err == nil {
		c.User = gothUser
		profile(w, r)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (c *Conf) callbackTwitter(w http.ResponseWriter, r *http.Request) {
	fmt.Println(`---------- c.callback ----------`)
	values := r.URL.Query()
	values.Add("provider", "twitter")
	r.URL.RawQuery = values.Encode()
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	c.User = gothUser
	profile(w, r)
}

func (c *Conf) logoutTwitter(w http.ResponseWriter, r *http.Request) {
	fmt.Println(`---------- c.logout ----------`)
	c.User = goth.User{}
	imgs = Images{}
	gothic.Logout(w, r)
	w.Header().Set("Location", help.BaseURL+"pintelest/guest")
	w.WriteHeader(http.StatusTemporaryRedirect)
	//http.Redirect(w, r, "https://brusbilis.com/freecodecamp/v1/apps/pintelest/guest", 301)
}

func (c *Conf) loadData(filePath string) { // parse JSON with MARSHALL
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	defer file.Close()
	// get file content
	chunk, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(chunk, &c)
}

func (c *Conf) isLogged() bool {
	return c.User.Name != ""
}
