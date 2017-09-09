# serve ⚡️

A simple binary written in go for serving your current directory as static content over HTTP from the terminal.  Just `go get` and use.

Free, open source, fast, and simple.


![serve](https://github.com/integrii/serve/blob/master/tutorial.gif?raw=true "Why did a static command line serve this good not exist yet?")


# Installation

- [Have go installed](https://golang.org) _(duh?)_
- Install with go: `go get github.com/integrii/serve`
- If your `$PATH` environment variable has `$GO?BIN` in it, then you can now simply use `serve`
  - Otherwise, you will need to run `$GOBIN/serve`

# Usage

- Run `serve`.
- See results at [http://127.0.0.1:8000](http://127.0.0.1:8000)
- Hit `Command-C` to stop gracefully


# Options

Run `serve --help` to see command line options:

```
Usage of serve:
  -l string
        The address for the server to listen on. Examples: :80, 127.0.0.1:8000 (default ":8000")
  -p string
        The path for the server to serve. (default "/Current/directory")
```
