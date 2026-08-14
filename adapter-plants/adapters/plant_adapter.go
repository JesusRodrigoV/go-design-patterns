package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"adapter-plants/models"
)

type plantNetEspecie struct {
	ID                          string   `json:"id"`
	CommonNames                 []string `json:"commonNames"`
	Genus                       string   `json:"genus"`
	ScientificNameWithoutAuthor string   `json:"scientificNameWithoutAuthor"`
	ScientificNameAuthorship    string   `json:"scientificNameAuthorship"`
}

type PlantNetAdapter struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NuevoPlantNetAdapter(apiKey string) *PlantNetAdapter {
	return &PlantNetAdapter{
		APIKey:  apiKey,
		BaseURL: "https://my-api.plantnet.org/v2/species",
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (a *PlantNetAdapter) ejecutarPeticion(url string) ([]models.Planta, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error construyendo petición: %w", err)
	}
	req.Header.Add("accept", "application/json")

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fallo de comunicación con el servidor: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("estado inesperado del servidor: %d", resp.StatusCode)
	}

	var especiesExternas []plantNetEspecie
	if err := json.NewDecoder(resp.Body).Decode(&especiesExternas); err != nil {
		return nil, fmt.Errorf("error interpretando los datos: %w", err)
	}

	var plantas []models.Planta
	for _, ext := range especiesExternas {
		nombreComun := "Sin nombre común registrado"
		if len(ext.CommonNames) > 0 {
			nombreComun = ext.CommonNames[0]
		}

		plantas = append(plantas, models.Planta{
			ID:               ext.ID,
			NombreComun:      nombreComun,
			Genero:           ext.Genus,
			NombreCientifico: fmt.Sprintf("%s %s", ext.ScientificNameWithoutAuthor, ext.ScientificNameAuthorship),
		})
	}
	return plantas, nil
}

func (a *PlantNetAdapter) GetPlantas(limite int, pagina int) ([]models.Planta, error) {
	url := fmt.Sprintf("%s?lang=en&type=kt&pageSize=%d&page=%d&api-key=%s",
		a.BaseURL, limite, pagina, a.APIKey)
	return a.ejecutarPeticion(url)
}

func (a *PlantNetAdapter) BuscarPorPrefijo(prefijo string, limite int, pagina int) ([]models.Planta, error) {
	url := fmt.Sprintf("%s?lang=en&type=kt&pageSize=%d&page=%d&prefix=%s&api-key=%s",
		a.BaseURL, limite, pagina, prefijo, a.APIKey)
	return a.ejecutarPeticion(url)
}
