package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custon form strcut, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializs a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form fied is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}

	return true
}