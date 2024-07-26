package scraper

// Scraper defines the interface for a site scraper.
type Scraper interface {
	FetchData() (string, error)
}
