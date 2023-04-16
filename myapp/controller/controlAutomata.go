package controller

import (
	"log"
	"myapp/model"
	"strings"
)

// ControlAutomatas es una estructura que representa un controlador de autómatas finitos.
type ControlAutomatas struct{}

// ReadFinalStates lee una cadena de estados finales separados por comas y devuelve una lista de strings sin espacios en blanco
func (c *ControlAutomatas) ReadFinalStates(finalStates string) []*model.State {
	stateFinalList := strings.Split(finalStates, ",")
	for i, state := range stateFinalList {
		stateFinalList[i] = strings.TrimSpace(state)
	}
	stateFinalSlice := c.CreateStateSlice(stateFinalList)
	return stateFinalSlice
}

// ReadInitialState lee un estado inicial y lo devuelve como un string sin espacios en blanco
func (c *ControlAutomatas) ReadInitialState(initialState string) *model.State {
	state := model.NewState(strings.TrimSpace(initialState))
	return state
}

// ReadSymbols lee una cadena de símbolos de entrada separados por comas y devuelve una lista de strings sin espacios en blanco
func (c *ControlAutomatas) ReadSymbols(symbols string) []string {
	symbolsList := strings.Split(symbols, ",")
	for i, symbol := range symbolsList {
		symbolsList[i] = strings.TrimSpace(symbol)
	}
	return symbolsList
}

// ReadState lee una cadena de estados separados por comas y devuelve un slice de states
func (c *ControlAutomatas) ReadState(state string) []*model.State {
	stateList := strings.Split(state, ",")
	for i, s := range stateList {
		stateList[i] = strings.TrimSpace(s)
	}
	stateSlice := c.CreateStateSlice(stateList)

	return stateSlice
}

// readTransitions recibe una cadena de transiciones en formato "from,input,to" y un sllice de estados, y devuelve una lista de objetos Transition.
func (c *ControlAutomatas) ReadTransitions(transitionString string, stateSlice []*model.State) []*model.Transition {
	transitions := make([]*model.Transition, 0)

	// Divide la cadena de entrada por corchetes para obtener una lista de transiciones
	transitionList := strings.Split(strings.Trim(transitionString, "[]"), "],[")

	for _, transition := range transitionList {
		// Ignora las transiciones vacías
		if transition == "" {
			continue
		}

		// Divide la transición por comas para obtener los campos "from", "input" y "to"
		parts := strings.Split(strings.TrimSpace(transition), ",")
		if len(parts) != 3 {
			log.Fatalf("Entrada de transición incorrecta. Debe ser 'from,input,to'.")
		}

		from, input, to := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), strings.TrimSpace(parts[2])

		var fromState *model.State
		for i, s := range stateSlice {
			if strings.EqualFold(s.GetName(), from){
				fromState = stateSlice[i]
			}
		}
		if fromState.GetName() == "" {
			log.Fatalf("El estado '%s' no existe en el autómata", from)
		}

		var toState *model.State
		for i, s := range stateSlice {
			if strings.EqualFold(s.GetName(), to){
				toState = stateSlice[i]
			}
		}
		if toState.GetName() == "" {
			log.Fatalf("El estado '%s' no existe en el autómata", to)
		}

		transitions = append(transitions, model.NewTransition(fromState, toState, input))
	}

	return transitions
}

// createStateSlice crea un slice de estados a partir de una lista de nombres de estados.
// Recibe: stateNames - una lista de nombres de estados.
// Retorna: Un slice de estados.
func (c *ControlAutomatas) CreateStateSlice(stateNames []string) []*model.State {
	stateSlice := make([]*model.State, len(stateNames))
	for i, name := range stateNames {
		stateSlice[i] = model.NewState(name)
	}
	return stateSlice
}

// CreateAutomata crea un autómata finito a partir de una lista de transiciones, un estado inicial,
// una lista de estados finales, una lista de estados y una lista de símbolos.
// Recibe: transitions - una lista de transiciones.
//
//	initialState - el estado inicial del autómata.
//	finalStatesSlice - un slice de estados finales.
//	stateSlice - un slice de estados.
//	symbols - una lista de símbolos.
//
// Retorna: Un autómata finito.
func (c *ControlAutomatas) CreateAutomata(transitions []*model.Transition, initialState *model.State,
	finalStatesSlice []*model.State, stateSlice []*model.State, symbols []string) *model.Automata {

	automata := model.NewAutomata(stateSlice, transitions, initialState, finalStatesSlice, symbols)

	return automata
}
