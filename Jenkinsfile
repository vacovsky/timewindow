import calpoly.golang.Go
import calpoly.sonarqube.SonarQube
import calpoly.artifact.Artifact


def go = new Go(env, steps)
def art = new Artifact(env, steps)
def sq = new SonarQube(env, steps)

node {

    stage('Preparation') {
        git repo
    }

    stage('Install Go and Dependencies') {
        def deps = [
            
        ] as String[]

        go.InstallDeps(deps) 
    }

    stage('Run Unit Tests') {
        go.Test()
    }

    stage('Static Code Analysis') {
        String inclusions = '*.go'
        String exclusions = 'go/*,*_test.go' // ,*_test.go,**/*_test.go'
        String testInclusions = '*_test.go'
        String testExclusions = 'go/*'
        sq.GoScan(inclusions, exclusions, testInclusions, testExclusions, project)
    }
    
    stage("Quality Gate") {

    }

    stage('Branch Merge') {

    }
}
