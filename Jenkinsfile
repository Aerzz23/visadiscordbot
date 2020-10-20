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
    stage('Docker Deploy') {
      when {
        branch 'master'
      }
      environment {
        DISCORD_BOT_TOKEN = credentials('discord_token')
      }
      steps {
        sshagent(credentials : ['rpi_ssh']) {
          sh '''if [ -z "$(docker ps -a -q)" ]
                then
                      echo "no docker containers found"
                else
                      echo "docker containers found - stopping and removing"
                      docker stop $(docker ps -a -q)
                      docker rm $(docker ps -a -q)
                fi
          '''
          withDockerRegistry([ credentialsId: "dockerhub_id", url: "" ]) {
            sh '''docker pull aerzz23/visadiscordbot:latest
                  export BUILD_ID=dontKillMe
                  export JENKINS_NODE_COOKIE=dontKillMe
                  nohup docker run --env VISA_BOT_TOKEN=${DISCORD_BOT_TOKEN} aerzz23/visadiscordbot:latest &'''
          }
        }
       
      }
    }
  }
}