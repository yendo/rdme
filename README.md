# rdme

Discover and run code snippets directly from your `README.md` or other markdowns (defaults to local `README.md`).

rdme makes a best effort approach to extracts all code snippets defined in code blocks and allowing to explore and execute them. rdme is currently in early alpha.

You can execute commands from a different directory using a `--chdir` flag.
To select a different file than `README.md`, use `--filename`.

## Installation

If you have Go developer tools installed, you can install it with `go install`:

```sh
$ go install github.com/yendo/rdme@latest
```

## Contributing & Feedback

Let us know what you think via GitHub issues or submit a PR. Join the conversation [on Discord](https://discord.gg/MFtwcSvJsk). We're looking forward to hear from you.

## LICENCE

Apache License, Version 2.0
