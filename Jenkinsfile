pipeline {
  agent any
  environment {
    name_final = "sgp-info-svc"
    DB_CREDS_DEV = credentials('db-creds-sgpinfosvc')
  }
  stages {
    stage('Docker Build') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          def result = sh(returnStdout: true, script: 'echo "$(docker ps -q --filter name=${name_final})"').trim()
          if (result != "") {
            sh '''
            docker stop ${name_final}
            docker rm -vf ${name_final}
            docker build . -t ${name_final}
            docker system prune -f
	    '''
          } else {
            sh '''
            docker build . -t ${name_final}
            docker system prune -f
	    '''
          }
        }
      }
    }
    stage('SonarQube Analysis') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        sh 'echo SonarQube'
      }
    }
    stage('RUN DB DEV') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          sh '''
          docker run --rm flyway/flyway:8.5.1 version
          docker run --rm -v $WORKSPACE/sql:/flyway_dev/sql -v $WORKSPACE/sql:/flyway_dev/conf flyway/flyway:8.5.1 -user=$DB_CREDS_DEV_USR -password=$DB_CREDS_DEV_PSW migrate
          docker run --rm -v $WORKSPACE/sql:/flyway_dev/sql -v $WORKSPACE/sql:/flyway_dev/conf flyway/flyway:8.5.1 -user=$DB_CREDS_DEV_USR -password=$DB_CREDS_DEV_PSW validate
          docker run --rm -v $WORKSPACE/sql:/flyway_dev/sql -v $WORKSPACE/sql:/flyway_dev/conf flyway/flyway:8.5.1 -user=$DB_CREDS_DEV_USR -password=$DB_CREDS_DEV_PSW info
	  '''
        }
      }
    }
    stage('Deploy to DEV') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          sh '''
          docker run -dt -p 30001:90 --name ${name_final} ${name_final}
          docker system prune -f
	  '''
        }
      }
    }
    stage('Cucumber Tests DEV') {
      agent {
        label 'dev'
      }
      when {
        anyOf {
          branch 'sgp*'
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        echo 'SonarQube'
      }
    }
    stage('RUN DB QA') {
      agent {
        label 'qa'
      }
      when {
        anyOf {
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          sh '''
          docker run --rm flyway/flyway:8.5.1 version
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:8.5.1 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW migrate
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:8.5.1 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW validate
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:8.5.1 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW info
	  '''
        }
      }
    }
    stage('Deploy to QA') {
      agent {
        label 'qa'
      }
      when {
        anyOf {
          branch 'sprint-*'
          branch 'master'
        }
      }
      steps {
        script {
          def result = sh(returnStdout: true, script: 'echo "$(docker ps -q --filter name=${name_final})"').trim()
          if (result != "") {
            sh '''
            docker stop ${name_final}
            docker rm -vf ${name_final}
            docker build . -t ${name_final}
            docker run -dt -p 30001:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          } else {
            sh '''
            docker build . -t ${name_final}
            docker run -dt -p 30001:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          }
        }
      }
    }
    stage('QA Approval') {
      agent {
        label 'prd'
      }
      when {
          branch 'master'
      }
      steps {
        input "Aprobacion Tester QA"
      }
    }
    stage('RUN DB PRD') {
      agent {
        label 'prd'
      }
      when {
          branch 'master'
      }
      steps {
        script {
          sh '''
          docker run --rm flyway/flyway:8.5.1 version
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:8.5.1 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW migrate
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:8.5.1 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW validate
          docker run --rm -v $WORKSPACE/sql:/flyway/sql -v $WORKSPACE/sql:/flyway/conf flyway/flyway:8.5.1 -user=$DB_CREDS_USR -password=$DB_CREDS_PSW info
	  '''
        }
      }
    }
    stage('Deploy to PRD') {
      agent {
        label 'prd'
      }
      when {
          branch 'master'
      }
      steps {
        script {
          def result = sh(returnStdout: true, script: 'echo "$(docker ps -q --filter name=${name_final})"').trim()
          if (result != "") {
            sh '''
            docker stop ${name_final}
            docker rm -vf ${name_final}
            docker build . -t ${name_final}
            docker run -dt -p 30001:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          } else {
            sh '''
            docker build . -t ${name_final}
            docker run -dt -p 30001:90 --name ${name_final} ${name_final}
            docker system prune -f
	    '''
          }
        }
      }
    }
  }
}
