package main

import (
	"context"
	"os"
	"log/slog"
	"fmt"
	"io"

	"github.com/wzshiming/diff-commit/prompts"
	"github.com/wzshiming/diff-commit/gpt"
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

	summary, err := gpt.Generate(ctx, prompts.SummarizeFileDiff(string(patch)))
	if err != nil {
		slog.Error("summarize file diff", "err", err)
		os.Exit(1)
	}

	title, err := gpt.Generate(ctx, prompts.SummarizeTitle(summary))
	if err != nil {
		slog.Error("summarize title", "err", err)
		os.Exit(1)
	}

	kind, err := gpt.Generate(ctx, prompts.ConventionalCommit(summary))
	if err != nil {
		slog.Error("conventional commit", "err", err)
		os.Exit(1)
	}

	fmt.Printf("%s: %s\n\n%s\n", kind, title, summary)
}
