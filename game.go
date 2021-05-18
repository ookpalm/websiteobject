package entity

import (
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/search"
)

type PosterPager struct {
	Entity     *[]Poster `datastore:"-"`
	Cursor     string    `datastore:"-"`
	Page       int       `datastore:"-"`
	TotalPages int       `datastore:"-"`
	PageUrl    string    `datastore:"-"`
	Gender     string    `datastore:"-"`
}

type OrphanFile struct {
	Entity  string
	Created time.Time
}

type Friend struct {
	Name           string
	SearchableHtml search.HTML
	PictureUrl     string
}

type Poster struct {
	Key                    *datastore.Key `datastore:"-"`
	LineId                 string         `datastore:",noindex"`
	Name                   string
	Age                    int    `datastore:",noindex"`
	AgeString              string `datastore:"-"`
	Gender                 int
	GenderString           string `datastore:"-"`
	SexualPreference       int
	SexualPreferenceString string `datastore:"-"`
	Province               int
	ProvinceString         string              `datastore:"-"`
	Message                string              `datastore:",noindex"`
	BeetalkId              string              `datastore:",noindex"`
	WechatId               string              `datastore:",noindex"`
	FacebookId             string              `datastore:",noindex"`
	MsnId                  string              `datastore:",noindex"`
	InstagramId            string              `datastore:",noindex"`
	Email                  string              `datastore:",noindex"`
	Phone                  string              `datastore:",noindex"`
	AgreementAcceptance    bool                `datastore:",noindex"`
	Pictures               []byte              `datastore:",noindex"`
	PictureBlobKeys        []appengine.BlobKey `datastore:",noindex"`
	PicturePrimary         string              `datastore:"-"`
	PictureStrings         []string            `datastore:"-"`
	PictureUrls            []string            `datastore:",noindex"`
	Effective              time.Time
	FriendlyEffective      string `datastore:"-"`
	Active                 bool   `datastore:",noindex"`
	RemoteAddress          string
}
