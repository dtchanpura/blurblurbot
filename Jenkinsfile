pipeline {
    agent any
    environment {
        NO_VAR = 'novar'
        // GOPATH and GOROOT already exists
        // GOROOT = ${root}
        // GOPATH = ${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/
    }
    stages {
        // stage('Checkout') {
        //     git url: 'https://github.com/dtchanpura/blurblurbot.git'
        // }
        stage('Initialize') {
            steps {
                sh 'printenv'
                sh 'ls -l'
                // sh 'go version'
                // sh 'go get -u github.com/golang/dep/...'
                // sh 'dep ensure'
            }
        }
    }
}
