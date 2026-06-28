set -e 

if [ -z $1 ]; then
  echo "Please provide lab number for image tag"
  exit 1
fi

lab=$1

docker build -t todo_app/frontend:${lab} ./todo_app/frontend/
docker build -t todo_app/backend:${lab} ./todo_app/backend/

k3d image import todo_app/frontend:${lab}
k3d image import todo_app/backend:${lab}

kubectl delete all --all

# remember to edit each deployment.yml with updated tag
kubectl apply -f ./persistant_volume
kubectl apply -f ./todo_app/manifests
