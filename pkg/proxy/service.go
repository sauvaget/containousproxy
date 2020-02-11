package proxy

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sauvaget/containousproxy/models"
)

type Service interface {
	ProcessRequest(*http.Request) (*models.Proxyresponse, error)
}

type service struct {
	client              *http.Client
	cacheitemRepository models.CacheitemRepository
	proxy               map[string]string
}

func NewService(client *http.Client, ciRepo models.CacheitemRepository, proxy map[string]string) *service {
	return &service{
		client:              client,
		cacheitemRepository: ciRepo,
		proxy:               proxy,
	}
}

func (s *service) ProcessRequest(r *http.Request) (*models.Proxyresponse, error) {

	requestURL := fmt.Sprintf("%s%s", s.proxy[r.Host], r.URL.Path)

	// only GET Requests are cached
	if r.Method == http.MethodGet {
		cacheitem, err := s.cacheitemRepository.Read(requestURL)
		if err == nil {
			log.Printf("Serving %s from cache", requestURL)
			return &models.Proxyresponse{
				Header: cacheitem.Header,
				Body:   cacheitem.Value,
			}, nil
		}
		// not checking for error, cause we could still get the live version
	}

	// not in cache
	req, err := http.NewRequest(r.Method, requestURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// now cache the response
	if r.Method == http.MethodGet && resp.StatusCode == http.StatusOK {
		cacheitem := models.Cacheitem{
			Key:    requestURL,
			Header: resp.Header,
			Value:  string(respBody),
		}
		err := s.cacheitemRepository.Write(cacheitem)
		if err != nil {
			log.Println(err)
		}
	}

	return &models.Proxyresponse{
		Header: resp.Header,
		Body:   string(respBody),
	}, nil
}
