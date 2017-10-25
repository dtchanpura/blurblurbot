pipeline {
  agent any
  stages {
    stage('Initialize') {
      agent any
      environment {
        GOPATH = '/var/lib/jenkins/go'
      }
      steps {
        sh 'mkdir -p $GOPATH/src/github.com/dtchanpura/blurblurbot'
        sh 'rsync -az $WORKSPACE $GOPATH/github.com/dtchanpura/blurblurbot'
      }
    }
    stage('Build') {
      steps {
        sh 'make build'
      }
    }
    stage('Test') {
      steps {
        sh 'make test'
      }
    }
    stage('Deploy') {
      steps {
        sh 'echo \'deploy\''
      }
    }
  }
  environment {
    NO_VAR = 'novar'
  }
}