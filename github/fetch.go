package source

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Github describes github repo of gmdata
type Github struct {
	user    string
	token   string
	repoURL string
}

// Fetch is to get details from github
func (g *Github) Fetch(fileName string) ([]byte, error) {
	fileURL := fmt.Sprintf("%s/%s", g.repoURL, fileName)
	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	req, err := http.NewRequest(http.MethodGet, fileURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.VERSION.raw")
	req.SetBasicAuth(g.user, g.token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		return nil, fmt.Errorf("file %s is empty", fileName)
	}

	return body, nil
}
