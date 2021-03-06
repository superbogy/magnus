package jwt

// header is a JOSE header narrowed down to the JWT specification from RFC 7519.
//
// Parameters are ordered according to the RFC 7515.
type header struct {
	Algorithm   string `json:"alg,omitempty"`
	KeyID       string `json:"kid,omitempty"`
	Type        string `json:"typ,omitempty"`
	ContentType string `json:"cty,omitempty"`
}
