![build-glyph](figures/build-glyph.png)
# Build

**Abstract**: Here you find informations about the `Eutherpe`'s build. This document
is intended to development folks.

## Topics

- [`How I can build the Eutherpe binary?`](#how-i-can-build-the-eutherpe-binary)
- [`How I run the tests?`](#how-i-run-the-tests)
- [`I have installed go in Raspberry Pi but I cannot run it`](#i-have-installed-go-in-raspberry-pi-but-i-cannot-run-it)
- [`Using the GNUMake rules`](#using-the-gnumake-rules)

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

### Using the GNUMake rules

It is a kind of uncommon a `software` written in `Golang` have automations based on `make`, but
`Eutherpe` it is about much more than just compiling, testing and installing. There are a bunch
of operations executed between the compilation and installation. Taking it into consideration,
I made up my mind to automate certain common operations that developers need to do.

This operations following automated in `src/Makefile`.

Until now there are four implemented `rules`:

- `eutherpe`
- `tests`
- `bootstrap`
- `update`

If you want to build the `eutherpe` application, being into the toplevel `src` directory, execute:

```
$ make eutherpe
```

This `rule` is capable of setting up your `GOENV` on its own.

Do you want to run the tests (ignoring the cached stuff), from the toplevel `src` directory, execute:

```
$ make tests
```

Does the development environment is clean, without a previous `Eutherpe` installation? You can
run `bootstrap` in order to install and configure everything. It is necessary to be `root` when
running this `rule`, being into the toplevel `src` directory, execute:

```
# make bootstrap
```

Have you done an adjustment in the application and you are wanting to update the `binary`
used by the `service` that is currently running in your development environment? Being `root` and
from the toplevel `src` directory, execute `update`:

```
# make update
```

> [!IMPORTANT]
> This update only refreshes the `Eutherpe` application and its `web assets` that compounds the
> `Eutherpe`'s `core`. If you have done changes in `services` or in `shell scripts` it is more
> indicated running `bootstrap`. Doing it, the `bootstrap` rule will detect that already exists
> a previous installation and it will not reinstall `bluealsa` nor recreate the `eutherpe` user
> and reboot the system neither. Even so, it will refresh the `services` and the `shell scripts`,
> too.

[`Back`](#topics)
