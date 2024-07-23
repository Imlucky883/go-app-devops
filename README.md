# End-to-End DevOps Implementation for a Golang Web Application 

![arhitecture](/static/architecture.png)

## Introduction

This project showcases a modern approach to building, deploying, and managing a Go web application using containerization, continuous integration, and continuous deployment (CI/CD) practices.

The application is designed to run in a Kubernetes environment, leveraging Docker for containerization and GitHub Actions for CI. We utilize Helm for managing Kubernetes manifests and ArgoCD for continuous deployment, ensuring that our application is always up to date.

## Project Features

**Containerization**: Utilizing Multi-Stage Docker Builds to create lightweight and efficient images.

**Continuous Integration**: Setting up GitHub Actions to automate the build and testing processes.

**Continuous Delivery**: Implementing Argo CD for streamlined deployment and management of applications on Kubernetes.

**Kubernetes Cluster Setup**: Launching a EKS CLuster in AWS

**Helm Chart Creation**: Designing Helm charts for easy configuration and management across multiple environments.

**Ingress Controller Configuration:** Setting up an Ingress controller to expose the application securely.

**DNS Mapping:** Configuring DNS for our domain to ensure smooth access to the application.

**End-to-End CI/CD Demonstration**: Showcasing the entire pipeline in action

## Getting Started

### Prerequisites

- Go installed on your local machine
- Docker installed
- Dockerhub account
- Access to a GitHub account
- AWS account
- kubectl installed
- aws-cli installed
- Helm installed

### Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/Imlucky883/go-web-app.git
   cd go-web-app