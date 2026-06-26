set -e 

if [ -z $1 ]; then
  echo "Please provide lab number for image tag"
  exit 1
fi

lab=$1

docker build -t log_output/producer:${lab} ./log_output/producer/
docker build -t log_output/consumer:${lab} ./log_output/consumer/
docker build -t pingpong:${lab} ./pingpong

k3d image import log_output/producer:${lab}
k3d image import log_output/consumer:${lab}
k3d image import pingpong:${lab}

kubectl delete all --all

# remember to edit each deployment.yml with updated tag
kubectl apply -f ./persistant_volume
kubectl apply -f ./log_output/manifests
kubectl apply -f ./pingpong/manifests