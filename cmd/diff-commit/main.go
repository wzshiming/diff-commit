package main

import (
	"context"
	"os"
	"log/slog"
	"fmt"
	"io"
	"strings"

	"github.com/wzshiming/diff-commit/prompts"
	"github.com/wzshiming/gh-gpt/pkg/run"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  diff-commit <patch>")
	fmt.Println("  git diff | diff-commit -")
	fmt.Println("  git diff --cached | diff-commit -")
	fmt.Println("  git show HEAD --patch | diff-commit -")
	os.Exit(0)
}

func openFile(f string) ([]byte, error) {
	if f == "-" {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(f)
}

func main() {
	ctx := context.Background()
	if len(os.Args) <= 1 {
		usage()
	}

	patch, err := openFile(os.Args[1])
	if err != nil {
		usage()
		slog.Error("read file", "err", err)
		os.Exit(1)
	}

	if len(patch) == 0 {
		usage()
	}

	summary, err := run.Run(ctx, prompts.SummarizeFileDiff(string(patch)))
	if err != nil {
		slog.Error("summarize file diff", "err", err)
		os.Exit(1)
	}
	summary = strings.TrimSpace(summary)

	title, err := run.Run(ctx, prompts.SummarizeTitle(summary))
	if err != nil {
		slog.Error("summarize title", "err", err)
		os.Exit(1)
	}
	title = strings.TrimSpace(title)

	kind, err := run.Run(ctx, prompts.ConventionalCommit(summary))
	if err != nil {
		slog.Error("conventional commit", "err", err)
		os.Exit(1)
	}
	kind = strings.TrimSpace(kind)

	fmt.Printf("%s: %s\n\n%s\n", kind, title, summary)
}
