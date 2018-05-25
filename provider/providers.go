package provider

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
)

// Provider defines the capabilities of webhook providers.
type Provider interface {
	Name() string
}

// GitHub defines the GitHub webhook provider.
type GitHub struct{}

// Name returns "github"
func (*GitHub) Name() string { return "github" }

// ServeHTTP responds to incoming webhooks
func (g *GitHub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		reason := "invalid request format, should be a map<string, Any>"
		http.Error(w, reason, http.StatusBadRequest)
		return
	}

	// no chance of error, since it was just marshalled
	body, _ := json.Marshal(data)

	// verify github signature
	if secret, ok := os.LookupEnv("GITHUB_SECRET"); ok {
		hash := hmac.New(sha1.New, []byte(secret))
		hash.Write(body)
		generated := "sha1=" + hex.EncodeToString(hash.Sum(nil))

		if r.Header.Get("X-Hub-Signature") != generated {
			http.Error(w, "signature not matched", http.StatusForbidden)
			return
		}
	}

	// TODO: identify the event type and then publish data to broker queue

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("webhook accepted"))
}

// GitLab defines the GitLab webhook provider.
type GitLab struct{}

// Name returns "gitlab"
func (*GitLab) Name() string { return "gitlab" }

// ServeHTTP responds to incoming webhooks
func (g *GitLab) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		reason := "invalid request format, should be a map<string, Any>"
		http.Error(w, reason, http.StatusBadRequest)
		return
	}

	// verify github signature
	if secret, ok := os.LookupEnv("GITLAB_SECRET"); ok {
		if r.Header.Get("X-Gitlab-Token") != secret {
			http.Error(w, "signature not matched", http.StatusForbidden)
			return
		}
	}

	// TODO: identify the event type and then publish data to broker queue

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("webhook accepted"))
}
