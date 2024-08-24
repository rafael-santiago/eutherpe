![Ancient Greek Pegasus Coin / Public Domain](figures/pegasus_coin.png "Ancient Greek Pegasus Coin / Public Domain")

# Guia de estilo de código

**Resumo**: Programar é um tipo de arte-artesanato-engenharia (quem sabe até expressão) que
envolve muitos traços culturais e idiomáticos. Devido a isso, similar a muitos outros segmentos,
`A Verdade` é um grande, graaaaaaaande `Unicórnio Alado`... De qualquer forma, esse texto intenta
descrever de forma objetiva os principais aspectos atuais do `meu` amado unicórnio.
Se você quer programar comigo, vai participar de uma festa que já começou, essa é a
minha festa.

# Tópicos

- [A ideia geral](#a-ideia-geral)
- [Construções básicas](#construções-básicas)
    - [if/else](#ifelse)
    - [for](#for)
    - [switches](#switches)
    - [func](#func)
- [Não polua a instalação go, esqueça download de packages](#não-polua-a-instalação-go-esqueça-download-de-packages)
- [Definition of done](#definition-of-done)
- [Use linguagem neutra e inclusiva](#use-linguagem-neutra-e-inclusiva)

## A ideia geral

**Curto e grosso**: `go to hell` com `go fmt` aqui. Não tenho intenção nenhuma de fazer
isso um pacote `go`, graças a isso eu não preciso engolir aqueles espaçamentos chatos
e impostos do `go fmt`.

Aqui use `quatro espaços` para indentar e espaço para alinhar. Coloque o seu editor para
substituir `tab` por espaço, assim todo aquele `bla-bla-bla` e polêmica sobre qual é melhor
**acaba**. Pois, pouco importa o que você usa para indentar, o problema é usar sempre espaço
para alinhar. Se o `tab` já vira espaço, pouco importa se você bate `tab` ou espaço. Vai sair
um espaço e nunca vai desalinhar em nenhum lugar, *yo!* (**Mic drop**). Porém, lembre-se **eu uso
`4 espaços` para representar o `tab`**.

Oitenta colunas como tamanho máximo de linhas é miserável. Procure não passar de cento e vinte
colunas, estamos no século 21 e não existe forma de voltar ao passado (`2024`). Se você programa
de polainas e adora octal, vai dar uma olhada naquele DeLorean cor de gelo parado lá na esquina,
você tem aí uma caneta para rebobinar minha fita cassete?? Não quero gastar minhas pilhas...

Falando em indentação, por favor evite abusar de código com `if` aninhado dentro de outros sem
necessidade real. Isso aumenta a complexidade ciclomática. Evite coisas feito isso:

```go
if i == 42 {
    if j == 84 {
        ...
    }
}
```

Olha que mais simples e elegante:

```go
if i == 42 && j == 84 {
    ...
}
```

Por fim, programe de forma determinística, `if` não é o fim da programação, nunca foi e nunca será,
`é caso de exceção`, se você programa orientando sua forma de pensar em "`if` isso e `if` aquilo"
bem certo que produz peças de engenharia propensas a se tornarem celeiros de `bug`. O uso abusivo
de `if` é sim, o Fim da programação, evite. Construa coisas, não bagunce nem destrua!

Falhe o mais cedo possível, pare de achar que o mundo é lindo enquanto programa, `verifique sempre
pela falha`, falhou? Tchau! A outra parte que chamou que se vire. Isso também evita aqueles
código anhinhados Poliana: "se deu certo aqui, faz isso, se deu certo aqui, faça mais isso... Zzz".
Aqui é Vietnã: "Se deu erro aqui, fim de conversa. Faz isso, deu erro aqui, não entra aqui e
bom dia Vietnã, tchau!". Falo isso e repito: `Edgar Alan Poe` se não tivesse sido escritor
teria sido um ótimo programador. `Edward A. Murphy` a mesma coisa.

É triste que a linguagem `Go` em certos casos orienta mesmo o programador produzir certos períodos
bem `noobs` mas tirando a falta de jeito `by-design` da linguagem, evite `branching` à toa,
obrigado! :wink:

Comece todo código com o copyright disclaimer. Organização. Sem isso, em pouco tempo o código
vira um ninho de rato, acredite!

Use `camelCase` para as variáveis e nomes de funções privadas. `EssaParaAsPublicas`. A convenção
básica de `Go` mesmo.

Se comentar algo, se comprometa, ponha:

- `INFO(SeuNomeOuApelido): ...`, para comentários informativos.
- `WARN(SeuNomeOuApelido): ...`, para algum aviso.
- `BUG(SeuNomeOuApelido): ...`, para sinalizar um bug conhecido.
- `TODO(SeuNomeOuApelido): ...`, para sinalizar algo por fazer, melhorar.

Procure resolver bugs da melhor forma: aplicando soluções gerais. Não implemente desvios de `bug`,
isso não resolve. Se você não está entendendo a essência da falha, não está em condições de propor
solução ainda, pense mais um pouco. Vá com calma! Isso aqui não é `hackathon` :wink:

Procure implementar testes. Coloque os testes para rodar durante o `build`. Você não está
disputando partidas de `build turfe`, foda-se se demorar, se você implementou algo precisa
se assegurar que aquilo funciona, isso é sobre engenharia, se não quer testar aquilo, remova
o código, pois não precisamos de áreas incertas no projeto, buscamos por certeza, o treco
aqui é sobre ser empírico e pragmático para evitar adubar a vida das pessoas, o mundo já
é muito incerto e zoado, nisso é legal entregar coisas bacanas e previsíveis nas mãos das
pessoas e tornar o mundo delas menos bagunçado nem que seja com uma humilde peça abstrata
de software (desculpa o pleonasmo) que simplesmente teima em ser contra a corrente e sempre
funcionar fazendo ela exclamar "porra, essa merdinha aqui funciona mesmo, cara que estranho...
eu curto isso!", obrigado! :wink:

Não utilize pacotes `Go` externos à biblioteca padrão. A ideia do projeto é ser minimalista mas
entregar o que se propõe. Acho essa ideia do `Go` de puxar fontes de bequinhos escuros da
`internet` e ir entulhando na instalação local, bem zoada. Não gosto de usar. Só uso o que
é padrão para me assegurar que sempre vou conseguir fazer o build do aplicativo em qualquer
momento do tempo, bastando apenas ter a versão exata do `Go` e acabou. Sou pragmático quanto
à programação, não gosto de surpresas e sou maluco, aficionado por ter controle do código
que produzo. A ideia de alguém impossibilitar o build do meu aplicativo porque mudou um repo
ou comportamento de algo ao bel-prazer dela, acho absurda e não opcional. Para mim é atestado
de incompetência. Nem considero. Meu objetivo é construir coisas que se mantenham em pé e que nem
precisem mais de mim quando estiverem por aí funcionando. Quanto menos as pessoas se lembrarem de
mim, me procurarem é porque o que eu fiz está funcionando. Em `T.I.` só vão lembrar de você para
reclamar. Se não estão te enchendo, é porque está bom.

Por fim mas não menos importante: **Busque status de roda nos seus construtos, por favor**.

Certo, não sabe o que diabos é "status de roda", imagine o mundo sem ela, ou tente melhorar
a essência dela... Entendeu?! :wink:

<center>
<h1>
<i>Mesmo assim, desconfie sempre de todo código que você produzir. Nada é perfeito, mas pode
se aproximar disso</i>.
</h1>
</center>

[`Voltar`](#tópicos)

## Construções básicas

A seguir você encontra a ideia geral de como se formata construções básicas
que vão compor suas implementações.

No geral:

- Procure deixar claro a precedência usando parênteses.
- Comente partes intrincadas do código, pare de achar que você é um gênio ou poeta, isso não existe
na programação, é tudo balela de gente que tem preguiça ou é inapto em aprender algo mais difícil
que requeira de fato algum nível de virtuosismo :smirk:.
- Programe em inglês.
- Mensagens de commit também são em inglês.
- Documentações podem ser em português-BR e/ou inglês.

[`Voltar`](#tópicos)

### if/else

Assim que você formata um `if/else`:

```go
if i == 42 {
    ...
} else if j == 42 {
    ...
} else {
    ...
}
```

Infelizmente `Go` não admite esse tipo de alinhamento:

```go
if i == 42
    && j == 42 {
    ...
}
```

Mas se quiser quebrar linha numa expressão lógica, pode fazer assim:

```go
if i == 42 &&
   j == 42 {
    ...
}
```

[`Voltar`](#tópicos)

### for

Assim que você formata um `for`:

```go
for i := 1; i < 42; i++ {
    ...
}
```

[`Voltar`](#tópicos)

### switches

Assim você formata `switches`:

```go
switch i {
    case 0, 1, 2:
        ...
        break

    case 3, 4, 5,
         6, 7, 8:
        break
}
```

[`Voltar`](#tópicos)

### func

Funções são assim:

```go
func privateFunc(i, j int, str string) ([]byte, error) {
    ...
}

func PublicFunc(i, j int, str string) ([]byte, error) {
    ...
}
```

[`Voltar`](#tópicos)

## Não polua a instalação go, esqueça download de packages

A ideia é sempre usar a biblioteca padrão `go`. No caso das `packages` internas `Eutherpe`, evite
instalá-las no seu `GOROOT`, esqueça toda essa baboseira.

Quando for iniciar uma nova package abaixo de `src/internal` rode esse comando para criar o
`go.mod` (vamos supor que você vai criar a package `shoobeedooblaublau`):

```
# mkdir src/internal/shoobeedooblaublau
# cd src/internal/shoobeedooblaublau
# go mod init github.com/rafael-santiago/eutherpe/shoobeedooblaublau
```
No `src/go.mod` (note, no `go.mod` imediatamente abaixo de `src`) adicione:

```
require internal/shoobeedooblaublau v1.0.0
replace internal/shoobeedooblaublau => ./internal/shoobeedooblaublau

```

Agora, por exemplo, se o pacote `shoobeedooblaublau` utilizar `mplayer`, no `src/internal/shoobeedooblaublau/go.mod`
você precisa adicionar:

```
require internal/mplayer v1.0.0
replace internal/mplayer => ../mplayer
```

Esse cuidado é bom pois torna tudo mais autocontido e não bagunça o seu `GOROOT` com toda aquela
suruba de pacotes baixados de qualquer lugar que `golang` se amarra em fazer. Eu acho essa
"feature" uma ideia de jerico total, mas, o lance das tecnologias é desviar das más ideias e usar
o que vale a pena, torcendo para que notem o quão idiota certas coisas são e purguem
definitivamente essas tranqueiras em versões futuras :wink:!

[`Voltar`](#tópicos)

## Definition of done

Uma nova feature é considerada feita quando:

1. Ela faz o que se propõe fazer.
2. Ela não adiciona bagunça, confusão ou mesmo instabilidade e nem bugs nas features prévias.
3. Ela entrega o que promete de um modo simples (mas não oco chamado de "simples"). Em outras
palavras você fez uso da navalha de Occam.
4. Ela é bem testada.
5. Ela não está amarrada em alguma dependência que não é da biblioteca padrão `Go`.
6. O `CI` está passando.
7. Ela é bem documentada.
8. O commit que adiciona a feature para o upstream é descritivo.
9. A mensagem de commit está em modo imperativo. Parecendo que você dá comandos ao sistema
de controle de versão. Então `Dando comandos ao controle de versão` é errado. `Dá comandos
ao controle de versão` também. `Dê comandos ao controle de versão`. Não tenha vergonha de
ser mandona/mandão com ele! :wink:

[`Voltar`](#tópicos)

## Use linguagem neutra e inclusiva

*Isso é o único ponto que **não há concessão**. Isso **NÃO É** sobre unicórnios, de verdade.
Siga isso ou adeus.*

Sempre tente usar termos/palavras neutras e inclusivas no seu código e documentações. Se você
encontrar algo que para você não parece tão correto, por favor, me deixe saber abrindo uma
issue me sugerindo as melhorias. Obrigado desde já!

Em geral evite usar cores para nomear o que deveria ser "bom" ou "mau". Termos ultrapassados
como `whitelist`/`blacklist` estão depreciados/banidos aqui. Você deveria usar `allowlist/denylist`
ou qualquer coisa mais relacionada ao que você realmente está fazendo. Termos como `master/slave`
estão fora também. Você poderia usar `main`, `secondary`, `next`, `trunk`, `current`, `supervisor`,
`worker` em substituição.

Não use termos sexistas e/ou machistas também.

De novo, se você encontrou algum(ns) termo(s) que para você não pareça(m) adequado(s), me deixe
[saber](https://github.com/rafael-santiago/eutherpe/issues) sugerindo edições, valeu!

*-- Rafael*

[`Voltar`](#tópicos)
