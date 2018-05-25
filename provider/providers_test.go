package provider_test

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gitlab.com/gitmate-micro/listen/provider"
)

func TestNames(t *testing.T) {
	gh, gl := provider.GitHub{}, provider.GitLab{}
	if gh.Name() != "github" {
		t.Error("GitHub provider does not have correct name.")
	}
	if gl.Name() != "gitlab" {
		t.Error("GitHub provider does not have correct name.")
	}
}

func TestGitHubInvalidRequestFormat(t *testing.T) {
	var ts = httptest.NewServer(&provider.GitHub{})
	defer ts.Close()

	rsp, err := http.Post(ts.URL, "application/json", nil)
	if err != nil {
		t.Error(err)
	}

	if rsp.StatusCode != http.StatusBadRequest {
		t.Error(rsp.Status)
		t.Error("invalid data is accepted!?")
	}
}

func TestGitLabInvalidRequestFormat(t *testing.T) {
	var ts = httptest.NewServer(&provider.GitLab{})
	defer ts.Close()

	rsp, err := http.Post(ts.URL, "application/json", nil)
	if err != nil {
		t.Error(err)
	}

	if rsp.StatusCode != http.StatusBadRequest {
		t.Error(rsp.Status)
		t.Error("invalid data is accepted!?")
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
	hash := hmac.New(sha1.New, []byte("secret"))
	hash.Write(reqBody)
	sign := hex.EncodeToString(hash.Sum(nil))

	var ts = httptest.NewServer(&provider.GitHub{})
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
	data = bytes.NewBuffer(reqBody)
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
	var ts = httptest.NewServer(&provider.GitLab{})
	defer ts.Close()

	// pass signature verification
	data := bytes.NewBuffer(reqBody)
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
	data = bytes.NewBuffer(reqBody)
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
