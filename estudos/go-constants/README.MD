# Go Constantes

Constantes referem-se a valores fixo que o programa não pdoe alterar durante a sua execução. esses valores fixo também são chamados de literais. 
As constantes podem ser de qualquer um dos tipos de dados básicos, como uma constante inteira, uma constante flutuante, uma constante de caractere ou uma string literal. Há também constantes de enumeração. 
Constantes são tratadas como variáveis regulares, exceto que seus valores não podem ser modificados após sua definição. 

## Literais inteiros

Um literal inteiro pode ser uma constante decimal, octal ou hexadecimal. Um prefixo especifica a base ou raiz: 0x ou 0x para hexadecimal, 0 para octal e nada para decimal. 
Um Literal inteiro também pode ter um sufixo que é uma combinação de U e L, para usingned e long, respectivamente. O sufixo pode ser maiúsculo ou minúsculo e pode estar em qualquer ordem. 
Aqui estão alguns exemplos de literais inteiros: 

> 212         /* Legal */ 
> 215u        /* Legal */ 
> 0xFeeL      /* Legal
> */ 078         /* Illegal: 
> 8 is not an octal digit */ 
> 032UU       /* Illegal: cannot repeat a suffix */

A seguir estão outros exemplos de vários tipos de literais inteiros: 

> 85         /* decimal */ 
> 0213       /* octal */ 
> 0x4b       /*> hexadecimal */ 
> 30         /* int */ 
> 30u        /* unsigned int */ 
> 30l > /* long */ 
> 30ul       /* unsigned long */

## Literais de ponto flutuante 
Um literal de ponto flutuante tem uma parte inteira, um ponto decimal, uma parte fracionária e uma parte expoente. Você pode representar literais de ponto flutuante na forma decimal ou exponencial. 
Ao representar usando a forma decimal, você deve incluir um ponto decimal, o expoente ou ambos e, ao representar usando a forma exponencial, você deve incluir a parte inteira, a parte fracionária ou ambas. O expoente com sinal é introduzido por e ou E. 
Aqui estão alguns exemplos de literais de ponto flutuante: 

> 3.14159       /* Legal */ 
> 314159E-5L    /* Legal */ 
> 510E          /* Illegal: incomplete exponent */ 210f          /* Illegal: no decimal
> or exponent */ 
> .e55          /* Illegal: missing integer or fraction

## Sequência de fuga 

Quando certos caracteres são precedidos por uma barra invertida, eles terão um significado especial em Go. Estes são conhecidos como código de sequência de sequencia de escape que são suadoa para representar nova linha (\n), tabulação (\t), retrocesso, etc. Aqui, você tem uma lista de alguns desses código de sequência de escape: 
|Sequência de fuga| Significado  |
|--|--|
|//| / personagem  |
|/'| ' personagem  |
|/"| " personagem  |
|/?| ? personagem  |
|/uma| Alerta ou campainha  |
|/b| Retrocesso |
|/f| Feed de formulário |
|/n| Nova linha |
|/r| Devolução de carro |
|/t| Guia horizontal |
|/v| Guia vertical |
|/oo| Número octal de um a três digitos |
|/xh...| Número hexadecimal de um ou mais dígitos |

Verificar o "exemplo01.go"

## Literais de string em Go

Literais ou constantes de string são colocados entre aspas duplas. Uma string contém caracteres semelhantes a literais de caracteres: caracteres simples, sequencias de escape e caracteres universais. 
Você quebrar uma linha longa em várias linhas usando literais de string e separando-as usando espaços em branco. 
Aqui estão alguns exemplos de literais de string. Todas as 3 formas são strings idênticas. 

> "hello, dear"
> "hello, \
dear"
"hello, " "d" "ear"

## A plavra-chave const

Você pode usar o prefixo const para declarar constantes com um tipo específicos da seguinte forma

    const variable type = value
Verificar o "exemplo02.go"
Observe que é uma boa prática de programação constantes em MAIÚSCULAS. 

