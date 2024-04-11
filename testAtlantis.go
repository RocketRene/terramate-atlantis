package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
)

func runCmd(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to execute %s: %v", name, err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomName := strings.ToLower(faker.FirstName())
	// Switch to the main branch, fetch, and pull the latest updates
	log.Println("Switching to the main branch and updating...")
	runCmd("git", "checkout", "main")
	runCmd("git", "fetch")
	runCmd("git", "pull")

	// Generate a random first name using faker

	// Append a random 3-digit number to ensure uniqueness
	randomNumber := rand.Intn(900) + 100 // Ensures a number in the range [100, 999]
	branchName := fmt.Sprintf("bucket-%s%d", randomName, randomNumber)

	log.Printf("Generated branch name: %s\n", branchName)

	// Create a new Git branch named `bucket-{name}{number}`
	runCmd("git", "checkout", "-b", branchName)

	// Open the `main.tf` file and replace the bucket name using regex
	filePath := "./stacks/s3/movies/main.tf"
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	content := string(contentBytes)

	regexPattern := `bucket = "terramate-rene-(.+?)-movies"`
	re := regexp.MustCompile(regexPattern)
	newBucketName := fmt.Sprintf(`bucket = "terramate-rene-%s%d-movies"`, randomName, randomNumber)
	newContent := re.ReplaceAllString(content, newBucketName)

	if err := os.WriteFile(filePath, []byte(newContent), 0666); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}

	// Commit the change
	runCmd("git", "add", filePath)
	commitMessage := fmt.Sprintf("Update bucket name to %s%d", randomName, randomNumber)
	runCmd("git", "commit", "-m", commitMessage)

	// Push the new branch to remote
	log.Println("Pushing changes to remote...")
	runCmd("git", "push", "--set-upstream", "origin", branchName)

	// Attempt to create a PR again
	log.Println("Creating a PR...")
	prTitle := fmt.Sprintf("Update bucket name to %s%d", randomName, randomNumber)
	prBody := "Automated PR to update S3 bucket name."
	runCmd(
		"gh",
		"pr",
		"create",
		"--base",
		"main",
		"--head",
		branchName,
		"--title",
		prTitle,
		"--body",
		prBody,
	)

	log.Println("PR created successfully.")
}
