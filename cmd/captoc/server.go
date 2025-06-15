package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// cachingFileServer adds caching headers to static file responses
func cachingFileServer(fs http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add cache headers for static assets (CSS, JS, images)
		if isStaticAsset(r.URL.Path) {
			w.Header().Add("Cache-Control", "public, max-age=31536000") // 1 year
		}
		
		fs.ServeHTTP(w, r)
	})
}

// gzipResponseWriter wraps http.ResponseWriter to provide gzip compression
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
	sniffDone bool
}

// Write implements io.Writer
func (w gzipResponseWriter) Write(b []byte) (int, error) {
	if !w.sniffDone {
		w.sniffDone = true
		if w.Header().Get("Content-Type") == "" {
			w.Header().Set("Content-Type", http.DetectContentType(b))
		}
	}
	return w.Writer.Write(b)
}

// Header returns the header map from the underlying ResponseWriter
func (w gzipResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// WriteHeader writes the HTTP header from the underlying ResponseWriter
func (w gzipResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

// compressionMiddleware adds gzip compression to responses
func compressionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			// No gzip support, pass through
			next.ServeHTTP(w, r)
			return
		}
		
		// Skip compression for small files or formats that are already compressed
		if !shouldCompress(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}
		
		// Create gzip writer
		gz, err := gzip.NewWriterLevel(w, gzip.DefaultCompression)
		if err != nil {
			// Fall back to uncompressed if error
			next.ServeHTTP(w, r)
			return
		}
		defer gz.Close()
		
		// Set compression headers
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Vary", "Accept-Encoding")
		
		// Create gzip response writer
		gzw := gzipResponseWriter{
			Writer:         gz,
			ResponseWriter: w,
		}
		
		// Serve with compression
		next.ServeHTTP(gzw, r)
	})
}

// shouldCompress checks if a path should be compressed
func shouldCompress(path string) bool {
	// Don't compress already compressed formats
	skipExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".zip", ".gz", ".mp3", ".mp4"}
	ext := filepath.Ext(path)
	
	for _, skipExt := range skipExtensions {
		if ext == skipExt {
			return false
		}
	}
	
	// Compress text-based formats
	compressExtensions := []string{".html", ".css", ".js", ".json", ".xml", ".txt", ".svg"}
	for _, compressExt := range compressExtensions {
		if ext == compressExt {
			return true
		}
	}
	
	// Default to not compressing
	return false
}

// isStaticAsset checks if a path is a static asset that should be cached
func isStaticAsset(path string) bool {
	staticExtensions := []string{".css", ".js", ".jpg", ".jpeg", ".png", ".gif", ".ico", ".svg"}
	ext := filepath.Ext(path)
	
	for _, staticExt := range staticExtensions {
		if ext == staticExt {
			return true
		}
	}
	
	return false
}

// startServer starts a local HTTP server to preview the output
func startServer(dir string, port int) error {
	// Create file server handler with caching
	fs := cachingFileServer(http.FileServer(http.Dir(dir)))
	
	// Add compression middleware
	handler := compressionMiddleware(fs)
	
	// Set up the HTTP server with read/write timeouts
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	
	// Print paths to all HTML files in the output directory
	fmt.Println("Available pages:")
	printHTMLFiles(dir)
	
	// Start the server in a goroutine
	go func() {
		fmt.Printf("Server started at http://localhost:%d\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v\n", err)
		}
	}()
	
	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	
	fmt.Println("\nShutting down server...")
	
	return nil
}

// printHTMLFiles prints all HTML files in the output directory
func printHTMLFiles(dir string) {
	// Walk the directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Only print HTML files
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			// Get relative path
			rel, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}
			
			// Print URL
			fmt.Printf("  http://localhost:8080/%s\n", filepath.ToSlash(rel))
		}
		return nil
	})
	
	if err != nil {
		fmt.Printf("Error listing HTML files: %v\n", err)
	}
} 