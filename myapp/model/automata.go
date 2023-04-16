package model

import (
	"fmt"
	// "os"
	"strings"
)

// Automata es la estructura que representa un autómata finito.
type Automata struct {
	states              []*State
	transitions         []*Transition
	initialState        *State
	currentState        *State
	finalStates         []*State
	alphabet            []string
	historyCurrentState []*State
}

// NewAutomata crea un nuevo objeto Automata con los estados, transiciones, estado inicial,
// estados finales y símbolos de entrada especificados.
func NewAutomata(states []*State, transitions []*Transition, initialState *State, finalStates []*State, alphabet []string) *Automata {
	return &Automata{
		states:              states,
		transitions:         transitions,
		initialState:        initialState,
		currentState:        initialState,
		finalStates:         finalStates,
		alphabet:            alphabet,
		historyCurrentState: []*State{initialState},
	}
}

// GetStates retorna el slice de estados del autómata.
func (a *Automata) GetStates() []*State {
	return a.states
}

// SetStates establece el slice de estados del autómata.
func (a *Automata) SetStates(states []*State) {
	a.states = states
}

// GetTransitions retorna la lista de transiciones del autómata.
func (a *Automata) GetTransitions() []*Transition {
	return a.transitions
}

// SetTransitions establece la lista de transiciones del autómata.
func (a *Automata) SetTransitions(transitions []*Transition) {
	a.transitions = transitions
}

// GetInitialState retorna el estado inicial del autómata.
func (a *Automata) GetInitialState() *State {
	return a.initialState
}

// SetInitialState establece el estado inicial del autómata.
func (a *Automata) SetInitialState(initialState *State) {
	a.initialState = initialState
}

// GetCurrentState retorna el estado actual del autómata.
func (a *Automata) GetCurrentState() *State {
	return a.currentState
}

// SetCurrentState establece el estado actual del autómata.
func (a *Automata) SetCurrentState(currentState *State) {
	a.currentState = currentState
}

// GetFinalStates retorna el mapa de estados finales del autómata.
func (a *Automata) GetFinalStates() []*State {
	return a.finalStates
}

// SetFinalStates establece el mapa de estados finales del autómata.
func (a *Automata) SetFinalStates(finalStates []*State) {
	a.finalStates = finalStates
}

// GetAlphabet retorna la lista de símbolos de entrada del autómata.
func (a *Automata) GetAlphabet() []string {
	return a.alphabet
}

// SetAlphabet establece la lista de símbolos de entrada del autómata.
func (a *Automata) SetAlphabet(alphabet []string) {
	a.alphabet = alphabet
}

func (a *Automata) GetHistoryCurrentState() []*State {
	return a.historyCurrentState
}

func (a *Automata) SetHistoryCurrentState(historyCurrentState []*State) {
	a.historyCurrentState = historyCurrentState
}

/*
El método toString() retorna una representación en cadena del objeto Automata.
Esta representación incluye la lista de estados, la lista de transiciones, el estado inicial,
los estados finales y el alfabeto.
*/
func (a *Automata) ToString() string {
	var builder strings.Builder

	builder.WriteString("Estados:\n")
	for _, state := range a.states {
		builder.WriteString(fmt.Sprintf("\t%s\n", state))
	}

	builder.WriteString("Transiciones:\n")
	for _, transition := range a.transitions {
		builder.WriteString(fmt.Sprintf("\t%s --(%s)--> %s\n", transition.GetFromState().GetName(), transition.GetInput(), transition.GetToState().GetName()))
	}

	builder.WriteString(fmt.Sprintf("Estado inicial: %s\n", a.initialState.GetName()))

	builder.WriteString("Estados finales:\n")
	for _, name := range a.finalStates {
		builder.WriteString(fmt.Sprintf("\t%s\n", name))
	}

	builder.WriteString("Alfabeto:\n")
	for _, symbol := range a.alphabet {
		builder.WriteString(fmt.Sprintf("\t%s\n", symbol))
	}

	return builder.String()
}

// ProcessInput procesa la entrada y actualiza el estado actual del autómata.
// También agrega el estado actual al historial de estados actuales.
// Recibe: input - una cadena de entrada que se procesará.
// No retorna ningún valor.
func (a *Automata) ProcessInput(input string) bool {
	a.currentState = a.initialState
	for _, char := range input {
		transitionFound := false
		inAlphabet := false
		for _, symbol := range a.alphabet {
			if symbol == string(char) {
				inAlphabet = true
				break
			}
		}
		for _, transition := range a.transitions {
			if transition.from.GetName() == a.currentState.GetName() && transition.input == string(char) /*string(rune(char))*/ {
				a.currentState = transition.to
				transitionFound = true
				a.historyCurrentState = append(a.historyCurrentState, a.currentState)
				break
			}
		}
		if !transitionFound || !inAlphabet {
			fmt.Println("Entrada no válida:", string(char))
			return false
			// os.Exit(1)
		}
	}
	return true
}

// IsAccepted determina si el autómata acepta la entrada basándose en su estado actual.
// No recibe ningún parámetro.
// Retorna: Un booleano que indica si la entrada es aceptada por el autómata.
func (a *Automata) IsAccepted() bool {
	// comprobamos si el estado es final
	isFinal := false
	for _, estFinal := range a.finalStates {
		if strings.EqualFold(estFinal.name, a.currentState.name) {
			isFinal = true
			break
		}
	}
	return isFinal
}
