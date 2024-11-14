pipeline {
    agent any
    
    tools {
        go 'golang'
    }

    stages {
        stage('Clean Workspace') {
            steps {
                cleanWs()
            }
        }
        
        stage('Checkout from Git') {
            steps {
                git branch: 'master', url: 'https://github.com/omidiyanto/golang-weatherapp-googleoauth.git'
            }
        }
        
        stage('Install Dependencies') {
            steps {
                sh "npm install"
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
                      ${tool 'snyk'}/snyk-linux code test --all-projects --severity-threshold=high > snyk-code-analysis-results.txt
                    """
                }
            }
        }
    }
    
    post {
        always {
            // Archive Snyk test results and code analysis artifacts (JSON or HTML)
            archiveArtifacts artifacts: '**/snyk-test*.json, snyk-code-analysis-results.txt, snyk-report.html', allowEmptyArchive: true
        }
        
        success {
            echo 'Snyk tests completed successfully.'
        }
        
        failure {
            echo 'Snyk tests failed. Please review the logs.'
        }
    }
}
