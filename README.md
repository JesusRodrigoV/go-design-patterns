# patrones-arq

Repositorio de práctica con implementaciones de patrones de diseño en Go.

## Estructura

| Carpeta | Patrón | Descripción |
| --- | --- | --- |
| `adapter-plants` | Adapter | CLI interactiva que consulta el catálogo de la API de PlantNet y adapta su respuesta al modelo del dominio. |
| `factory-notifications` | Factory Method | Sistema de notificaciones multicanal (email, SMS, push) que delega la creación del notificador en una fábrica. |
| `observer-weather` | Observer | Estación meteorológica que notifica a múltiples observadores (display, estadísticas, alertas, logger) ante cada medición. |

## Requisitos

- Go 1.25+
- Para `adapter-plants`: una clave de API de PlantNet

## Cómo ejecutar

Cada proyecto es independiente. Entrá a la carpeta y ejecutá:

```bash
# Adapter (requiere la API key)
cd adapter-plants
go run main.go

# Factory Method
cd factory-notifications
go run main.go

# Observer
cd observer-weather
go run main.go
```

## Configuración de la API key

`adapter-plants` lee la clave desde el entorno:

```bash
# Linux/macOS
export PLANTNET_API_KEY=tu_key_aqui

# Windows (PowerShell)
$env:PLANTNET_API_KEY = "tu_key_aqui"
```

También podés copiar `adapter-plants/.env.example` a `adapter-plants/.env` y completar la clave ahí. El archivo `.env` está en `.gitignore` y no debe subirse al repositorio.