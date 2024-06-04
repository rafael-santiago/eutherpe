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

`Eutherpe` is a kind of jukebox ethereal wiring. Using `Eutherpe` you will be able to
listen to your `MP3` without needing to log on in a `stream` service external to your `LAN`.
Here you are who `100%` rules.

The basic idea is about having your tunes recored in a `pendrive`, so you plug this `pendrive`
to the computer is running `Eutherpe`, you connect `Eutherpe` to a `Bluetooth` speaker or headphone
and done. Go listening to your music and period.

You are able to control `Eutherpe` from any `web browser`. Being only necessary to connect to the
computer address that is serving `Eutherpe` in your `LAN`.

[`Back`](#topics)

## Features

When writing `Eutherpe` my main concern was `minimalism`. The idea is let you the maximum
possible far from screens and more imerse in the simple act of listening to your music.
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
- The connection to the `player` (done via `web browser`) can be `HTTP` or `HTTPS`. All of it
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
- You can connect `Eutherpe` to an output `Bluetooth` sound device (speakears or headphones).
- `Eutherpe` supports `MP3`, `MP4` and `M4A`.
- The last usage session is saved, so your selections stay remain between listening sessions.
  It does mean that you can continue your audition from the point that you previously stopped.
- The `web` interface renders pretty well in `desktop` or `mobile` devices.
- It works based on `Linux`.
- On `Windows` you can use `Eutherpe` from a `virtual machine`.
- You still can use `Eutherpe` embedded in a `raspeberry-pi` and in this case it will be called
  `Euther-PI`.

[`Back`](#topics)

## Bootstrapping

### Bootstrap who?

If you are not from the crazy flying gang of people that program computers, maybe it will be
nice to understand what a hell means `bootstrapping`.

Well, it is a term that we use to indicate that we will provide all necessary to start playing.
It would be similar to `free solo` on climbing by fixing all crampons by making easy and safer
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
things, the whole system will start be busy by taking care of those white elephants and,
this waste of resources can impact the listening of your beloved tunes.

I indicate `Debian 11` (only `text-based`, with networking, because when `bootstraping` we
need to download and install some specific packages).

The `Debian 11` was my distro of choice to develop `Eutherpe`. Previously I tried `12` but
it sucked as hell and I was also using `Gnome` and all those trinkets *- that for `Eutherpe`'s
goal are essentially pointless*.

Now maybe you are thinking: *"- Gosh to do the 'bootstrapping' shall be hard as a rock..."*.
Take it easy, it is not!

[`Back`](#topics)

### ...wow! an ethereal jukebox has really sprouted into my LAN!

After you set up a minimal `Debian` installation, it is necessary you login in the system as `root`.

Once logged as `root`, you need to download `Eutherpe's` source and to be doing it you will need
to `git` application.

In order to install it you need the following command (if you installed `Debian 11` it is necessary
that you insert the `CD-ROM` or its .iso image):

```
# apt install git -y
```

After installing `git` it is time to download `Eutherpe's` source-code, by running the following:

```
# git https://github.com/rafael-santiago/eutherpe -b v1
```

A directory called `eutherpe` will be created, let's access it using the following commnad:

```
# cd eutherpe
```

Done! Now it is only about doing the `bootstrapping` :boot: by running the following command:

```
# ./bootstrapping.sh
```

You need to confirm with `y` or `Y` and all will be done. The dependencies will be installed,
system fine tunes will be done. You will get an output similar to it:

```
#########################
       ,|_|,   ,|_|,
       |===|   |===|
       |   |   |   |
       /  &|   |&  \
  _.-'`  , )* *( ,  `'-._ [ Eutherpe's Bootstrap ]
   `"""""`"`   `"`"""""`
#########################

Hi there! I am the Eutherpe's bootstrap! What I am intending to do: 

- Create an user "eutherpe";
- Add it to sudo's group;
- Install some system dependencies required to you play your beloved tunes;
- Install Golang to actually build Eutherpe's app;
- Install kernel headers to make easy any specific system tune that you may want to do;
- Create the default's USB mount point in /media/USB;
- Build up Eutherpe's app;
- Install whole Eutherpe's package it;


=== Okay, you are root user :) let's start...
=== Nice, eutherpe user already exists.
=== bootstrap info: Adding eutherpe to sudo group...
=== bootstrap info: granting eutherpe some nopasswd privileges...
=== bootstrap info: Done.
=== bootstrap info: Installing system dependencies...
*-- sudo already installed.
*-- git already installed.
*-- mc already installed.
*-- pulseaudio already installed.
*-- bluez already installed.
*-- pulseaudio-module-bluetooth already installed.
*-- ffmpeg already installed.
*-- alsa-utils already installed.
*-- wpasupplicant already installed.
*-- wireless-tools already installed.
=== bootstrap info: Done.
=== bootstrap info: Installing golang...
=== bootstrap info: Done.
=== bootstrap info: Setting up golang environment...
=== bootstrap info: Done.
=== bootstrap info: Installing kernel headers...
+-- linux-headers-5.10.0-27-amd64 installed.
+-- gcc installed.
+-- make installed.
+-- perl installed.
=== bootstrap info: Done.
=== bootstrap info: Creating USB storage mount point...
=== bootstrap info: Done.
=== bootstrap info: Now building Eutherpe...
=== bootstrap info: Done.
=== bootstrap info: Now installing Eutherpe...
=== bootstrap info: Done.
```
Once the `bootstrapping` done you will be able to acces your ethereal `jukebox`
by using a nearest `web browser` at [`http://eutherpe.local:8080/eutherpe`](http://eutherpe.local:8080/eutherpe).

*Voi lá!* :sunglasses:

Now it is time to know your no-frills-minimalist `juke`!

[`Back`](#topics)

## Knowing the screens

The initial `Eutherpe's` configuration it will be pretty basic:

- `HTTP` connection;
- No password authenticated;
- A straightforward `Jukebox` ready-to-go;

When connecting to `http://eutherpe.local:8080/eutherpe` you will see the screen illustrated by **Figure 1**.

Notice that it is about a straightforward screen. By having a left-sided menu where you access
the functions and configurations from your ethereal `juke`.

[`Back`](#topics)

### The Music screen

In the **Figure 2** you can see the `player` screen layout. It is pretty self explanatory.
Basically the screen offers the basic functions such as play, stop the music, go to next,
go to prior, move songs up/down in the repoduction list, activate `shuffle`, remove selected
songs from reproduction list and clear all reproduction list.

You still have a `slider` to control the volume and the possibility of activating the repetion
modes (all or one song).

In order to access the reproduction list, you need to click on `UP NEXT`. By the way, it is
time to teach you about a convention that I decided to follow.

If you are really crazy about music, maybe you have already noticed that, if not yet, now
you will go: all that is preceded by the sharp symbol, means that when you click it, more info
will be shown. Once it shown, by its side the flat symbol will turn into a flat symbol. Now
when clickin on this sharp symbol, the information will be hidden. The relationship between
`sharp/flat` is equals to `+/-`, :wink:!

Take a look at **Figura 3** a reproduction list shown in detailing mode. Noticed each song
have by their side a checkbox that you will check to apply some functions over those songs.
Functions like: remove and move. If you want to play a specific song just select it and
click `play`.

Done! Being `Eutherpe` a no-frills `jukebox`, you already know every single thing about how
to pilot its `player`. Nonetheless, you may have be intrigued: **how can I select songs to play?*

[`Back`](#topics)

## A tela Collection

Nessa tela você tem todo acesso às músicas que `Eutherpe` conseguiu escanear de seu dispositivo
de armazenamento. Entenda essa tela como a sua *estante*, onde todos seus álbuns seguem organizados.
Note, organizados. E isso mesmo se você for uma criatura desorganizada! :wink:

Dá uma conferida na **Figura 4** no `layout` dessa tela. Note que aqui eu também lancei mão daquela
convenção pirada *sustenido/bemol*...

Ainda em relação à tela ilustrada na **Figura 4**, perceba que a coleção fica organizada seguindo
`Artista/Álbum/Músicas` e os álbuns mais recentes vão sendo listados antes. Artistas, álbums e
músicas podem ser usados para compor uma seleção (pois eles possuem caixas de checagem ao lado
dos nomes, eu tenho certeza absoluta que você entendeu!).

Na parte inferior da tela você vai notar que existem botões bem autoexplicativos:

- `ADD TO NEXT` (adiciona a sua seleção ao final da fila de reprodução)
- `ADD TO UP NEXT` (fura a fila de reprodução e adiciona sua seleção atrás da posição atual de reprodução)
- `ADD TO PLAYLIST...` (adiciona a sua seleção a uma nova `playlist` ou a uma prévia)
- `ADD TAGS...` (marca sua seleção com `tags` fornecidas por você)
- `DEL TAGS...` (remove tags previamente adicionadas da seleção)
- `PLAY TAGGED...` (toca uma quantidade de músicas levando em conta `tags` fornecidas por você)

Aqui temos uma outra convenção. Tudo que for botão com reticências significa que ao clicar,
uma tela te pedindo mais info vai ser apresentada, nesse ponto lembre do `Guia do mochileiro
das galáxias`: `DON'T PANIC`! :wink:

A **Figura 5** exibe a tela que é apresentada ao clicar em `ADD TO PLAYLIST...`, você precisa
fornecer o nome da `playlist` a qual deseja adicionar a seleção e clicar em `ADD`. Se quiser
desistir da ideia, apenas clique em `BACK` que está tudo certo `Eutherpe` não brigará com você...

A **Figura 6** exibe uma tela similar a da **Figura 5**, porém aqui é para taguear a seleção.

Na **Figura 7** você confere a tela para remover `tags` previamente adicionadas à seleção. A ideia
geral é que você desmarque as `tags` que deseja remover e clique em `SAVE`. Se mudou de ideia e
não quiser remover nada, clica em `BACK`, `Eutherpe` é compreensiva...

A **Figura 8** ilustra o que você encontra ao clicar em `PLAY TAGGED...`. Você precisa fornecer
uma lista de `tags` separada por vírgula e ainda indicar o total de músicas que deseja escutar.
Para confirmar você clica em `PLAY` ou se desistiu e não quer mais clica em `BACK`, `Eutherpe`
vai entender e não reclamar...

[`Voltar`](#tópicos)

## A tela Playlists

A **Figura 9** traz a tela `Playlists`. A partir dessa tela você poderá pôr uma `playlist` ou
apenas algumas músicas dela para tocar, além de editá-la.

Note que cada `playlist` listada esconde suas músicas e que quando você clica sobre o nome dela
as músicas são listadas, confira a **Figura 10** onde a lista `HOW-LOU` (sim, eu curto
`Lou Rawls`!) está sendo detalhada.

As funções oferecidas por essa tela são:

- `REMOVE...` (você seleciona a playlist que deseja remover).
- `CLEAR...` (você seleciona a playlist que deseja limpar).
- `SONGS UP` (você moverá um ponto para cima dentro da listagem as músicas selecionadas).
- `SONGS DOWN` (você moverá um ponto para baixo dentro da listagem as músicas selecionadas).
- `REMOVE SONGS...` (você seleciona a(s) música(s) que deseja remover).
- `REPRODUCE` (você colocará para tocar a `playlist` selecionada).
- `REPRODUCE SELECTION` (você colocará para tocar as músicas específicas que escolheu a partir da
   listagem geral de uma playlist).

Pronto! Você já se formou na escola de `DJs Eutherpe`, congrats!

[`Voltar`](#tópicos)

## A tela Storage

A tela `Storage` é onde você sinaliza à `Eutherpe` a partir de onde tentar ler suas músicas.
A **Figura 11** ilustra o que você encontra nessa tela. A tela oferece três operações:

- `LIST` (todos os dispositivos `USB` encontrados são listados, se você plugou um USB é necessário
clicar nesse botão para ter acesso ao seu novo dispositivo).
- `SET` (o dispositivo selecionado na lista de dispositvos encontrados é de fato selecionado para
uso).
- `SCAN` (o dispositivo é escaneado e a coleção extraída e organizada é listada na tela
`Collection`, no primeiro escaneamento o processo pode demorar um pouco mais, isso é por conta de
como o dispositivos de armazenamento `USB` funcionam no `Linux`).

Uma vez o dispositivo escolhido e escaneado. Na próxima sessão, caso `Eutherpe` detecte que o
dispositivo `USB` já esteja plugado, a sua `juke` etérea é esperta o suficiente para já te
apresentar a coleção escaneada da última vez. Se você adicionou ou removeu mais músicas ou
alterou a localização delas, vai precisar clicar em `SCAN` de novo.

[`Voltar`](#tópicos)

## A tela Bluetooth

A **Figura 12** traz a tela `Bluetooth`. Na linha das outras telas é bem direta ao assunto.
Possui três funções disparadas por três botões:

- `PROBE` (você clica nele para sondar o local e ver se encontra o dispositivo bluetooth de saída
de áudio do seu interesse).
- `PAIR...` (você clica nele para parear com o dispositivo que você selecionou da lista que a
sondagem te retornou, vai te exibir a tela da **Figura 13**).
- `UNPAIR..` (você clica nele para desparear o pareamento prévio, o som vai parar de emanar via o
dispositivo `Bluetooth` que você havia pareado antes).

Após parear com um dispositivo de saída, o `ID` desse dispositivo é indicado, dê uma olhada
na **Figura 14**.

[`Voltar`](#tópicos)

## A tela Settings

A tela `Settings` é onde você pode deixar a sua `juke` no jeito que mais lhe aprouver (Uau hein?
Sound portuguese babe, I love my mother tongue). A cara geral dessa tela pode ser conferida via
**Figura 15**.

A tela possui várias funcionalidades mas todas são bem diretas de configurar:

- **a.** Quando marcado uma tela pedindo senha de acesso será apresentada. A senha `default` é `music` (essa eu tenho certeza que você nem eu esquecemos!!!).
- **b.** Quando marcado, a conexão da interface `Web` será via `HTTPS`.
- **c.** Clique aqui e você gera um certificado autoassinado para ser utilizado com o `HTTPS`.
- **d.** Clique aqui e você faz o download do certificado, caso queira instalá-lo no dispositivo que
acessará a interface `Web`.
- **e.** Aqui você abre a interface para configurar a conexão `Wi-Fi/WLAN` para `Eutherpe` ingressar
no seu `Wi-Fi` e você poder controlar o `player` via `cel`, `tablet` ou mesmo `desktop`. Se a rede
da sua casa for sem-fio (quase certo que é), você precisa indicar o nome da sua rede (`SSID`) e a
senha. Relaxa, a senha ficará guardada encriptada, do jeito que deve ser :wink: Eu sou o chato
militante/ilitante da criptografia (e com orgulho!) :smile:
- **f.** Você confirma o nome para acessar a interface `Web Eutherpe`. É importante que seja nos moldes
`<algo>.local`, se você digitar apenas `<algo>` `Eutherpe` vai completar isso com `.local` para
terminar de configurar.
- **g.** Clique aqui e você desliga o dispositivo que está rodando `Eutherpe` (isso é bom para quando
você estiver rodando `Eutherpe` embarcada [na verdade essa é a ideia base] e não quiser colocar
um botão liga/desliga na sua placa).
- **h.** Clique aqui e você reinicia o dispositivo que está rodando `Eutherpe`, você sabe, em computação
sempre tem momentos em que acabamos precisando de um `reboot`. Embora procure deixar `Eutherpe`
o mais estável possível para você :wink:!
- **i.** Se você quiser alterar a senha de acesso à interface `Web` clica aqui. Se esquecer a
senha vai ter que fazer umas madracarias para resetar. Mais para frente te mostro!

[`Voltar`](#tópicos)
