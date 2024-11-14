<div align="center">
	<h1> 🌟 DevSecOps Project - Golang Weather Application🚀 </h1>
</div>
<div align="center">
    <img src="https://img.shields.io/badge/kubernetes-blue?style=for-the-badge&logo=kubernetes&logoColor=white">
    <img src="https://img.shields.io/badge/docker-white.svg?style=for-the-badge&logo=docker&logoColor=blue">
    <img src="https://img.shields.io/badge/golang-cyan.svg?style=for-the-badge&logo=go&logoColor=black">
    <img src="https://img.shields.io/badge/snyk-purple.svg?style=for-the-badge&logo=snyk&logoColor=white">
    <img src="https://img.shields.io/badge/sonarqube-white.svg?style=for-the-badge&logo=sonarqube&logoColor=blue">
    <img src="https://img.shields.io/badge/trivy-yellow.svg?style=for-the-badge&logo=trivy&logoColor=white">
    <img src="https://img.shields.io/badge/prometheus-orange.svg?style=for-the-badge&logo=prometheus&logoColor=white">
    <img src="https://img.shields.io/badge/grafana-grey.svg?style=for-the-badge&logo=grafana&logoColor=white">
	  <img src="https://img.shields.io/badge/jenkins-maroon.svg?style=for-the-badge&logo=jenkins&logoColor=white">
    <img src="https://img.shields.io/badge/cloudflare-red.svg?style=for-the-badge&logo=cloudflare&logoColor=white">
</div>
<br>

## CI/CD Pipeline Architecture
<img src="https://github.com/user-attachments/assets/c38799ff-abd5-4997-8834-d96423311558"></img>

## Overview

This Simple **Golang Weather Application** is part of a DevSecOps project, designed to implement modern CI/CD practices with security and monitoring in mind. The application retrieves weather data from an external API, processes it, and displays the results to users. It is built using **Go (Golang)**, dockerized with **Docker**, and deployed with **Kubernetes**. The project leverages various DevSecOps tools to ensure code quality, security, and performance throughout the development lifecycle.

The project integrates a **CI/CD pipeline** that automates build, testing, security testing and analysis (SCA and SAST), and deployment to the Kubernetes. It also includes monitoring and alerting capabilities with **Prometheus** and **Grafana**, ensuring the application is highly available and secure in production environments.

### Key Features:
- **Go (Golang)**: The backend is built using Go, a statically typed, compiled language known for its efficiency and performance.
- **Dockerized Application**: Easily deployable using Docker containers.
- **Kubernetes**: Orchestrates and scales the application efficiently.
- **DevSecOps Integration**:
  - **SonarQube** as Static Application Security Testing (SAST) tools for static code analysis and quality checks.
  - **Snyk** as Software Composition Analysis (SCA) tools for security scanning of dependencies.
  - **Trivy** for container image vulnerability scanning.
  - **Prometheus & Grafana** for real-time monitoring and visualization.
  - **Jenkins** for automated CI/CD pipeline execution.
  - **Cloudflare** for enhanced security and domain provider.

## Installation
To get started with the Golang Weather Application and DevSecOps pipeline, follow these steps:

### Prerequisites:
- Kubernetes Cluster
- Jenkins
- Prometheus & Grafana

### Steps:
1. Run Sonarqube and Install Trivy

    Run sonarqube as container:
    ```bash
    docker run -d --name sonar -p 9000:9000 sonarqube:lts-community
    ```

    Installing trivy:
    ```bash
    sudo apt-get install wget apt-transport-https gnupg lsb-release
    wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | sudo apt-key add -
    echo deb https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main | sudo tee -a /etc/apt/sources.list.d/trivy.list
    sudo apt-get update
    sudo apt-get install trivy        
    ```

2. Install Jenkins Plugins
    - Golang
    - Github integration
    - SonarQube Scanner
    - Snyk
    - Docker
    - Kubernetes

3. Configure Jenkins Tools
    - Golang
    - Sonarqube Scanner
    - Snyk
    - Docker

4. Configure Jenkins Global System
    - Add Sonarqube Server
    - Configure E-Mail Notification
    - Configure Extended E-Mail Notification

5. Get Snyk, Sonarqube, Quay.io token/API_KEY

6. Configure Required Secrets/Credentials in Jenkins

7. Run the Pipeline using Jenkinsfile

## Run Locally

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
