Desafio Técnico - Estágio DevOps
Esse repositório tem a solução completa para o desafio de estágio DevOps. A ideia aqui é criar um serviço em Golang, colocar ele dentro de contêineres e deixar o processo de colocar no ar (deploy) automático.

Como funciona a estrutura do projeto
O projeto usa contêineres Docker que rodam juntos com ajuda do Docker Compose.

http-server-projeto-korp: é um serviço feito em Go que responde com JSON na porta 8080.

nginx: é um contêiner NGINX que funciona como uma "ponte", recebendo os acessos na porta 80 e repassando para o serviço Go.

projeto-korp-network: é uma rede Docker (do tipo bridge) que deixa os dois contêineres conversarem entre si.

Automatizando com Ansible
Todo o processo de preparar o ambiente e subir o projeto foi feito de forma automática com o Ansible. O arquivo principal é o ansible/playbook.yml, que cuida de:

Instalar o Docker e o que mais for necessário.

Configurar o repositório oficial do Docker.

Copiar os arquivos do projeto para a máquina que vai rodar.

Apagar os contêineres antigos (se tiver) pra garantir que tudo comece do zero.

Rodar o docker compose pra montar as imagens e iniciar os serviços.

Fazer um teste no final pra ver se o serviço tá funcionando direitinho.

Como rodar
O que você precisa ter instalado:
Ubuntu 24.04 (ou parecido)

Ansible (sudo apt install ansible)

Coleção do Ansible pra trabalhar com Docker:
ansible-galaxy collection install community.docker

Passo a passo pra rodar:
Clone o repositório:

bash
Copiar
Editar
git clone <link-do-repo>
Vá até a pasta do Ansible:

bash
Copiar
Editar
cd ansible/
Rode o playbook:

bash
Copiar
Editar
ansible-playbook -i inventory.ini playbook.yml
Depois disso, o serviço vai estar rodando no endereço:
http://localhost:80

Organização das pastas
arduino
Copiar
Editar
.
├── ansible/
│   ├── inventory.ini
│   └── playbook.yml
├── app/
│   ├── Dockerfile
│   └── main.go
├── docker-compose.yml
├── nginx-config/
│   └── http-server-projeto-korp.conf
└── README.md
