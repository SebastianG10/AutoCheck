package controller

import (
	"log"
	"myapp/model"
	"strings"
)

// ControlAutomatas es una estructura que representa un controlador de autómatas finitos.
type ControlAutomatas struct{}

// ReadFinalStates lee una cadena de estados finales separados por comas y devuelve una lista de strings sin espacios en blanco
func (c *ControlAutomatas) ReadFinalStates(finalStates string) map[string]*model.State {
	stateFinalList := strings.Split(finalStates, ",")
	for i, state := range stateFinalList {
		stateFinalList[i] = strings.TrimSpace(state)
	}
	stateFinalMap := c.CreateStateMap(stateFinalList)
	return stateFinalMap
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

// ReadState lee una cadena de estados separados por comas y devuelve un mapa de states
func (c *ControlAutomatas) ReadState(state string) map[string]*model.State {
	stateList := strings.Split(state, ",")
	for i, s := range stateList {
		stateList[i] = strings.TrimSpace(s)
	}
	stateMap := c.CreateStateMap(stateList)

	return stateMap
}

// readTransitions recibe una cadena de transiciones en formato "from,input,to" y un mapa de estados, y devuelve una lista de objetos Transition.
func (c *ControlAutomatas) ReadTransitions(transitionString string, stateMap map[string]*model.State) []*model.Transition {
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

		fromState, ok := stateMap[from]
		if !ok {
			log.Fatalf("El estado '%s' no existe en el autómata", from)
		}

		toState, ok := stateMap[to]
		if !ok {
			log.Fatalf("El estado '%s' no existe en el autómata", to)
		}

		transitions = append(transitions, model.NewTransition(fromState, toState, input))
	}

	return transitions
}

// createStateMap crea un mapa de estados a partir de una lista de nombres de estados.
// Recibe: stateNames - una lista de nombres de estados.
// Retorna: Un mapa de estados.
func (c *ControlAutomatas) CreateStateMap(stateNames []string) map[string]*model.State {
	stateMap := make(map[string]*model.State)
	for _, name := range stateNames {
		stateMap[name] = model.NewState(name)
	}
	return stateMap
}

// CreateAutomata crea un autómata finito a partir de una lista de transiciones, un estado inicial,
// una lista de estados finales, una lista de estados y una lista de símbolos.
// Recibe: transitions - una lista de transiciones.
//
//	initialState - el estado inicial del autómata.
//	finalStatesMap - un mapa de estados finales.
//	stateMap - un mapa de estados.
//	symbols - una lista de símbolos.
//
// Retorna: Un autómata finito.
func (c *ControlAutomatas) CreateAutomata(transitions []*model.Transition, initialState *model.State,
	finalStatesMap map[string]*model.State, stateMap map[string]*model.State, symbols []string) *model.Automata {

	automata := model.NewAutomata(stateMap, transitions, initialState, finalStatesMap, symbols)

	return automata
}
