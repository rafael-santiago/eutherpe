![bluetooth-blues](figures/bluetooth-blues.png "'I woke up this morning feeling unpaired all my audio tools'... (c) Creative Commons Rafael Santiago. Gostou do trocadalho? Se aproprie... ;)")
> I woke up this morning...
> feeling unpaired all my audio tools!
# The Linux Bluetooth Blues

**Resumo**: Nesse texto você vai encontrar um pouco sobre minhas lições aprendidas
passando raiva com `Bluetooth` no `Linux`. Caso você esteja passando pela mesma raiva e
já tentou diversas coisas, talvez esse texto te ajude renovar as esperanças, recalcular
a sua rota ou confirmar certas desconfianças que você talvez tenha tido e ficar com
mais raiva ainda. Aqui vou me permitir ter menos papas na língua, já aviso.

## Tópicos

- [Onde o nosso Blues começa](#onde-o-nosso-blues-começa)
- [Into the blue ou o que tristemente eu descobri](#into-the-blue-ou-o-que-tristemente-eu-descobri)
- [Eis que nesse Blues vem o turnaround](#eis-que-nesse-blues-vem-o-turnaround)

## Onde o nosso Blues começa

Há muito tempo eu era um ávido colecionador de `CD`s, resistia a todo custo às `MP3`, até que um
dia resolvi dar uma chance para o `iPod Shuffle`. Longa história curta: não era que aquele trequinho
era bacana? Nisso, eu logo adquiri um `iPod Touch` e durante anos ele foi meu principal tocador
de música (e sim, o tenho e o uso até hoje!)

Quando a `Apple` resolveu descontinuar o `iPod`, sendo eu uma criatura pragmática e que gosta
sempre de ter um plano na manga, já antevendo que um belo dia irei ficar sem meu amado `iPod`,
resolvi que queria embarcar numa placa um `software` para ser o meu tocador de `MP3` para o
dia a dia. Com a vantagem de não sofrer com decisões de terceiro em descontinuar o treco que eu
uso e curto.

O treco não precisava ser intricado, mas não precisava ser porcamente mambembe, enfim...

Decidi que escreveria a maior parte do `software` em `C`, mas depois ponderando
um pouco, optei por `Golang`. O `OS`, levando em conta que escolhi o `Raspberry Pi 4B` como
minha placa alvo, seria `Linux` (distro `Debian`, por conta de acabar embarcado num `Raspbian`).

Decidi que poria suporte para `Bluetooth` pois de fato eu hoje uso bastante esse tipo de
tecnologia para ouvir música. Eu sou aficionado em mitologia grega então eu chamaria o
projeto de `Eutherpe`. Com "`th`" para fazer um trocadalho do carilho com `ethereal` pois
a minha ideia principal era ter um `player` mais onipresente, etéreo via minha `LAN`.
Podendo acessar o `player` por meio de qualquer `web browser` a partir de algum dispositivo com
acesso a minha `LAN` (leia-se meu `smartphone` ou meu `laptop`, `desktop` etc).

Dei uma esquadrinhada nas possibilidades e pensei comigo: "Poxa, é simples, serão mais
`wrappers` para disparar certas funcionalidades no sistema e vai rodar liso"... **Ledo engano!...**

[`Voltar`](#tópicos)

## Into the blue ou o que tristemente eu descobri

Para iniciar o desenvolvimento eu escolhi `Debian 12 (com ambiente desktop)` e foi horroroso.
Quando comecei testar a reprodução das músicas aquilo atrasava a reprodução o tempo todo,
parava de funcionar do nada. Era um inferno. Eu ainda estava usando `PulseAudio`.

No desespero, enxuguei um `Debian 11`, `text-based` e o negócio melhorou um pouco ao ponto
de conseguir uma versão `beta` de `Eutherpe`.

Daí eu constatei na pele que `PulseAudio` é um lixo. Muito ruim. Não serve para nada. Para
nada mesmo, por ser abstrato, nem para peso de papel. Se você vai apenas ficar tocando
notificação de sistema, beleza. Entretanto, se você quer usar o computador exclusivamente para
áudio (mesmo a coisa mais básica: **tocar música**), ele é uma bela de uma bosta fedida escabrosa.
Desculpa, mas nesse texto vou ser sincerão, ele é uma merda mesmo! Se você não gosta de
ler a palavra **merda**, para de ler o texto, pois vai ter muito relato dessa **merda** e de
outras...

Com o `PulseAudio`, encontrei problemas diversos: do som um dia funcionar super bem, no
outro ficar reiniciando e reiniciando o sistema e nada do som via `Bluetooth` ficar bom.
Reinstalava o sistema e ficava bom, voltava ficar horrível depois de uns dias. Acho que
ele variava baseado no dólar, vai saber...

Existe ainda quilos e quilos de `sets` crípticos que você vai coletando da `Internet` e tentando
agregar nos seus próprios `confs` do `PulseAudio` para ver se funciona melhor. A verdade é que
não é que funciona, o que eu desconfio é que o comportamento do `PulseAudio` é tão anômalo que as
pessoas depois de ficarem horas e horas tentando, colocam um `set` lá e o caiau "funciona" daí
pensam: `- Aha! É isso!`. Mas não, um belo dia o problema volta da mesma forma. No `core` do
`pulseaudio` deve ter um pombo caminhando entre duas áreas, uma verde e uma vermelha, quando
o pombo pisa na verde tudo funciona, quando pisa na vermelha, você corre na `Internet` buscando
passar seu tempo buscando um `set` para resolver enquanto o pombo com sorte se move para
a área verde e seu `set` *"resolva"* milagrosamente o problema, até que o pombo maquiavélico
escolha pelo vermelho de novo...

O `Bluetooth` com o `PulseAudio` certas vezes conecta mas não flui som, pois ele não consegue
colocar como `sink default` a do dispositivo `Bluetooth`. **Mesmo ele sendo listado como um
sink**. Às vezes (só *às vezes*, não sempre) reiniciar o serviço do `pulseaudio` "resolvia"
comigo.

Outra coisa extremamente irritante é o `switch` automático para `HFP`!!! Arrrrgh. Se você
não sabe o que é `HFP` em `Bluetooth`, é um acrônimo para `Hands-free Profile`. `Headphones`
com microfone, os `Headsets` se enquandram nesse perfil. Por padrão ao conectar o `pulseaudio`
a um `headset` ele vai pôr no jeito para que você consiga usar o microfone e funções `hands-free`
para controlar o dispositivo e isso quase sempre cancela o áudio de uma forma que você não
consegue usar a porção `headphone` dele. Você precisa dizer para o `pulseaudio` não fazer isso.
Sim, mais um `set` de configuração.....

Aqui cabe falta de malícia e capacidade de inferência, se o `pulseaudio` roda de um `desktop`
é bem certo que o usuário não está dando a mínima para funções `hands-free`.

Cheguei num ponto que `Eutherpe` já tinha tanta gambiarra para conviver com a ruindade do
`PulseAudio` que eu constatei que se você usa uma coisa tão ruim como contrapartida para seu
`software`, bem certo que aquela ruindade, desleixo, falta de jeito de resolver de forma
geral as coisas, vai transbordar dessa parte zoada para todas as outras. O que estiver
bom ou no mínimo aceitável vai se tornar um lixo irrecuperável igual a parte zoada que vai
impregnar e fazer feder tudo.

É aquela lei da entropia do Schopenhauer, cara:

>Se você puser uma colher de esgoto em um barril cheio de vinho, você obtém esgoto.

O inverso também é interessante de observar:

>Se você puser uma colher de vinho em um barril cheio de esgoto, você ainda obtém esgoto.

Nisso, procure programar bem a sua parte e escolha bem as suas dependências. Pois elas podem
melhorar ou pôr a perder todo o seu esforço de entregar algo utilizável. Ainda vou mais além,
ao meu ver nada é perfeito e tudo tem `bug`, quanto menos coisa junto, menos `bug`. Seus próprios
`bugs` são mais fáceis de conviver do que com os `bugs` dos outros. Procure depender o mínimo
possível dos outros para entregar seu `software` e o que precisar depender, se certifique que
é no mínimo aceitável e que tem um bom senso de engenharia. Sim, tem muito `software` por
aí que é feito por gente que não conseguiria construir uma cabaninha com `lego` sem explodir
tudo... A pessoa se puder cagar no chão e passar a bunda em cima da merda para limpar, faria e,
acharia isso crível e aceitável. Esse tipo de engenharia é inaceitável, tem muito por aí,
infelizmente! Embora os seres humanos sejam inventivos, pouquíssimos são capazes de entregar
coisas utilizáveis e manuteníveis, que possam se manter de pé durante tanto tempo.
Veja as pouquíssimas ruínas que temos por aí que contam um pouco da nossa história enquanto
civilização, que tal procurar tirar `insights` a partir do que pessoas espertas antes de nós
fizeram ao invés de ficar inventando moda furada?...

A boa notícia é que o `PulseAudio` já foi descontinuado. O `Debian 12` ainda o utiliza mas
quase nenhuma `distro` parece que utilizará essa droga novamente. Foi tarde essa bagaceira!
Deveriam fazer um `data wiping` elevado à enésima potência em todo o servidor que ainda hospeda
uma cópia dessa porra para que a humanidade nunca mais consiga buildar essa bela peça de merda.
Amém!

Aí resolvi dar uma chance para o `PipeWire`. Ah o `PipeWire`/`WirePlumber`... Ele FUNCIONOU!!!
...bem durante uma semana.... Mas a verdade é que eu testei ele em dois fones, um do tipo `earbud`
e um `headphone` mesmo. O `earbud` funcionou de cara, o `headphone` não, pois ele tinha microfone
e o `PipeWire` também faz o `switch` automático para `HFP`. A solução? Mais um chato arquivo
de configuração. Descobri que o `PipeWire` e o `WirePlumber` levaram os remendos via `conf files`
para outro nível, chegaram ao ponto de usar `Lua` (!!!) depois voltaram atrás. Existe um padrão
de ir entulhando os arquivos prefixando-os com números (sem comentários) para ordenar o remendo
que vem primeiro (patético)... Logo vi que ele não reduziria em nada a maior complexidade do som
no `Linux`: simplesmente funcionar sem ficar enchendo o saco na frente com promessas que funciona,
mas espera do usuário quilos e quilos de contrapartidas via remendos em arquivos de configuração.
Porra, não é possível simplesmente funcionar?! No `Windows` vai de boa, no `macOS` vai de boa, no
`Android` vai de boa, no `iOS` vai de boa. Porque no `Linux` o usuário simplesmente não consegue
ter um `Bluetooth` decente e liso? É o que me perguntava...

> [!NOTE]
> Um fato interessante e para pensar e, talvez juntar os pontos: o `Bluetooth` do meu
>`iPod Touch`, cuja a última atualização parou no `iOS 9` (o atual é `17.6`[2024]), dá um
> banho no `PulseAudio`, `PipeWire` e `WirePlumber` juntos. O subsistema `bluetooth` dele
> conecta em qualquer dispositivo `Bluetooth` de saída, não engasga o som e faz o que se propõe,
> toca música sem enrolação. Uma perfeição! `Apple`, hein? Que dizem ter fama de só funcionar com `Apple`...
> Pensa aí...

Você já entendeu: descobri que o `PipeWire` também não me atende. Talvez ele tenha
atendido usuários que precisem gravar coisas mas para mim que precisava tocar som direito sem
engasgos e conectar de forma confiável e estável via `Bluetooth` **SEMPRE**! Falhou miseravelmente,
não ficou melhor que o `PulseAudio`. No final, para mim foi merda por bosta, 6 por meia dúzia, você
escolhe... Pois a minha caixa de som `Bluetooth` conectava mas o som era péssimo. E o som
picotava bastante às vezes. Pedir ajuda em fórum, abrir `issue`? Preguiça e quase certo que
iriam me sugerir mais uma `conf` "sacerdotal"... blargh! Quase todo projeto `opensource`
você relata `bugs`, expõe problemas e uma casta meio chatonilda e *blasé* fica fazendo que não ouve,
é uma palhaçada total e eu cansei dessa dinâmica chata... Hoje pego as coisas por mim e resolvo,
quer copiar, usar, pega e usa e não me enche o saco. `Pull request`?! Não perco mais meu tempo,
desculpa, mas pronto, falei...

Você não está entendendo, eu fiquei por quase 6 meses dando chance para essas merdas... Seis
longos meses!!!! Toda a parte que me propus fazer ficou bem funcional, facilmente embarcável
na placa mas o sistema de som cagava na retranca de forma magistral! Não aguentava o tranquinho
básico de sempre conectar no `Bluetooth` e não engasgar, tocar a porra do caralho da música.
Absolutamente ridículo! Hoje em dia acho que todo sistema operacional digno de ser usado em
`desktop` é obrigatório que tenha um sistema de som excepcional. Quase todo usuário final
supervaloriza som! Eu, desenvolvedor de `software`, amo música, não programo sem, um computador
pessoal que não toca música para mim é um caiau bizarro!

> [!IMPORTANT]
> Com o `PipeWire` que diz ser o substituto do `PulseAudio`, reiniciar os serviços às vezes resolvia
> as coisas, mas não sempre. Curioso que com o `PulseAudio` era assim também... Enfim, eu não gosto
> de coisa que funciona mas não sempre... Sou pragmático cara, se é para fazer algo é para que te
> adiante e resolva a vida e não o contrário... Solução que cria problema ou esconde problema, não
> é solução. É porqueira. É melhor não funcionar nunca e me libertar para buscar algo melhor ou
> funcionar sempre. Nisso é preciso ser 8 ou 80. Coisa ruim precisa ser purgada, eliminada, assim
> se promove qualidade. Caso contrário você corre o risco de virar um amante de ganache orgânica.

[`Voltar`](#tópicos)

## Eis que nesse Blues vem o turnaround

Depois de muito rodar lâmpada com `PulseAudio`, `PipeWire` e `WirePlumber` resolvi voltar as
minhas origens: `minimalismo`. Ficar com o simples, porque o simples funciona. E quando não
funciona de cara, você tem poucas partes para revisar e encontrar o problema. Porém, o pulo
do gato é que o simples, nem sempre é simplório, com o tempo, nos meus 1/4 de século programando
(`2024`) eu aprendi isso...

Entendeu, né? O simples, nem sempre vai resolver via um `apt-get`... Para sair do senso comum
na maioria das vezes você vai ter que fazer as coisas por você mesmo, pois a maioria só quer
mesmo sentar no pudim. Quer que o mundo acabe em melado para morrer doce.

Num desses meus doces périplos pela `Internet` buscando `sets` de configuração para `Pulseaudio` e
`PipeWire` (é uma `bad trip` viciante, cuidado!...) li um `post` de um cara (soterrado em `posts`
com `sets` e mais `sets`) que dizia mais ou menos isso:

> Esqueça essas coisas. Use `ALSA` e `Bluealsa`.

Lacônico assim. Eu não liguei... Continuei confiante que transformaria xixi em ouro com algum
`set` filosofal. Até que intoxicado de tanto `set` e indignado por ter que jogar uns meses
de trabalho fora e enconstar um `Raspberry-Pi` que havia comprado só para isso, decidi por
tentar `ALSA` e `Bluealsa`.

Nos meus áureos tempos utilizando `Slackware` já havia utilizado `ALSA` sem muita porcaria
na frente e tinha tido uma boa impressão. Mas naquela época eu usava para ouvir `CDs` via
o `workbone`. Nunca havia tentado `Bluetooth` "direto" com `ALSA`...

Depois de buscar um pouco sobre, descobri que o `Bluealsa` foi descontinuado e depois continuado
de novo via um `rebirth` e que agora era `Bluez-ALSA`. Dei uma olhada no projeto e já me agradou
o fato de ter um documento direto que ensina como buildar os fontes, sem enrolação, direto ao ponto.

Porém, nesse ponto, para preparar o ambiente de `build` deixo aqui o meu protesto sobre a
insanidade que é a nomenclatura dessas `libs` via gerenciadores de pacote. Cada uma segue um
padrão e varia de distro para distro. E o nome da `lib` via o gerenciador de pacotes às vezes
nada tem a ver com o nome da `lib`. Por que não se cria um padrão para nomear essas coisas?!!!!
Porra é muita falta de organização, muita deficiência em sensos básicos de engenharia, não sei
como esse pessoal que idealiza as pantanosas "bases" dessas coisas conseguem cagar dentro do vaso
(espero). Depois de descobrir o nome exato de cada dependência que precisei instalar no meu
ambiente de `build` (o que me consumiu mais tempo), rapidamente buildei o `Bluez-Alsa`, fiz os
ajustes com o `DBus-1` e o serviço `systemd` no meu `Raspbian` simplesmente subiu. Lindamente.
Mas não estava muito confiante. Algo que o pessoal já abandonou deve ser muito ruim... Fiquei uns
*30 minutos conectando e desconectando meus dispositivos `Bluetooth`*, **mas magicamente sempre
conectava**. Sem `reboot` ou `restart` no serviço!!!! Beleza: aí depois foi tocar som... No caso
encontrei o jeito de tocar `MP3` utilizando `mpg123` (o qual você precisa habilitar o suporte
durante o build). Cara, é lindo! Tocou sem engasgo. A conexão `Bluetooth` é muito mais estável.
Você simplesmente inicia o serviço `bluealsa` e conecta. Igual você esperaria fazer num
`smartphone` etc. Você não tem aquelas mensagens idiotas de não poder bindar um treco lá na fuckin
porra do `bluetooth`. Com `bluealsa` deu `systemctl start bluealsa` o serviço ativa e você vai embora
ouvir sua música e tocar sua vida, o computador simplesmente faz a parte dele e não fica na frente
te tomando tempo, te apurrinhando. O computador colabora, se torna de fato uma ferramenta, um meio
para você atingir o seu objetivo. Pelo visto, purificamos um barril que estava lotado de merda!!!!

Com o `Bluez-Alsa` já embarcado eu notei que é muito mais leve. O `Raspberry-Pi` parece esquentar
menos! Até agora só vi vantagem. Finalmente encontrei a melhor forma de tocar música via
`Bluetooth` para o meu estado de coisas: `ALSA` e `Bluez-ALSA`. Mais uma vez o simples
é que chegou e salvou o dia. Foi ou vai ser descontinuado no `Raspbian`? **Foda-se**. Eu buildo
o negócio por mim mesmo e purgo o lixo que estiver na frente me atrapalhando. `ALSA` sempre
vai estar lá. Não é perfeito, **mas é menos pior**.

Simplicidade é tudo. Se você precisar andar para frente, dê um passo, evite cambalhotas logo
de cara. Complexidade se precisar existir, é algo que precisa ser legitimamente composto.

Se tivesse me mantido nessa, eu teria economizado seis meses e já estaria utilizando minha placa
para ouvir minhas músicas. `Do not worry, stay simple and be happy`... :wink:

Agora, se você precisa fazer coisas como gravar música, produzir música no `Linux`... Cara,
sinceridade? Boa sorte e que você passe o mínimo possível de raiva.
É o que eu te desejo! :four_leaf_clover:

Ainda, se você quiser `100%` do tempo um som sem nenhum tipo de interferência, evite tecnologias
baseadas em rádio, use cabos :electric_plug:. Porém, mesmo assim não se adapte a camadas de
`software` instáveis que te prometem um `bluetooth` utilizável e não te dão nem um mínimo decente.

Para mim, nós seres humanos temos duas super habilidades: o pensamento e a adaptabilidade. Com toda
habilidade vem de quebra o seu efeito rebote. O rebote do pensamento é ficar pensando demais e não
agir. O rebote da adaptabilidade é se adaptar a situações inaceitáveis, normalizá-las, comer merda
e depois de um tempo sentir que está se deliciando com uma ganache dos deuses. Reconheça sua
natureza humana e suas exclusivas habilidades e: **cuidado com elas!** Se algo é ruim e não atende
os seus requisitos, jogue fora e vá procurar algo melhor que te atenda! :dart:

[`Voltar`](#tópicos)
