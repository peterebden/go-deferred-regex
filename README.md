# go-deferred-regex

This is a wrapper around the standard `regex` package to help defer initialisation
of the regex from init time to first use.

Regexes are commonly created at global scope with `MustCompile`, which is very useful and
convenient, but front-loads the regex parsing & creation to init time which must complete
before `main()` can run. This can add up over time since even moderately complex regexes
can cause hundreds of allocations, which may actually never be used if the relevant code
paths are never hit.

This package offers a `DeferredRegex` type which mimics the interface of `regexp.Regexp`
but initialises itself on first use.
Usage is simple:
```
var re = deferredregex.DeferredRegex{Re: `([0-9]+)\.([0-9]+)\.([0-9]+)`}

// later...
matches := re.FindStringSubmatch("1.2.3")
```
