pipeline {
  agent any
 
  stages {
    stage('Go Build'){
      docker {
      image 'golang:1.15.3'
      args '--mount src=/home/pi/go-build-cache,target=/.cache,type=bind'}
    }
      steps {
        sh 'go get -u ./...'
        sh 'ls -ltra'
        sh 'go build ./...'
      }
    }
    stage('Go Test') {
      docker {
      image 'golang:1.15.3'
      args '--mount src=/home/pi/go-build-cache,target=/.cache,type=bind'}
    }
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