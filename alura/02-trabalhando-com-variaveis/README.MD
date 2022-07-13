# Exibindo o menu

Agora, após as mensagens de boas vinda e de verão do programa, queremos dar as opções para o usuário, se ele quer monitorar os sites, exibir os logs, etc, e vamos mostrar tudo isso em um menu. 
Disponibilizaremos o menu imprimindo as suas opções para o usuário, ele irá digitar a opção deseja, e nós capturaremos o que foi digitado. Então, primeiramente devemos imprimir as opções:

main.go

## Capturando a entrada do usuário

para capturar o que o usuário escrever, existe a função "Scanf", tam´bem do pacote "fmt". Ela recebe dois arumentos, um modificador (o que queremos receber como entrada, um inteiro, string etc) e um ponteiro para a variável que guardará a entrada do usuário. 
Então, vamos declara essa variável: 

main.go

E agora nós passamos para a função "Scanf" o valor "%d", que significa que esperamos receber um número inteiro, e a variável "comando":

main.go

Daqui a pouco falaremos sobre o `&`, que se encontrada à frente da variável `comando`. Ou seja, a função salará na variável o que o usu´rio digitar. E para verificar se a entrada está mesmo sendo salva na variável `comando`, vamos imprimi-la: 

main.go


## Entendendo o ponteiro

Sobre o `&` visto antes, ele significa o endereço da vairável que queremos salvar a entrada, pois a função "Scanf" não espera uma variável, e sim o seu endereço, um ponteiro para a variável. 
A variável nada mais é do que uma "caixa", onde guardamos dados. Essa "caixa" está em algum lugar da memória do nosso computador, e esse lugar, o endereço da nossa "caixa", é o que chamamos de ponteiro. 
Para descobrir o endereço da virável, basta colocar o `&` à frente dela. 

## Devemos especificar o que esperamos ser digitado? 

A variável `comando` é do tipo inteiro, logo, só pode receber números inteiros. Se sabemos disso, por que devemos especificar que esperamos receber um número inteiro na função `Scanf`, através do modificador `"d"`? Na verdade, nós não precisamos. 
Alguém do Go também pensou nisso, para isso criou a função `Scan` (sem letra f). Nela, nós não precisamos especificar o modificador que esperamos:

main.go

Ou seja, estamos esperando  um daod, e quando recebermos, colocamos dentro de uma variável inteira. Como a variável é inteira, a função que deve esperar, receber valores inteiros. 
Ao executar o programa, caso digitemos um valor não-inteiro, como uma letra, por exemplo, não será guardado nenhum valor no endereço da variável `comando`, logo ela continuará com o valor padrão atribuído e ela no momento da sua inicialização, que é 0. Ou seja, o ideia seria que detectássemos quando um valor não-inteiro fosse digitado. 

Agora, após as mensagem de boas vindas e de versão do programa, queremos dar as opções para o usuário, se ele quer monitorar os sites, exibir os logs, etc, e vamos mostrar tudo isso em um menu!
Disponibilizaremos o menu imprimindo as suas opções para o usuário, ele irá digitar a opção desejada, e nós capturaremos o que foi digitado. Então, primeiramente devemos imprimir as opções:

main.go

Ao executar o programa, ele fica travado esperando que digitemos algo. Ao digitar. 

## Entendendo o ponteiro

Sobre o `&` visto antes, ele significa o endereço da variável que queremos salvar a entrada, pois a função `Scanf` não espera uma variável, e sim o seu endereço, um ponteiro para a variável. 
A variável nada mais é do que uma "caixa", onde guardamos dados. Essa "caixa" está em algum lugar da memória do nosso computador, e esse lugar, o endereço da nossa "caixa", é o que chamamos de ponteiro. 
Para descobrir o endereço da variável, basta colocar o `&` à frente dela. 

## Devemos especificar o que esperamos ser digitado? 

A variável `comando` é do tipo inteiro, logo, só pode receber números inteiros, Se sabemos disso, por que devemos especificar que esperamos receber um número inteiro na função `Scanf`, através do modificador `"%d"`? Na verdade, nós não precisamos. 
Alguém do Go também pensou nisso, para  isso criou a função `Scan` (sem a letra `f`). Nela, nós não precisamos especificar o modificador que esperamos:

sem-esperar-entrada.go

Ou seja, estamos esperando um dado, e quando o recebermos, colocamos dentro de uma variável inteira. Como a variável é inteira, a função sabe que deve esperar, receber valores inteiros. 
Ao executar o programa, caso digitemos um valor não-inteiro, como uma letra, por exemplo, não será guardado nenhum valor no endereço da variável `comando`, logo ela continuará com o valor padrão atribuído a ela no momento da sua inicialização, que é 0. Ou seja, o ideia seria que detectássemos quando um valor não-inteiro fosse digitado. 

O que aprendemos? 
* Declaração de variáveis 
	* Valor padrão das variáveis sem valor inicial 
* Inferência de tipos de veriáveis 
* Descobrir o tipo da variável 
* Através da função `TypeOf`, do pacote `reflect`
* Declaração curta de variáveis  
* Ler dados digitados do usuário 
* Através das funções `Scanf` e `Scan`, do pacote `fmt`
* Mais convenções do Go
* Variáveis e "imports" não utilizados são deletado

