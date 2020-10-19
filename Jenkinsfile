pipeline {
  agent any
  stages {
    stage('Go Build'){
     agent{ 
      docker {
        image 'golang:1.15.3'
        args '--mount src=/home/pi/go-build-cache,target=/.cache,type=bind'
        }
      }
      steps {
        sh 'go get -u ./...'
        sh 'ls -ltra'
        sh 'go build ./...'
      }
    }
    stage('Go Test') {
      agent{
        docker {
        image 'golang:1.15.3'
        args '--mount src=/home/pi/go-build-cache,target=/.cache,type=bind'
        }
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
    stage('Docker Publish') {
      when {
        branch 'master'
      }
      steps {
        withDockerRegistry([ credentialsId: "dockerhub_id", url: "" ]) {
          sh "docker push aerzz23/visadiscordbot:latest"
          }
      }
    }
  }
}