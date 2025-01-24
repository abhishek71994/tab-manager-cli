// internal/models/tab.go
package models

type Tab struct {
	TabID string `json:"tabId"`
	Title string `json:"title"`
	URL   string `json:"url"`
}
