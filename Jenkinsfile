pipeline {
  agent {
    kubernetes {
      yamlFile './deployments/go.pod.slave.yaml'
    }
  }

  environment {
    DOCKER_REGISTRY_HOST="docker-registry.192-168-33-10.nip.io"
    DOCKER_REGISTRY_LOGIN=credentials("docker-registry-login")
  }

  stages {
    stage('编译api...') {
      steps {
        container('golang') {
          sh 'go get -d -v'
          // fix from: https://www.cloudreach.com/en/resources/blog/containerize-this-how-to-build-golang-dockerfiles/
          sh '''
          GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
          go build -a -installsuffix cgo \
          -ldflags '-extldflags "-static"' -o bin/app
          '''
        }
      }
    }

    stage('构建镜像并上传仓库') {
      steps {
        container('docker') {
          sh 'docker build -f Dockerfile -t ${DOCKER_REGISTRY_HOST}/fox-api:${GIT_COMMIT} .'
          sh 'docker tag ${DOCKER_REGISTRY_HOST}/fox-api:${GIT_COMMIT} ${DOCKER_REGISTRY_HOST}/fox-api:latest'
          sh 'docker login ${DOCKER_REGISTRY_HOST} -u ${DOCKER_REGISTRY_LOGIN_USR} -p ${DOCKER_REGISTRY_LOGIN_PSW}'
          sh 'docker push ${DOCKER_REGISTRY_HOST}/fox-api:${GIT_COMMIT}'
          sh 'docker push ${DOCKER_REGISTRY_HOST}/fox-api:latest'
        }
      }
    }

    stage('发布到开发环境') {
      steps {
        container('kubectl') {
          sh 'kubectl apply -o name --force -f ./deployments/dev/service.yaml'
          sh '''
          DEPLOYMENT_NAME=$(kubectl apply -o name --force -f ./deployments/dev/deployment.yaml)
          kubectl rollout status $DEPLOYMENT_NAME
          '''.stripIndent()
        }
      }
    }
  }
}
