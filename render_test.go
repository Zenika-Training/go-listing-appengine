package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	w := httptest.NewRecorder()
	transformAndDisplay(appEngineServicesStruct{}, w)
	b := w.Body.String()
	for _, s := range []string{`class="logo"`, `class="sun"`, `class="moon"`, `href="/assets/css/style.css"`} {
		if !strings.Contains(b, s) {
			t.Errorf("missing %s", s)
		}
	}
	if strings.Contains(b, "<style>") {
		t.Error("inline style left")
	}
}
