# Golang Weather App
<img src="https://github-production-user-asset-6210df.s3.amazonaws.com/90853331/386301561-7e92780a-4b90-4140-b7c9-2d7fd9c63743.svg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVCODYLSA53PQK4ZA%2F20241114%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20241114T172402Z&X-Amz-Expires=300&X-Amz-Signature=e219bb1d39f73f5f02f8e0529905b6039aeafe0e4d9f28115ef5f1f1552077b3&X-Amz-SignedHeaders=host"></img>
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
docker build -t <IMAGE_NAME>:<TAG> .
```

### Run Command
Run the following command in your terminal to start running as container
```bash
docker run -d -p 5000:5000 \
  -e GOOGLE_CLIENT_ID=<YOUR_GOOGLE_CLIENT_ID> \
  -e GOOGLE_CLIENT_SECRET=<YOUR_GOOGLE_CLIENT_SECRET> \
  -e OPENWEATHER_API_KEY=<YOUR_OPENWEATHER_API_KEY> \
  -e REDIRECT_URI=<YOUR_REDIRECT_URI> \
   --name GO_WEATHERAPP <IMAGE_NAME>:<TAG> 
```
