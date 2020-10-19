pipeline {
  agent {
    docker {
      image 'golang:1.15.3'
      args '--mount src=/home/pi/go-build-cache,target=/.cache,type=bind'}

  }
  stages {
    stage('Go Build'){
      steps {
        sh 'export GOROOT='
        sh 'go get -u ./...'
        sh 'ls -ltra'
        sh 'go build ./...'
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