# diff-commit

This is a simple tool to generate commit messages from the diff.

## Requirements

- Refer to [gh-gpt](https://github.com/wzshiming/gh-gpt)

## Usage

Generate commit message from the diff.

``` bash
diff-commit <patch>
```

Generate commit message from the diff and append to the commit message.

``` bash
git commit -m "$(git diff --cached | diff-commit -)"
```

Generate commit message from last commit and amend to the commit message.

``` bash
git commit --amend -m "$(git show HEAD --patch | diff-commit -)"
```
