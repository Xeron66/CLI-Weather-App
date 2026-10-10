# CLI-Weather-App

A command-line weather application built with **Go** that retrieves weather information for a city using the [Open-Meteo API](https://open-meteo.com/).

The application uses the Open-Meteo Geocoding API to find a city's latitude and longitude, then uses those coordinates to fetch weather information from the Forecast API and display it in the terminal.

## Features

- **City-based search:** Retrieve weather information by providing a city name.
- **Geocoding:** Convert city names into geographic coordinates using the Open-Meteo Geocoding API.
- **Weather data:** Fetch current weather conditions and daily weather information.
- **Dynamic API requests:** Use the coordinates returned by geocoding to construct the weather API request.
- **CLI output:** Display the retrieved weather information directly in the terminal.
- **JSON parsing:** Decode API responses into Go structs using Go's `encoding/json` package.
- **HTTP requests:** Communicate with the APIs using Go's `net/http` package.

## APIs Used

This project uses the following free APIs from Open-Meteo:

| API | Purpose |
|---|---|
| [Geocoding API](https://open-meteo.com/en/docs/geocoding-api) | Finds a location and returns its latitude and longitude. |
| [Weather Forecast API](https://open-meteo.com/en/docs) | Retrieves current weather conditions and daily weather data using the coordinates. |

No API key is required for the standard Open-Meteo API endpoints used by this project.

## How It Works

The application follows this workflow:

1. **Accept a city name:** Specify the city whose weather information you want to retrieve.
2. **Geocode the location:** Send a request to the Geocoding API to obtain the city's latitude and longitude.
3. **Extract coordinates:** Parse the JSON response and extract the geographic coordinates.
4. **Request weather data:** Construct a Forecast API URL using the retrieved coordinates and the required weather parameters.
5. **Parse the response:** Decode the weather API's JSON response into Go structs.
6. **Display the results:** Format and print the weather information in the CLI.

### Application Workflow

```text
City Name
    |
    v
Open-Meteo Geocoding API
    |
    v
Latitude and Longitude
    |
    v
Open-Meteo Forecast API
    |
    v
JSON Weather Response
    |
    v
Go Structs and JSON Parsing
    |
    v
Formatted CLI Output
```

## Weather Information

Depending on the API parameters requested, the application can display information such as:

- Current temperature
- Apparent temperature (feels-like temperature)
- Relative humidity
- Precipitation and rain
- Wind speed and direction
- Daily maximum temperature
- Sunrise and sunset times
- Daylight duration
- Maximum UV index
- Maximum precipitation probability

## Technologies Used

- **Language:** Go (Golang)
- **APIs:** Open-Meteo Geocoding API and Forecast API
- **HTTP client:** `net/http`
- **JSON decoding:** `encoding/json`
- **String processing and formatting:** `strings` and `fmt`

## Prerequisites

- [Go](https://go.dev/dl/) installed on your system.
- An internet connection to access the Open-Meteo APIs.

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Xeron66/Go-with-test.git
```

### 2. Navigate to the Project Directory

Navigate to the directory containing the weather application's `go.mod` file and Go source code.

```bash
cd Go-with-test
```

If the weather application is in a subdirectory, navigate into that directory instead.

### 3. Run the Application

```bash
go run .
```

Follow any input prompts implemented by the application to retrieve weather information for your chosen city.

### 4. Build an Executable

To compile the application into a Windows executable:

```bash
go build -o weather.exe .
```

Run it using:

```powershell
.\weather.exe
```

## Error Handling

The application should account for common API and input errors, including:

- Network connectivity failures.
- Non-successful HTTP status codes.
- Invalid or empty API responses.
- Cities that cannot be found by the Geocoding API.
- JSON decoding errors.

## Learning Objectives

This project is part of my journey learning Go and applying its fundamentals to a practical application. It provides hands-on experience with:

- Making HTTP requests in Go.
- Working with external REST APIs.
- Defining structs that map to JSON responses.
- Parsing JSON data.
- Handling slices and nested data structures.
- Building dynamic URLs from API response data.
- Formatting output for command-line applications.

## Future Improvements

Potential enhancements include:

- Improving input validation and error handling.
- Displaying weather conditions in a more readable CLI layout.
- Supporting multiple locations and disambiguating cities with the same name.
- Allowing users to select temperature and wind-speed units.
- Adding automated tests for API response parsing and application logic.
- Separating geocoding, weather retrieval, and presentation into separate functions or packages.

## Acknowledgements

Weather data and geocoding services are provided by [Open-Meteo](https://open-meteo.com/).

---

**Note:** This is a learning project built to practise Go fundamentals through a real-world API integration.