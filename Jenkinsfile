pipeline {
  agent {
    dockerfile {
      filename 'Dockerfile'
    }

  }
  stages {
    stage('Build') {
      steps {
        sh '''go test ./...
go build api/main.go
docker build -t aerzz23/visadiscordbot:latest .'''
      }
    }

  }
}