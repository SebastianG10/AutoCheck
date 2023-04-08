package main

import (
	// "fmt"
	// "myapp/controller"

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
)

// paleta de colores
var blue = color.NRGBA{R: 50, G: 119, B: 168, A: 0xff}

// main es la función principal del programa.
func main() {

	interfaz()

	logicaAutomata()

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
	// control := &controller.ControlAutomatas{}

	// stateMap := control.ReadState("q0,q1")
	// symbols := control.ReadSymbols("0,1")
	// transitionsList := control.ReadTransitions("[q0,0,q1],[q0,1,q0],[q1,0,q1],[q1,1,q0]", stateMap)
	// initialState := control.ReadInitialState("q0")
	// finalStatesMap := control.ReadFinalStates("q1")

	// automata := control.CreateAutomata(transitionsList, initialState, finalStatesMap, stateMap, symbols)

	// // Prueba para ProcessInput
	// input := "1010"
	// automata.ProcessInput(input)
	// fmt.Printf("Después de procesar la entrada '%s', el estado actual es: %s\n", input, automata.GetCurrentState().GetName())

	// fmt.Println("Historial de estados actuales:")
	// for _, state := range automata.GetHistoryCurrentState() {
	// 	fmt.Printf("%s -> ", state.GetName())
	// }

	// // Prueba para IsAccepted
	// if automata.IsAccepted() {
	// 	fmt.Println("La entrada es aceptada por el autómata.")
	// } else {
	// 	fmt.Println("La entrada no es aceptada por el autómata.")
	// }

	//------------------------------------------------------------------------------------------------------------------------------
	instrucciones := canvas.NewText("Ingrese la quintupla del autómata siguiendo las indicaciones que se den", blue)

	//ingreso de estados
	statesInstruc := widget.NewLabel("Ingrese los estados de su autómata separandolos con comas\nEjemplo: q0,q1,q2...")
	statesInput := widget.NewEntry()
	statesCont := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), statesInstruc, statesInput)

	//ingreso de alfabeto
	symbolsInstruc := widget.NewLabel("Ingrese los simbolos del alfabeto separandolos con comas\nEjemplo: 0,1...")
	symbolsInput := widget.NewEntry()
	symbolsCont := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), symbolsInstruc, symbolsInput)

	//ingreso de transiciones
	transitionsInstruc := widget.NewLabel("Ingrese las transiciones del automata con el sigiente formato: [from,input,to]\nEjemplo: [q0,1,q0],[q1,0,q1],...")
	transitionsInput := widget.NewEntry()
	transitionsCont := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), transitionsInstruc, transitionsInput)

	//ingreso de alfabeto
	initialInstruc := widget.NewLabel("Ingrese el estado inicial de su autómata\nEjemplo: q0")
	initialInput := widget.NewEntry()
	initialCont := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), initialInstruc, initialInput)

	//ingreso de transiciones
	finalInstruc := widget.NewLabel("Ingrese las transiciones del automata con el sigiente formato: [from,input,to]\nEjemplo: [q0,1,q0],[q1,0,q1],...")
	finalInput := widget.NewEntry()
	finalCont := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), finalInstruc, finalInput)

	//botón para construir autómata
	constAutomata := widget.NewButtonWithIcon("Construir Autómata", theme.ConfirmIcon(), func() {

		// stateMap := control.ReadState(statesInput.SelectedText())
		// symbols := control.ReadSymbols(symbolsInput.SelectedText())
		// transitionsList := control.ReadTransitions(transitionsInput.SelectedText(), stateMap)
		// initialState := control.ReadInitialState(initialInput.SelectedText())
		// finalStatesMap := control.ReadFinalStates(finalInput.SelectedText())

		// automata := control.CreateAutomata(transitionsList, initialState, finalStatesMap, stateMap, symbols)

	})

	// se puede intentar añadir fondo con un max y un rectangulo
	cargarAutomatacont := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		instrucciones,
		statesCont,
		symbolsCont,
		transitionsCont,
		initialCont,
		finalCont,
		constAutomata,
	)

	content := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		cargarAutomatacont,
		// renderizar automata en la parte derecha de la ventana
		layout.NewSpacer(),
	)
	return content
}

func welcomeContent(app fyne.App) *fyne.Container {

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
		logicWindow.Resize(fyne.NewSize(1000, 750))
		logicWindow.SetContent(logicContent())
		logicWindow.Show()
	})

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

	cards := container.NewHBox(layout.NewSpacer(), gcard, sCard, jCard, layout.NewSpacer())

	content := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
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
