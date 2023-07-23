pipeline {
    agent { 
        node {
            label 'docker-agent-golang'
            }
	}
    stages {
        stage('install mod') {
            steps {
                echo "Building.."
                sh '''
                cd app
                go mod download
				go mod verify
                '''
            }
        }
        stage('build') {
            steps {
                echo "Testing.."
                sh '''
                cd app
				go build -o binary
                '''
            }
        }
        stage('testing') {
            steps {
                echo 'Deliver....'
                sh '''
                ./binary
                '''
            }
        }
    }
}