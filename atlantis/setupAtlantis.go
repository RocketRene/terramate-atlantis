package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/go-github/v61/github"
	"github.com/joho/godotenv"
)

type ngrokResp struct {
	Tunnels []struct {
		PublicURL string `json:"public_url"`
	} `json:"tunnels"`
}

func startNgrok() string {
	cmd := exec.Command("sh", "-c", "ngrok http 4141 > /dev/null &")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("starting Ngrok on http port 4141")

	time.Sleep(2 * time.Second)

	resp, err := http.Get("http://localhost:4040/api/tunnels")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var respone ngrokResp

	err = json.Unmarshal(body, &respone)
	if err != nil {
		log.Fatal(err)
	}

	pubUrl := respone.Tunnels[0].PublicURL

	return pubUrl
}

func getClient(token string) *github.Client {
	client := github.NewClient(nil).WithAuthToken(token)
	return client
}

func listWebhooks(client *github.Client, owner, repo string) ([]*github.Hook, error) {
	ctx := context.Background()
	hooks, _, err := client.Repositories.ListHooks(ctx, owner, repo, nil)
	return hooks, err
}

func updateWebhook(
	client *github.Client,
	owner, repo string,
	hookID int64,
	newURL string,
	secret string,
) error {
	ctx := context.Background()
	contentType := "json"
	// Initialize HookConfig as a map, then convert it to *github.HookConfig.
	config := github.HookConfig{
		URL:         &newURL,
		ContentType: &contentType,
		Secret:      &secret,
	}

	hook := &github.Hook{
		Config: &config,
	}

	_, _, err := client.Repositories.EditHook(ctx, owner, repo, hookID, hook)
	if err != nil {
		log.Printf("Failed to update webhook: %v", err)
		return err
	}

	log.Println("Webhook updated successfully")
	return nil
}

func startAtlantis(url string) {
	cmdString := fmt.Sprintf(`atlantis server \
--atlantis-url="%s" \
--gh-user="%s" \
--gh-token="%s" \
--gh-webhook-secret="%s" \
--repo-allowlist="%s" \
--repo-config="%s" `, url, os.Getenv("ATLANTIS_GH_USER"), os.Getenv("ATLANTIS_GH_TOKEN"), os.Getenv("ATLANTIS_GH_WEBHOOK_SECRET"), os.Getenv("ATLANTIS_REPO_ALLOWLIST"), os.Getenv("ATLANTIS_REPO_CONFIG"))

	cmd := exec.Command("sh", "-c", cmdString)
	cmd.Stdout = os.Stdout // Log stdout to os.Stdout
	cmd.Stderr = os.Stderr // Log stderr to os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start Atlantis: %v", err)
		os.Exit(1)
	}
	log.Println("Atlantis server started")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup signal handling to gracefully shutdown Ngrok and Atlantis
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println("Shutting down Ngrok and Atlantis...")
		// Commands to stop Ngrok and Atlantis
		exec.Command("pkill", "-f", "ngrok").Run()
		exec.Command("pkill", "-f", "atlantis").Run()
		os.Exit(0)
	}()

	pubUrl := startNgrok()
	pubUrlWh := pubUrl + "/events"
	token := os.Getenv("ATLANTIS_GH_TOKEN")
	owner := os.Getenv("ATLANTIS_GH_USER")
	repo := os.Getenv("GITHUB_REPO")
	gh := getClient(token)
	secret := os.Getenv("ATLANTIS_GH_WEBHOOK_SECRET")

	hooks, err := listWebhooks(gh, owner, repo)
	if err != nil {
		log.Fatal(err)
	}
	hookID := hooks[0].GetID()
	err = updateWebhook(gh, owner, repo, hookID, pubUrlWh, secret)
	if err != nil {
		log.Fatal(err)
	}

	startAtlantis(pubUrl)
}
