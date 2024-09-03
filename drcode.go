package drcode

import (
	"fmt"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
)

// Initialize sets up the Sentry client with the provided project ID and public key
func Initialize(projectID, publicKey string) error {
	dsn := fmt.Sprintf("https://%s@pulse.drcode.ai:443/%s", publicKey, projectID)
	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
	})
	if err != nil {
		return fmt.Errorf("sentry initialization failed: %v", err)
	}
	return nil
}

// ErrorHandler is a middleware that captures panics and reports them to Sentry
func ErrorHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				sentry.CurrentHub().Recover(err)
				sentry.Flush(2 * time.Second)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}
}

// ReportError sends an error to Sentry
func ReportError(err error) {
	sentry.CaptureException(err)
}