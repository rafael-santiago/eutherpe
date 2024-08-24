![build-glyph](figures/build-glyph.png)
# Build

**Abstract**: Here you find informations about the `Eutherpe`'s build. This document
is intended to development folks.

## Topics

- [`How I can build the Eutherpe binary?`](#how-i-can-build-the-eutherpe-binary)
- [`How I run the tests?`](#how-i-run-the-tests)
- [`I have installed go in Raspberry Pi but I cannot run it`](#i-have-installed-go-in-raspberry-pi-but-i-cannot-run-it)

### How I can build the Eutherpe binary?

The `Eutherpe` binary is about a program written in `Golang`. Thus, the `build`
is pretty straightforward. Being into the toplevel directory `src`, run:

```
# go build
```

[`Back`](#topics)

### How I run the tests?

Wow! I are minding about the tests. It is good! It is a characteristic of pragmatic,
perfectionist developers.

Being the `Eutherpe`'s core a `Golang` application, the tests are runned by calling
the following command (being intro the toplevel `src` directory):

```
# go test internal/...
```

If you want to get more information about what is actually happening during the process, execute:

```
# go test internal/... -v
```

> [!TIP]
> `Golang` has a kind of annoying behavior that is caching the tests results.
> If you want to force it execute the tests every time, before running them, clear
> the cache in the following way:
>
> ```
> # go clean -testcache
> ```

[`Back`](#topics)

### I have installed go in Raspberry Pi but I cannot run it

If you accomplished the [`bootstrapping`](MANUAL-EN.md#bootstrapping), give it a try to the
following command:

```
# source /etc/profile.d/goenv.sh
```

[`Back`](#topics)
