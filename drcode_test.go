package drcode

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitialize(t *testing.T) {
	err := Initialize("173", "214b3b0f665048df9e5880fba63478be")
	if err != nil {
		t.Errorf("Initialize failed: %v", err)
	}
}

func TestErrorHandler(t *testing.T) {
	err := Initialize("173", "214b3b0f665048df9e5880fba63478be")
	if err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}

	handler := ErrorHandler(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestReportError(t *testing.T) {
	err := Initialize("173", "214b3b0f665048df9e5880fba63478be")
	if err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}

	testError := errors.New("test error")
	ReportError(testError)

	// Note: We can't easily verify if the error was actually sent to Sentry in a unit test.
	// In a real-world scenario, you'd check your Sentry dashboard to confirm.
}