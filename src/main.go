package main

import (
    "fmt"
    "math/rand"
    "sort"
    "sync"
    "time"
)

// -------------------------------------------------------------------
// Definição de cores para saída formatada no terminal
const (
    Reset  = "\033[0m"   // Reset das cores para padrão
    Cyan   = "\033[36m"  // Ciano para títulos e divisórias
    Green  = "\033[32m"  // Verde para processos iniciando
    Red    = "\033[31m"  // Vermelho para processos finalizando
    Yellow = "\033[33m"  // Amarelo para tempos e cálculos matemáticos
    Purple = "\033[35m"  // Roxo para pausas de tarefas
    Blue   = "\033[34m"  // Azul para estados prontos
)

// -------------------------------------------------------------------
// Definição dos estados que um processo pode assumir

type ProcessState string

const (
    Ready     ProcessState = "Pronto"       // Processo está pronto para execução
    Running   ProcessState = "Executando"  // Processo está em execução
    Paused    ProcessState = "Pausado"     // Processo foi pausado temporariamente
    Completed ProcessState = "Finalizado"  // Processo foi concluído
    Error     ProcessState = "Erro"        // Processo encontrou um erro
)

// -------------------------------------------------------------------
// Estrutura de um processo

type Process struct {
    ID        int          // Identificador único do processo
    BurstTime int          // Tempo de execução do processo
    Priority  int          // Prioridade do processo (quanto menor, mais prioritário)
    State     ProcessState // Estado atual do processo
}

// -------------------------------------------------------------------
// Função para exibir a lista de processos
func printProcesses(processes []Process) {
    fmt.Print("\033[H\033[2J") // Limpa a tela para melhor visualização
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

// -------------------------------------------------------------------
// Função para exibir uma legenda explicativa dos estados dos processos
func printLegend() {
    fmt.Println(Cyan + "\n==================== Legenda ====================" + Reset)
    fmt.Println(Blue + "Pronto" + Reset + ": O processo está pronto para ser executado.")
    fmt.Println(Yellow + "Executando" + Reset + ": O processo está atualmente em execução.")
    fmt.Println(Purple + "Pausado" + Reset + ": O processo foi pausado temporariamente.")
    fmt.Println(Green + "Finalizado" + Reset + ": O processo foi concluído.")
    fmt.Println(Red + "Erro" + Reset + ": O processo encontrou um erro.")
    fmt.Println(Cyan + "================================================" + Reset)
}

// -------------------------------------------------------------------
// Função para exibir explicações detalhadas sobre os algoritmos de escalonamento
func printExplanation(algorithm string) {
    switch algorithm {
    case "FIFO":
        typeWriterPrint(Cyan + "\n[Escalonador FIFO]" + Reset)
        typeWriterPrint("O algoritmo FIFO (First-Come, First-Served) executa os processos na ordem em que chegam.")
    case "Round-Robin":
        typeWriterPrint(Cyan + "\n[Escalonador Round-Robin]" + Reset)
        typeWriterPrint("O algoritmo Round-Robin distribui o tempo de CPU entre os processos de forma igualitária, usando um quantum.")
    case "Prioridade":
        typeWriterPrint(Cyan + "\n[Escalonador por Prioridade]" + Reset)
        typeWriterPrint("O algoritmo de Prioridade executa os processos com base na prioridade, onde processos com menor valor de prioridade são executados primeiro.")
    }
}

// -------------------------------------------------------------------
// Função para exibir um quiz ao final da execução para reforçar o aprendizado
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
        fmt.Println("Iniciando quiz...")
    case 2:
        fmt.Println(Green, "Quiz pulado.", Reset)
    case 3:
        fmt.Println("Reiniciando execução...")
    default:
        fmt.Println(Red, "Opção inválida!", Reset)
    }
}

// -------------------------------------------------------------------
// Função para gerar processos aleatórios com tempos de burst e prioridades variadas
func generateProcesses() []Process {
    return []Process{
        {ID: 1, BurstTime: rand.Intn(500) + 500, Priority: rand.Intn(10), State: Ready},
        {ID: 2, BurstTime: rand.Intn(500) + 500, Priority: rand.Intn(10), State: Ready},
    }
}

// -------------------------------------------------------------------
// Função auxiliar para encontrar o mínimo entre dois valores
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// -------------------------------------------------------------------
// Menu interativo principal para escolha do algoritmo de escalonamento
func main() {
    rand.Seed(time.Now().UnixNano())
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
            fmt.Println("Executando FIFO...")
        case 2:
            fmt.Println("Executando Round Robin...")
        case 3:
            fmt.Println("Executando Prioridade...")
        case 4:
            fmt.Println(Green + "Saindo..." + Reset)
            return
        default:
            fmt.Println(Red + "Opção inválida!" + Reset)
        }
    }
}
