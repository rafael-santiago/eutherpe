# Eutherpe Manual

# Topics

- [What is it?](#what-is-it)
- [Features](#features)
- [Bootstrapping](#bootstrapping)
    - [Bootstrap who?!](#bootstrap-who)
    - [...wow! an ethereal jukebox has really sprouted into my LAN!](#wow-an-ethereal-jukebox-has-really-sprouted-into-my-lan)
- [Knowing the screens](#knowing-the-screens)
    - [The Music screen](#the-music-screen)
    - [The Collection screen](#the-collection-screen)
    - [The Playlist screen](#the-playlist-screen)
    - [The Storage screen](#the-storage-screen)
    - [The Bluetooth screen](#the-bluetooth-screen)
    - [The Settings screen](#the-settings-screen)

## What is it?

`Eutherpe` is a kind of jukebox ethereal wiring. Using `Eutherpe` you will be capable of
listening your `MP3` without needing to log on in a `stream` service external to your `LAN`.
Here you are who `100%` rules.

The basic idea is about you having your tunes recored in a `pendrive`, so you plug this `pendrive`
to the computer is running `Eutherpe`, you connect `Eutherpe` to a `Bluetooth` speaker or headphone
and done. Go listening to your music and period.

You are able to control `Eutherpe` from any `web browser`, just needing to connect to the
computer address that are serving `Eutherpe` in your `LAN`.

[`Back`](#topics)

## Features

When writing `Eutherpe` my main concern was `minimalism`. The idea is let you the maximum
possible far from screens and more imerse in the simple act of being listen to your music.
From some years until now (2024), due to some much rush of information, we have forgetting
about what is listening, inclusive music. F_cking `FOMO` I just want to slow down surrounded
by what I already love.

So with `Eutherpe` you do not have ads, suggestions, no interruptions. It is only about you
and your tunes. You have your `MP3` collection, you want to listen to it and done.

The main `Eutherpe` features are:

- Minimalistic `Web` stuff. There is no `pharaonic` frameworks here. `HTML`, `javascript`,
  `Golang` and period. I did not bring and **I do not want to bring** the `Desktop` to your
  `web` browser.
- The connection to the `Web` interface can be password authenticated or not. All of it can
  be easily configured from the `Web` interface itself.
- The connection to the `player` (done via `web browser`) can be `HTTP` or `HTTPS`, All of it
  can be easily configured from the `Web` interface itself.
- You are able to connecto to your `Eutherpe` device without needing to run crazy commands to
  figure out the `IP` address of your ethereal `jukebox`.
- You can give tidiness to your music collection. Even it being untidy. When scanning the
  `pendrive` seeking to music files, `Eutherpe` is capable of reading some `metatags` from your
 `MP3` by listing it automagically tidied up by `Artist/Album`. Albums are listed from the latest
 to the oldies.
- You can create `playlists`.
- The `playlists` are associated with the `pendrive`, if you plug another one, these playlists
  will vanish away, but if you plug the associated `pendrive` back, they will show up back, too.
- Each tune can be tagged up.
- `Tags` are basically keywords: `80s`, `best solos`, `stoner rock`, `mojo`, `blues`, `jazz`,
  `friday night`, `grumpy monday` etc.
- You can ask to `Eutherpe` to play `n` songs that fit up to a list of `tags` given by you.
- The essential dynamics of `Eutherpe`'s player is of a `jukebox`, then you actively select
  what you want to listen to.
- You can connect `Eutherpe` to a output `Bluetooth` sound device (speakears or headphones).
- `Eutherpe` supports `MP3`, `MP4` and `M4A`.
- The last usage session is save, so your selections stay saved between listening sessions.
  It does mean that you can continue your audition from the point that you stopped.
- `Web` interface renders pretty well in `desktop` or `mobile` devices.
- It works based on `Linux`.
- On `Windows` you can use `Eutherpe` from a `virtual machine`.
- You still can use `Eutherpe` embedded in a `raspeberry-pi` and in this case it will be called
  `Euther-PI`.

[`Back`](#topics)

## Bootstrapping

### Bootstrap who?

If you are not from the crazy gang of people that program computers, maybe it will be nice to
understand what a hell means `bootstrapping`.

Well, it is a term that we use to indicate that we will provide all necessary to start playing.
It would be similar to `free solo` on climbing by fixing all crampons by making easy and safe
future ascendings. It seems that the term is original from a story called *The surprising Adventures
of Baron Munchausen*, written by Rudolf Erich Raspe. In this story the baron was saved himself
and your horse from a swamp by pulling itself and his horse by his hair and its pigtail
(so true...). Some folks arguments about other origins of this term, but here it is not the case.
The essence of the idea is that in the `bootstrapping` the goal is to overcome a swampy
part, dry in resources to at the end to arrive in a solid ground. All by using practically
nothing or few resources.

In the `Eutherpe`'s case, the `jukebox` uses some dependencies to make all its ecosystem. Those
dependencies are applications and resources that `Eutherpe` uses under the hood to provide
all infrastructure to you nod your head, dance naked at your living room or simply having an
ambient music.

`Eutherpe` was developed based on `Debian Linux` distribution. So the `bootstrapping`
**needs to be done from a `Debian` installation (I advise)**.

The indicated is you make a `minimal` installation (without `desktop` resources). Just
a shell based environment. Why? Well, what happens is that when you install so much unecessary
things, the whole system will start be busing by taking care of those white elephants and,
this waste can impact the listening of your beloved tunes.

I indicate `Debian 11` (only text-based, with networking, because when `bootstraping` we
need to download and install some specific packages).

The `Debian 11` was my distro of choice to develop `Eutherpe`. Previously I tried `12` but
it sucked and I was also using `Gnome` and all those trinkets *- that for `Eutherpe`'s goal
are essentially pointless*.

Now maybe you are thinking: *"- Gosh to do the 'bootstrapping' shall be hard as a rock..."*.
Take it easy, it is not!

[`Back`](#topics)

### ...wow! an ethereal jukebox has really sprouted into my LAN!

[`Back`](#topics)
