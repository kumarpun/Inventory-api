agentName = "master"
pipeline {
  agent {label agentName}
    
   stages {
        
    stage('Cloning Git') {
      steps {
        git branch: 'develop',
            credentialsId: 'a7e32b0a-b2af-426b-8042-af214734f21e',
            url: 'http://build@10.1.1.10:8080/scm/git/HRMS-Inventory'
      }
    }
     stage('Initilization') {
       steps{
            sh 'echo initialization is being skipped since tools server need to set GoPath.'
    //       sh  'go get -v github.com/gorilla/handlers'
    //       sh  'go get -v github.com/gorilla/mux'
    //       sh  'go get -v github.com/jinzhu/gorm'
    //       sh  'go get -v github.com/jinzhu/gorm/dialects/sqlite'
       }
     } 
   stage('Build') {
      steps {
       // sh 'go build main.go controller.go models.go'
       sh 'echo build is being skipped since tools server need to set GoPath.'
           }
    }   
   stage('Test') {
      steps {
        sh 'echo Test is being skipped since there is no test.'
           }
    }
	
    stage('Deploy'){
        steps{
            	sh 'sshpass -p "sevalg2014" ssh root@antlet14 -o StrictHostKeyChecking=no "rm -rf /var/www/temp-deploy/dist/HRMS-Inventory"'
		sh 'sshpass -p "sevalg2014" ssh root@antlet14 "mkdir -p /var/www/temp-deploy/dist/HRMS-Inventory"'
		sh 'sshpass -p "sevalg2014" scp -r * root@antlet14:/var/www/temp-deploy/dist/HRMS-Inventory/'
		sh 'sshpass -p "sevalg2014" ssh root@antlet14 "rm -rf /var/www/HRMS/HRMS-Inventory/"'
		sh 'sshpass -p "sevalg2014" ssh root@antlet14 "mv /var/www/temp-deploy/dist/HRMS-Inventory/ /var/www/HRMS/"'
		sh 'sshpass -p "sevalg2014" ssh root@antlet14 "cd /var/www/HRMS/HRMS-Inventory/ && source getPackagesAndStartService.sh"'
        }
    }
  }
   post{
      success{
          echo 'This will run only if successful' 
      }
      failure {
          emailext body: '$DEFAULT_CONTENT', recipientProviders: [culprits(), developers()], subject: '$DEFAULT_SUBJECT', to: '$DEFAULT_RECIPIENTS'
          emailextrecipients([developers()])
      }
      unstable {  
             echo 'This will run only if the run was marked as unstable'  
         }  
         changed {  
             echo 'This will run only if the state of the Pipeline has changed'  
             echo 'For example, if the Pipeline was previously failing but is now successful'  
         }  
  }
}