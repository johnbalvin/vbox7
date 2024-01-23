package vbox7

import (
	"net/url"
	"strings"

	"github.com/johnbalvin/vbox7/dependencies"
	"github.com/johnbalvin/vbox7/download"
	"github.com/johnbalvin/vbox7/trace"
	"github.com/johnbalvin/vbox7/utils"
)

func GetFinalVideoURLFromPageID(id string, proxyURL *url.URL) (string, error) {
	return dependencies.GetbaseURL(id, nil)
}
func GetFinalVideoURLFromPageURL(pageURL string, proxyURL *url.URL) (string, error) {
	splited := strings.Split(pageURL, ":")
	if len(splited) != 3 {
		return "", trace.NewOrAdd(1, "dependencies", "GetFinalVideoURLFromPageURL", trace.ErrParameter, "")
	}
	id := utils.RemoveSpace(splited[2])
	return dependencies.GetbaseURL(id, nil)
}

// it's the final videoURL
func DowloadVideoFromFinalVideoURL(finalVideoURL string, proxyURL *url.URL) ([]byte, string, error) {
	return download.Video(finalVideoURL, nil)
}
