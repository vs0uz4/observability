# EXTRAS

## Health Check

Implementei um `endpoint` extra na API de nossos serviços onde teremos como obter informações a respeito da saúde dos mesmos, o famoso `health check`. O mesmo atenderá no seguinte endereço abaixo:

```plain-text
http://localhost:{PORTA_DO_SERVICO}/health
```

> Como teremos dois microsserviços/serviços, então o `endpoint` de `health-check` irá funcionar nas portas `8000` e `8001` respectivamente as portas em que os microsserviços serão levantados a princípio.

E terá como resposta um `payload` bem rico, contendo informações sobre:

- Uso de cpu;
- Uso de memória;
- Horário do servidor;
- Tempo de funcionamento;
- Tempo da requisição.

Abaixo segue um exemplo de como será disponibilizado o `payload` na API

```json
{
    "cpu": {
        "cores": 4,
        "percent_used": [
            37.7,
            34.1,
            30,
            26.7,
        ]
    },
    "memory": {
        "total": 17179869184,
        "used": 13563052032,
        "free": 153239552,
        "available": 3616817152,
        "percent_used": 78.9
    },
    "uptime": "4.990172625s",
    "duration": "342.917µs",
    "status": "pass",
    "message": "Alive and kicking!",
    "time": "2024-12-10T16:56:00-03:00"
}
```

> [!TIP]
> Alguns pontos a serem destacados sobre o Health Check
>
> Caso ocorra falha na obtenção de informações, sejam elas de CPU ou Memória seus respectivos dados serão informados como vázio
> o `status` será retornado como `fail` e o campo `message` será exibido como `Still alive, but not kicking!`, caso contrário todas
> as informações irão preenchidas e o `status` e `message` serão retornados conforme o modelo apresentado logo acima.

### Suite de Testes

Para executar as suites de testes de ambos os projetos, estando na pasta raiz, basta executar o seguinte comando:

```shell
❯ make tests
```

> As suítes de testes de ambos os projetos serão executadas e relatórios contendo informações tanto referente a execução quanto a taxa de cobertura dos testes serão apresentadas no console. Abaixo seguem os relatórios gerados mais recentes que disponibilizei.

Você deverá ver em sua tela, algo parecido com as informações apresentadas abaixo:

```plain-text
❯ make tests
Running tests for InputValidate...
        github.com/vs0uz4/observability/ms-inputvalidate/cmd/inputvalidate              coverage: 0.0% of statements
        github.com/vs0uz4/observability/ms-inputvalidate/doc/swagger                    coverage: 0.0% of statements
        github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/utils           coverage: 0.0% of statements
?       github.com/vs0uz4/observability/ms-inputvalidate/internal/infra/web/health      [no test files]
?       github.com/vs0uz4/observability/ms-inputvalidate/internal/service/contract      [no test files]
?       github.com/vs0uz4/observability/ms-inputvalidate/internal/usecase/contract      [no test files]
...
```

Caso você queira executar as suítes de testes apenas de um dos projetos, você também pode. Para isto basata você executar um dos comandos disponíveis abaixo:

- Executando suíte de testes do ms-inputvalidate

```shell
❯ make test-inputvalidate
```

- Executando suíte de testes do ms-weatherzip

```shell
❯ make test-weatherzip
```

Abaixo seguem os relatórios de cobertura mais atuais que foram executados.

- [ms-inputvalidate](../ms-inputvalidate/doc/COVERAGE.md)
- [ms-weatherzip](../ms-weatherzip/doc/COVERAGE.md)

## Documentação da API com Swagger

Adicionei também a documentação da API utilizando o Swagger, desta forma os `endpoints` de ambos os serviços podem ser testados também diretamente do navegador, através das suas respectivas rotas. Para acessar a documentação dos serviços, basta acessa as seguintes URL's abaixo:

- http://localhost:8000/swagger/ - Para acessar a documentação do serviço A (Input Validate);
- http://localhost:8001/swagger/ - Para acessar a documentação do serviço B (Weather Zip).

## Makefile

Adicionei um `makefile` para automatizar algumas tarefar, tais como:

- setup dos serviços;
- geração de documentações;
- excução de suite de testes;

### Setup dos Serviços

Será realizada a cópia dos arquivos `.env.example` para `.env` na para `../cmd/SERVICO/` fazendo assim com que o serviço possa carregar as configurações necessárias, como **URL**, **PORTAS** e etc.

### Geração de Documentações

Será gerada e salva na pasta `doc` de cada serviço toda a sua documentação com Swagger UI. A mesma será salva nos formatos `.json` e `.xml` por padrão.

### Execução das Suítes de Testes

Serão executadas ambas as súites de testes, e gerado relatório de cobertura de código, o mesmo será salvo na pasta `report` dentro da pasta de cada respectivo serviço.
