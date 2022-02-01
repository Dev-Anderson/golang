# Go Data Types

Na linguagem de programçaão Go, os tipos de dados refere-se a um sistema extensivo usado para declarar variáveis ou funções de diferentes tipos. O tipo de uma varia´vel determina quanto espaço ela ocupa no armazenamento e como o padrão de bits armazenado é interpretado. 
Os tipo em Go pode ser classificados da seguinte forma: 

 - Tipos booleanos 
	 - Eles são tipos booleanos e consistem em duas constantes predefinidas: (a) true (b) false
- Tipo numéricos
	- Eles são novamente tipo aritméticos e representam, tipos inteiros ou valores de ponto flutuante em todo o programa
- Tipo de string
	- Um tipo de string representa o conjunto de valores de string. seu valor é uma sequencia de bytes. Strings são tipos imutáveis que uma vez criados, não é possível alterar o conteúdo de uma string. O tipo de string pré-declarado é string. 
- Tipos derivados 
	- Eles incluem tipos de ponteiros, tipos de matriz, tipos de estrutura, tipos de união e tipos de função, tipos de fatia, tipos de interface, tipos de mapa e tipos de canal. 

Tipos de matriz e tipos de estrutura são chamados coletivamente de tipos agregados. O tipo de uma função especifica o conjunto de toda as funções com os mesmos tipos de parâmetro e resultado. Discutiremos os tipos básicos na seção a seguir, enquanto outros tipo serão abordados nos próximos capítulos. 

## Tipos inteiros

Os tipos inteiros independentes de arquitetura predefinidos são: 
Sr. Não.

Tipos e Descrição

1

**uint8**

Inteiros de 8 bits sem sinal (0 a 255)

2

**uint16**

Inteiros de 16 bits sem sinal (0 a 65535)

3

**uint32**

Inteiros de 32 bits sem sinal (0 a 4294967295)

4

**uint64**

Inteiros de 64 bits não assinados (0 a 18446744073709551615)

5

**int8**

Inteiros de 8 bits com sinal (-128 a 127)

6

**int16**

Inteiros de 16 bits assinados (-32768 a 32767)

7

**int32**

Inteiros de 32 bits assinados (-2147483648 a 2147483647)

8

**int64**

Inteiros de 64 bits assinados (-9223372036854775808 a 9223372036854775807)

## Tipos flutuantes 

Os tipos float independentes de arquitetura predefinidos são: 
Sr. Não.

Tipos e Descrição

1

**float32**

Números de ponto flutuante de 32 bits IEEE-754

2

**float64**

Números de ponto flutuante IEEE-754 de 64 bits

3

**complexo64**

Números complexos com partes reais e imaginárias float32

4

**complexo 128**

Números complexos com partes reais e imaginárias float64

O valor de um inteiro de n bits é n bits e é representando usado operações aritméticas de complemento de dois. 

## Outros tipos numéricos

Há também um conjunto de tipos numéricos com tamanhos específicos de implantação: 
Sr. Não.

Tipos e Descrição

1

**byte**

o mesmo que uint8

2

**runa**

o mesmo que int32

3

**uint**

32 ou 64 bits

4

**int**

mesmo tamanho que uint

5

**uintptr**

um inteiro sem sinal para armazenar os bits não interpretados de um valor de ponteiro