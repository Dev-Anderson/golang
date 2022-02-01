# Go Operadores 

Um Operador é um símbolo que diz ao compilador para realizar manipulações matemáticas ou lógicas específicas. A a linguagem go é rica em operadores integrados e fornece os seguintes tipos de operadores: 

- Operadores aritméticos 
- Operadores Relacionais
- Operadores lógicos 
- Operadores bit a bit 
- Operadores de atribuição 
- Operadores diversos 

## Operadores aritméticos 
A tabela a seguir mostra todos os operadores aritméticos suportados pela linguagem Go. Suponha que a variável A tenha 10 e a variável B tenha 20 então: 
|Operador | Descrição  | Exemplo  |
|--|--| --|
| + | Adiciona dois operadores | A + B dá 30 |
| - | Subtrai o segundo operando do primeiro | A - B dá -10 |
| * | Multiplica os dois operadores | A * B dá 200 |
| / | Divide o numerador pelo denominador | B / A dá 2 |
| % | Operador de módulo dá o resto após uma divisão inteira | B % A dá 0 |
| ++ | Operador de incremento. Aumenta o valor inteiro em um | A++ dá 11 |
| -- | Operador de decremento. Ele diminui o valor inteiro em um | A-- dá 9 |


## Operadores relacionais 

A tabela a seguir lista todos os operadores relacionais suportados pela Linguagem Go. Suponha que a variável A tenha 10 e a variável B tenha 20, então: 
|Operador | Descrição  | Exemplo  |
|--|--| --|
| == | Verifica se os valores de dois operandos são iguais ou não, se sim, a condição se torna verdadeira. | (A==B) não é verdade |
| != | Verifica se os valores de dois operandos são iguais ou não, se os valores não forem iguais, a condição se tornará verdadeira | (A!=B) é verdadeira |
| > | Verifica se o valor do operando esquerdo é maior que o valor do operando direito, se sim, a condição se torna verdadeira | (A>B) não é verdade |
| < | Verifica se o valor do operando esquerdo é menor que o valor do operando direito, se sim, a condição se torna verdadeira | (A<B) é verdadeira |
| >= | Verifica se o valor do operando esquerdo é maior ou igual ao valor do operando direito, se sim, a condição se torna verdadeira | (A>=B) não é verdade |
| <= | Verifica se o valor do operando esquerdo é menor ou igual ao valor do operando direito, se sim, a condição se torna verdadeira | (A<=B) é verdadeira |

## Operadores lógicos 

A tabela a seguir lista todos os operadores lógicos suportados pela linguagem Go. Suponha que a variável A tenha 1 e a variável B tenha 0, então: 
|Operador | Descrição  | Exemplo  |
|--|--| --|
| && | Chamado operador lógico AND. Se ambos os operandos forem diferentes de zero, então a condição se torna verdadeira | (A && B) é falso |
| 2 barras |Chamado operador OU lógico. Se qualquer um dos dois operadores for diferente de zero, então a condição se torna verdadeira  | (A || B) é verdadeira |
| ! | Chamado operador NOT lógico. Use para reverter o estado lógico de seu operando. Se uma condição for verdadeira, o operador NOT lógico tornará falso | !(A && B) é verdadeira |

A tabela a seguir mostra todos os operadores lógicos suportados pela linguagem Go. Suponha que a variável A seja verdadeira e a variável B seja falsa, então:
|Operador | Descrição  | Exemplo  |
|--|--| --|
| && | Chamado operador lógicos AND. Se ambos os operadores forem falsos, a condição se tornará falsa | (A && B) é falso |
| 2 barras | Chamado operador OU lógicos. Se qualquer um dos dois operadores for verdadeiro, a condição se tornará verdadeira | (A  B) é verdadeira |
| ! | Chamado operador NOT lógico. Use para reverter o estado lógico de seu operando. Se uma condição for verdadeira, o operador lógico NOT a tornará falsa | !(A && B) é verdadeira |

## Operadores bit a bit

Os operadores bit a bit trabalham em bits e realizam operações bit a bit. As tabelas verdade para &, | e ^. 

## Operadores de Atribuição

A tabela a seguir lista todos os operadores de atribuição suportados pela linguagem Go
|Operador | Descrição  | Exemplo  |
|--|--| --|
| = |Operador de atribuição simples, Atribui valores dos operados do lado direito ao operado ao lado esquerdo | C = A + B atribuirá valor de A + B em C |
| += |Adicionar operador de atribuição AND< adiciona o operado direito ao operando esquerdo e atribui o resultado ao operando esquerdo | C += A é equivalente a C = C + A |
| -= |Subtrair o operador de atribuição AND, subtrai o operando direito do operando esquerdo e atribui o resultado ao operando esquerdo | C -= A é equivalente a C = C - A |
| *= |Multiply AND operador de atribuição, multiplica o operando direito pelo operando esquerdo e atribui o resultado ao operando esquerdo |C *= A é equivalente a C = C * A  |
| /= |Operador de atribuição de divisão AND, ele divide o operando esquerdo com o operado direito e atribui o resultado ao operado esquerdo | C /= A é equivalente a C = % A |
| %= |Operador de atribuição de módulo AND, leva módulo usado dois operados e atribui o resultado ao operando esquerdo | C %= A é equivalente a C = % A |
| <<= |Deslocamento à esquerda E operador de atribuição | C <<= 2 é igual a C = C << 2 |
| >>= |Deslocamento à direita E operador de atribuição | C >>= 2 é igual a C = C >> 2 |
| &= |Operador de atribuição AND bit a bit | C &= 2 é o mesmo que C = & 2 |
| ^= |OR exclusivo bit a bit e operador de atribuição | C ^= 2 é o mesmo que C = C ^ 2 |
| ]= |OR inclusivo bit a bit e operador de atribuição | C barra= e igual a C = C }2 |

## Precedência de operadores em Go

A precedência do operador determina o agrupamento de termos em uma expressão. Isso afeta como uma expressão é avaliada. Certos operadores têm precedência mais alta do que outros, por exemplo, operador de multiplicação tem maior precedência do que o operador de adição. 
Aqui, os operadores com maior precedência aparecem na parte superior da tabela, aqueles com menor prioridade aparecem na parte inferior. Dentro de uma expressão, os operadores de maior precedência serão avaliados primeiro. 

|Categoria| Operador  | Associatividade  |
|--|--| --|
| Pós-fixação |() [] . ++ -- | Da esquerda para direita |
| Unário |+ - ! ~ ++ -- (tipo) * & size of| Direta para esquerda |
| Multiplicativo|* / %| Da esquerda para direita |
| Aditivo|+ - | Da esquerda para direita |
| Mudança|<<>> | Da esquerda para direita |
| Relacional|<<=>>= | Da esquerda para direita |
| Igualdade|==!= | Da esquerda para direita |
| E bit a bit|& | Da esquerda para direita |
| Xor bit a bit|^ | Da esquerda para direita |
| OU bit a bit| barra | Da esquerda para direita |
| E Lógico|&& | Da esquerda para direita |
| OU lógico |barra barra | Da esquerda para direita |
| Atribuição|= += -= *= /= %=>>= <<= &= ^= barra = |Direita para esquerda |
| Virgula|,|Da esquerda para direta |
