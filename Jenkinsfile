pipeline {
    agent any
    environment {
<<<<<<< HEAD
        NO_VAR = 'novar'
=======
>>>>>>> ebf3823e8d803917fb2f06ea39db29900f42de5b
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
