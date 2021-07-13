package api

//go:generate moq -out mock/contentstore.go -pkg mock . ContentStore

// ContentStore defines the required methods from the data store of content
type ContentStore interface {
}
