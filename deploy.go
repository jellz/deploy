package main

import _ "github.com/joho/godotenv/autoload" // Load .env variables

import (
	"github.com/GitbookIO/go-github-webhook"
	"os"
	"os/exec"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"strings"
)

func main() {
	log.Print(os.Getenv("WEBHOOK_SECRET"))
	log.Print("Listening on port ", os.Getenv("PORT"))
	if err := http.ListenAndServe(":" + os.Getenv("PORT"), HandleWebhooks(os.Getenv("WEBHOOK_SECRET"))); err != nil {
		log.Fatalf("Error: %s", err)
	}
}

// HandleWebhooks : Receive webhooks
func HandleWebhooks(secret string) http.Handler {
	return github.Handler(secret, func (event string, payload *github.GitHubPayload, req *http.Request) error {
		log.Printf("Received %s for %s", event, payload.Repository.FullName)

		if event == "push" {
			projects := getProjects()
			for _, p := range projects {
				if (strings.ToLower(p.Repo) == strings.ToLower(payload.Repository.FullName)) {
					log.Printf("[%s] Received push event, attempting to pull...", p.Repo)

					cmdReset := exec.Command("git", "-C", p.Directory, "reset", "--hard")
					errReset := cmdReset.Run()
					if errReset != nil {
						log.Fatalf("[%s] Error resetting: %s", p.Repo, errReset)
					}
					log.Printf("[%s] Reset", p.Repo)

					cmdPull := exec.Command("git", "-C", p.Directory, "pull")
					errPull := cmdPull.Run()
					if errPull != nil {
						log.Fatalf("[%s] Error pulling: %s", p.Repo, errPull)
					}
					log.Printf("[%s] Pulled", p.Repo)

					log.Printf("[%s] Pulled latest code", p.Repo)
				}
			}
			
		}

		return nil
	})
}

// getProjects : Get JSON from file and turn it into []Project
func getProjects() []Project {
	var projects []Project
	file, err := ioutil.ReadFile("projects.json")
	if err != nil { log.Fatal(err) }
	err = json.Unmarshal(file, &projects)
	if err != nil { log.Fatal(err) }
	return projects
}

// Project : Project that contains repo name and directory
type Project struct {
	Repo      string `json:"repo"`
	Directory string `json:"directory"`
}