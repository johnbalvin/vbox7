package download

import (
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/johnbalvin/vbox7/trace"
)

func Video(videoURL string, proxyURL *url.URL) ([]byte, string, error) {
	req, err := http.NewRequest("GET", videoURL, nil)
	if err != nil {
		return nil, "", trace.NewOrAdd(1, "download", "Video", err, "")
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Sec-Ch-Ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	transport := &http.Transport{
		MaxIdleConnsPerHost: 30,
		DisableKeepAlives:   true,
	}
	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	client := &http.Client{
		Timeout: time.Minute * 20,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", trace.NewOrAdd(2, "download", "Video", err, "")
	}
	if resp.StatusCode != 200 {
		return nil, "", trace.NewOrAdd(3, "download", "Video", trace.ErrParameter, "")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", trace.NewOrAdd(4, "download", "Video", err, "")
	}
	contentType := resp.Header.Get("Content-Type")
	return body, contentType, nil
}
