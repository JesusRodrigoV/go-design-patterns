package services

import "adapter-plants/models"

type PlantaService interface {
	GetPlantas(limite int, pagina int) ([]models.Planta, error)
	BuscarPorPrefijo(prefijo string, limite int, pagina int) ([]models.Planta, error)
}
