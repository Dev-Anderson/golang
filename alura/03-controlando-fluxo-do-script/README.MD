# 03 Controlando o fluxo do script

Dando sequencia à construção da nossa aplicação, agora que já conseguimos capturar o comando escolhido pelo usuário, podemos fazer algo de acordo com o comando digitado. 
Se o usuário escolheu o comando 1, vamos fazer algo, se ele escolheu o comando 2, fazemos outra coisa, e assim por diante. Quando queremos saber se o usuário escolheu tal comando, precisamos utilizar a instrução `if` da programação. 
A condição do `if` no Go, não fica entre parênteses, como já é prática de outras linguagens: 

    if comadnod == 1 {}
E a condição deve sempre retornar um booleano, ou seja, `true` ou `false`. 
Como queremos testar os 3 comandos, vamos colocá-lo no `if`:

add-condicao.go

E por último, se não for nenhum comando conhecido, vamos imprimir uma mensagem, demonstrando que não conhecemos o comando digitado. Vamos aproveitar e colocar as mensagens dos outros comandos: 

add-condicao.go

Com isso, concluímos os `if` s, que não possui muito mistério, as diferenças para as outras linguagens é o não uso de parênteses na condição, que leve sempre retornar um booleano.

Além do `if`, sabemos que existem outras instruções para controlarmos o fluxo da nossa aplicação E uma que se adéqua melhor ao nosso código é a instrução `switch`. 
Essa instrução recebe uma variável e dá possíveis situações para cada valor dessa variável. Nós dizemos quais são só casos, e para cada um há uma situação. 
Caso a variável valha 0, acontece algo, caso valha 1, acontece outra, e assim por diante.

    switch comando {
    case 1: 
	    fmt.Println("Monitorando..")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")


mas se o valor da variável não estiver em nenhum dos casos listados? Para isso, existe o caso `default`, que é o que será executado se o nosso casos não forem atendidos: