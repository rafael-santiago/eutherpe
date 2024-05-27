# Eutherpe Manual

## Tópicos

- [O que é?](#o-que-é)
- [Características](#características)
- [Bootstrapping](#bootstrapping)
    - [Ãh?!](#ãhm)
    - [...e não é que jukebox etérea brotou na minha LAN!](#e-não-é-que-uma-jukebox-etérea-brotou-na-minha-lan)
- [Conhecendo as telas](#conhecendo-as-telas)
    - [A tela Music](#a-tela-music)
    - [A tela Collection](#a-tela-collection)
    - [A tela Playlists](#a-tela-playlists)
    - [A tela Storage](#a-tela-storage)
    - [A tela Bluetooth](#a-tela-bluetooth)
    - [A tela Settings](#a-tela-settings)

## O que é?

`Eutherpe` é uma espécie de fiação etérea para `jukebox`. Utilizando `Eutherpe` você será
capaz de ouvir suas `MP3` sem necessidade de se logar em nenhum serviço de `stream` externo
a sua rede local. Aqui é você quem `100%` manda.

A ideia básica é você ter suas músicas gravadas em um `pendrive`, plugar esse `pendrive`
ao computador rodando `Eutherpe`, conectar `Eutherpe` a uma caixa de som ou `headphone`
`Bluetooth` e feito! Vá escutar sua música e acabou.

Você consegue controlar `Eutherpe` a partir de qualquer `web browser`, bastando se conectar
ao endereço do computador que está servindo `Eutherpe` em sua rede.

[`Voltar`](#tópicos)

## Características

Ao escrever `Eutherpe` minha principal preocupação foi `minimalismo`. A ideia é deixar você
o máximo possível longe de telas e mais imerso na audição da sua música.

Então com `Eutherpe` você não tem propagandas, nem sugestões, nada disso. É só você e suas
músicas. Você tem sua coleção de `MP3` quer ouvir suas músicas e só.

As principais características de `Eutherpe` são:

- Minimalismo na interface `Web`. Nada de `frameworks` faraônicos e tal. `HTML`, `javascript`,
  `Golang` e acabou. Não trouxe e nem quero trazer o `Desktop` para dentro de seu navegador `Web`.
- Conexão à interface `web` pode ser autenticada por senha ou não. Tudo isso facilmente
  configurável via a própria interface `Web`.
- Conexão ao `player` via `web browser` pode ser `HTTP` ou `HTTPS`. Tudo isso facilmente
  configurável via a própria interface `Web`.
- Você consegue se conectar ao seu dispositivo `Eutherpe` facilmente, sem precisar ficar
  rodando comandos malucos para descobrir o endereço `IP` da sua `jukebox` etérea,
  pois isso nada realmente tem a ver com o que você quer fazer: *ouvir a f_cking p_rra da
  sua música, ora! :wink:!*
- Você consegue dar organização para sua coleção mesmo estando ela armazenada de forma
  desorganizada. Ao escanear o `pendrive` em busca de músicas, `Eutherpe` é capaz de ler
  as `metatags` das suas `MP3` e listá-las automagicamente organizadas por `Artista/Álbum` e
  organizados ainda dos lançamentos mais recentes para os mais antigos.
- No `Eutherpe` você consegue criar `playlists`.
- As `playlists` ficam associadas ao `pendrive`, se você puser outro no lugar, elas sumirão,
  mas se você voltar a conectar o `pendrive` associado as suas `playlists`, elas reaparecerão.
- Cada música pode ser marcada com `tags`.
- `Tags` são basicamente palavras-chave: `80s`, `solos favoritos`, `stoner rock`, `mojo`,
  `blues`, `jazz`, `soul`, `sextou`, `segundou que tédio` etc.
- Você pode pedir para o `Eutherpe` tocar `n` músicas que se encaixem numa lista de `tags`
  que você passar.
- A dinâmica essencial do `player` do `Eutherpe` é de uma `jukebox`, então você seleciona
  o que quer ouvir.
- Você pode se conectar a um dispositivo de saída `Bluetooth` (caixas de som ou `headphones`).
- `Eutherpe` dá suporte para `MP3`, `MP4` e `M4A`.
- A última sessão é salva, então suas seleções ficam gravadas entre um uso e outro do aplicativo.
  Você pode continuar a audição do ponto que parou.
- Interface `Web` renderiza bem tanto no `Desktop` quanto em dispositivos móveis.
- Funciona baseado na plataforma `Linux`.
- No `Windows` você pode utilizar `Eutherpe` a partir de uma `máquina virtual`.
- Você ainda pode usar `Eutherpe` dentro de um `raspberry-pi` e nesse caso ele será carinhosamente
  chamado de `Euther-PI`.

[`Voltar`](#tópicos)

## Bootstrapping

### Ãhm?!

Se você não é do bonde dos malucos (leia-se gente que programa computadores), talvez seja legal
entender o que cargas d'água significa `bootstrapping`.

Bem, isso é um termo que utilizamos para indicar que vamos provisionar todo o necessário para
começar de fato a brincadeira. Seria um escalar sem corda colocando os grampos, para que
posteriores escaladas sejam mais fáceis. Ao que parece esse termo veio de uma história de
*"As supreendentes Aventuras do Barão Munchausen"* (tradução livre), escrito por Rudolf Erich
Raspe. Onde nessa história o barão teria saído de um pântano puxando a si mesmo pelos cabelos
e seu cavalo pelo rabo (é vdd esse bilete [sic]...). Há controvérsias mas o sentido se mantém:
no `bootstrapping`, o objetivo é passar por uma parte complicada, pantanosa, árida em recursos
para então chegar do outro lado onde tudo será mais sólido sob os nossos pés e para isso,
usando praticamente nada ou parcos recursos.

No caso do `Eutherpe`, a `jukebox` utiliza algumas dependências para compor todo o seu
ecossistema. Essas dependências no caso são aplicativos e recursos que ele usa por baixo dos
panos para prover toda a infraestrutura necessária para você bater cabeça, dançar pelada(o)
na sala ou simplesmente pôr uma música na conversa.

`Eutherpe` foi desenvolvido tendo como contrapartida a distribuição `Linux` `Debian`. Então
o `bootstrapping` **precisa ser feito a partir de uma instalação `Debian` (eu aconselho)**.

O ideal é você fazer uma instalação `Debian` mínima (sem recursos para `desktop`). Apenas
console mesmo. Por que? Bem, o que acontece é que quando você coloca muita coisa desnecessária,
o sistema vai começar perder tempo e recursos tomando conta desses elefantes brancos e, esse
desperdício pode impactar na audição das suas amadas músicas.

Te indico instalar o `Debian 11` (somente console, com a rede configurada, pois o `bootstrapping`
irá baixar pacotes específicos da Internet para deixar o seu `Debian` no jeito para rodar
`Eutherpe`).

O `Debian 11` foi o meu sistema de escolha para desenvolver `Eutherpe`, tentei o `12` mas ficou
bem ruim e estava utilizando um sistema com `Gnome` e todas aquelas tranqueiras *- que para o
objetivo do `Eutherpe` são essencialmente inúteis*.

Você deve estar pensando: *"- Poxa fazer esse tal 'boostrapping' deve ser difícil bagarai..."*.
Relaxa, não é!

[`Voltar`](#tópicos)

### ...e não é que uma jukebox etérea brotou na minha LAN!

Após você preparar uma versão mínima `Debian` é necessário que você faça login como usuário `root`.

Uma vez a sessão `root` iniciada, você precisa baixar o código-fonte do `Eutherpe` e para isso você
precisará do aplicativo `git`.

Então vamos instalá-lo, emitindo o seguinte comando (se você instalou o `Debian 11` vai precisar
inserir o `CD-ROM` ou imagem do .iso dele):

```
# apt install git -y
```

Com o `git` instalado é hora de baixar o código-fonte do `Eutherpe`, usando o seguinte comando:

```
# git clone https://github.com/rafael-santiago/eutherpe -b v1
```

Um diretório chamado `eutherpe` será criado, vamos acessá-lo usando o seguinte comando:

```
# cd eutherpe
```

Pronto! Agora é só fazer o `bootstrapping` :boot: emitindo o seguinte comando:

```
# ./bootstrapping.sh
```

Você precisa confirmar com `y` ou `Y` e tudo será feito. As dependências instaladas, ajustes
finos no sistema serão feitos. Você terá uma saída similar a essa:

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

Uma vez o `bootstrapping` feito você poderá acessar sua `jukebox` etérea a partir
de um `web browser` mais próximo via o endereço: [`http://eutherpe.local:8080/eutherpe`](http://eutherpe.local:8080/eutherpe).

*Voi lá!* :sunglasses:

É hora de conhecer a sua `juke` minimalista e sem frescuras!

[`Voltar`](#tópicos)

## Conhecendo as telas

A configuração inicial da sua `jukebox` `Eutherpe` vai ser bem básica:

- Conexão `HTTP`;
- Sem autenticação por senha;

Ao se conectar em `http://eutherpe.local:8080/eutherpe` você verá a tela ilustrada pela **Figura 1**.

Note que se trata de uma tela bem direta ao assunto, possuindo um menu do lado esquerdo onde
você acessa as funções e configurações da sua `juke` etérea.

[`Voltar`](#tópicos)

### A tela Music

Na **Figura 2** você confere o *layout* da tela do `player`. Ela é bem autoexplicativa. Basicamente
a tela te oferece funções básicas como tocar, parar a música, ir para próxima da lista de
reprodução, ir para a anterior, mover música(s) para cima ou para baixo da lista de reprodução,
ativar o `shuffle` (que vai misturar a sua lista de reprodução), remover músicas selecionadas da
sua lista de reprodução e limpar toda a lista de reprodução.

Você ainda tem um `slider` para controlar o volume e a possibilidade de ativar os modos de
repetição geral ou repetição de apenas uma música (da que estiver atualmente em reprodução).

Para acessar a lista de reprodução, você precisa clicar sobre `UP NEXT`. Vou aproveitar e já
te ensinar sobre uma convenção que adotei.

Se você é realmente um ser musicófilo, já deve ter notado, caso não, agora vai notar:
tudo que estiver precedido pelo símbolo de *sustenido*, significa que se você clicar sobre
lhe será exibido mais. Uma vez esse conteúdo exibido, do lado dele o sustenido se tornará um
*bemol*, o que indica que se você clicar novamente, o conteúdo será ocultado. A relação aqui
entre `sustenido/bemol` é idêntico a entre `+/-`, :wink:!

Confira na **Figura 3** uma lista de reprodução em modo de detalhamento. Note que do lado de
cada música existe uma caixa de checagem que você poderá marcar para aplicar sobre ela certas
funções como remover e também movê-la para cima ou para baixo da lista de reprodução. Caso
queira tocar uma música específica, selecione ela e dê um `play`.

Pronto! Sendo `Eutherpe` uma `jukebox` sem frescura, você já sabe tudo sobre como pilotar o
seu `player`. Entretanto, você deve ter ficado com uma pulga atrás da orelha: *como eu seleciono
o que quero tocar?*

[`Voltar`](#tópicos)

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
as músicas são listadas, confira a **Figura 10** onde a lista `HOW-LOU` está sendo detalhada.

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
