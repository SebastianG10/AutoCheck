package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myapp/controller"
	"myapp/model"
	"os"
	"os/exec"
	"strconv"
	"strings"

	// "time"

	// Imports para interfaz gráfica
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

// paleta de colores
var blue = color.NRGBA{R: 50, G: 119, B: 168, A: 0xff}

// var red = color.NRGBA{R: 242, G: 80, B: 80, A: 0xff}
// var gray = color.NRGBA{R: 170, G: 170, B: 170, A: 0xff}

// main es la función principal del programa.
func main() {

	interfaz()

	logicaAutomata()

	leerEntrada()

}

func interfaz() {
	myApp := app.New()
	welcomeWindow := myApp.NewWindow("Bienvenido")
	welcomeWindow.Resize(fyne.NewSize(900, 750))

	welcomeWindow.SetContent(welcomeContent(myApp))
	welcomeWindow.Show()

	myApp.Run()
}

func logicContent() *fyne.Container {
	instrucciones := canvas.NewText("Ingrese la quintupla del autómata siguiendo las indicaciones que se den", blue)
	instrucciones.TextStyle = fyne.TextStyle{Bold: true}

	//ingreso de estados
	estadosLabel := canvas.NewText("Estados", color.White)
	estadosLabel.TextStyle = fyne.TextStyle{Bold: true}
	statesInstruc := widget.NewLabel("Ingrese los estados de su autómata separandolos con comas\nEjemplo: q0,q1,q2...")
	statesInput := widget.NewEntry()
	statesCont := container.New(layout.NewVBoxLayout(), estadosLabel, statesInstruc, statesInput)

	//ingreso de alfabeto
	alfabetoLabel := canvas.NewText("Alfabeto", color.White)
	alfabetoLabel.TextStyle = fyne.TextStyle{Bold: true}
	symbolsInstruc := widget.NewLabel("Ingrese los simbolos del alfabeto separandolos con comas\nEjemplo: 0,1...")
	symbolsInput := widget.NewEntry()
	symbolsCont := container.New(layout.NewVBoxLayout(), alfabetoLabel, symbolsInstruc, symbolsInput)

	//ingreso de transiciones
	transicionesLabel := canvas.NewText("Transiciones", color.White)
	transicionesLabel.TextStyle = fyne.TextStyle{Bold: true}
	transitionsInstruc := widget.NewLabel("Ingrese las transiciones del automata con el sigiente formato: [from,input,to]\nEjemplo: [q0,1,q0],[q1,0,q1],...")
	transitionsInput := widget.NewEntry()
	transitionsCont := container.New(layout.NewVBoxLayout(), transicionesLabel, transitionsInstruc, transitionsInput)

	//ingreso del estado inicial
	inicialLabel := canvas.NewText("Estado Inicial", color.White)
	inicialLabel.TextStyle = fyne.TextStyle{Bold: true}
	initialInstruc := widget.NewLabel("Ingrese el estado inicial de su autómata\nEjemplo: q0")
	initialInput := widget.NewEntry()
	initialCont := container.New(layout.NewVBoxLayout(), inicialLabel, initialInstruc, initialInput)

	//ingreso de estados finales
	finalesLabel := canvas.NewText("Estados Finales", color.White)
	finalesLabel.TextStyle = fyne.TextStyle{Bold: true}
	finalInstruc := widget.NewLabel("Ingrese el estado final de su autómata:\nEjemplo: q1")
	finalInput := widget.NewEntry()
	finalCont := container.New(layout.NewVBoxLayout(), finalesLabel, finalInstruc, finalInput)
	//botón para construir autómata
	automataContainer := container.NewCenter()

	//creamos un automata vacio de manera global para luego porder acceder a él
	control := &controller.ControlAutomatas{}
	dstateMap := control.ReadState("")
	dsymbols := control.ReadSymbols("")
	dtransitionsList := control.ReadTransitions("", dstateMap)
	dinitialState := control.ReadInitialState("")
	dfinalStatesMap := control.ReadFinalStates("")
	automata := control.CreateAutomata(dtransitionsList, dinitialState, dfinalStatesMap, dstateMap, dsymbols)
	fmt.Println("Automata por defecto\n" + automata.ToString())

	// botón para construir el automata
	constuirAutomata := widget.NewButtonWithIcon("Construir Autómata", theme.ConfirmIcon(), func() {
		// pasamos los datos ingresados al automata
		stateMap := control.ReadState(statesInput.Text)
		symbols := control.ReadSymbols(symbolsInput.Text)
		transitionsList := control.ReadTransitions(transitionsInput.Text, stateMap)
		initialState := control.ReadInitialState(initialInput.Text)
		finalStatesMap := control.ReadFinalStates(finalInput.Text)
		// se actualiza el automata global
		automata = control.CreateAutomata(transitionsList, initialState, finalStatesMap, stateMap, symbols)
		// añadimos la imagen graphviz retornada por renderizarAutomata al automataContainer
		fmt.Println("Automata construido\n" + automata.ToString())
		automataContainer.Add(renderizarAutomata(automata))
	})

	//container izquierdo donde se ingresa la quintupla del autómata
	cargarAutomatacont := container.New(layout.NewVBoxLayout(),
		instrucciones,
		layout.NewSpacer(),
		statesCont,
		layout.NewSpacer(),
		symbolsCont,
		layout.NewSpacer(),
		transitionsCont,
		layout.NewSpacer(),
		initialCont,
		layout.NewSpacer(),
		finalCont,
		layout.NewSpacer(),
		constuirAutomata,
	)

	// elementos del container derecho de entradas
	entradasTitulo := canvas.NewText("Entradas                                                                                 ", blue)
	entradasTitulo.TextStyle = fyne.TextStyle{Bold: true}

	entradasInstrucciones := widget.NewLabel("En este apartado puede seleccionar alguna entrada para probar su automata")
	entradasInstrucciones.Wrapping = fyne.TextWrapBreak

	entradasDisp := widget.NewLabel("Entradas disponibles:")

	// Leer el archivo JSON con las entradas
	file, _ := ioutil.ReadFile("../resources/entradas.json")

	// Decodificar el archivo JSON
	var entradas []map[string]string
	json.Unmarshal([]byte(file), &entradas)

	// Imprimir las opciones del menú
	listEntrStr := ""
	for i := range entradas {
		entradaStr := entradas[i]["Entrada"]
		listEntrStr += fmt.Sprintf("%d. %s\n", i+1, entradaStr)
	}

	listEntrLabel := widget.NewLabel(listEntrStr)

	// Seleccionar una entrada
	opcionInstruccion := widget.NewLabel("Seleccione una entrada: ")
	ingresoOpcion := widget.NewEntry()

	respuestaCont := container.New(layout.NewMaxLayout())

	probarEntradaBoton := widget.NewButtonWithIcon("Probar Entrada", theme.NavigateNextIcon(), func() {
		fmt.Println("Automata xd\n" + automata.ToString())

		opcion, _ := strconv.ParseInt(ingresoOpcion.Text, 0, 64)
		var input string = entradas[opcion-1]["Entrada"]
		// Imprimir la entrada seleccionada
		fmt.Printf("Entrada seleccionada: %s\n", input)

		entradaAceptada := automata.ProcessInput(input)
		fmt.Printf("Después de procesar la entrada '%s', el estado actual es: %s\n", input, automata.GetCurrentState().GetName())

		fmt.Println("Historial de estados actuales:")
		for _, state := range automata.GetHistoryCurrentState() {
			fmt.Printf("%s -> ", state.GetName())
		}
		// se limpria el historial de estados para el ingreso de nuevas entradas
		automata.SetHistoryCurrentState([]*model.State{})

		icono := theme.InfoIcon()
		var msg string
		// Prueba para IsAccepted
		if automata.IsAccepted() && entradaAceptada {
			msg = "Entrada Aceptada"
			icono = theme.ConfirmIcon()
			fmt.Println("La entrada es aceptada por el autómata.")
		} else {
			msg = "Entrada Rechazada"
			icono = theme.CancelIcon()
			fmt.Println("La entrada no es aceptada por el autómata.")
		}
		
		botonRespuesta := widget.NewButtonWithIcon(msg, icono, func() {})
		respuestaCont.Add(botonRespuesta)
	})

	//container derecho para las entradas del automata
	entradasCont := container.New(layout.NewVBoxLayout(),
		entradasTitulo,
		entradasInstrucciones,
		entradasDisp,
		listEntrLabel,
		opcionInstruccion,
		ingresoOpcion,
		probarEntradaBoton,
		layout.NewSpacer(),
		respuestaCont,
		layout.NewSpacer(),
	)

	// container horizontal principal
	content := container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		cargarAutomatacont,
		widget.NewSeparator(),
		automataContainer,
		widget.NewSeparator(),
		entradasCont,
		layout.NewSpacer(),
	)
	return content
}

func leerEntrada() {
	/*
		// Leer el archivo JSON
		file, _ := ioutil.ReadFile("../resources/entradas.json")

		// Decodificar el archivo JSON
		var entradas []map[string]string
		json.Unmarshal([]byte(file), &entradas)

		// Imprimir las opciones del menú
		for i, entrada := range entradas {
			fmt.Printf("%d. %s\n", i+1, entrada)
		}

		// Seleccionar una entrada
		var opcion int
		fmt.Print("Seleccione una entrada: ")
		fmt.Scan(&opcion)

		var cadena string = entradas[opcion-1]["Entrada"]
		// Imprimir la entrada seleccionada
		fmt.Printf("Entrada seleccionada: %s\n", cadena)
	*/
}

func welcomeContent(app fyne.App) *fyne.Container {
	// creamos los elementos que se mostrarán en la ventana de bienvenida
	infinite := widget.NewProgressBarInfinite()

	welcomeLabel := canvas.NewText("¡AUTÓMATAS!", blue)
	welcomeLabel.TextStyle = fyne.TextStyle{Bold: true}
	welcome := container.NewCenter(welcomeLabel)

	image := canvas.NewImageFromFile("../resources/welcomeimage.png")
	image.SetMinSize(fyne.NewSize(458, 210))
	imageCont := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), image, layout.NewSpacer())

	descriptionLabel := widget.NewLabelWithStyle("Bienvenido al software donde podrá probar entradas para algunos autómatas básicos.\n", fyne.TextAlignCenter, fyne.TextStyle{})
	instructionsLabel := widget.NewLabel("Pulse el botón en cualquier momento para empezar!")

	// la función del parámetro del botón abre la ventana donde se desarrolla la lógica de la aplicaciónx
	startButton := widget.NewButtonWithIcon("Empezar", theme.MediaPlayIcon(), func() {
		var logicWindow = app.NewWindow("Autómatas")
		logicWindow.Resize(fyne.NewSize(1450, 750))
		logicWindow.SetContent(logicContent())
		logicWindow.Show()
	})

	// elementos de la zona inferior
	developed := widget.NewLabel("Desarrollado por:")
	//datos para el card de Garcias
	gUrl, _ := url.Parse("https://github.com/SebastianG10")
	gLink := widget.NewHyperlink("https://github.com/SebastianG10", gUrl)
	gimage := canvas.NewImageFromFile("../resources/garcias.png")
	gcard := widget.NewCard("Sebastían", "García", gLink)
	gcard.SetImage(gimage)
	//datos para el card de Seider
	sUrl, _ := url.Parse("https://github.com/vanebrina")
	sLink := widget.NewHyperlink("https://github.com/vanebrina", sUrl)
	sImage := canvas.NewImageFromFile("../resources/seider.png")
	sCard := widget.NewCard("Seider", "Sanín", sLink)
	sCard.SetImage(sImage)
	//datos para el card de Jhon
	jUrl, _ := url.Parse("https://github.com/ElJhones18")
	jLink := widget.NewHyperlink("https://github.com/ElJhones18", jUrl)
	jImage := canvas.NewImageFromFile("../resources/jhon.png")
	jCard := widget.NewCard("Jhon Esteban", "Rodríguez", jLink)
	jCard.SetImage(jImage)
	// metemos los cards en un container horizontal que a su vez añadiremos al contenido principal que es un container vertical
	cards := container.NewHBox(layout.NewSpacer(), gcard, sCard, jCard, layout.NewSpacer())
	// container vertical para el contenido principal de la ventana de bienvenida
	content := container.New(layout.NewVBoxLayout(),
		layout.NewSpacer(),
		infinite,
		welcome,
		imageCont,
		descriptionLabel,
		instructionsLabel,
		startButton,
		developed,
		cards,
		layout.NewSpacer(),
	)
	return content
}

func logicaAutomata() {
	/* control := &controller.ControlAutomatas{}

		stateMap := control.ReadState("q0,q1")
		symbols := control.ReadSymbols("0,1")
		transitionsList := control.ReadTransitions("[q0,0,q1],[q0,1,q0],[q1,0,q1],[q1,1,q0]", stateMap)
		initialState := control.ReadInitialState("q0")
		finalStatesMap := control.ReadFinalStates("q1")

		automata := control.CreateAutomata(transitionsList, initialState, finalStatesMap, stateMap, symbols)

	    // Prueba para ProcessInput
	    input := "1010"
	    automata.ProcessInput(input)
	    fmt.Printf("Después de procesar la entrada '%s', el estado actual es: %s\n", input, automata.GetCurrentState().GetName())

		fmt.Println("Historial de estados actuales:")
		for _, state := range automata.GetHistoryCurrentState() {
			fmt.Printf("%s -> ", state.GetName())
		}

		// Prueba para IsAccepted
		if automata.IsAccepted() {
			fmt.Println("La entrada es aceptada por el autómata.")
		} else {
			fmt.Println("La entrada no es aceptada por el autómata.")
		} */
}

func renderizarAutomata(automata *model.Automata) *canvas.Image {
	stateHash := func(s string) string {
		return s
	}
	// creamos el grafo que leerá graphviz
	g := graph.New(stateHash, graph.Directed())

	// for para añadir los estados a la interfaz
	for _, state := range automata.GetStates() {
		// comprobamos si el estdo es inicial
		esInicial := false
		if state.GetName() == automata.GetInitialState().GetName() {
			esInicial = true
		}

		// comprobamos si el estado es final
		esFinal := false
		for _, estFinal := range automata.GetFinalStates() {
			if strings.EqualFold(estFinal.GetName(), state.GetName()) {
				esFinal = true
				break
			}
		}

		//si el estado es inicial lo creamos con color azul, si es final como circulo doble y si es final con ambas características
		if esInicial && !esFinal {
			_ = g.AddVertex(state.GetName(), graph.VertexAttribute("colorscheme", "blues3"), graph.VertexAttribute("style", "filled"), graph.VertexAttribute("color", "2"), graph.VertexAttribute("fillcolor", "1"))
		} else if esFinal && !esInicial {
			_ = g.AddVertex(state.GetName(), graph.VertexAttribute("shape", "doublecircle"))
		} else if esInicial && esFinal {
			_ = g.AddVertex(state.GetName(), graph.VertexAttribute("colorscheme", "blues3"), graph.VertexAttribute("style", "filled"), graph.VertexAttribute("color", "2"), graph.VertexAttribute("fillcolor", "1"), graph.VertexAttribute("shape", "doublecircle"))
		} else {
			_ = g.AddVertex(state.GetName())
		}
	}

	//for para añadir las transiciones a la interfaz
	for _, t1 := range automata.GetTransitions() {
		inputs := t1.GetInput()
		for _, t2 := range automata.GetTransitions() {
			if t1.GetFromState().GetName() == t2.GetFromState().GetName() && t1.GetInput() != t2.GetInput() && t1.GetToState().GetName() == t2.GetToState().GetName() {
				inputs += ","
				inputs += t2.GetInput()
			}

		}
		_ = g.AddEdge(t1.GetFromState().GetName(), t1.GetToState().GetName(), graph.EdgeAttribute("label", inputs))
	}

	//escribimos el archivo que leerá graphviz
	file, _ := os.Create("my-graph.gv")
	_ = draw.DOT(g, file)

	// corremos el comando que genera la imagen del grafo
	out, err := exec.Command("dot", "-Tpng", "-O", "my-graph.gv").Output()
	println(&out, err)
	// obtenemos la imagen generada y le damos un tamaño minimo
	imagen := canvas.NewImageFromFile("my-graph.gv.png")
	imagen.FillMode = canvas.ImageFillContain
	imagen.SetMinSize(fyne.NewSize(600, 600))
	// retornamos la imagen
	return imagen
}
