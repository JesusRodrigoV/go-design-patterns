package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"adapter-plants/adapters"
	"adapter-plants/models"
	"adapter-plants/services"

	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: No se encontró el archivo .env.")
	}

	apiKey := os.Getenv("PLANTNET_API_KEY")
	if apiKey == "" {
		pterm.Error.Println("Fallo crítico: Falta PLANTNET_API_KEY en el entorno.")
		os.Exit(1)
	}

	var servicio services.PlantaService = adapters.NuevoPlantNetAdapter(apiKey)

	tamanoPagina := 10

	for {
		pterm.DefaultHeader.WithFullWidth().Println("Plantas")

		opcion, err := pterm.DefaultInteractiveSelect.
			WithOptions([]string{
				"1. Ver Catálogo General",
				"2. Buscar Plantas por Prefijo",
				"3. Configurar Resultados por Página",
				"4. Salir",
			}).
			Show("Seleccione una operación")

		if err != nil {
			pterm.Error.Println("Error leyendo la selección:", err)
			return
		}

		switch opcion {
		case "1. Ver Catálogo General":
			gestionarPaginacion(servicio, tamanoPagina, "")

		case "2. Buscar Plantas por Prefijo":
			prefijo := leerEntrada("Introduzca el prefijo de búsqueda (ej. 'Ama'): ")
			if prefijo != "" {
				gestionarPaginacion(servicio, tamanoPagina, prefijo)
			}

		case "3. Configurar Resultados por Página":
			input := leerEntrada(fmt.Sprintf("Cantidad actual (%d). Nueva cantidad: ", tamanoPagina))
			nuevoTamano, err := strconv.Atoi(input)
			if err != nil || nuevoTamano <= 0 {
				pterm.Warning.Println("Entrada inválida. Se mantiene el tamaño actual.")
			} else {
				tamanoPagina = nuevoTamano
				pterm.Success.Printf("Tamaño de página actualizado a %d.\n", tamanoPagina)
			}

		case "4. Salir":
			pterm.Info.Println("Cerrando")
			os.Exit(0)
		}

		fmt.Println()
	}
}

func gestionarPaginacion(servicio services.PlantaService, limite int, prefijo string) {
	paginaActual := 1

	for {
		spinner, _ := pterm.DefaultSpinner.Start(fmt.Sprintf("Cargando página %d...", paginaActual))

		var plantas []models.Planta
		var err error

		if prefijo == "" {
			plantas, err = servicio.GetPlantas(limite, paginaActual)
		} else {
			plantas, err = servicio.BuscarPorPrefijo(prefijo, limite, paginaActual)
		}

		if err != nil {
			spinner.Fail("Fallo al obtener registros: ", err.Error())
			return
		}
		spinner.Success(fmt.Sprintf("Mostrando página %d", paginaActual))

		if len(plantas) == 0 {
			pterm.Warning.Println("No hay más resultados para mostrar.")
			if paginaActual > 1 {
				paginaActual--
				continue
			}
			return
		}

		mostrarTabla(plantas)

		fmt.Printf("\n--- [S] Siguiente Página | [A] Página Anterior | [V] Volver al Menú Principal ---\n")
		accion := strings.ToUpper(leerEntrada("Su elección: "))

		switch accion {
		case "S":
			paginaActual++
		case "A":
			if paginaActual > 1 {
				paginaActual--
			} else {
				pterm.Warning.Println("Ya estás en la primera página.")
			}
		case "V":
			return
		default:
			pterm.Error.Println("Comando no reconocido.")
		}
	}
}

func mostrarTabla(plantas []models.Planta) {
	datosTabla := pterm.TableData{{"ID", "Nombre Común", "Género", "Nombre Científico"}}
	for _, p := range plantas {
		datosTabla = append(datosTabla, []string{p.ID, p.NombreComun, p.Genero, p.NombreCientifico})
	}
	pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(datosTabla).Render()
}

func leerEntrada(mensaje string) string {
	fmt.Print(mensaje)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
