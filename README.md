
# **Projeto de Sistema Operacional Simulado em Go**

## Índice
1. [Introdução](#introdução)
2. [Objetivo](#objetivo)
3. [Instalação](#instalação)
4. [Exemplo de Uso](#exemplo-de-uso)
5. [Documentação do Código](#documentação-do-código)
   1. [Funcionamento](#funcionamento)
   2. [Partes Importantes do Código](#partes-importantes-do-código)
   3. [Estrutura de Dados](#estrutura-de-dados)
   4. [Algoritmos de Escalonamento Implementados](#algoritmos-de-escalonamento-implementados)
6. [Conclusão](#conclusão)

## Introdução
Este projeto simula o funcionamento básico de um sistema operacional, implementado em Go, com foco no gerenciamento de processos, escalonamento de CPU e simulação de diferentes algoritmos de escalonamento. O objetivo principal é fornecer uma visão geral sobre como os sistemas operacionais gerenciam os recursos do sistema, especificamente os processos.

## Objetivo
O projeto tem como propósito demonstrar os principais conceitos de um sistema operacional, como o escalonamento de processos e a interação entre processos e CPU. Utilizando Go, foram implementados três algoritmos de escalonamento: **FIFO**, **Round-Robin** e **Prioridade**, todos com uma interface interativa para visualização e compreensão de seu funcionamento.

## Instalação
Para instalar e executar este projeto, siga os passos abaixo:
-
1. **Clone o repositório**:
   ```sh
   git clone https://github.com/seu-usuario/projeto-de-so.git
   ```

2. **Acesse o diretório do projeto**:
   ```sh
   cd projeto-de-so
   ```

3. **Compile o projeto**:
   ```sh
   go build
   ```

4. **Execute o projeto**:
   ```sh
   ./projeto-de-so
   ```

## Exemplo de Uso
Após a instalação, você pode executar o programa com o comando abaixo:

```sh
./projeto-de-so --help
```

Isso irá exibir as opções e o menu de escalonamento do sistema operacional simulado.

## Documentação do Código

### Funcionamento
O sistema operacional simulado implementa três algoritmos de escalonamento de processos:

- **FIFO (First-Come, First-Served)**: O primeiro processo a chegar é o primeiro a ser executado.
- **Round-Robin**: Cada processo recebe uma fatia de tempo (quantum) para ser executado. Se não terminar dentro do tempo, é interrompido e colocado no final da fila.
- **Prioridade**: Os processos são escalonados de acordo com sua prioridade, com processos de maior prioridade sendo executados primeiro.

### Partes Importantes do Código

1. **`main.go`**: O ponto de entrada do sistema operacional, onde a execução dos algoritmos é gerenciada.
    ```go
    // main.go
    package main

    func main() {
        // Inicializa o sistema operacional
        initOS()
        // Executa o loop principal do sistema operacional
        runOS()
    }
    ```

2. **`process.go`**: Responsável pelo gerenciamento de processos, incluindo a criação, escalonamento e terminação dos mesmos.
    ```go
    // process.go
    package main

    type Process struct {
        PID     int
        BurstTime int
        Priority int
        State   ProcessState
    }

    func createProcess() *Process {
        // Lógica para criar um novo processo
    }

    func scheduleProcess() {
        // Lógica para escalonar processos
    }

    func terminateProcess(pid int) {
        // Lógica para terminar um processo
    }
    ```

3. **`memory.go`**: Gerenciamento da memória, incluindo alocação e desalocação de memória para processos.
    ```go
    // memory.go
    package main

    func allocateMemory(size int) {
        // Lógica para alocar memória
    }

    func freeMemory(address int) {
        // Lógica para desalocar memória
    }
    ```

4. **`filesystem.go`**: Gerenciamento do sistema de arquivos, com funções para criar, ler, escrever e deletar arquivos.
    ```go
    // filesystem.go
    package main

    func createFile(name string) {
        // Lógica para criar um arquivo
    }

    func readFile(name string) {
        // Lógica para ler um arquivo
    }

    func writeFile(name string, data []byte) {
        // Lógica para escrever em um arquivo
    }

    func deleteFile(name string) {
        // Lógica para deletar um arquivo
    }
    ```

### Estrutura de Dados

A estrutura de dados principal utilizada no projeto é a **`Process`**, que armazena informações sobre os processos, como o identificador (PID), o tempo de execução (burst time), a prioridade e o estado do processo.

```go
type Process struct {
    ID        int
    BurstTime int
    Priority  int
    State     ProcessState
}
```

O estado de cada processo é gerenciado por uma constante `ProcessState`, com os seguintes valores:

- `Ready` (Pronto)
- `Running` (Executando)
- `Paused` (Pausado)
- `Completed` (Finalizado)
- `Error` (Erro)

### Algoritmos de Escalonamento Implementados

#### FIFO (First-Come, First-Served)
O algoritmo FIFO executa os processos na ordem em que chegam. O primeiro processo a ser iniciado será o primeiro a ser completado, o que pode resultar em problemas de inanição quando processos curtos ficam bloqueados devido a processos longos.

#### Round-Robin
No algoritmo Round-Robin, cada processo recebe uma fatia de tempo (quantum). Se o processo não concluir sua execução dentro do tempo alocado, ele é pausado e colocado novamente no final da fila. Este método busca evitar a inanição, mas pode gerar overhead devido ao grande número de trocas de contexto.

#### Prioridade
O algoritmo de escalonamento por Prioridade executa os processos com base em sua prioridade. Processos com menor valor de prioridade são executados primeiro. Isso pode gerar problemas de inanição para processos de baixa prioridade se o sistema constantemente tiver processos de alta prioridade.

## Conclusão
Este projeto oferece uma simulação interativa dos principais algoritmos de escalonamento de processos em sistemas operacionais, permitindo que o usuário visualize, entenda e experimente cada um desses algoritmos de forma prática. O código é modular e extensível, proporcionando uma excelente base para o estudo de sistemas operacionais e a implementação de outros conceitos, como gerenciamento de memória e sistemas de arquivos.

