package liveinfo

import (
	"bytes"
	"log/slog"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogCarriesNoSecrets(t *testing.T) {
	var logs bytes.Buffer
	h := &Handler{Hive: Hive{PollMinutes: 120}, Log: slog.New(slog.NewTextHandler(&logs, &slog.HandlerOptions{Level: slog.LevelDebug}))}
	r := httptest.NewRequest("POST", PathConfig+"?token=QUERYSECRET", strings.NewReader("BODYSECRET"))
	r.Header.Set("Authorization", "HEADERSECRET")
	r.Header.Set("User-Agent", "UASECRET")
	h.ServeHTTP(httptest.NewRecorder(), r)
	assertClean(t, logs.String())
}

func assertClean(t *testing.T, out string) {
	t.Helper()
	if out == "" {
		t.Fatal("nothing logged")
	}
	for _, s := range []string{"QUERYSECRET", "BODYSECRET", "HEADERSECRET", "UASECRET", "424f44595345435245", "token="} {
		if strings.Contains(out, s) {
			t.Fatalf("%q leaked: %s", s, out)
		}
	}
}
