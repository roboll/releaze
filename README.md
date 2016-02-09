# releaze

Embed static (compile time) build info in go executables.

## How It Works

_releaze_ embeds static build info in executables using the `-ldflags` linker argument. To embed,
use `releaze ldflags` at compile time to generate the necessary flags.

`go build -ldflags $(releaze ldflags) entrypoint.go`

or with the version, too.

`go build -ldflags "$(releaze version -v v1.0) $(releaze ldflags)" entrypoint.go`

Later, at runtime, variables are accessible through the _releaze_ api.

```
info := releaze.Get()
version := info.Version()
commit := info.Scm().Commit()
branch := info.Scm().Branch()
```

You'll notice that __all data is accessed via function calls__. _releaze_ exposes a read-only view
of the embedded data - it is not exposed as a struct to avoid modifications.

Data is captured in the call to `releaze ldflags`, where a set of 'collectors' are used to pull
available data from the environment. For example, the default data set includes the scm commit and
the branch name (git support now, more easily pluggable).

## Get It

`go get github.com/roboll/releaze`

_releaze_ uses `GO15VENDOREXPERIMENT=1` vendoring via [gvt](https://github.com/FiloSottile/gvt).

## Extra Fun

### cli

The [`codegangsta/cli`](https://github.com/codegangsta/cli) package is great, here is a command to
dump build info from _releaze_.

```
app := cli.NewApp()
app.Commands = append(app.Commands, releaze.CliCommand())
app.Run(os.Args)
```


```
$ go run -ldflags "$(releaze ldflags) $(releaze version -v version)" main.go buildinfo
{
  "version": "version",
  "scm": {
    "commit": "commit",
    "branch": "branch"
  }
}
```

### http

Here is a simple http handler to respond with the build info in json.

```
http.HandleFunc("/buildinfo", releaze.HttpHandler)
http.ListenAndServe(":8000", nil)
```
