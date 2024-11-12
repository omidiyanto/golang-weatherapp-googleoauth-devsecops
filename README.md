# Golang Weather App

## Overview

This project is a web application that integrates Google OAuth for authentication and fetches weather data from the OpenWeather API.

## Prerequisites

Before building the Docker image, ensure you have the following:

- [Docker](https://www.docker.com/get-started) installed on your machine.
- Your Google Cloud project set up with OAuth credentials.
- An OpenWeather API key.

## Build Docker Image

To build the Docker image for this application, you need to provide several build arguments. These arguments include your Google Client ID, Google Client Secret, OpenWeather API Key, and the Redirect URI.

### Build Command

Run the following command in your terminal, replacing the placeholders with your actual credentials:

```bash
docker build \
  --build-arg ARG_GOOGLE_CLIENT_ID=<YOUR_GOOGLE_CLIENT_ID> \
  --build-arg ARG_GOOGLE_CLIENT_SECRET=<YOUR_GOOGLE_CLIENT_SECRET> \
  --build-arg ARG_OPENWEATHER_API_KEY=<YOUR_OPENWEATHER_API_KEY> \
  --build-arg ARG_REDIRECT_URI=<YOUR_REDIRECT_URI> \
  -t <IMAGE_NAME>:<TAG> .
```

### Run Command
Run the following command in your terminal to start running as container
```bash
docker run -d --name GO_WEATHERAPP -p 5000:5000 <IMAGE_NAME>:<TAG> 
```