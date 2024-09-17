![build-glyph](figures/build-glyph.png)
# Build

**Resumo**: Aqui você encontra informações sobre o `build` de `Eutherpe`. Esse
documento é mais voltado para o pessoal do desenvolvimento.

## Tópicos

- [`Como construo o binário Eutherpe?`](#como-construo-o-binário-eutherpe)
- [`Como executo os testes?`](#como-executo-os-testes)
- [`Instalei go no Raspberry Pi mas não consigo rodá-lo`](#instalei-go-no-raspberry-pi-mas-não-consigo-rodá-lo)
- [`Utilizando as rules GNUMake`](#utilizando-as-rules-gnumake)

### Como construo o binário Eutherpe?

O binário `Eutherpe` se trata de um programa escrito em `Golang`. Nisso, o `build`
dele é extremamente simples e direto. Estando dentro do diretório `src` execute:

```
# go build
```

[`Voltar`](#tópicos)

### Como executo os testes?

Gostei, hein? Está preocupando-se com os testes. Isso é bom! Sinal de desenvolvedores
pragmáticos, perfeccionistas.

Sendo o núcleo de `Eutherpe` uma aplicação `Golang`, os testes são executados chamando
o seguinte comando (estando dentro do diretório `src`):

```
# go test internal/...
```

Se quiser mais informações sobre o que está acontecendo no processo, execute:

```
# go test internal/... -v
```

> [!TIP]
> `Golang` tem um comportamento meio irritante de cachear o resultado dos testes,
> se quiser forçar a execução de todos os testes, antes de dispará-los, limpe o
> cache da seguinte forma:
>
> ```
> # go clean -testcache
> ```

[`Voltar`](#tópicos)

### Instalei go no Raspberry Pi mas não consigo rodá-lo

Se você fez o processo de [`bootstrapping`](MANUAL-PT.md#bootstrapping), tente rodar o seguinte comando:

```
# source /etc/profile.d/goenv.sh
```

[`Voltar`](#tópicos)

### Utilizando as rules GNUMake

É pouco usual um `software` escrito em `Golang` dispor de automações via `make`, porém, `Eutherpe`
vai muito além de compilar, testar e instalar. Existem muitas outras operações que se desenrolam
entre compilar e instalar. Levando isso em consideração, resolvi automatizar algumas operações
comuns dos desenvolvedores precisarem fazer.

Essas operações seguem automatizadas em `src/Makefile`.

Até o momento existem quatro `rules` implementadas:

- `eutherpe`
- `tests`
- `bootstrap`
- `update`

Se você quiser gerar o aplicativo `eutherpe`, estando no diretório `src`, rode:

```
$ make eutherpe
```

Essa `rule` é capaz de setar seu `GOENV` por conta própria.

Quer rodar os testes (ignorando os caches)? A partir do diretório `src`, rode:

```
$ make tests
```

O ambiente de desenvolvimento está limpo, sem uma instalação de `Eutherpe`? Você pode rodar `bootstrap`
para instalar e configurar tudo. Para rodar essa `rule` é necessário ser `root`, estando em `src`, rode:

```
# make bootstrap
```

Você fez um ajuste no aplicativo e quer atualizar o `binário` utilizado pelo serviço que
está rodando em seu ambiente de desenvolvimento? Sendo `root` e estando no diretório `src`, rode
`update`:

```
# make update
```

> [!IMPORTANT]
> Esse update apenas atualiza a aplicação `Eutherpe` e os seus `web assets` que
> compõem o `core` de `Eutherpe`. Caso você tenha feito alterações nos `serviços` ou nos `shell scripts`
> é mais indicado rodar `bootstrap`. Nesse caso, ele vai detectar que já existe uma instalação prévia
> e não vai reinstalar o `bluealsa`, não vai recriar o usuário `eutherpe` e nem `rebootar` o sistema.
> Porém, vai atualizar também os `serviços` e os `shell scripts`.

[`Voltar`](#tópicos)
