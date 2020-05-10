# Petssenger

Resolução do sistema proposto em [hashlab/hiring](https://github.com/hashlab/hiring/blob/5ae82743d1afd7f741d59ee63ffa8149ffa12660/challenges/pt-br/backend-finance-challenge.md)

### Getting Started

- Em um sistema POSIX, certifique-se que o [docker](https://docs.docker.com/install/) e o [docker-compose](https://docs.docker.com/compose/install/) esteja instalado corretamente
- Na pasta raíz da aplicação execute o comando `docker-compose up --build -d`
- Com **todos** os containeres online, execute o comando `yarn populate` para preencher os bancos de dados com os dados iniciais
- Por fim, teste a aplicação executando o comando `yarn test`

### HTTP Routes

- POST http://localhost:3000/ride/estimate - Estima o preço de uma corrida baseado na cidade escolhida, na distância e no tempo da viagem. A estimativa possui uma validade de apenas **20 segundos** por causa da taxa dinâmica da cidade, nesse tempo a estimativa pode ser confirmada, ou apenas expirar.

```sh
$ curl --location --request POST 'http://localhost:3000/ride/estimate' --header 'Content-Type: application/json' --header 'X-User-ID: 08842beb-a4fc-4cb2-9f87-d80f1a2d5045' --data-raw '{ "city": "RIO_DE_JANEIRO", "distance": 7.23, "time": 19.6 }'
```

- POST http://localhost:3000/ride/perform - Confirma a última viagem estimada pelo usuário. Ao confirmar, a taxa dinâmica da cidade é incrementada e só volta ao seu valor anterior após **5 minutos**.

```sh
$ curl --location --request POST 'http://localhost:3000/ride/perform' --header 'X-User-ID: 08842beb-a4fc-4cb2-9f87-d80f1a2d5045'
```

- POST http://localhost:3002/user - Cria um usuário através de um dado email.

```sh
curl --location --request POST 'http://localhost:3002/user' --header 'Content-Type: application/json' --data-raw '{ "email": "next@petssenger.com" }'
```

### gRPC Functions

- _GetPricingFeesByCity_ - Retorna as informações de precificação (taxa base, taxa de distância, taxa de tempo e taxa de serviço) de uma corrida para uma determinada cidade

- _GetDynamicFeesByCity_ - Retorna a taxa dinâmica de uma cidade. Não foi fundida à função acima por questões de _caching_

- _IncreaseDynamicFeesByCity_ - Incrementa a taxa dinâmica de uma cidade

- _AuthUser_ - Determina se um usuário existe ou não através de seu UUID (X-User-ID)

### Considerações

- A conexão do banco de dados nos serviços escritos em Golang, estão num contexto global (dentro do pacote _models_) pois não consegui repassa-lá utilizando _closure_, já que nas funções executadas pelo [_taskq_](https://github.com/vmihailenco/taskq), a conexão sofria um erro de _dereference_.
