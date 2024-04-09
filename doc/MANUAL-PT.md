# Eutherpe Manual

## Tópicos

- [O que é?](#o-que-é)
- [Características](#características)
- Conhecendo as telas

## O que é?

`Eutherpe` é uma espécie de fiação etérea para `jukebox`. Utilizando `Eutherpe` você será
capaz de ouvir suas `MP3` sem necessidade de se logar em nenhum serviço de `stream` externo
a sua rede local. Aqui é você quem `100%` manda.

A ideia básica é você ter suas músicas gravadas em um `pendrive`, plugar esse `pendrive`
ao computador rodando `Eutherpe`, conectar `Eutherpe` a uma caixa de som ou `headphone`
`Bluetooth` e pronto! Vai escutar sua música e acabou.

Você consegue controlar `Eutherpe` a partir de qualquer `web browser` bastando se conectar
ao endereço do computador que está servindo `Eutherpe` em sua rede.

[`Voltar`](#tópicos)

## Características

Ao escrever `Eutherpe` minha principal preocupação foi `minimalismo`. A ideia é deixar você
o máximo possível longe de telas e mais imerso na audição da sua música.

Então com `Eutherpe` você não tem propagandas, nem sugestões, nada disso. É só você e suas
músicas. Você tem sua coleção de `MP3` quer ouvir suas músicas e pronto.

As principais características de `Eutherpe` são:

- Minimalismo na interface `Web`. Nada de `frameworks` faraônicos etc. `HTML`, `javascript`,
  `Golang` no `backend` e acabou. Não trouxe e nem quero trazer o `Desktop` para dentro de seu
  navegador `Web`.
- Conexão à interface `web` pode ser autenticada por senha ou não. Tudo isso facilmente
  configurável via a própria interface `Web`.
- Conexão ao `player` via `web browser` pode ser `HTTP` ou `HTTPS`. Tudo isso facilmente
  configurável via a própria interface `Web`.
- Você consegue dar organização para sua coleção mesmo estando ela armazenada de forma
  desorganizada. Ao escanear o `pendrive` em busca de músicas, `Eutherpe` é capaz de ler
  as `metatags` das suas `MP3` e listá-las automagicamente organizadas por `Artista/Álbum` e
  organizados ainda dos lançamentos mais recentes para os mais antigos.
- No `Eutherpe` você consegue criar `playlists`.
- As `playlists` ficam associadas ao `pendrive`, se você puser outro no lugar elas sumirão,
  mas se você voltar a conectar o `pendrive` associado a sua `playlist` elas reaparecerão.
- Cada música pode ser marcada com `tags`.
- `Tags` são basicamente palavras-chave: `80s`, `solos favoritos`, `stoner rock`, `mojo`,
  `blues`, `jazz`, `soul` etc.
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
