package internal

import (
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"golang.org/x/net/html"
)

var _ = Describe("URL Shortener test", func() {

	Context("Handlers for API", func() {
		When("/shortly is hit and then hit /shortgo for redirection to the original url", func() {
			var(
				us *URLShortener
				req *http.Request
				w *httptest.ResponseRecorder
			)
			BeforeEach(func(){
				us = &URLShortener{
					Urls:      make(map[string]string),
					UrlHashes: make(map[string]string),
					DomainFreq: make(map[string]int),
				}

				originalURL := "http://localhost:8080/shortly?url=http://browserstack.com/swedcfrvgbthnjymkumyjnhgtsdcbfvdcdefrtgcvdxcfvcfvgdcfvgbcdfgcdfvgfcvg"
				form := strings.NewReader("url=" + originalURL)

				req = httptest.NewRequest(http.MethodPost, "/shorten", form)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				
				w = httptest.NewRecorder()
			})

			It("should generate a short url and give status code 200.", func() {
				us.HandleShorten(w, req)
				Expect(w.Code).To(Equal(http.StatusOK))

				// Parse HTML response body
				doc, err := html.Parse(strings.NewReader(w.Body.String()))
				Expect(err).To(BeNil())

				shortenedURL1 := fetchHref(doc)
				Expect(shortenedURL1).ToNot(Equal(""))

				//Case 2
				//Again hit the API with the same url. Should return the same shortened url as previous.
				us.HandleShorten(w, req)
				Expect(w.Code).To(Equal(http.StatusOK))

				// Parse HTML response body
				doc, err = html.Parse(strings.NewReader(w.Body.String()))
				Expect(err).To(BeNil())

				shortenedURL2 := fetchHref(doc)
				Expect(shortenedURL2).To(Equal(shortenedURL1))

				//Case 3
				//check redirection
				request := "/shortgo"+shortenedURL1
				req = httptest.NewRequest(http.MethodGet, request,nil)
				w = httptest.NewRecorder()

				us.HandleRedirect(w,req)
				Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))

			})


		})

		When("if /metrics is hit, should return top 3 domains", func() {
			var(
				us *URLShortener
				req *http.Request
				w *httptest.ResponseRecorder
			)
			BeforeEach(func(){
				us = &URLShortener{
					Urls:      make(map[string]string),
					UrlHashes: make(map[string]string),
					DomainFreq: make(map[string]int),
				}

				originalURL := "http://localhost:8080/shortly?url=http://browserstack.com/swedcfrvgbthnjymkumyjnhgtsdcbfvdcdefrtgcvdxcfvcfvgdcfvgbcdfgcdfvgfcvg"
				form := strings.NewReader("url=" + originalURL)

				req = httptest.NewRequest(http.MethodPost, "/shorten", form)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				
				w = httptest.NewRecorder()
			})

			It("should generate a short url and give status code 200.", func() {
				us.HandleShorten(w, req)
				Expect(w.Code).To(Equal(http.StatusOK))

				// Parse HTML response body
				doc, err := html.Parse(strings.NewReader(w.Body.String()))
				Expect(err).To(BeNil())

				shortenedURL1 := fetchHref(doc)
				Expect(shortenedURL1).ToNot(Equal(""))

				
				request := "/metrics"
				req = httptest.NewRequest(http.MethodGet, request,nil)
				w = httptest.NewRecorder()
				us.HandleTop3Domains(w,req)
				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))

			})


		})

	})

	Context("Utils tests", func() {
		When("GenerateUniqueHash is called", func() {
			var(
				us *URLShortener
			)

			BeforeEach(func(){
				us = &URLShortener{
					Urls:      make(map[string]string),
					UrlHashes: make(map[string]string),
					DomainFreq: make(map[string]int),
				}
			})

			It("should return 5 digit hash value", func() {
				url := "http://browserstack.com/swedcfrvgbthnjymkumyjnhgtsdcbfvdcdefrtgcvdxcfvcfvgdcfvgbcdfgcdfvgfcvg"
				hashValue := us.GenerateUniqueHash(url)
				Expect(hashValue).ToNot(Equal(""))
			})
			It("should return same hash on recalling", func() {
				url := "test"
				hashValue := us.GenerateUniqueHash(url)
				Expect(hashValue).ToNot(Equal(""))

				//call GenerateUniqueHash method again, should return same hash
				hashValue = us.GenerateUniqueHash(url)
				Expect(hashValue).ToNot(Equal(""))
			})
		})
	
	})
})
