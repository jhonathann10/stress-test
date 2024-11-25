# stress-test
Este serviço visa criar um CLI em Go para realizar testes em um serviço WEB.

## Requisitos
- Go
- Docker
- A URL deve ter o protocolo http ou https.

### Parametros
- `--url`: URL do serviço que será testado.
- `--requests`: Quantidade de requisições que serão realizadas.
- `--concurrency`: Quantidade de requisições concorrentes.

### Execução
```shell
docker build -t stress-test:latest .
```

#### Exemplo de execução
```shell
docker run --rm stress-test:latest --url=http://google.com --requests=1000 --concurrency=100
```