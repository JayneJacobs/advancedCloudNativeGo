# advancedCloudNativeGo
minikube start
  ls
  kubectl apply -f simplek8serverdeployment.yaml
  kubectl apply -f simplek8serverservices.yaml
  kubectl apply -f simplek8configmap.yaml
  kubectl apply -f simplek8client.yaml
  kubectl get pods