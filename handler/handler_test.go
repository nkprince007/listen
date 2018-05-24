package handler_test

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"os"
	"syscall"
	"testing"

	"gitlab.com/nkprince007/listen/handler"
)

func generateGhSign(secret string, body []byte) string {
	hash := hmac.New(sha1.New, []byte(secret))
	hash.Write(body)
	return hex.EncodeToString(hash.Sum(nil))
}

func TestCaptureWrongMethod(t *testing.T) {
	var ts = httptest.NewServer(http.HandlerFunc(handler.Capture))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Error(res.Status)
		t.Error("methods other than post should be rejected")
	}
}

func TestNoMatchingProvider(t *testing.T) {
	var ts = httptest.NewServer(http.HandlerFunc(handler.Capture))
	defer ts.Close()

	// no explicit X-*-Event header has been set on request
	data := bytes.NewBuffer([]byte("{}"))
	res, err := http.Post(ts.URL, "application/json", data)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Error(res.Status)
		t.Error("matching provider should not be found")
	}
}

func TestGitHubSignatureVerification(t *testing.T) {
	oldGhSecret := os.Getenv("GITHUB_SECRET")
	os.Setenv("GITHUB_SECRET", "secret")

	defer func() {
		if oldGhSecret != "" {
			os.Setenv("GITHUB_SECRET", oldGhSecret)
		}
	}()

	reqBody := []byte("{}")
	sign := generateGhSign("secret", reqBody)

	var ts = httptest.NewServer(http.HandlerFunc(handler.Capture))
	defer ts.Close()

	data := bytes.NewBuffer(reqBody)

	// pass signature verification
	req, err := http.NewRequest("POST", ts.URL, data)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("X-GitHub-Event", "ping")
	req.Header.Add("X-Hub-Signature", "sha1="+sign)

	res, err := ts.Client().Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(res.Status)
		t.Error("github signature verification success, but still rejected")
	}

	// failed signature verification
	req, err = http.NewRequest("POST", ts.URL, data)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("X-GitHub-Event", "ping")
	req.Header.Add("X-Hub-Signature", "sha1=somethingelse")

	res, err = ts.Client().Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusForbidden {
		t.Error(res.Status)
		t.Error("github signature match failed, but still forwarded")
	}
}

func TestGitLabSignatureVerification(t *testing.T) {
	oldGlSecret := os.Getenv("GITLAB_SECRET")
	os.Setenv("GITLAB_SECRET", "secret")

	defer func() {
		if oldGlSecret != "" {
			os.Setenv("GITLAB_SECRET", oldGlSecret)
		}
	}()

	reqBody := []byte("{}")
	var ts = httptest.NewServer(http.HandlerFunc(handler.Capture))
	defer ts.Close()

	data := bytes.NewBuffer(reqBody)

	// pass signature verification
	req, err := http.NewRequest("POST", ts.URL, data)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("X-Gitlab-Event", "ping")
	req.Header.Add("X-Gitlab-Token", "secret")

	res, err := ts.Client().Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(res.Status)
		t.Error("gitlab signature verification success, but still rejected")
	}

	// failed signature verification
	req, err = http.NewRequest("POST", ts.URL, data)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("X-Gitlab-Event", "ping")
	req.Header.Add("X-Gitlab-Token", "somethingelse")

	res, err = ts.Client().Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusForbidden {
		t.Error(res.Status)
		t.Error("gitlab signature match failed, but still forwarded")
	}
}

func TestNoSignatureVerification(t *testing.T) {
	oldGhSecret := os.Getenv("GITHUB_SECRET")
	oldGlSecret := os.Getenv("GITLAB_SECRET")
	syscall.Unsetenv("GITHUB_SECRET")
	syscall.Unsetenv("GITLAB_SECRET")

	defer func() {
		if oldGhSecret != "" {
			os.Setenv("GITHUB_SECRET", oldGhSecret)
		}
		if oldGlSecret != "" {
			os.Setenv("GITLAB_SECRET", oldGlSecret)
		}
	}()

	var ts = httptest.NewServer(http.HandlerFunc(handler.Capture))
	defer ts.Close()

	data := bytes.NewBuffer([]byte("{}"))

	// GitHub
	req, err := http.NewRequest("POST", ts.URL, data)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("X-GitHub-Event", "ping")

	res, err := ts.Client().Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(res.Status)
		t.Error("github signature verification disabled, but still rejected")
	}

	// GitLab
	req, err = http.NewRequest("POST", ts.URL, data)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("X-Gitlab-Event", "ping")

	res, err = ts.Client().Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error(res.Status)
		t.Error("gitlab signature verification disabled, but still rejected")
	}
}
