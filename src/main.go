package main

import (
    "fmt"
    "math/rand"
    "sort"
    "sync"
    "time"
)

// Definição de cores para saída formatada
const (
    Reset  = "\033[0m"
    Cyan   = "\033[36m"
    Green  = "\033[32m"
    Red    = "\033[31m"
    Yellow = "\033[33m"
    Purple = "\033[35m"
    Blue   = "\033[34m"
)

// Definição dos estados de um processo
type ProcessState string

const (
    Ready     ProcessState = "Pronto"
    Running   ProcessState = "Executando"
    Paused    ProcessState = "Pausado"
    Completed ProcessState = "Finalizado"
    Error     ProcessState = "Erro"
)

// Estrutura de um processo
type Process struct {
    ID        int
    BurstTime int
    Priority  int
    State     ProcessState
}

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

// Função para gerar processos aleatórios
func generateProcesses() []Process {
    return []Process{
        {ID: 1, BurstTime: rand.Intn(500) + 500, Priority: rand.Intn(10), State: Ready},
        {ID: 2, BurstTime: rand.Intn(500) + 500, Priority: rand.Intn(10), State: Ready},
        {ID: 3, BurstTime: rand.Intn(500) + 500, Priority: rand.Intn(10), State: Ready},
        {ID: 4, BurstTime: rand.Intn(500) + 500, Priority: rand.Intn(10), State: Ready},
    }
}

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

// Função para impressão de texto com efeito de máquina de escrever
func typeWriterPrint(text string) {
    for _, char := range text {
        fmt.Print(string(char))
        time.Sleep(50 * time.Millisecond)
    }
    fmt.Println()
}

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

// Função auxiliar para encontrar o mínimo de dois valores
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// Menu interativo
func main() {
    rand.Seed(time.Now().UnixNano())
    var wg sync.WaitGroup
    var choice int
    for {
        fmt.Println(Yellow + "\n===== Menu de Escalonamento =====" + Reset)
        fmt.Println("1. FIFO")
        fmt.Println("2. Round Robin (Quantum = 200ms)")
        fmt.Println("3. Prioridade")
        fmt.Println("4. Sair")
        fmt.Print("Escolha uma opção: ")
        fmt.Scan(&choice)
        switch choice {
        case 1:
            wg.Add(1)
            go func() {
                defer wg.Done()
                fifoScheduler(generateProcesses())
            }()
        case 2:
            wg.Add(1)
            go func() {
                defer wg.Done()
                roundRobinScheduler(generateProcesses(), 200)
            }()
        case 3:
            wg.Add(1)
            go func() {
                defer wg.Done()
                priorityScheduler(generateProcesses())
            }()
        case 4:
            fmt.Println(Green + "Saindo..." + Reset)
            return
        default:
            fmt.Println(Red + "Opção inválida!" + Reset)
        }
        wg.Wait()
    }
}
