You are an expert programmer, and you are trying to summarize a code change.
You went over every file that was changed in it.
For some of these files changes where too big and were omitted in the files diff summary.
Determine the best label for the commit.

Here are the labels you can choose from:

- BUILD: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
- CHORE: Updating libraries, copyrights or other repo setting, includes updating dependencies.
- CI: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, GitHub Actions)
- DOCS: Non-code changes, such as fixing typos or adding new documentation (example scopes: Markdown file)
- FEATURE: a commit of the type feat introduces a new feature to the codebase
- FIX: A commit of the type fix patches a bug in your codebase
- PERFORMANCE: A code change that improves performance
- REFACTOR: A code change that neither fixes a bug nor adds a feature
- STYLE: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
- TEST: Adding missing tests or correcting existing tests


THE FILE SUMMARIES:
###
{{ .summaryPoints }}
###

Based on the changes described in the file summaries, What's the best label for the commit? Your answer must be one of the labels above. Don't describe the changes, just write the label.
