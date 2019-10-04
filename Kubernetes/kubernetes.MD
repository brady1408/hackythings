
# My Kubernetes Cheat Sheet

## Kubectl delete

### Delete all pods in a service

    kubectl get pods --no-headers=true | awk '/myServiceName/ {print $1}' | xargs kubectl delete pods

remember to replace the service above with the service you want to delete the pods for, this can be a partial name
