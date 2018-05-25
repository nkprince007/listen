package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Provider string

const (
	GitHub Provider = "github"
	GitLab Provider = "gitlab"
	None   Provider = ""
)

type response struct {
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func encodeResponse(desc string, status bool) []byte {
	rsp := &response{desc, status}
	enc, _ := json.Marshal(rsp)
	return enc
}

func recognizeWebhook(h http.Header) (string, Provider) {
	event := h.Get("X-GitHub-Event")
	if len(event) > 0 {
		return event, GitHub
	}

	event = h.Get("X-Gitlab-Event")
	if len(event) > 0 {
		return event, GitLab
	}

	return "", None
}

func handleGitHub(r *http.Request, event string) {

}

func handleGitLab(r *http.Request, event string) {

}

// Capture accepts webhooks and multicasts events.
func Capture(w http.ResponseWriter, r *http.Request) {
	// verify whether it is a POST request
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write(encodeResponse("only POST requests allowed", false))
		return
	}

	event, provider := recognizeWebhook(r.Header)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write(encodeResponse("could not decode request body", false))
		return
	}

	switch provider {
	case GitHub:
		requestSign := r.Header.Get("X-Hub-Signature")
		if githubSecret, ok := os.LookupEnv("GITHUB_SECRET"); ok {
			hash := hmac.New(sha1.New, []byte(githubSecret))
			hash.Write(body)
			generatedSign := hex.EncodeToString(hash.Sum(nil))

			if requestSign[5:] != generatedSign {
				w.WriteHeader(403)
				w.Write(encodeResponse("signature not matched", false))
				return
			}
		}

		handleGitHub(r, event)
	case GitLab:
		requestSign := r.Header.Get("X-Gitlab-Token")
		if gitlabSecret, ok := os.LookupEnv("GITLAB_SECRET"); ok {
			if requestSign != gitlabSecret {
				w.WriteHeader(403)
				w.Write(encodeResponse("signature not matched", false))
				return
			}
		}

		handleGitLab(r, event)
	default:
		// no provider matched
		w.WriteHeader(400)
		w.Write(encodeResponse("no matching providers implemented", false))
		return
	}

	w.WriteHeader(200)
	w.Write(encodeResponse("webhook accepted", true))
}
