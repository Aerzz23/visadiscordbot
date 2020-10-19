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
        sh '''curl -fsSLO https://get.docker.com/builds/Linux/x86_64/docker-17.04.0-ce.tgz \
  && tar xzvf docker-17.04.0-ce.tgz \
  && mv docker/docker /usr/local/bin \
  && rm -r docker docker-17.04.0-ce.tgz'''
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