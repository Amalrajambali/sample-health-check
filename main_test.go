package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var body HealthResponse
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if body.Status != "UP" {
		t.Errorf("expected status UP, got %s", body.Status)
	}
}

func TestRootHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	rootHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "running") {
		t.Errorf("expected body to contain 'running', got %q", rec.Body.String())
	}
}

func TestOpenAPISpecIsValidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/swagger/openapi.json", nil)
	rec := httptest.NewRecorder()

	openAPIHandler(rec, req)

	var spec map[string]any
	if err := json.NewDecoder(rec.Body).Decode(&spec); err != nil {
		t.Fatalf("invalid OpenAPI JSON: %v", err)
	}
	paths, _ := spec["paths"].(map[string]any)
	for _, p := range []string{"/", "/health"} {
		if _, ok := paths[p]; !ok {
			t.Errorf("spec missing path %s", p)
		}
	}
}

func TestSwaggerUIHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/swagger/", nil)
	rec := httptest.NewRecorder()

	swaggerUIHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "/swagger/openapi.json") {
		t.Error("swagger UI page does not reference the spec URL")
	}
}
