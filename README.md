# minicloud-openapi-client-go

Example OpenApi based Client-Go for MiniCloud REST API

## Generate the OpenAPI based client

1) Navigate to the REST API code repo
2) Check and generate the Swagger based code

```bash
swag init
```

3) Generate the OpenAPI based client code

```bash
openapi-generator-cli generate -i docs/swagger.yaml -g go -o ../minicloud-openapi-client-go/minicloud --additional-properties=packageName=minicloud
```

4) Generate the go.mod

```bash
go mod init github.com/odeeka/minicloud-openapi-client-go
```

5) Write example go code to call the client
