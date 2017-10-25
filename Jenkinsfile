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
        sh 'rsync -az $WORKSPACE/* $GOPATH/src/github.com/dtchanpura/blurblurbot/'
      }
    }
    stage('Build') {
      steps {
        sh 'cd $GOPATH/src/github.com/dtchanpura/blurblurbot; make build'
      }
    }
    stage('Test') {
      steps {
        sh 'cd $GOPATH/src/github.com/dtchanpura/blurblurbot; make test'
      }
    }
    stage('Deploy') {
      steps {
        sh 'echo \'deploy\''
      }
    }
  }
}
