package handler

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

// A Request represents an HTTP request received by a server
// or to be sent by a client.
//
// The field semantics differ slightly between client and server
// usage. In addition to the notes on the fields below, see the
// documentation for Request.Write and RoundTripper.
type Request struct {
	Method           string
	URL              *url.URL
	Proto            string // "HTTP/1.0"
	ProtoMajor       int    // 1
	ProtoMinor       int    // 0
	Header           http.Header
	Body             string
	ContentLength    int64
	TransferEncoding []string
	Host             string
	Form             url.Values
	PostForm         url.Values
	MultipartForm    *multipart.Form
	Trailer          http.Header
	RemoteAddr       string
	RequestURI       string
	TLS              *tls.ConnectionState
	Response         *http.Response
}

func Handler(w http.ResponseWriter, r *http.Request) {
	req := new(Request)
	req.Method = r.Method
	req.URL = r.URL
	req.Proto = r.Proto
	req.ProtoMajor = r.ProtoMajor
	req.ProtoMinor = r.ProtoMinor
	req.Header = r.Header
	req.ContentLength = r.ContentLength
	req.TransferEncoding = r.TransferEncoding
	req.Host = r.Host
	req.Form = r.Form
	req.PostForm = r.PostForm
	req.MultipartForm = r.MultipartForm
	req.Trailer = r.Trailer
	req.RemoteAddr = r.RemoteAddr
	req.RequestURI = r.RequestURI
	req.TLS = r.TLS
	req.Response = r.Response

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Reading from request body error: " + err.Error()))
		return
	}

	req.Body = string(bodyBytes)

	bb, err := json.MarshalIndent(req, "", "\t")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("JSON Marshalling error: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	_, err = w.Write(bb)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("An error occurred while trying to write to the ResponseWriter: " + err.Error()))
		return
	}
}
