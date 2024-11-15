# Anti Grass Weather App

A simple Go app that tells you if it’s safe to go outside near grass based on the weather.

## What It Does
- Finds your location using your IP address.
- Gets the current weather (temperature and wind speed) for your area.

# Installation

## Using the pre-compiled exe (Windows users only):
1. Head over to the [releases](https://github.com/anti-grass/anti-grass-weather-app/releases/).
2. Download the .exe file.
3. Run the .exe once installed.

## Self-compiling (Non windows users, or advanced users):
1. Clone the repo:
   ```bash
   git clone https://github.com/anti-grass/anti-grass-weather-app.git
   cd anti-grass-weather-app
   ```
2. Install dependencies:
    ```go
    go mod tidy
    ```
3. Run the app:
    ```go
    go run main.go
    ```

## How it works
- **Location**: Uses the ipinfo.io API to get your latitude and longitude.
- **Weather**: Fetches data from Open-Meteo.
