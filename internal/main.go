package main

import (
	"compress/gzip"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	targetURL := "http://localhost:8000" // Replace with your backend server URL
	proxyURL, err := url.Parse(targetURL)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	// Customize the Director function to handle request modifications
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = proxyURL.Scheme
		req.URL.Host = proxyURL.Host
		req.Host = proxyURL.Host
		req.Header.Add("Accept-Encoding", "gzip")
	}
	// Customize the ModifyResponse function to handle responses
	proxy.ModifyResponse = func(resp *http.Response) error {
		// Check if the response is already gzip-encoded
		if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
			return nil
		}

		// Remove Content-Length because the content length will change after gzip compression
		resp.Header.Del("Content-Length")
		// Set the Content-Encoding header to indicate that the response is gzip-compressed
		resp.Header.Set("Content-Encoding", "gzip")
		// Add the Vary header to indicate that the response varies based on the Accept-Encoding header
		resp.Header.Add("Vary", "Accept-Encoding")

		pr, pw := io.Pipe()
		gw := gzip.NewWriter(pw)

		// Store the original body to close it later
		origBody := resp.Body
		resp.Body = io.NopCloser(pr)

		// Use a goroutine to compress the response body and write to the pipe
		go func() {
			defer origBody.Close()
			defer pw.Close()
			defer gw.Close()

			buf := make([]byte, 1024)
			for {
				n, err := origBody.Read(buf)
				if n > 0 {
					if _, err := gw.Write(buf[:n]); err != nil {
						log.Printf("Error writing to gzip writer: %v", err)
						return
					}
					if err := gw.Flush(); err != nil {
						log.Printf("Error flushing gzip writer: %v", err)
						return
					}
				}
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					if !errors.Is(err, context.Canceled) {
						log.Printf("Error reading from response body: %v", err)
					}
					return
				}
			}
		}()

		return nil
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Println("Starting proxy server on :9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func flushResponse(w http.ResponseWriter) {
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	n, err := w.Writer.Write(b)
	if err == nil {
		flushResponse(w.ResponseWriter)
	}
	return n, err
}

func newGzipResponseWriter(w http.ResponseWriter) http.ResponseWriter {
	gw := gzip.NewWriter(w)
	return gzipResponseWriter{Writer: gw, ResponseWriter: w}
}
