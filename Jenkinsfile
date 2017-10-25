pipeline {
  agent {
    docker {
      image 'golang:alpine'
    }
    
  }
  stages {
    stage('Build') {
      steps {
        sh 'make'
      }
    }
  }
}