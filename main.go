package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type State struct {
	name string
}

type Transition struct {
	from  *State
	to    *State
	input rune
}

type Automata struct {
	states       map[string]*State
	transitions  []*Transition
	currentState *State
	finalStates  map[string]*State
}

func createStateMap(stateNames []string) map[string]*State {
	stateMap := make(map[string]*State)
	for _, name := range stateNames {
		stateMap[name] = &State{name: name}
	}
	return stateMap
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingresa los estados separados por comas (ejemplo: q0,q1): ")
	statesInput, _ := reader.ReadString('\n')
	statesInput = strings.TrimSpace(statesInput)

	stateNames := strings.Split(statesInput, ",")
	stateMap := createStateMap(stateNames)

	fmt.Print("Ingresa los símbolos de entrada separados por comas (ejemplo: 0,1): ")
	symbolsInput, _ := reader.ReadString('\n')
	symbolsInput = strings.TrimSpace(symbolsInput)

	symbols := strings.Split(symbolsInput, ",")

	fmt.Print("Ingresa el estado inicial: ")
	initialState, _ := reader.ReadString('\n')
	initialState = strings.TrimSpace(initialState)

	fmt.Print("Ingresa los estados finales separados por comas (ejemplo: q0): ")
	finalStatesInput, _ := reader.ReadString('\n')
	finalStatesInput = strings.TrimSpace(finalStatesInput)

	finalStateNames := strings.Split(finalStatesInput, ",")
	finalStateMap := createStateMap(finalStateNames)

	fmt.Print("Ingresa el número de transiciones: ")
	var numTransitions int
	fmt.Scanln(&numTransitions)

	transitions := make([]*Transition, numTransitions)
	for i := 0; i < numTransitions; i++ {
		var from, input, to string
		fmt.Printf("Transición %d (from,input,to): ", i+1)
		transitionInput, _ := reader.ReadString('\n')
		transitionInput = strings.TrimSpace(transitionInput)

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

	automata := &Automata{
		states:       stateMap,
		transitions:  transitions,
		currentState: stateMap[initialState],
		finalStates:  finalStateMap,
	}

	fmt.Print("Ingresa una cadena de entrada: ")
	var input string
	fmt.Scanln(&input)

	for _, char := range input {
		symbolFound := false
		for _, symbol := range symbols {
			if string(char) == symbol {
				symbolFound = true
				break
			}
		}
		if !symbolFound {
			fmt.Printf("Entrada no válida: solo se aceptan los símbolos de entrada %v\n", symbols)
			os.Exit(1)
		}
	}

	automata.Process(input)

	if automata.IsAccepted() {
		fmt.Println("La cadena de entrada es aceptada.")
	} else {

		fmt.Println("La cadena de entrada no es aceptada.")
	}
}

func (a *Automata) Process(input string) {
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

func (a *Automata) IsAccepted() bool {
	_, isFinal := a.finalStates[a.currentState.name]
	return isFinal
}
