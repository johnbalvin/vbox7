package dependencies

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"
	"strings"

	"github.com/johnbalvin/vbox7/requests"

	"github.com/johnbalvin/vbox7/trace"

	"github.com/PuerkitoBio/goquery"
)

func GetbaseURL(id string, proxyURL *url.URL) (string, error) {
	finalURL, vidID, err := getbaseURL(id, proxyURL)
	if err != nil {
		return "", trace.NewOrAdd(1, "dependencies", "GetbaseURL", err, "")
	}
	if finalURL != "" {
		return finalURL, nil
	}
	finalURL, _, err = getbaseURL(vidID, proxyURL)
	if err != nil {
		return "", trace.NewOrAdd(2, "dependencies", "GetbaseURL", err, "")
	}
	return finalURL, nil
}

func getbaseURL(id string, proxyURL *url.URL) (string, string, error) {
	urlToUse := fmt.Sprintf("https://www.vbox7.com/aj/player/item/options?vid=%s", id)
	body, err := requests.Get(urlToUse, proxyURL)
	if err != nil {
		return "", "", trace.NewOrAdd(1, "dependencies", "getbaseURL", err, "")
	}
	var data Data
	if err := json.Unmarshal(body, &data); err != nil {
		return "", "", trace.NewOrAdd(2, "dependencies", "getbaseURL", err, "")
	}
	if !data.Sucess {
		urlToUse := "https://www.vbox7.com/article:" + id
		body, err := requests.Get(urlToUse, proxyURL)
		if err != nil {
			return "", "", trace.NewOrAdd(3, "dependencies", "getbaseURL", err, "")
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			return "", "", trace.NewOrAdd(4, "dependencies", "getbaseURL", err, "")
		}
		src := doc.Find("iframe").AttrOr("src", "")
		urlParsed, err := url.Parse(src)
		if err != nil {
			return "", "", trace.NewOrAdd(5, "dependencies", "getbaseURL", err, "")
		}
		vidID := urlParsed.Query().Get("vid")
		if vidID == "" {
			return "", "", trace.NewOrAdd(6, "dependencies", "getbaseURL", trace.ErrEmpty, "")
		}
		return "", vidID, nil
	}
	splited := strings.Split(data.Option.Src, "/")
	if len(splited) < 2 {
		return "", "", trace.NewOrAdd(7, "dependencies", "getbaseURL", trace.ErrEmpty, "")
	}
	baseEnpoint := strings.Join(splited[:len(splited)-1], "/")
	body, err = requests.Get(data.Option.Src, proxyURL)
	if err != nil {
		return "", "", trace.NewOrAdd(8, "dependencies", "getbaseURL", err, "")
	}
	var metadata MPD
	if err := xml.Unmarshal(body, &metadata); err != nil {
		return "", "", trace.NewOrAdd(9, "dependencies", "getbaseURL", err, "")
	}
	if len(metadata.Periods) == 0 || len(metadata.Periods[0].AdaptationSets) == 0 || len(metadata.Periods[0].AdaptationSets[0].Representations) == 0 {
		return "", "", trace.NewOrAdd(10, "dependencies", "getbaseURL", trace.ErrEmpty, "")
	}
	name := metadata.Periods[0].AdaptationSets[0].Representations[0].BaseURL
	finalURL := baseEnpoint + "/" + name
	return finalURL, "", nil
}
