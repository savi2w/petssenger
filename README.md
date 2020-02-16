# Petssenger

Resolução do sistema proposto em [hashlab/hiring](https://github.com/hashlab/hiring/blob/5ae82743d1afd7f741d59ee63ffa8149ffa12660/challenges/pt-br/backend-finance-challenge.md)

### Arquitetura

A arquitetura do produto foi desenhada para permitir uma fácil escalabilidade e substituição.
Ao estabelecer uma separação entre os serviços houveram ganhos consideráveis de performance e de isolamento de responsabilidades, o que torna o sistema mais conciso, seguro e performático em comparação a sistemas monolíticos. Este benefício pôde ser constatado **ainda no desenvolvimento**, quando houve a necessidade de criar mais microserviços do que o previsto[¹](https://imgur.com/a/llpJ6Ir). Tudo ficou muito "plugável".

### Getting Started

- Em um sistema POSIX, certifique-se que o [docker](https://docs.docker.com/install/) e o [docker-compose](https://docs.docker.com/compose/install/) esteja instalado corretamente
- Na pasta raíz da aplicação execute o comando `docker-compose up --build`
- Com **todos** os containeres online, execute o comando `yarn populate` (ou execute o script `./POPULATE.sh`) para preencher os bancos de dados com os dados iniciais
- Por fim, teste a aplicação executando o comando `yarn test` (ou execute o script `./TEST.sh`)

### HTTP Routes

- POST _http://localhost:3001/ride/estimate_ - Estima o preço de uma corrida baseado na cidade escolhida, na distância e no tempo da viagem. A estimativa possui uma validade de apenas **20 segundos** por causa da taxa dinâmica da cidade, nesse tempo a estimativa pode ser confirmada, ou apenas expirar. Exemplo de requisição: `curl --location --request POST 'http://localhost:3001/ride/estimate' --header 'Content-Type: application/json' --header 'X-User-ID: 08842beb-a4fc-4cb2-9f87-d80f1a2d5045' --data-raw '{ "city": "RIO_DE_JANEIRO", "distance": 7.23, "time": 19.6 }'`
- POST http://localhost:3001/ride/perform - Confirma uma viagem estimada pelo usuário. Ao confirmar, a taxa dinâmica da cidade é incrementada e só volta ao seu valor anterior após **5 minutos**. Exemplo de requisição: `curl --location --request POST 'http://localhost:3001/ride/perform' --header 'X-User-ID: 08842beb-a4fc-4cb2-9f87-d80f1a2d5045'`
- POST http://localhost:3002/user - Cria um usuário através de um dado email. Exemplo de requisição: `curl --location --request POST 'http://localhost:3002/user' --header 'Content-Type: application/json' --data-raw '{ "email": "next@petssenger.com" }'`

### gRPC Functions

- GetPricingFeesByCity - Retorna as informações de precificação (taxa base, taxa de distância, taxa de tempo e taxa de serviço) de uma corrida para uma determinada cidade
- GetDynamicFeesByCity - Retorna a taxa dinâmica de uma cidade. Não foi fundida à função acima por questões de _caching_
- IncreaseDynamicFeesByCity - Incrementa a taxa dinâmica de uma cidade
- AuthUser - Determina se um usuário existe ou não através de seu UUID (X-User-ID)

### Considerações

- Foi a minha primeira vez utilizando Golang em algo palpável, se cometi algum crime escrevendo os códigos, abra uma issue 😄, o mesmo vale para melhorias e sugestões.
- A conexão do banco de dados nos serviços escritos em Golang, estão num contexto global (dentro do pacote _models_) pois não consegui repassa-lá utilizando _Closure_, já que nas funções executadas pelo [_taskq_](https://github.com/vmihailenco/taskq), a conexão sofria um erro de _dereference_.
- Nas _models_ que consultam um microserviço gRPC do serviço _ride_, foi necessário criar uma interface para as funções assíncronas pois a função `bluebird.promisifyAll` não consegue inferir o tipo das novas funções criadas por ele[²](https://github.com/Microsoft/TypeScript/issues/8685#issuecomment-240201897).
