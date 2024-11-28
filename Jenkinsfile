pipeline {
    agent any
    
    tools {
        go 'golang'
        dockerTool 'docker'
    }
    
    environment {
        IMAGE_TAG = "v1.0.${env.BUILD_NUMBER}"
        REPOSITORY = 'quay.io/omidiyanto' 
        APP_NAME = 'golang-weatherapp' 
        YAML_PATH = 'Kubernetes' 
        GOOGLE_CLIENT_ID = credentials('GOOGLE_CLIENT_ID')
        GOOGLE_CLIENT_SECRET = credentials('GOOGLE_CLIENT_SECRET')
        OPENWEATHER_API_KEY = credentials('OPENWEATHER_API_KEY')
        REDIRECT_URI='https://weatherapp.omidiyanto.my.id/callback'
    }

    stages {
        stage('Clean Workspace') {
            steps {
                cleanWs()
            }
        }
        
        stage('Checkout from Git') {
            steps {
                git branch: 'master', url: 'https://github.com/omidiyanto/golang-weatherapp-googleoauth-devsecops.git'
            }
        }
        
        stage('Install Dependencies') {
            steps {
                sh "go mod download"
                sh "go mod tidy"
            }
        }
        
        stage('SAST (sonarqube'){
            steps{
                withSonarQubeEnv('sonarqube') {
                    sh """
                        ${tool 'sonar-scanner'}/bin/sonar-scanner \
                            -Dsonar.projectKey=omi-organization_golang-weatherapp 
                    """
                }
                sh "sleep 5"
            }
        }
        
        stage("Quality Gate") {
            steps {
                script {
                    waitForQualityGate abortPipeline: false, credentialsId: 'SONAR_TOKEN'
                }
            }
        }

        stage('SCA (Dependency Check)') {
            steps {
                snykSecurity(
                    severity: 'high',
                    snykInstallation: 'snyk',
                    snykTokenId: 'SNYK_TOKEN'
                )
            }
        }

        stage('SCA (Static Code Analysis)') {
            steps {
                withCredentials([string(credentialsId: 'MY_SNYK_TOKEN', variable: 'MY_SNYK_TOKEN')]) {
                    sh """
                      ${tool 'snyk'}/snyk-linux auth ${MY_SNYK_TOKEN}
                      ${tool 'snyk'}/snyk-linux code test --all-projects --severity-threshold=high | tee snyk-code-analysis.txt
                    """
                }
            }
        }

        stage("Docker Build & Push") {
            steps {
                script {
                    withDockerRegistry(credentialsId: 'quayio-account', url: 'https://quay.io', toolName: 'docker') {
                        sh "docker build -t ${env.REPOSITORY}/${env.APP_NAME}:${env.IMAGE_TAG} ."
                        sh "docker push ${env.REPOSITORY}/${env.APP_NAME}:${env.IMAGE_TAG}"
                    }
                }
            }
        }

        stage("CONTAINER IMAGE SCAN") {
            steps {
                sh "trivy image ${env.REPOSITORY}/${env.APP_NAME}:${env.IMAGE_TAG} > trivy-image-scan.txt"
            }
        }

        stage("Deploy to k8s") {
            steps {
                script {
                    // 1) Update Image Value 
                    sh """
                        yq eval '.spec.template.spec.containers[0].image = "${env.REPOSITORY}/${env.APP_NAME}:${env.IMAGE_TAG}"' -i ${env.YAML_PATH}/deployment.yml
                    """
                    
                    // 2) Update Deployment Environment Variable
                    sh """
                        yq eval '.spec.template.spec.containers[0].env |= map(select(.name == "GOOGLE_CLIENT_ID").value = "${GOOGLE_CLIENT_ID}")' -i ${env.YAML_PATH}/deployment.yml
                        yq eval '.spec.template.spec.containers[0].env |= map(select(.name == "GOOGLE_CLIENT_SECRET").value = "${GOOGLE_CLIENT_SECRET}")' -i ${env.YAML_PATH}/deployment.yml
                        yq eval '.spec.template.spec.containers[0].env |= map(select(.name == "OPENWEATHER_API_KEY").value = "${OPENWEATHER_API_KEY}")' -i ${env.YAML_PATH}/deployment.yml
                        yq eval '.spec.template.spec.containers[0].env |= map(select(.name == "REDIRECT_URI").value = "${REDIRECT_URI}")' -i ${env.YAML_PATH}/deployment.yml
                    """

                    withKubeConfig(caCertificate: '', clusterName: '', contextName: '', credentialsId: 'kubeconfig', namespace: '', restrictKubeConfigAccess: false, serverUrl: '') {
                        sh "kubectl apply -f ${env.YAML_PATH}"
                    }
                }
            }
        }

        
    }
    
    post {
        always {
            archiveArtifacts artifacts: 'snyk-code-analysis.txt, trivy-image-scan.txt', allowEmptyArchive: true
            emailext(
                attachLog: true,
                subject: "Jenkins Report - ${env.JOB_NAME}#${env.BUILD_NUMBER}",
                body: "Results: '${currentBuild.result}'",
                to: 'midiyanto26@gmail.com',
                attachmentsPattern: 'snyk-code-analysis.txt,trivy-image-scan.txt'
            )
        }
    }
}
