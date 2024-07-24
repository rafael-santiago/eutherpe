# Running Eutherpe from a virtual machine

**Abstract**: This text is intended to show the main steps of how to configure `Eutherpe`
to be executed from a virtual machine. Here we take into consideration that the reader
has known already all main aspects of application by reading the [manual](MANUAL-EN.md).

## Topics

- [To whom is intended this text](#to-whom-is-intended-this-text)
- [Paraphrasing Dee-Dee Ramone: 1-2-3-4!](#paraphrasing-dee-dee-ramone-1-2-3-4)
- [The general idea of what we are going to do](#the-general-idea-of-what-we-are-going-to-do)
- [Creating the virtual machine from OVA](#creating-the-virtual-machine-from-ova)

### To whom is intended this text

I made my mind write this text thinking in users that do not have a computer running `Linux` natively,
but even so they are wanting to run `Eutherpe`. In general `Windows` users can take advantage
of this document, because until now `Eutherpe` is not `Windows` compatible. I believe that `OSX`
users, too.

> [!NOTE]
> If you already know about virtual machines, you already use them, it is pretty sure that this
> text will not provide anything new. The general idea is creating a virtual machine based on
> `Debian Linux` and then running the `Eutherpe bootstrapping`. If you already know how to do it
> on your own, this text is not for you. :wink:

[`Back`](#topics)

### Paraphrasing Dee-Dee Ramone: 1-2-3-4!

In order to create our `Eutherpe` `VM` (**V**irtual **M**achine) we will use the [`Vitualbox`](https://www.virtualbox.org/Downloads)
application.

You will need to install `Virtualbox` in your system and once it done, we will create a virtual
machine based on an `OVA` that I previously prepared, available at [here](https://drive.google.com/file/d/1oow6zf-eEzRe6ySEQTkC9orsYNtQ3ZjJ/view?usp=sharing).

Done! If you have installed `Virtualbox` and downloaded the `OVA` you have already everyhing you
need to continue.

> [!TIP]
> **Remark and tip**: I am not going to give you details about the virtualization concept norb
> virtual machines. Anyway, this is a very neat subject that you could find out a bunch of ideas
> and utilities to solve your day to day needs, I would suggest you to take `Eutherpe` as an excuse
> to go deeper in this subject. :dart:

[`Back`](#topics)

### The general idea of what we are going to do

Bear in mind `OVA` as a `zip` file or an installation program. This file gathers all files that
compounds the virtual machine. Imagine the computer wrapped into a box and then you only have
to get it out of the box, plug in some cables, power it on and done!

The `OVA Eutherpe` is a vitual machine based on `Debian 12` where I installed all basic dependencies
to have an operating system where I could be able to run `Eutherpe`. Yes, after installing this
well basic `Debian 12` I downloaded the `Eutherpe` sources and executed the `bootstrap` (if you have
not read the manual yet, [read it](MANUAL-EN.md) and *voilá*, I powered off the virtual machine
and then I created the `OVA`.

Taking this `OVA`, you will import it from your `Virualbox` and you will have the exact virtual
machine that I made, being needed to apply minor tunings to allow you access the `Eutherpe's
miniplayer` from your `web browser`. After that you will be able to nod your head or dance naked, 
or both :notes: :headphones: :guitar: :microphone: :dancer: :notes:

[`Back`](#topics)
