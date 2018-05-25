node {
  stage('Preparation') { // for display purposes
      // Get some code from a GitHub repository
      git 'https://github.com/vacoj/timewindow.git'
  }
  stage('Build Test + SonarQube analysis') {

     // Install the desired Go version
    sh 'wget https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz > /dev/null'
    sh 'tar -zxvf go1.10.2.linux-amd64.tar.gz > /dev/null'
 
    // Export environment variables pointing to the directory where Go was installed
    sh 'go/bin/go test -coverprofile=coverage.out'
    sh 'wget https://sonarsource.bintray.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-3.2.0.1227-linux.zip > /dev/null'
    sh 'unzip -o sonar-scanner-cli-3.2.0.1227-linux.zip > /dev/null'
    // sh 'mv sonar-scanner-cli-3.2.0.1227-linux sqs'
    sh "sonar-scanner-3.2.0.1227-linux/bin/sonar-scanner -e \
        -Dsonar.projectKey=timewindow \
        -Dsonar.sources=. \
        -Dsonar.inclusions='*.go' \
        -Dsonar.exclusions='*_test.go' \
        -Dsonar.tests='.' \
        -Dsonar.test.inclusions='*_test.go' \
        -Dsonar.host.url=${sonarqubeServer} \
        -Dsonar.login=${sonarqubeKey}"
  } 
  stage("Quality Gate") {}
  stage('BranchMerge') {}
  stage('StoreArtifact') {}
}
