package book

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type bookList struct {
	Units         []unit
	ActiveUser    string
	MyRequests    []request
	ForMeRequests []request
}

type unit struct {
	BookID    int
	Title     string
	Thumbnail string
	Owner     string
}

type request struct {
	Candidate string
	BookID    int
	Title     string
}

type bookGoogle struct {
	Kind       string `json:"-"`
	TotalItems int    `json:"-"`
	Items      []struct {
		Kind       string `json:"-"`
		ID         string `json:"-"`
		Etag       string `json:"-"`
		SelfLink   string `json:"-"`
		VolumeInfo struct {
			Title               string   `json:"title"`
			Authors             []string `json:"-"`
			Publisher           string   `json:"-"`
			PublishedDate       string   `json:"-"`
			Description         string   `json:"-"`
			IndustryIdentifiers []struct {
				Type       string `json:"-"`
				Identifier string `json:"-"`
			} `json:"-"`
			ReadingModes struct {
				Text  bool `json:"-"`
				Image bool `json:"-"`
			} `json:"-"`
			PageCount        int      `json:"-"`
			PrintType        string   `json:"-"`
			Categories       []string `json:"-"`
			AverageRating    float64  `json:"-"`
			RatingsCount     int      `json:"-"`
			MaturityRating   string   `json:"-"`
			AllowAnonLogging bool     `json:"-"`
			ContentVersion   string   `json:"-"`
			ImageLinks       struct {
				SmallThumbnail string `json:"smallThumbnail"`
				Thumbnail      string `json:"thumbnail"`
			} `json:"imageLinks"`
			Language            string `json:"-"`
			PreviewLink         string `json:"-"`
			InfoLink            string `json:"-"`
			CanonicalVolumeLink string `json:"-"`
		} `json:"volumeInfo"`
		SaleInfo struct {
			Country     string `json:"-"`
			Saleability string `json:"-"`
			IsEbook     bool   `json:"-"`
			ListPrice   struct {
				Amount       float64 `json:"-"`
				CurrencyCode string  `json:"-"`
			} `json:"-"`
			RetailPrice struct {
				Amount       float64 `json:"-"`
				CurrencyCode string  `json:"-"`
			} `json:"-"`
			BuyLink string `json:"-"`
			Offers  []struct {
				FinskyOfferType int `json:"-"`
				ListPrice       struct {
					AmountInMicros float64 `json:"-"`
					CurrencyCode   string  `json:"-"`
				} `json:"-"`
				RetailPrice struct {
					AmountInMicros float64 `json:"-"`
					CurrencyCode   string  `json:"-"`
				} `json:"-"`
				Giftable bool `json:"-"`
			} `json:"-"`
		} `json:"-"`
		AccessInfo struct {
			Country                string `json:"-"`
			Viewability            string `json:"-"`
			Embeddable             bool   `json:"-"`
			PublicDomain           bool   `json:"-"`
			TextToSpeechPermission string `json:"-"`
			Epub                   struct {
				IsAvailable  bool   `json:"-"`
				AcsTokenLink string `json:"-"`
			} `json:"-"`
			Pdf struct {
				IsAvailable  bool   `json:"-"`
				AcsTokenLink string `json:"-"`
			} `json:"-"`
			WebReaderLink       string `json:"-"`
			AccessViewStatus    string `json:"-"`
			QuoteSharingAllowed bool   `json:"-"`
		} `json:"-"`
		SearchInfo struct {
			TextSnippet string `json:"-"`
		} `json:"-"`
	} `json:"items"`
}

func getBooksFromGoogle(w http.ResponseWriter) {
	pet, _ := http.Get("https://www.googleapis.com/books/v1/volumes?q=science")
	decoder := json.NewDecoder(pet.Body)
	var book bookGoogle
	err := decoder.Decode(&book)
	if err != nil {
		panic(err)
	}
	defer pet.Body.Close()
	for i := range book.Items {
		fmt.Println(book.Items[i].VolumeInfo.Title)
		fmt.Println(book.Items[i].VolumeInfo.ImageLinks.SmallThumbnail)
		//fmt.Println(book.Items[i].VolumeInfo.ImageLinks.Thumbnail)
	}
	// Convert struct to JSON and send to the client
	//w.Header().Set("Content-Type", "application/json")
	js, _ := json.MarshalIndent(book, "", " ")
	w.Write(js)
}

type bookGoogleOriginal struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []struct {
		Kind       string `json:"kind"`
		ID         string `json:"id"`
		Etag       string `json:"etag"`
		SelfLink   string `json:"selfLink"`
		VolumeInfo struct {
			Title               string   `json:"title"`
			Authors             []string `json:"authors"`
			Publisher           string   `json:"publisher"`
			PublishedDate       string   `json:"publishedDate"`
			Description         string   `json:"description"`
			IndustryIdentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
			ReadingModes struct {
				Text  bool `json:"text"`
				Image bool `json:"image"`
			} `json:"readingModes"`
			PageCount        int      `json:"pageCount"`
			PrintType        string   `json:"printType"`
			Categories       []string `json:"categories"`
			AverageRating    float64  `json:"averageRating"`
			RatingsCount     int      `json:"ratingsCount"`
			MaturityRating   string   `json:"maturityRating"`
			AllowAnonLogging bool     `json:"allowAnonLogging"`
			ContentVersion   string   `json:"contentVersion"`
			ImageLinks       struct {
				SmallThumbnail string `json:"smallThumbnail"`
				Thumbnail      string `json:"thumbnail"`
			} `json:"imageLinks"`
			Language            string `json:"language"`
			PreviewLink         string `json:"previewLink"`
			InfoLink            string `json:"infoLink"`
			CanonicalVolumeLink string `json:"canonicalVolumeLink"`
		} `json:"volumeInfo"`
		SaleInfo struct {
			Country     string `json:"country"`
			Saleability string `json:"saleability"`
			IsEbook     bool   `json:"isEbook"`
			ListPrice   struct {
				Amount       float64 `json:"amount"`
				CurrencyCode string  `json:"currencyCode"`
			} `json:"listPrice"`
			RetailPrice struct {
				Amount       float64 `json:"amount"`
				CurrencyCode string  `json:"currencyCode"`
			} `json:"retailPrice"`
			BuyLink string `json:"buyLink"`
			Offers  []struct {
				FinskyOfferType int `json:"finskyOfferType"`
				ListPrice       struct {
					AmountInMicros float64 `json:"amountInMicros"`
					CurrencyCode   string  `json:"currencyCode"`
				} `json:"listPrice"`
				RetailPrice struct {
					AmountInMicros float64 `json:"amountInMicros"`
					CurrencyCode   string  `json:"currencyCode"`
				} `json:"retailPrice"`
				Giftable bool `json:"giftable"`
			} `json:"offers"`
		} `json:"saleInfo"`
		AccessInfo struct {
			Country                string `json:"country"`
			Viewability            string `json:"viewability"`
			Embeddable             bool   `json:"embeddable"`
			PublicDomain           bool   `json:"publicDomain"`
			TextToSpeechPermission string `json:"textToSpeechPermission"`
			Epub                   struct {
				IsAvailable  bool   `json:"isAvailable"`
				AcsTokenLink string `json:"acsTokenLink"`
			} `json:"epub"`
			Pdf struct {
				IsAvailable  bool   `json:"isAvailable"`
				AcsTokenLink string `json:"acsTokenLink"`
			} `json:"pdf"`
			WebReaderLink       string `json:"webReaderLink"`
			AccessViewStatus    string `json:"accessViewStatus"`
			QuoteSharingAllowed bool   `json:"quoteSharingAllowed"`
		} `json:"accessInfo"`
		SearchInfo struct {
			TextSnippet string `json:"textSnippet"`
		} `json:"searchInfo,omitempty"`
	} `json:"items"`
}
