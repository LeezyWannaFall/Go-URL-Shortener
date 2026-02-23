package handler

import (
	"context"
	"errors"
	"testing"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"github.com/go-chi/chi/v5"
)

type mockService struct {
	AddShortUrlFunc func(ctx context.Context, full string) (string, error)
	GetFullUrlFunc func(ctx context.Context, short string) (string, error)
}

func (m *mockService) AddShortUrl(ctx context.Context, full string) (string, error) {
	return m.AddShortUrlFunc(ctx, full)
}

func (m *mockService) GetFullUrl(ctx context.Context, short string) (string, error) {
	return m.GetFullUrlFunc(ctx, short)
}

func TestAddShortUrl(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		mockResponse   string
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "success",
			body:           `{"full": "https://github.com/LeezyWannaFall"}`,
			mockResponse:   "lol123",
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedBody:   `"lol123"` + "\n",
		},
		{
			name:           "invalid body",
			body:           `invalid`, // Некорректный JSON
			mockResponse:   "",
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "invalid body\n",
		},
		{
			name:           "service error",
			body:           `{"full": "https://github.com/LeezyWannaFall"}`,
			mockResponse:   "",
			mockError:      errors.New("service error"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "service error\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &mockService{
				AddShortUrlFunc: func(ctx context.Context, full string) (string, error) {
					return tt.mockResponse, tt.mockError
				},
			}

			handler := NewHandler(service)

			req := httptest.NewRequest(http.MethodPost, "/shorten", strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			handler.AddShortUrl(w, req)

			res := w.Result()
			body, _ := io.ReadAll(res.Body)

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}
			if string(body) != tt.expectedBody {
				t.Errorf("expected body %s, got %s", tt.expectedBody, string(body))
			}
		})
	}
}

func TestRedirect(t *testing.T) {
	tests := []struct {
		name           string
		short          string
		mockResponse   string
		mockError      error
		expectedStatus int
		expectedLocation string
	}{
		{
			name:           "success",
			short:          "lol123",
			mockResponse:   "https://github.com/LeezyWannaFall",
			mockError:      nil,
			expectedStatus: http.StatusFound,
			expectedLocation: "https://github.com/LeezyWannaFall",
		},
		{
			name:           "not found",
			short:          "lol123",
			mockResponse:   "",
			mockError:      errors.New("not found"),
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mock := &mockService{
				GetFullUrlFunc: func(ctx context.Context, short string) (string, error) {
					return tt.mockResponse, tt.mockError
				},
			}

			handler := NewHandler(mock)

			r := chi.NewRouter()
			r.Get("/{short}", handler.Redirect)

			req := httptest.NewRequest(http.MethodGet, "/"+tt.short, nil)
			rec := httptest.NewRecorder()

			r.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			if tt.expectedStatus == http.StatusFound {
				location := rec.Header().Get("Location")
				if location != tt.expectedLocation {
					t.Errorf("expected location %s, got %s", tt.expectedLocation, location)
				}
			}
		})
	}
}
