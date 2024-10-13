![Ancient Greek Pegasus Coin / Public Domain](figures/pegasus_coin.png "Ancient Greek Pegasus Coin / Public Domain")

# Codestyling guide

**Abstract**: Coding is a kind of art-craft-engineering (who knows experssion) which involves
many cultural and idiomatics features. Due to it, similar to many other fields, `The Truth` is
such a big, biiiiiiiiiiig `Winged Unicorn`... Anyway, this tet intends to describer in an objective
way the main aspects of `my` current beloved unicorn. If you are up to programme with me, you
will enroll in a party that has already begun, this is my party.

# Topics

- [The general idea](#the-general-idea)
- [Basic constructions](#basic-constructions)
    - [if/else](#ifelse)
    - [for](#for)
    - [switches](#switches)
    - [func](#func)
- [Do not polute the go installation, forget about downloading packages](#do-not-polute-the-go-installation-forget-about-downloading-packages)
- [Do not code the build logics in the CI's yaml](#do-not-code-the-build-logics-in-the-cis-yaml)
- [A build needs documentation](#a-build-needs-documentation)
- [Definition of done](#definition-of-done)
- [Use inclusive and neutral language](#use-inclusive-and-neutral-language)

## The general idea

**Short and sweet**: `go to hell` with `go fmt` here. I do not have any intention of make
it a `go` package, thanks to that I do not need to swallow those annoying and imposed
formatting patterns from `go fmt`.

Here use `four spaces` for indentation and space for aligning, too. Configure your text editor to
replace `tab` for space, in this way all that `blah-blah-blah` and polemics about which is better,
**ends up**. Because, it does not matter which you use when indenting, the problem is always use
space when aligning. If the `tab` turns into space, it does not matter if you hit `tab` key or
space. At the end you will record a space in the file and it will never screw up the alignment.
(**Mic drop**). But, remember that **I use `4 space` to represent a `tab`**.

Eighty columns as the maximum is stingy and miserable. Do not exceed a hundred twenty columns,
we are in the 21st century and there is no way of back in time (`2024`).

Talking about indentation, avoid abusing `if` branching, please. You know, bad vices like the
following:

```go
if i == 42 {
    if j == 84 {
        ...
    }
}
```

Follows a more simple and ellegant option:

```go
if i == 42 && j == 84 {
    ...
}
```

Finally, write code in a deterministic way, `if` is not the purpose of programming, it never was
and never will be, `it is an exception case`, if you code by driving you way of thinking like
"`if` this and `if` that" it is pretty sure that you produce pieces of engineering that are
bug prone. The abusive user of the `if` is The End of Programming, avoid it. Build stuff, do
not mess up with nor destroy!

Fail as soon as possible, stop thinking that world is beautiful when programming, `always check on
fails`, Did fail? See ya! Who called the code has the responsibility of handle the error.


It is sad that `Go` language sometimes pushes the programmer to produce certain snippets a kind
of `noobs`, anyway avoid `branching` with no necessity, thanks! :wink:

Start on the code with the copyright disclaimer. Tidiness. Without it, sooner the code can
become a messy hell.

Use `camelCase` for variables and private function. `ThisOneForPublics`. The basic `Go` convention.

When commenting, commit yourself, use:

- `INFO(YourNameOrNickname): ...`, for information.
- `WARN(YourNameOrNickname): ...`, for warnings.
- `BUG(YourNameOrNickname): ...`, for known bugs.
- `TODO(YourNameOrNickname): ...`, for things to be done.

Fix bugs in the best way: by applying general solutions. Do not implement `bug` diversions,
it does not solve problems. If you are not understanding the fail's essency, you are not in
conditions of proposing any solution by now, think more about. Calm down! It is not a `hackathon`
:wink:

Implement tests. Run the tests during the `build`. You are not in a race, it does not mind
if it takes longer, if you have implemented something it needs to be tested, it is about
engineering, if do not want to test something, remove it from the code, we do not need
imprecise parts in the project, we seek for certainty.

Do not use third-party `Go` packages. Just its standard library, please. I do not want to
be unable of build my stuff because something vanished aways from that dark alley from in the
internet. Ha-ha-ha!

Last but not least: **Try to achieve wheel status on your stuff, please**.

Right, you do not understand what a hell is "wheel status", try to imagine the World without
wheels, or still try to improve on the essence of what is a wheel... Dig it? :wink: 

<center>
<h1>
<i>Even so, be suspicious of all code you produce. Nothing is perfect, but it can be near to be.</i>.
</h1>
</center>

[`Back`](#topics)

## Basic constructions

Follows the general idea of how you should format the basic constructions that
will compose your implementations.

In general:

- Make precedence explicit by using parentheses.
- Comment intrincate parts of your code, stop thinking you are a genious or poet, it does not exist
in programming, it is small talk of people that does not have guts on learning hard things that
really demands some virtuosity level :smirk:.
- Use english language when programming.
- Commit messages are in english, too.
- Documentations can be in portuguese-BR and/or english.

[`Back`](#topics)

### if/else

It is the way of formatting an `if/else`:

```go
if i == 42 {
    ...
} else if j == 42 {
    ...
} else {
    ...
}
```

It is a pity that `Go` does not allow this kind of alignment/formatting:

```go
if i == 42
    && j == 42 {
    ...
}
```

But if you want to break line during a logical expression, you can do in this way:

```go
if i == 42 &&
   j == 42 {
    ...
}
```

[`Back`](#topics)

### for

This is the way of formatting a `for`:

```go
for i := 1; i < 42; i++ {
    ...
}
```

[`Back`](#topics)

### switches

This is the way of formatting `switches`:

```go
switch i {
    case 0, 1, 2:
        ...
        break

    case 3, 4, 5,
         6, 7, 8:
        ...
        break
}
```

[`Back`](#topics)

### func

Functions follow this way:

```go
func privateFunc(i, j int, str string) ([]byte, error) {
    ...
}

func PublicFunc(i, j int, str string) ([]byte, error) {
    ...
}
```

[`Back`](#topics)

## Do not polute the go installation, forget about downloading packages

The idea is always using the `go` standard libraries. In the case of internal `Eutherpe` `packages`,
avoid installing them in your `GOROOT`, forget about this bullshit.

When initialising a new package under `src/internal` run the command to create the `go.mod`
(let's suppose that you will create the package `shoobeedooblaublau`):

```
# mkdir src/internal/shoobeedooblaublau
# cd src/internal/shoobeedooblaublau
# go mod init github.com/rafael-santiago/eutherpe/shoobeedooblaublau
```

In `src/go.mod` (notice, in `go.mod` immediately under `src`) add:

```
require internal/shoobeedooblaublau v1.0.0
replace internal/shoobeedooblaublau => ./internal/shoobeedooblaublau
```

Now, as instance, if the package `shoobeedooblaublau` uses `mplayer`, in `src/internal/shoobeedooblaublau/go.mod`
you need to add:

```
require internal/mplayer v1.0.0
replace internal/mplayer => ../mplayer
```

This care it is good because it makes everything more self-contained avoid messing your `GOOROOT`
with all that package-downloading-from-anywhere orgy that `golang` is hooked up on doing. I think
that this "feature" is a such bad idea, but, the secret on using technologies is divert from
the bad ideas and using only what matters, hoping that maintainers figure out how idiot is some
utilities (utils that really are inutils) and so purge this trinket out from new versions :wink:!

[`Back`](#topics)

## Do not code the build logics in the CI's yaml

It is sloppy and not so smart. Due to a simple reason: to be able to run the `software` build, it
would be needed all CI infrastructure to interpret and executing the `yaml`. Do you notice how
stupid it is? It takes from you all possibilities of getting the artifact asap.

A well written build is able to run where it should run. The programmer do not need to carrying
the build environment on its shoulders like a snail.

The build logics `must` be written in `scripts` that can be called from any environment with
all tools necessary to build the project up. Writing things out in this way, from the point of
view of the `CI`, it would be only about calling the script from some `yaml`.

Still, `yamls` are only minor scripts, used to provide the `CI` infrastructure, based on the
`status quo` of the source hosting service in use, `nothing more than it`. Nonetheless, the
`project` can be hosted in another service in the future, where all `yaml` trinket become
pointless (due to it `status quo`) Basing your build in a technology that can be easily
abandoned is an ode to the stupidity, naiveless and sloppiness. Do not do it, do not try,
I will not accept.

[`Back`](#topics)

## A build needs documentation

If you are unable to write down in natural language `how to run the build` of a `sofware`
taking into consideration the created infrastructure, for sure that it would be a mess.
Make it better until the point that you can describe in natural language how to run the build.

Suggesting users to read the `script` or even `CI` `yamls` is unacceptable, it would be
a proof of how incompetent we would be on maintaining the `software` `build`.

If the idea is to make something, let's make the things in the right way, if you do not know
how to do it, describe what you want to someone able and let's get the things done. Follow
up the process, so in this way you learn how to do.

[`Back`](#topics)

## Definition of done

A new feature is 

considered done when:

1. It does what it must do.
2. It does not add mess, confusion or even unstability nor bugs in the previous stuff.
3. It ships what it promises in a simple (but not hollow claimed "simple") way. In other words,
you have used Ockham razor principle on it.
4. It is well tested.
5. It is not tied up with some dependency that is out of the `Go` standard library.
6. The `CI` is passing.
7. It is well documented.
8. The commit that adds this new feature to the upstream is descriptive.
9. The commit message must use imperative form. Acting like you are giving commands to the version
control system. So `Giving commands to the version control system` is wrong. `Gives commands
to the version control system` still. `Give commands to the version control system`. Do not
be shy of being bossy with it! :wink:

[`Back`](#topics)

## Use inclusive and neutral language

*This is the only point that **there is no concession**. It is **NOT** about unicorns, really.
Follow it or farewell*

Always try o use inclusive and neutral language words/terms in your source codes and
documentations. If you find something that for you seems not to be so correct, please, let me know
by opening an issue and suggesting improvements. Thank you in advance!

In general avoid using colors to name what should be "good" or "bad". Outdated terms such as
`whitelist`/`blacklist` are deprecated/banned here. You should use `allowlist/denylist` or
anything more related to what you are really doing. Terms like `master/slave` are out, too.
You could use `main`, `secondary`, `next`, `trunk`, `current`, `supervisor`, `worker` in
replacement.

Do not use sexist and/or machist terms, too.

Again, if you have found some term(s) that for you is not much suitable, let me [know](https://github.com/rafael-santiago/eutherpe/issues)
by suggesting some edition(s), thank you in advance!

*-- Rafael*

[`Back`](#topics)
