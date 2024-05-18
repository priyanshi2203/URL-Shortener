package internal

import (
    "fmt"
    "net/http"
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

    // Generate a unique short key for the original URL
    shortKey := us.GenerateUniqueHash(originalURL)
    us.Urls[shortKey] = originalURL

    //Final url
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


