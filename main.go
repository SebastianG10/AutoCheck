package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// State representa un estado en el autómata.
type State struct {
	name string
}

// Transition representa una transición en el autómata.
type Transition struct {
	from  *State
	to    *State
	input rune
}

// Automata es la estructura que representa un autómata finito.
type Automata struct {
	states       map[string]*State
	transitions  []*Transition
	initialState *State
	currentState *State
	finalStates  map[string]*State
	alphabet     []string
}

// createStateMap crea un mapa de estados a partir de una lista de nombres de estados.
// Recibe: stateNames - una lista de nombres de estados.
// Retorna: Un mapa de estados.
func createStateMap(stateNames []string) map[string]*State {
	stateMap := make(map[string]*State)
	for _, name := range stateNames {
		stateMap[name] = &State{name: name}
	}
	return stateMap
}

// readUserInput lee una línea de entrada del usuario y elimina los espacios en blanco alrededor.
// Recibe: reader - un bufio.Reader para leer la entrada del usuario.
// Retorna: La entrada del usuario sin espacios en blanco y un error si ocurre.
func readUserInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

// readStates lee los estados a partir de la entrada del usuario.
// Recibe: reader - un bufio.Reader para leer la entrada del usuario.
// Retorna: Una lista de nombres de estados.
func readStates(reader *bufio.Reader) []string {
	fmt.Print("Ingresa los estados separados por comas (ejemplo: q0,q1): ")
	statesInput, _ := readUserInput(reader)
	return strings.Split(statesInput, ",")
}

// readSymbols lee los símbolos de entrada a partir de la entrada del usuario.
// Recibe: reader - un bufio.Reader para leer la entrada del usuario.
// Retorna: Una lista de símbolos de entrada.
func readSymbols(reader *bufio.Reader) []string {
	fmt.Print("Ingresa los símbolos de entrada separados por comas (ejemplo: 0,1): ")
	symbolsInput, _ := readUserInput(reader)
	return strings.Split(symbolsInput, ",")
}

// readInitialState lee el estado inicial a partir de la entrada del usuario.
// Recibe: reader - un bufio.Reader para leer la entrada del usuario.
// Retorna: El nombre del estado inicial.
func readInitialState(reader *bufio.Reader) string {
	fmt.Print("Ingresa el estado inicial: ")
	initialState, _ := readUserInput(reader)
	return initialState
}

// readFinalStates lee los estados finales a partir de la entrada del usuario.
// Recibe: reader - un bufio.Reader para leer la entrada del usuario.
// Retorna: Una lista de nombres de estados finales.
func readFinalStates(reader *bufio.Reader) []string {
	fmt.Print("Ingresa los estados finales separados por comas (ejemplo: q0): ")
	finalStatesInput, _ := readUserInput(reader)
	return strings.Split(finalStatesInput, ",")
}

// readTransitions lee las transiciones a partir de la entrada del usuario.
// Recibe: reader - un bufio.Reader para leer la entrada del usuario.
//
//	numTransitions - el número de transiciones.
//	stateMap - un mapa de estados.
//
// Retorna: Una lista de transiciones.
func readTransitions(reader *bufio.Reader, numTransitions int, stateMap map[string]*State) []*Transition {
	transitions := make([]*Transition, numTransitions)
	for i := 0; i < numTransitions; i++ {
		var from, input, to string
		fmt.Printf("Transición %d (from,input,to): ", i+1)
		transitionInput, _ := readUserInput(reader)

		parts := strings.Fields(transitionInput)
		if len(parts) != 3 {
			fmt.Println("Entrada de transición incorrecta. Debe ser 'from,input,to'.")
			os.Exit(1)
		}

		from, input, to = parts[0], parts[1], parts[2]
		transitions[i] = &Transition{
			from:  stateMap[from],
			to:    stateMap[to],
			input: rune(input[0]),
		}
	}
	return transitions
}

// createAutomata crea un autómata a partir de la entrada del usuario.
// Recibe: reader - un bufio.Reader para leer la entrada del usuario.
// Retorna: Un puntero a un objeto Automata.
func createAutomata(reader *bufio.Reader) *Automata {
	stateNames := readStates(reader)
	stateMap := createStateMap(stateNames)

	symbols := readSymbols(reader)

	initialStateName := readInitialState(reader)
	initialState := stateMap[initialStateName]

	finalStateNames := readFinalStates(reader)
	finalStateMap := createStateMap(finalStateNames)

	fmt.Print("Ingresa el número de transiciones: ")
	var numTransitions int
	fmt.Scanln(&numTransitions)

	transitions := readTransitions(reader, numTransitions, stateMap)

	return &Automata{
		states:       stateMap,
		transitions:  transitions,
		initialState: initialState,
		currentState: initialState,
		finalStates:  finalStateMap,
		alphabet:     symbols,
	}
}

// Process procesa la entrada y actualiza el estado actual del autómata.
// Recibe: input - una cadena de entrada que se procesará.
// No retorna ningún valor.
func (a *Automata) Process(input string) {
	a.currentState = a.initialState
	for _, char := range input {
		transitionFound := false
		for _, transition := range a.transitions {
			if transition.from == a.currentState && transition.input == char {
				a.currentState = transition.to
				transitionFound = true
				break
			}
		}
		if !transitionFound {
			fmt.Println("Entrada no válida:", string(char))
			os.Exit(1)
		}
	}
}

// IsAccepted determina si el autómata acepta la entrada basándose en su estado actual.
// No recibe ningún parámetro.
// Retorna: Un booleano que indica si la entrada es aceptada por el autómata.
func (a *Automata) IsAccepted() bool {
	_, isFinal := a.finalStates[a.currentState.name]
	return isFinal
}

// main es la función principal del programa.
func main() {
	reader := bufio.NewReader(os.Stdin)
	automata := createAutomata(reader)

	for {
		// Leer la cadena de entrada del usuario
		fmt.Print("Ingresa una cadena de entrada (o presiona 'q' para salir): ")
		input, _ := readUserInput(reader)

		if input == "q" {
			break
		}

		// Comprobar si la cadena de entrada contiene solo símbolos válidos
		for _, char := range input {
			symbolFound := false
			for _, symbol := range automata.alphabet {
				if string(char) == symbol {
					symbolFound = true
					break
				}
			}

			if !symbolFound {
				fmt.Printf("Entrada no válida: solo se aceptan los símbolos de entrada %v\n", automata.alphabet)
				os.Exit(1)
			}
		}

		// Procesar la cadena de entrada utilizando el autómata
		automata.Process(input)

		// Determinar si la entrada es aceptada o no
		if automata.IsAccepted() {
			fmt.Println("Aceptado.")
		} else {
			fmt.Println("Incorrecto.")
		}
	}
}
