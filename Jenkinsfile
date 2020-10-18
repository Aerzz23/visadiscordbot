pipeline {
  agent {
    docker {image 'golang:1.15.3'}

  }
  stages {
    stage('Go Build'){
      steps {
        sh 'export XDG_CACHE_HOME=/home/pi/go-build-cache'
        sh 'go get -u ./...'
        sh 'go build api/'
      }
    }
    stage('Go Test') {
      steps {
        sh 'go test ./...'
      }
    }
    stage('Docker Build') {
      steps {
        sh 'docker build -t aerzz23/visadiscordbot:latest .'
      }
    }
     stage('Docker Push') {
      steps {
        sh 'docker push aerzz23/visadiscordbot:latest'
      }
    }

  }
}