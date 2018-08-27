node {
	
	checkout scm

	sh "git rev-parse --short HEAD > commit-id"

	tag = readFile('commit-id').replace("\n", "").replace("r", "")
	appName = "hello-world-app"
	registryHost = "pittcontainerreg.azurecr.io/"
	imageName = "${registryHost}${appName}:${tag}"

	stage('Build') {
		sh "docker build -t ${imageName} ."
	}

	stage('Push') {
		withCredentials([usernamePassword(credentialsId: 'azure_acr', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
			sh "docker login pittcontainerreg.azurecr.io -u $USERNAME -p $PASSWORD"
			sh "docker push ${imageName}"
		}
	}

	stage('Deliver') {
		gitRepo = "https://github.com/twc17/k8s-infrastructure"
		sh 'git config --global user.email "twc17@pitt.edu"'
		sh 'git config --global user.name "Jenkins Automation"'
		git "${gitRepo}"
		sh """
		cat <<EOF > patch.yaml
spec:
  template:
    spec:
      containers:
        - name: hello-world-app
          image: ${imageName}"""
		
		sh 'kubectl patch --local -o yaml -f apps/hello-world-app/deployments/hello-world-app.yaml -p "$(cat patch.yaml)" > output.yaml'
		sh 'mv output.yaml apps/hello-world-app/deployments/hello-world-app.yaml'
		sh 'git add apps/hello-world-app/deployments/hello-world-app.yaml'
		sh """
		git commit -F- <<EOF
Update the hello-world-app application

This commit updates the hello-world-app deployment container image to:

	${imageName}
	https://github.com/twc17/${appName}/${tag}"""
		sh 'git config --global credential.https://github.com.helper /usr/local/bin/hub-credential-helper'
		sh 'git push origin master'
	}
}
