pipeline {
    agent any
    environment {
        GOROOT = '${root}'
        GOPATH= '${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/'
    }
    stages {
        // stage('Checkout') {
        //     git url: 'https://github.com/dtchanpura/blurblurbot.git'
        // }
        stage('Initialize') {
            steps {
                sh 'printenv'
                // sh 'go version'
                // sh 'go get -u github.com/golang/dep/...'
                // sh 'dep ensure'                
            }
        }
    }
}
