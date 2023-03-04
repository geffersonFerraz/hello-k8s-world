# Cluster privado K8S

## Tutorial de um novato no mundo k8s para quem é ainda mais novato que eu <3

Considerando um cluster cru, vamos prosseguir a configuração até a etapa de que fazemos um deploy de uma aplicação.

## Comando kubectl, mas o que é isso?

Utilizamos o kubectl para auxiliar na configuração do cluster. Se quiser maiores detalhes ou tenha curiosidade veja a [documentação](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/).

- Instale o kubectl com os seguintes comandos:

```console
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
```

```console
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

```console
kubectl version --short
```

- O resultado esperado é:

```console
Client Version: v1.26.0
Kustomize Version: v4.5.7
Server Version: v1.22.9
```

Agora, para que o comando funcione corretamente, voce precisa criar ter o arquivo `config` dentro do diretório `~/.kube/`. Maiores detalhes de como configura-lo estão [aqui](https://kubernetes.io/pt-br/docs/concepts/configuration/organize-cluster-access-kubeconfig/).

### Testando se o `kubectl` está funcionando

- No console digite `kubectl get nodes` para ver os nodes da sua aplicação, o resultado esperado é semelhante a este:

```console
NAME                                  STATUS   ROLES                  AGE   VERSION
cluster-geff-ws-control-plane   Ready    control-plane,master   33h   v1.22.9
cluster-geff-ws-control-plane   Ready    control-plane,master   33h   v1.22.9
cluster-geff-ws-control-plane   Ready    control-plane,master   33h   v1.22.9
nodep1                          Ready    <none>                 31h   v1.22.9
```

- Se o resultado nao for semelhante, certifique-se que os dados do `config` estão corretos e também valide se o seu cluster está criado/saudavel.

## Instalando o Dashboard K8S

- Instale o dashboard k8s com o seguinte comando, mais detalhes [nessa doc](https://github.com/kubernetes/dashboard#install):

```console
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
```

- Verifique se funcionou:

```console
kubectl proxy
```

- Abra no navegador:

[http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login](http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login)

- Se nao abrir uma tela pedindo o Token, deu ruim... Pesquise onde errou e refaça tudo se necessário *Se deu certo, continue.*

## Configurando novo usuario para admin no dashboard

- Crie o arquivo `create_role_dashboard.yaml` e adicione o conteudo:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: default
```

- Crie o arquivo `create_user_dashboard.yaml` e adicione o conteudo:

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: default
```

- Após ter criado os arquivos, com o comando mágico `kubectl apply -f xxx.yaml` vamos executá-los:

```console
➜  kubectl apply -f create_role_dashboard.yaml
clusterrolebinding.rbac.authorization.k8s.io/admin-user created

➜  kubectl apply -f create_user_dashboard.yaml 
serviceaccount/admin-user created
```

- Feito isto, vamos gerar o token para acessar o dashboard:

```console
kubectl -n default create token admin-user
```

Tendo como resultado o token gerado para voce fazer o login no dash:

```console
eyJhbGciOiJSUzI1Nxxxxx..........xxxxxxxxxxnplGS-EOUuPw
```

## Instalando o NGINX e suas parafernalhas, os arquivos de deploys estão no [aqui](https://docs.nginx.com/nginx-ingress-controller/installation/installation-with-manifests/#prerequisites) repo do nginx junto com os detalhes de cada item

1. Configure RBAC

`kubectl apply -f k8s/installing-nginx/ns-and-sa.yaml`

`kubectl apply -f k8s/installing-nginx/rbac.yaml`

`kubectl apply -f k8s/installing-nginx/ap-rbac.yaml`

`kubectl apply -f k8s/installing-nginx/apdos-rbac.yaml`

2. Create Common Resources

`kubectl apply -f k8s/installing-nginx/default-server-secret.yaml`

`kubectl apply -f k8s/installing-nginx/nginx-config.yaml`

`kubectl apply -f k8s/installing-nginx/ingress-class.yaml`

3. Create Custom Resources

`kubectl apply -f k8s/installing-nginx/k8s.nginx.org_virtualservers.yaml`

`kubectl apply -f k8s/installing-nginx/k8s.nginx.org_virtualserverroutes.yaml`

`kubectl apply -f k8s/installing-nginx/k8s.nginx.org_transportservers.yaml`

`kubectl apply -f k8s/installing-nginx/k8s.nginx.org_policies.yaml`

`kubectl apply -f k8s/installing-nginx/k8s.nginx.org_globalconfigurations.yaml`

`kubectl apply -f k8s/installing-nginx/appprotect.f5.com_aplogconfs.yaml`

`kubectl apply -f k8s/installing-nginx/appprotect.f5.com_apusersigs.yaml`

`kubectl apply -f k8s/installing-nginx/appprotect.f5.com_appolicies.yaml`

`kubectl apply -f k8s/installing-nginx/appprotectdos.f5.com_apdoslogconfs.yaml`

`kubectl apply -f k8s/installing-nginx/appprotectdos.f5.com_apdospolicy.yaml`

`kubectl apply -f k8s/installing-nginx/appprotectdos.f5.com_dosprotectedresources.yaml`

`kubectl apply -f k8s/installing-nginx/appprotect-dos-arb.yaml`

`kubectl apply -f k8s/installing-nginx/appprotect-dos-arb-svc.yaml`

`kubectl apply -f k8s/installing-nginx/nginx-ingress.yaml`

4. Get Access to the Ingress Controller

`kubectl apply -f k8s/installing-nginx/loadbalancer.yaml`

5. Get the "ingress ip", ou o seu ip publico para ser utilizado no seu DNS =)

`kubectl get svc nginx-ingress --namespace=nginx-ingress`

## Pré deploy, prerarar o SSL / HTTPS

Instale o `CertManager`:

```console
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.7.1/cert-manager.yaml
```

Crie o arquivo `letsencrypt-production.yaml` com o conteudo:

```yaml
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt-production
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: cert_prd@seumail.com.br
    privateKeySecretRef:
      name: letsencrypt-production
    solvers:
    - http01:
        ingress:
          class: nginx
```

Execute a configuração do LetsEncryp:

```cosole
kubectl apply -f letsencrypt-production.yaml
```

## Deployando aplicação

Execute algumas configurações (sim ainda tem mais)

```cosole
kubectl apply -f k8s/networkpolice-egreess.yaml
```

```cosole
kubectl apply -f k8s/networkpolice-ingress.yaml
```

```cosole
kubectl apply -f k8s/configmap-http2.yaml
```

Crie um arquivo deployment da sua aplicação, no caso é o `geffws-deployment.yml`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: geffws
  name: geffws-deployment
  namespace: geffws
spec:
  replicas: 1
  selector:
    matchLabels:
      app: geffws
  template:
    metadata:
      labels:
        app: geffws
    spec:
      containers:
      - name: geffws
        command:
        - ./geffws
        image: docker.io/geffws/hello-k8s-world:v0.4.3
        imagePullPolicy: IfNotPresent
        resources:
          limits:
            memory: 60Mi
            cpu: 500m
        ports:
        - containerPort: 8083
          name: https
          protocol: TCP
        env:
        - name: GO_ENV
          value: prd
        - name: GIN_MODE
          value: release
        - name: PORT
          value: '8083'
```

Agora execute o deploy:

```console
kubectl apply -f geffws-deployment.yml
```

Crie um arquivo ingress da sua aplicação, no caso é o `geffws-ingress.yml`:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: geffws-ingress
  namespace: geffws
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-production
    acme.cert-manager.io/http01-edit-in-place: "true"
spec:
  rules:
  - host: api.geff.ws
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: geffws-service
            port:
              number: 443
  tls:
  - hosts:
    - api.geff.ws
    secretName: api-geff-ws-domain-cert-prod
```

Agora execute a criação do ingres:

```console
kubectl apply -f geffws-ingress.yml
```

Por ultimo, crie um arquivo service da sua aplicação, no caso é o `geffws-service.yml`:

```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    app: geffws
  name: geffws-service
spec:
  ports:
  - name: https
    protocol: TCP
    port: 443
    targetPort: 8083
  selector:
    app: geffws
```

E finalmente, execute a criação do service:

```console
kubectl apply -f geffws-service.yml
```

Se tudo der certo, sua aplicação estará deployada e pronto para ser acessada.
No meu caso, ficou: `https://api.geff.ws/v1/ping`

```json
{
  "message": "pong"
}
```

Se estiver recebendo erro 502, exclua o service da sua aplicação e execute a criação dele novamente.
Se ainda sim persistir o erro, de um google e qualquer coisa, manda um email [me.ajuda@gefferson.com.br](mailto:me.ajuda@gefferson.com.br) que te passo algumas dicas e podemos nos ajudar.
