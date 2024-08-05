# Guia de estilo de código

**Resumo**: Programar é um tipo de arte-artesanato-engenharia que envolve muitos traços
culturais e idiomáticos. Devido a isso, similar a muitos outros segmentos, `A Verdade`
é um grande, graaaaaaaande `Unicórnio Alado`... De qualquer forma, esse texto intenta
descrever de forma objetiva os principais aspectos atuais do `meu` amado unicórnio.
Se você quer programar comigo, vai participar de uma festa que já começou, essa é a
minha festa.

# Tópicos

- [A ideia geral](#a-ideia-geral)

## A ideia geral

**Curto e grosso**: `go to hell` com `go fmt` aqui. Não tenho intenção nenhuma de fazer
isso um pacote `go`, graças a isso eu não preciso engolir aqueles espaçamentos chatos
e impostos que usam `tab` no `go fmt`.

Aqui use `quatro espaços` para indentar. Mas por favor evite abusar de código com `if`
aninhado dentro de outros sem necessidade. Sabe, vícios bem ruins como:

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
de `if` é sim, o Fim da programação, evite.

Falhe o mais cedo possível, pare de achar que o mundo é lindo enquanto programa, `verifique sempre
pela falha`, falhou? Tchau, a outra parte que chamou que se vire. Isso também evita aqueles
código anhinhados Poliana: "se deu certo aqui, faz isso, se deu certo aqui, faça mais isso... Zzz".
Aqui é vietnã: "Se deu erro aqui, fim de conversa. Faz isso, deu erro aqui, não entra aqui e
bom dia vietnã, tchau.". Falo isso e repito: `Edgar Alan Poe` se não tivesse sido escritor
teria sido um ótimo programador. `Edward A. Murphy` a mesma coisa.

É triste que a linguagem `go` em certos casos orienta mesmo o programador produzir certos períodos
bem `noobs` mas tirando a falta de jeito `by-design` da linguagem, evite `branching` à toa,
obrigado! :wink:

Comece todo código com o copyright disclaimer. Organização. Sem isso, em pouco tempo o código
vira um ninho de rato, acredite!

Use `camelCase` para as variáveis. Nomes de função privadas e `EssaParaAsPublicas`. A convenção
básica de `go` mesmo.

Se comentar algo, se comprometa, ponha:

- `INFO(SeuNomeApelido): ...`, para comentários informativos.
- `WARN(SeuNomeApelido): ...`, para algum aviso.
- `BUG(SeuNomeApelido): ...`, para sinalizar um bug conhecido.
- `TODO(SeuNomeApelido): ...`, para sinalizar algo por fazer, melhorar.

Procure resolver bugs da melhor forma: aplicando soluções gerais. Não implemente desvios de `bug`,
isso não resolve. Se você não está entendendo a essência da falha, não está em condições de propor
solução ainda, pense mais um pouco. :wink:

Procure implementar testes. Coloque os testes para rodar durante o `build`. Você não está
disputando partidas de `build turfe`, foda-se se demorar, se você implementou algo precisa
se assegurar que aquilo funciona, isso é sobre engenharia, se não quer testar aquilo, remova
o código, pois não precisamos de áreas incertas no projeto, buscamos por certeza, o treco
aqui é sobre ser empírico e pragmático para evitar adubar a vida das pessoas, o mundo já
é muito incerto e zoado, nisso é legal entregar coisas bacanas e previsíveis nas mãos das
pessoas e tornar o mundo delas menos bagunçado nem que seja com uma humilde peça abstrata
de software que simplesmente teima em ser contra a corrente e sempre funcionar fazendo ela
exclamar "porra, essa merdinha aqui funciona mesmo, cara que estranho...", obrigado! :wink:

Não utilize pacotes `go` externos à biblioteca padrão. A ideia do projeto é ser minimalista mas
entregar o que se propõe. Acho essa ideia do `go` de puxar fontes de bequinhos escuros da
`internet` e ir entulhando na instalação local, bem zoada. Não gosto de usar. Só uso o que
é padrão para me assegurar que sempre vou conseguir fazer o build do aplicativo em qualquer
momento do tempo, bastando apenas ter a versão exata do `go` e acabou. Sou pragmático quanto
a programação, não gosto de surpresas e sou maluco, aficionado por ter controle do código
que produzo. A ideia de alguém impossibilitar o build do meu aplicativo porque mudou um repo
ou comportamento de algo ao bel prazer dela, acho absurda e não opcional. Nem considero.
Meu objetivo é construir coisas que se mantenham em pé e que nem precisem mais de mim
quando estiverem por aí funcionando.

Por fim mas não menos importante: **Busque status de roda nos seus construtos, por favor**.

Certo, não sabe o que diabos é "status de roda", imagine o mundo sem ela, ou tente melhorar
a essência dela... Entendeu?! :wink:

*Mesmo assim, desconfie sempre de todo código que você produzir. Nada é perfeito, mas pode
se aproximar disso*.

[`Voltar`](#tópicos)

