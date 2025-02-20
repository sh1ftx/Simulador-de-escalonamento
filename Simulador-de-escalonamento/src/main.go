package main

import (
	"fmt"       // Pacote para formatação de entrada e saída
	"math/rand" // Pacote para geração de números aleatórios
	"sort"      // Pacote para ordenação de slices e coleções
	"sync"      // Pacote para sincronização de goroutines
	"time"      // Pacote para manipulação de tempo e datas
)

/*
Este script simula o funcionamento de um sistema operacional básico, focando no gerenciamento de processos e escalonamento de CPU.
Ele implementa três algoritmos de escalonamento: FIFO (First-Come, First-Served), Round-Robin e Prioridade.
O objetivo é fornecer uma compreensão prática de como esses algoritmos funcionam e como eles afetam a execução dos processos.
*/

// ------------------------------------------
// Definição de cores para saída formatada
const (
	Reset  = "\033[0m"  // Reset da cor
	Cyan   = "\033[36m" // Cor ciano
	Green  = "\033[32m" // Cor verde
	Red    = "\033[31m" // Cor vermelha
	Yellow = "\033[33m" // Cor amarela
	Purple = "\033[35m" // Cor roxa
	Blue   = "\033[34m" // Cor azul
)

/*
As constantes acima são usadas para definir cores para a saída formatada no terminal.
Isso ajuda a visualizar melhor o estado dos processos durante a execução do script.
*/

// ------------------------------------------
// Definição dos estados de um processo
type ProcessState string

const (
	Ready     ProcessState = "Pronto"     // Estado pronto para execução
	Running   ProcessState = "Executando" // Estado em execução
	Paused    ProcessState = "Pausado"    // Estado pausado
	Completed ProcessState = "Finalizado" // Estado finalizado
	Error     ProcessState = "Erro"       // Estado de erro
)

/*
A definição dos estados de um processo é feita usando um tipo string.
Isso permite que o estado de um processo seja representado de forma clara e legível.
*/

// ------------------------------------------
// Estrutura de um processo
type Process struct {
	ID        int          // Identificador do processo
	BurstTime int          // Tempo de execução do processo
	Priority  int          // Prioridade do processo
	State     ProcessState // Estado atual do processo
}

/*
A estrutura Process define as propriedades de um processo, incluindo seu ID, tempo de execução, prioridade e estado atual.
Essa estrutura é usada para criar e gerenciar processos no sistema operacional simulado.
*/

// ------------------------------------------
// Função para exibir a lista de processos
func printProcesses(processes []Process) {
	fmt.Print("\033[H\033[2J") // Limpa a tela
	fmt.Println(Cyan + "\n==================== Lista de Processos ====================" + Reset)
	fmt.Println("ID | Tempo | Prioridade | Estado")
	for _, p := range processes {
		var color string
		switch p.State {
		case Ready:
			color = Blue
		case Running:
			color = Yellow
		case Paused:
			color = Purple
		case Completed:
			color = Green
		case Error:
			color = Red
		}
		fmt.Printf("%2d | %5d | %10d | %s%s%s\n", p.ID, p.BurstTime, p.Priority, color, p.State, Reset)
	}
	fmt.Println(Cyan + "==========================================================" + Reset)
}

/*
A função printProcesses exibe a lista de processos no terminal.
Ela usa cores para diferenciar os estados dos processos, tornando a visualização mais intuitiva.
*/

// ------------------------------------------
// Função para exibir uma legenda explicativa
func printLegend() {
	fmt.Println(Cyan + "\n==================== Legenda ====================" + Reset)
	fmt.Println(Blue + "Pronto" + Reset + ": O processo está pronto para ser executado.")
	fmt.Println(Yellow + "Executando" + Reset + ": O processo está atualmente em execução.")
	fmt.Println(Purple + "Pausado" + Reset + ": O processo foi pausado temporariamente.")
	fmt.Println(Green + "Finalizado" + Reset + ": O processo foi concluído.")
	fmt.Println(Red + "Erro" + Reset + ": O processo encontrou um erro.")
	fmt.Println(Cyan + "================================================" + Reset)
}

/*
A função printLegend exibe uma legenda explicativa para os estados dos processos.
Isso ajuda o usuário a entender o que cada cor representa.
*/

// ------------------------------------------
// Função para exibir explicações detalhadas sobre os algoritmos
func printExplanation(algorithm string) {
	switch algorithm {
	case "FIFO":
		typeWriterPrint(Cyan + "\n[Escalonador FIFO]" + Reset)
		typeWriterPrint("O algoritmo FIFO (First-Come, First-Served) executa os processos na ordem em que chegam.")
		typeWriterPrint("Isso significa que o primeiro processo a chegar é o primeiro a ser executado.")
		typeWriterPrint("Este algoritmo é simples, mas pode levar a problemas de inanição, onde processos com maior tempo de execução podem atrasar a execução de processos menores.")
	case "Round-Robin":
		typeWriterPrint(Cyan + "\n[Escalonador Round-Robin]" + Reset)
		typeWriterPrint("O algoritmo Round-Robin distribui o tempo de CPU entre os processos de forma igualitária, usando um quantum.")
		typeWriterPrint("Cada processo recebe uma fatia de tempo (quantum) para ser executado. Se o processo não for concluído dentro desse tempo, ele é pausado e colocado no final da fila.")
		typeWriterPrint("Este algoritmo é justo e evita a inanição, mas pode ter um overhead maior devido ao contexto de troca frequente.")
	case "Prioridade":
		typeWriterPrint(Cyan + "\n[Escalonador por Prioridade]" + Reset)
		typeWriterPrint("O algoritmo de Prioridade executa os processos com base na prioridade, onde processos com menor valor de prioridade são executados primeiro.")
		typeWriterPrint("Se dois processos tiverem a mesma prioridade, eles são executados na ordem de chegada.")
		typeWriterPrint("Este algoritmo pode levar a problemas de inanição se processos de baixa prioridade nunca tiverem a chance de serem executados.")
	}
}

/*
A função printExplanation exibe explicações detalhadas sobre os algoritmos de escalonamento.
Isso ajuda o usuário a entender como cada algoritmo funciona e quais são suas vantagens e desvantagens.
*/

// ------------------------------------------
// Função para exibir um quiz ao final da execução
func runQuiz(algorithm string) {
	var choice int
	fmt.Println(Green + "\n===== Quiz de Avaliação =====" + Reset)
	fmt.Println("1. Responder perguntas sobre o algoritmo")
	fmt.Println("2. Pular o quiz")
	fmt.Println("3. Assistir a execução novamente")
	fmt.Print("Escolha uma opção: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		quizQuestions(algorithm)
	case 2:
		fmt.Println(Green, "Quiz pulado.", Reset)
	case 3:
		switch algorithm {
		case "FIFO":
			fifoScheduler(generateProcesses())
		case "Round-Robin":
			roundRobinScheduler(generateProcesses(), 200)
		case "Prioridade":
			priorityScheduler(generateProcesses())
		}
	default:
		fmt.Println(Red, "Opção inválida!", Reset)
	}
}

/*
A função runQuiz exibe um quiz ao final da execução do algoritmo.
Isso permite que o usuário teste seu conhecimento sobre o algoritmo que acabou de ser executado.
*/

// ------------------------------------------
// Função para gerar processos aleatórios
func generateProcesses() []Process {
	return []Process{
		{ID: 1, BurstTime: rand.Intn(500) + 500, Priority: rand.Intn(10), State: Ready},
		{ID: 2, BurstTime: rand.Intn(500) + 500, Priority: rand.Intn(10), State: Ready},
		{ID: 3, BurstTime: rand.Intn(500) + 500, Priority: rand.Intn(10), State: Ready},
		{ID: 4, BurstTime: rand.Intn(500) + 500, Priority: rand.Intn(10), State: Ready},
	}
}

/*
A função generateProcesses gera uma lista de processos aleatórios.
Isso é útil para simular diferentes cenários de escalonamento.
*/

// ------------------------------------------
// Função para perguntas do quiz
func quizQuestions(algorithm string) {
	var answer int
	switch algorithm {
	case "FIFO":
		questions := []struct {
			question string
			options  []string
			answer   int
		}{
			{"O que significa FIFO?", []string{"First In First Out", "First In First Over", "Fast In Fast Out"}, 1},
			{"Qual é a principal desvantagem do algoritmo FIFO?", []string{"Inanição", "Overhead", "Processos longos podem atrasar processos curtos"}, 3},
		}
		for _, q := range questions {
			fmt.Println(q.question)
			correctAnswer := q.options[q.answer-1]
			rand.Shuffle(len(q.options), func(i, j int) { q.options[i], q.options[j] = q.options[j], q.options[i] })
			for i, option := range q.options {
				fmt.Printf("%d. %s\n", i+1, option)
			}
			fmt.Print("Resposta: ")
			fmt.Scan(&answer)
			if q.options[answer-1] == correctAnswer {
				fmt.Println(Green, "Correto!", Reset)
			} else {
				fmt.Println(Red, "Errado. A resposta correta é:", correctAnswer, Reset)
			}
		}
	case "Round-Robin":
		questions := []struct {
			question string
			options  []string
			answer   int
		}{
			{"O que é um quantum no algoritmo Round-Robin?", []string{"O tempo total de execução", "A fatia de tempo que cada processo recebe", "A prioridade do processo"}, 2},
			{"Qual é a principal vantagem do algoritmo Round-Robin?", []string{"Simplicidade", "Evita inanição", "Menor overhead"}, 2},
		}
		for _, q := range questions {
			fmt.Println(q.question)
			correctAnswer := q.options[q.answer-1]
			rand.Shuffle(len(q.options), func(i, j int) { q.options[i], q.options[j] = q.options[j], q.options[i] })
			for i, option := range q.options {
				fmt.Printf("%d. %s\n", i+1, option)
			}
			fmt.Print("Resposta: ")
			fmt.Scan(&answer)
			if q.options[answer-1] == correctAnswer {
				fmt.Println(Green, "Correto!", Reset)
			} else {
				fmt.Println(Red, "Errado. A resposta correta é:", correctAnswer, Reset)
			}
		}
	case "Prioridade":
		questions := []struct {
			question string
			options  []string
			answer   int
		}{
			{"Como o algoritmo de Prioridade decide a ordem de execução dos processos?", []string{"Pela ordem de chegada", "Pelo tempo de burst", "Pela prioridade"}, 3},
			{"Qual é a principal desvantagem do algoritmo de Prioridade?", []string{"Inanição de processos de baixa prioridade", "Overhead", "Complexidade"}, 1},
		}
		for _, q := range questions {
			fmt.Println(q.question)
			correctAnswer := q.options[q.answer-1]
			rand.Shuffle(len(q.options), func(i, j int) { q.options[i], q.options[j] = q.options[j], q.options[i] })
			for i, option := range q.options {
				fmt.Printf("%d. %s\n", i+1, option)
			}
			fmt.Print("Resposta: ")
			fmt.Scan(&answer)
			if q.options[answer-1] == correctAnswer {
				fmt.Println(Green, "Correto!", Reset)
			} else {
				fmt.Println(Red, "Errado. A resposta correta é:", correctAnswer, Reset)
			}
		}
	}
}

/*
A função quizQuestions exibe perguntas sobre o algoritmo de escalonamento.
Isso ajuda a reforçar o conhecimento do usuário sobre o algoritmo.
*/

// ------------------------------------------
// Função para introdução e confirmação antes da execução
func introduction(algorithm string) {
	printExplanation(algorithm)
	var choice int
	fmt.Println(Green + "\nPressione 1 para continuar para a execução ou 2 para voltar ao menu principal." + Reset)
	fmt.Print("Escolha uma opção: ")
	fmt.Scan(&choice)
	if choice != 1 {
		main()
	}
}

/*
A função introduction exibe uma introdução sobre o algoritmo e pede confirmação do usuário antes de continuar.
Isso garante que o usuário entenda o algoritmo antes de vê-lo em ação.
*/

// ------------------------------------------
// Função para exibir uma barra de carregamento
func printProgressBar(progress int) {
	bar := "["
	color := Red
	if progress > 50 {
		color = Yellow
	}
	if progress > 75 {
		color = Green
	}
	for i := 0; i < 50; i++ {
		if i < progress/2 {
			bar += "="
		} else {
			bar += " "
		}
	}
	bar += "]"
	fmt.Printf("\r%s%s %d%%%s\n", color, bar, progress, Reset)
}

/*
A função printProgressBar exibe uma barra de carregamento no terminal.
Isso ajuda a visualizar o progresso da execução de um processo.
*/

// ------------------------------------------
// Função para impressão de texto com efeito de máquina de escrever
func typeWriterPrint(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()
}

/*
A função typeWriterPrint imprime texto com um efeito de máquina de escrever.
Isso torna a saída do terminal mais interessante e envolvente.
*/

// ------------------------------------------
// Simulação do algoritmo FIFO (First-Come, First-Served)
func fifoScheduler(processes []Process) {
	introduction("FIFO")
	printLegend()
	totalBurstTime := 0
	for _, p := range processes {
		totalBurstTime += p.BurstTime
	}
	elapsedTime := 0
	for i := range processes {
		processes[i].State = Running
		printProcesses(processes)
		typeWriterPrint(fmt.Sprintf(Yellow+"Processo %d está executando. (FIFO: Executa na ordem de chegada)"+Reset, processes[i].ID))
		for j := 0; j < processes[i].BurstTime; j += 100 {
			time.Sleep(100 * time.Millisecond)
			elapsedTime += 100
			printProgressBar(min(elapsedTime*100/totalBurstTime, 100))
		}
		processes[i].State = Completed
		printProgressBar(100)
		typeWriterPrint(fmt.Sprintf(Green+"Processo %d finalizado! (FIFO: Executa na ordem de chegada)"+Reset, processes[i].ID))
		printProcesses(processes)
	}
	runQuiz("FIFO")
}

/*
A função fifoScheduler simula o algoritmo FIFO.
Ela executa os processos na ordem em que chegam e exibe o progresso no terminal.
*/

// ------------------------------------------
// Simulação do algoritmo Round-Robin
func roundRobinScheduler(processes []Process, quantum int) {
	introduction("Round-Robin")
	printLegend()
	totalBurstTime := 0
	for _, p := range processes {
		totalBurstTime += p.BurstTime
	}
	elapsedTime := 0
	queue := append([]Process{}, processes...)
	for len(queue) > 0 {
		p := &queue[0]
		p.State = Running
		printProcesses(queue)
		typeWriterPrint(fmt.Sprintf(Yellow+"Processo %d está executando por %d ms. (Round-Robin: Cada processo recebe uma fatia de tempo)"+Reset, p.ID, quantum))
		for j := 0; j < min(p.BurstTime, quantum); j += 100 {
			time.Sleep(100 * time.Millisecond)
			elapsedTime += 100
			printProgressBar(min(elapsedTime*100/totalBurstTime, 100))
		}
		p.BurstTime -= quantum
		if p.BurstTime > 0 {
			p.State = Paused
			printProgressBar(100)
			typeWriterPrint(fmt.Sprintf(Purple+"Processo %d pausado! (Round-Robin: Processo pausado e colocado no final da fila)"+Reset, p.ID))
			queue = append(queue[1:], *p)
		} else {
			p.State = Completed
			printProgressBar(100)
			typeWriterPrint(fmt.Sprintf(Green+"Processo %d finalizado! (Round-Robin: Cada processo recebe uma fatia de tempo)"+Reset, p.ID))
			queue = queue[1:]
		}
		printProcesses(queue)
	}
	runQuiz("Round-Robin")
}

/*
A função roundRobinScheduler simula o algoritmo Round-Robin.
Ela distribui o tempo de CPU entre os processos de forma igualitária, usando um quantum.
*/

// ------------------------------------------
// Simulação do algoritmo Priority Scheduling
func priorityScheduler(processes []Process) {
	introduction("Prioridade")
	printLegend()
	totalBurstTime := 0
	for _, p := range processes {
		totalBurstTime += p.BurstTime
	}
	elapsedTime := 0
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].Priority < processes[j].Priority
	})
	for i := range processes {
		processes[i].State = Running
		printProcesses(processes)
		typeWriterPrint(fmt.Sprintf(Yellow+"Processo %d está executando. (Prioridade: Executa com base na prioridade)"+Reset, processes[i].ID))
		for j := 0; j < processes[i].BurstTime; j += 100 {
			time.Sleep(100 * time.Millisecond)
			elapsedTime += 100
			printProgressBar(min(elapsedTime*100/totalBurstTime, 100))
		}
		processes[i].State = Completed
		printProgressBar(100)
		typeWriterPrint(fmt.Sprintf(Green+"Processo %d finalizado! (Prioridade: Executa com base na prioridade)"+Reset, processes[i].ID))
		printProcesses(processes)
	}
	runQuiz("Prioridade")
}

/*
A função priorityScheduler simula o algoritmo de escalonamento por Prioridade.
Ela executa os processos com base em sua prioridade, com processos de maior prioridade sendo executados primeiro.
*/

// ------------------------------------------
// Função auxiliar para encontrar o mínimo de dois valores
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
A função min retorna o menor de dois valores inteiros.
Isso é útil para calcular o progresso da execução dos processos.
*/

// ------------------------------------------
// Menu interativo
func main() {
	rand.Seed(time.Now().UnixNano()) // Inicializa a semente do gerador de números aleatórios
	var wg sync.WaitGroup            // Cria um grupo de espera para sincronizar goroutines
	var choice int                   // Variável para armazenar a escolha do usuário
	for {
		fmt.Println(Yellow + "\n===== Menu de Escalonamento =====" + Reset)
		fmt.Println("1. FIFO")
		fmt.Println("2. Round Robin (Quantum = 200ms)")
		fmt.Println("3. Prioridade")
		fmt.Println("4. Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&choice) // Lê a escolha do usuário
		switch choice {
		case 1:
			wg.Add(1) // Adiciona uma goroutine ao grupo de espera
			go func() {
				defer wg.Done()                    // Marca a goroutine como concluída ao final
				fifoScheduler(generateProcesses()) // Executa o escalonador FIFO
			}()
		case 2:
			wg.Add(1)
			go func() {
				defer wg.Done()
				roundRobinScheduler(generateProcesses(), 200) // Executa o escalonador Round-Robin com quantum de 200ms
			}()
		case 3:
			wg.Add(1)
			go func() {
				defer wg.Done()
				priorityScheduler(generateProcesses()) // Executa o escalonador por Prioridade
			}()
		case 4:
			fmt.Println(Green + "Saindo..." + Reset)
			return // Sai do programa
		default:
			fmt.Println(Red + "Opção inválida!" + Reset)
		}
		wg.Wait() // Espera todas as goroutines concluírem
	}
}

/*
A função main exibe um menu interativo para o usuário escolher o algoritmo de escalonamento.
Ela usa goroutines para executar os algoritmos de forma assíncrona e sincroniza a execução usando um grupo de espera.
*/
