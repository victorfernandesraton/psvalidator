# psvalidator
Projeto no qual é implementado um validador de senha conforme requisitos informados no [LINK DO ARQUIVO]

## Arquitetura
Projeto foi implementado conforme os principios de SOLID, seguindo como base o Pattern CQRS, o qual separa a leitrua de escrita de dados, assim temos as seguintes camadas

### Domain
Onde define-se as estruturas de dôminio de negócio como entidades e funções básicas

### Command
Camada onde se encontra as ações que o sistema é capaz de executar conforme uma entrada de dados

### Query (Não usada)
Camada de leitura de dados que consegue abstrair as responsabilidades e características unicas dos serviços de persistência tais como banco de dados ou sistema de arquivos.

### Infra
Onde abstraimos as unidades do sistema responsavél pela comunicação, aqui fica por exemplo as resoluções dos controllers http, visto que os mesmos acessa as demais camadas de forma indireta por meio de dependências

## Quick start
Este projeto utiliza do docker e do docker compose para que este possa ser reproduzido por containeres, assim por meio dos comandos abaixo é possivél reproduzir o anbiente sem se preocupar com coisas como versão de linguagem e afins

### Build de containers
```
docker-compose build
```
### Iniciar servidor
```
docker-compose up --build -d api
```
### Interromper containeres
```
docker-compose stop
```
### Executar testes unitários
```
docker-compose run --rm unit-test
```