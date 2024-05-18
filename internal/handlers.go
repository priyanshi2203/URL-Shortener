package internal

import (
    "fmt"
    "net/http"
	"sort"
)

func (us *URLShortener) HandleShorten(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Fetch the original url from the request, throw error if missing
    originalURL := r.FormValue("url")
    if originalURL == "" {
        http.Error(w, "URL is missing in the request!", http.StatusBadRequest)
        return
    }

    //fetch domain from the url and store the increment the count of each domain
    domain := getDomain(originalURL)
	us.DomainFreq[domain]+= 1

    // Generate a unique short key for the original URL
    shortKey := us.GenerateUniqueHash(originalURL)
    us.Urls[shortKey] = originalURL

    //final url
    shortenedURL := fmt.Sprintf("http://localhost:8080/shortly/%s", shortKey)

    // Render the HTML response with the shortened URL
    w.Header().Set("Content-Type", "text/html")
    responseHTML := fmt.Sprintf(`
        <h1>SHORTLY</h1>
        <p>Original URL: %s</p>
        <p>Shortened URL: <a href="%s">%s</a></p>
    `, originalURL, shortenedURL, shortenedURL)
    fmt.Fprintf(w, responseHTML)
}

func (us *URLShortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
    shortKey := r.URL.Path[len("/shortgo/"):]
    if shortKey == "" {
        http.Error(w, "Shortened key is missing", http.StatusBadRequest)
        return
    }

    // Retrieve the original URL from the `urls` map using the shortened key
    originalURL, found := us.Urls[shortKey]
    if !found {
        http.Error(w, "Shortened key not found", http.StatusNotFound)
        return
    }

    // Redirect the user to the original URL
    http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

func (us *URLShortener) HandleTop3Domains(w http.ResponseWriter, r *http.Request) {

	// Convert the map to a slice of key-value pairs
	var pairs []Pair
	var topDomains []string

	for domain, freq := range us.DomainFreq {
		pairs = append(pairs, Pair{domain, freq})
	}

	// Sort the slice by values in descending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	 // Get the top 3 domains
	 for i := 0; i < 3 && i < len(pairs); i++ {
        topDomains = append(topDomains, pairs[i].Key)
    }

	// Initialize placeholders for the top domains
	first := "No first domain"
	second := "No 2nd domain"
	third := "No 3rd domain"

	if len(topDomains) > 0 {
    	first = topDomains[0]
	}
	if len(topDomains) > 1 {
    	second = topDomains[1]
	}
	if len(topDomains) > 2 {
    	third = topDomains[2]
	}	

     // Render the HTML response with the top 3 domains
	 w.Header().Set("Content-Type", "text/html")
	 responseHTML := fmt.Sprintf(`
		 <h2>Top 3 domains:</h2>
		 <p>1st- %s : %d times </p>
		 <p>2nd- %s : %d times </a></p>
		 <p>3rd- %s : %d times </a></p>
	 `, first,us.DomainFreq[first], second,us.DomainFreq[second], third,us.DomainFreq[third])
	 fmt.Fprintf(w, responseHTML)
}