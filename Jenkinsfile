pipeline {
  agent any
  stages {
    stage('Steps') {
      parallel {
        stage('Print Message') {
          steps {
            echo 'Hello'
          }
        }

        stage('Another One') {
          steps {
            echo 'Again'
          }
        }

      }
    }

  }
}