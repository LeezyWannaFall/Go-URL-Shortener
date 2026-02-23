package service

import (
	"context"
	"errors"
	"testing"

	"github.com/LeezyWannaFall/Go-URL-Shortener/internal/model"
)

type mockRepo struct {
	GetByShortFunc func(short string) (*model.URL, error)
	GetByFullFunc  func(full string) (*model.URL, error)
	SaveFunc       func(url *model.URL) error	
}

func (m *mockRepo) Save(ctx context.Context, url *model.URL) error {
	return m.SaveFunc(url)
}

func (m *mockRepo) GetByFullUrl(ctx context.Context, full string) (*model.URL, error) {
	return m.GetByFullFunc(full)
}

func (m *mockRepo) GetByShortUrl(ctx context.Context, short string) (*model.URL, error) {
	return m.GetByShortFunc(short)
}

func TestGetFullUrl(t *testing.T) {
	tests := []struct{
		name			string
		shortInput		string
		mockResponse	*model.URL
		mockError		error
		expectedFull	string
		expectedError	bool
	}{
		{
			name: 				"get url success",
			shortInput: 		"lol123",
			mockResponse:		&model.URL{Full: "https://finance.ozon.ru", Short: "lol123"},
			mockError:			nil,
			expectedFull:		"https://finance.ozon.ru",
			expectedError:		false,
		},
		{
			name:       "url not found",
			shortInput: "notfound",
			mockResponse:   nil,
			mockError:    errors.New("not found"),
			expectedFull: "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockRepo{
				GetByShortFunc: func(short string) (*model.URL, error) {
					return tt.mockResponse, tt.mockError
				},
			}
			
			svc := NewUrlService(repo)

			res, err := svc.GetFullUrl(context.Background(), tt.shortInput)

			if (err != nil) != tt.expectedError {
				t.Fatalf("expected error: %v, got: %v", tt.expectedError, err)
			}

			if res != tt.expectedFull {
				t.Errorf("expected: %s, got: %s", tt.expectedFull, res)
			}
		})
	}
}

func TestAddShortUrl(t *testing.T) {
    test := []struct {
        name            string
        fullInput       string
        mockGetFull     func(string) (*model.URL, error)
        mockSave        func(*model.URL) error
        expectedError   bool
    }{
        {
            name:      "success: link already exists",
            fullInput: "https://finance.ozon.ru",
            mockGetFull: func(full string) (*model.URL, error) {
                return &model.URL{Full: full, Short: "exist123"}, nil
            },
            mockSave: func(url *model.URL) error {
                t.Error("Save should NOT be called when link exists")
                return nil
            },
            expectedError: false,
        },
        {
            name:      "success: create new link",
            fullInput: "https://google.com",
            mockGetFull: func(full string) (*model.URL, error) {
                return nil, errors.New("not found")
            },
            mockSave: func(url *model.URL) error {
                return nil
            },
            expectedError: false,
        },
        {
            name:      "invalid url",
            fullInput: "not-a-url",
            mockGetFull: func(full string) (*model.URL, error) {
                return nil, nil
            },
            mockSave: func(url *model.URL) error {
                return nil
            },
            expectedError: true,
        },
    }

    for _, tt := range test {
        t.Run(tt.name, func(t *testing.T) {
            repo := &mockRepo{
                GetByFullFunc: tt.mockGetFull,
                SaveFunc:      tt.mockSave,
            }

            svc := NewUrlService(repo)
            _, err := svc.AddShortUrl(context.Background(), tt.fullInput)

            if (err != nil) != tt.expectedError {
                t.Fatalf("expected error: %v, got: %v", tt.expectedError, err)
            }
        })
    }
}