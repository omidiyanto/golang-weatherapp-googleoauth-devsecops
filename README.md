# Golang Weather App
<img src="https://github-production-user-asset-6210df.s3.amazonaws.com/90853331/386306430-c38799ff-abd5-4997-8834-d96423311558.svg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVCODYLSA53PQK4ZA%2F20241114%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20241114T174037Z&X-Amz-Expires=300&X-Amz-Signature=488ba6f386f807c5495c76725c95af6f54b2d58bad4371160543864445ae9a25&X-Amz-SignedHeaders=host"></img>

<img src="https://github.com/user-attachments/assets/c38799ff-abd5-4997-8834-d96423311558"></img>
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
