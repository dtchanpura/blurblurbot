pipeline {
  agent {
    docker {
      image 'go/alpine'
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