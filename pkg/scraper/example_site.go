package scraper

// ExampleSiteScraper is an example implementation of the Scraper interface.
type ExampleSiteScraper struct{}

// FetchData fetches data from an example site.
func (s ExampleSiteScraper) FetchData() (string, error) {
	// Example implementation, replace with real scraping logic.
	return "Example firmware data", nil
}

// NewExampleSiteScraper creates a new instance of ExampleSiteScraper.
func NewExampleSiteScraper() Scraper {
	return &ExampleSiteScraper{}
}
