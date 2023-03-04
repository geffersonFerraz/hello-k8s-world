#!/bin/bash

file="geffws.version" #the file where you keep your string name

current_version=$(cat "$file")        #the output of 'cat $file' is assigned to the $name variable

GIN_MODE=release CGO_ENABLED=0 go build
new_version=$(echo $current_version | awk -F. -v OFS=. 'NF==1{print ++$NF}; NF>1{if(length($NF+1)>length($NF))$(NF-1)++; $NF=sprintf("%0*d", length($NF), ($NF+1)%(10^length($NF))); print}')

echo $new_version > $file
new_image="geffws/hello-k8s-world:v$new_version"
echo "Deploying $new_image"

deployment_file="k8s/geffws-deployment.yml"

sudo docker build -t $new_image .
sudo docker push $new_image

python -c 'import yaml;f=open("k8s/geffws-deployment.yml");ymldoc=yaml.safe_load(f);ymldoc["spec"]["template"]["spec"]["containers"][0]["image"] = "docker.io/'"$new_image"\"';new_ymldoc=yaml.dump(ymldoc, default_flow_style=False, sort_keys=False);f.close();f=open("k8s/geffws-deployment.yml","w");f.write(new_ymldoc)'

kubectl apply -f k8s/geffws-deployment.yml
kubectl apply -f k8s/geffws-ingress.yml
kubectl apply -f k8s/geffws-service.yml