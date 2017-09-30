package search

import (
	"encoding/json"
	"fmt"
	"freeCodeCamp/v1/back-end/apis/_help"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func init() {
	//fmt.Println("Init from Package Search")
}

const searchPath = "./search/search.json"
const configPath = "./secret.json"
const registryLimit = 5000

// RouterSearch ...
func RouterSearch(w http.ResponseWriter, r *http.Request) {
	option := strings.Split(r.URL.Path, "/")[3]
	switch option {
	case "recent":
		showRecentSearches(w, r)
	case "search":
		doSearch(w, r)
	default:
		http.Redirect(w, r, help.BaseURL+"search/search.html", 301)
	}
}

type records []record

type record struct {
	Title       string `json:"title"`
	ID          string `json:"id"`
	Snippet     string `json:"snippet"`
	DisplayLink string `json:"displayLink"`
	ThumbSrc    string `json:"thumbSrc"`
	ImageSrc    string `json:"imageSrc"`
}

type quests []quest

type quest struct {
	IDImage int    `json:"idimage"`
	Search  string `json:"search"`
	When    string `json:"when"`
}

func (qs quests) Len() int {
	return len(qs)
}

func (qs quests) Less(i, j int) bool {
	return qs[i].IDImage > qs[j].IDImage
}

func (qs quests) Swap(i, j int) {
	qs[i], qs[j] = qs[j], qs[i]
}

func showRecentSearches(w http.ResponseWriter, r *http.Request) {
	var d help.Data
	var qs quests
	help.GetJSONDataFromFile(searchPath, &d)
	convertInterfaceToQuests(&d, &qs)
	sort.Sort(qs)
	help.StructToJSON(w, r, qs)
}

func doSearch(w http.ResponseWriter, r *http.Request) {
	var c help.Conf
	help.LoadConfig(configPath, &c)
	urlData, err := url.Parse("https://www.googleapis.com/customsearch/v1?q=")
	if err != nil {
		fmt.Println(err)
	}
	params := url.Values{}
	if r.URL.Query().Get("q") == "" {
		fmt.Println(`error`)
	} else {
		params.Add("q", r.URL.Query().Get("q"))
	}
	params.Add("cx", c.APIImage.CseID)
	params.Add("key", c.APIImage.Key)
	num, err := strconv.Atoi(r.URL.Query().Get("num"))
	if err != nil || num < 0 {
		fmt.Println(`error or num<0`)
	} else {
		params.Add("num", r.URL.Query().Get("num"))
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil || offset < 0 {
		fmt.Println(`error or offset<0`)
	} else {
		params.Add("offset", r.URL.Query().Get("offset"))
	}
	urlData.RawQuery = params.Encode()
	fmt.Println(urlData.String())
	var gq googleQuest

	getAPIJSONDataToStruct(w, urlData.String(), &gq)
	help.StructToJSON(w, r, gq)

	// now do we have to insert the search ?
	var d help.Data
	var q quest
	var qs quests
	help.GetJSONDataFromFile(searchPath, &d)
	convertInterfaceToQuests(&d, &qs)
	sort.Sort(qs)
	q.Search = strings.ToLower(params.Get("q"))
	for _, v := range qs {
		if v.Search == q.Search {
			return
		}
	}
	layout := "2006-01-02 15:04:05"
	q.When = time.Now().Format(layout)
	// Take the first free number
	q.IDImage = qs[0].IDImage + 1
	qs = append(qs, q)
	writeJSONtoFile(&qs, searchPath)
}

func writeJSONtoFile(qs *quests, pathToFile string) {
	f, err := os.Create(pathToFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	e := json.NewEncoder(f)
	e.Encode(qs)
}

func getAPIJSONDataToStruct(w http.ResponseWriter, urlData string, gq *googleQuest) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get(urlData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HOLA", resp.StatusCode)
	if resp.StatusCode != 200 {
		w.Write([]byte("Not FOUND"))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &gq)
	if err != nil {
		log.Fatalln(err)
	}
}

func convertInterfaceToQuests(d *help.Data, qs *quests) {
	var q quest
	fields := []string{"idimage", "search", "when"}
	for _, v := range *d {
		values := v.(map[string]interface{})
		for _, t := range fields { // convert nil to "" to avoid exceptions
			if values[t] == nil {
				//fmt.Println(t, values[t] == nil)
				values[t] = ""
			}
		}
		q.IDImage = int(values["idimage"].(float64))
		q.Search = values["search"].(string)
		/*layout := "2006-01-02 15:04:05"
		dateSearch, err := time.Parse(layout, values["when"].(string))
		if err != nil {
			q.When = ""
		}
		q.When = dateSearch.Format(time.RFC1123)*/
		q.When = values["when"].(string)
		*qs = append(*qs, q)
	}
}

func convertInterfaceToRecords(d *help.Data, recs *records) {
	var rec record
	fields := []string{"title", "id", "snippet", "displayLink", "thumbSrc", "imageSrc"}
	for _, v := range *d {
		values := v.(map[string]interface{})
		for _, t := range fields { // convert nil to "" to avoid exceptions
			if values[t] == nil {
				fmt.Println(t, values[t] == nil)
				values[t] = ""
			}
		}
		rec.Title = values["title"].(string)
		rec.ID = values["id"].(string)
		rec.Snippet = values["snippet"].(string)
		rec.DisplayLink = values["displayLink"].(string)
		rec.ThumbSrc = values["thumbSrc"].(string)
		rec.ImageSrc = values["imageSrc"].(string)
		*recs = append(*recs, rec)
	}
}

type googleQuest struct {
	Kind string `json:"kind"`
	URL  struct {
		Type     string `json:"type"`
		Template string `json:"template"`
	} `json:"url"`
	Queries struct {
		Request []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"request"`
		NextPage []struct {
			Title          string `json:"title"`
			TotalResults   string `json:"totalResults"`
			SearchTerms    string `json:"searchTerms"`
			Count          int    `json:"count"`
			StartIndex     int    `json:"startIndex"`
			InputEncoding  string `json:"inputEncoding"`
			OutputEncoding string `json:"outputEncoding"`
			Safe           string `json:"safe"`
			Cx             string `json:"cx"`
		} `json:"nextPage"`
	} `json:"queries"`
	Context struct {
		Title string `json:"title"`
	} `json:"context"`
	SearchInformation struct {
		SearchTime            float64 `json:"searchTime"`
		FormattedSearchTime   string  `json:"formattedSearchTime"`
		TotalResults          string  `json:"totalResults"`
		FormattedTotalResults string  `json:"formattedTotalResults"`
	} `json:"searchInformation"`
	Items []struct {
		Kind             string `json:"kind"`
		Title            string `json:"title"`
		HTMLTitle        string `json:"htmlTitle"`
		Link             string `json:"link"`
		DisplayLink      string `json:"displayLink"`
		Snippet          string `json:"snippet"`
		HTMLSnippet      string `json:"htmlSnippet"`
		CacheID          string `json:"cacheId"`
		FormattedURL     string `json:"formattedUrl"`
		HTMLFormattedURL string `json:"htmlFormattedUrl"`
		Pagemap          struct {
			CseThumbnail []struct {
				Width  string `json:"width"`
				Height string `json:"height"`
				Src    string `json:"src"`
			} `json:"cse_thumbnail"`
			Review []struct {
				Ratingstars string `json:"ratingstars,omitempty"`
				Ratingcount string `json:"ratingcount,omitempty"`
				URL         string `json:"url,omitempty"`
				Publishdate string `json:"publishdate,omitempty"`
			} `json:"review"`
			Person []struct {
				URL  string `json:"url"`
				Name string `json:"name"`
			} `json:"person"`
			Hreviewaggregate []struct {
				Count    string `json:"count"`
				Votes    string `json:"votes"`
				LinkUrls string `json:"link_urls"`
			} `json:"hreviewaggregate"`
			Book []struct {
				Image          string `json:"image"`
				Name           string `json:"name"`
				Bookformattype string `json:"bookformattype"`
				Inlanguage     string `json:"inlanguage"`
			} `json:"book"`
			Aggregaterating []struct {
				Ratingvalue string `json:"ratingvalue"`
				Ratingcount string `json:"ratingcount"`
			} `json:"aggregaterating"`
			Metatags []struct {
				CsrfParam          string `json:"csrf-param"`
				CsrfToken          string `json:"csrf-token"`
				RequestID          string `json:"request-id"`
				TwitterCard        string `json:"twitter:card"`
				TwitterSite        string `json:"twitter:site"`
				TwitterTitle       string `json:"twitter:title"`
				TwitterDescription string `json:"twitter:description"`
				TwitterImage       string `json:"twitter:image"`
				TwitterImageAlt    string `json:"twitter:image:alt"`
				AppleItunesApp     string `json:"apple-itunes-app"`
				OgTitle            string `json:"og:title"`
				OgType             string `json:"og:type"`
				OgURL              string `json:"og:url"`
				OgSiteName         string `json:"og:site_name"`
				OgImage            string `json:"og:image"`
				FbAppID            string `json:"fb:app_id"`
				OgDescription      string `json:"og:description"`
				GoodReadsPageNum   string `json:"good_reads:page_num"`
				GoodReadsAuthor    string `json:"good_reads:author"`
				AlIosURL           string `json:"al:ios:url"`
				AlIosAppStoreID    string `json:"al:ios:app_store_id"`
				AlIosAppName       string `json:"al:ios:app_name"`
			} `json:"metatags"`
			CseImage []struct {
				Src string `json:"src"`
			} `json:"cse_image"`
		} `json:"pagemap"`
	} `json:"items"`
}
