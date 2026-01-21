package demo

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"strings"

	"github.com/google/go-github/v71/github"
)

//go:embed all:helm_guestbook
var helmGuestbookFS embed.FS

// CopyEmbeddedDirToRepo copies the embedded helm_guestbook directory to a GitHub repository
func CopyEmbeddedDirToRepo(
	ctx context.Context,
	client *github.Client,
	destOwner, destRepo, destPath string,
) error {
	err := fs.WalkDir(helmGuestbookFS, "helm_guestbook", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("walk error at %s: %w", path, err)
		}
		if d.IsDir() {
			return nil
		}

		content, readErr := helmGuestbookFS.ReadFile(path)
		if readErr != nil {
			return fmt.Errorf("failed to read embedded file %s: %w", path, readErr)
		}

		// Convert path: "helm_guestbook/Chart.yaml" → "helm-guestbook/Chart.yaml"
		relativePath := strings.TrimPrefix(path, "helm_guestbook/")
		destFilePath := destPath + "/" + relativePath

		opts := &github.RepositoryContentFileOptions{
			Message: github.Ptr("Add " + destFilePath),
			Content: content,
		}
		_, _, createErr := client.Repositories.CreateFile(
			ctx, destOwner, destRepo, destFilePath, opts,
		)
		if createErr != nil {
			return fmt.Errorf("failed to create %s: %w", destFilePath, createErr)
		}

		fmt.Printf("✓ %s\n", destFilePath)
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to walk embedded directory: %w", err)
	}
	return nil
}
