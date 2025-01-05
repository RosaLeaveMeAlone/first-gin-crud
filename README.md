### First Gin CRUD

Para correr este proyecto en local, es necesario seguir los siguientes pasos:

1. Copiar el archivo .env.example y renombrarlo como .env

```
cp .env.example .env
```

2. Asegurarse de que este corriendo un servicio de base de datos de postgresql, en caso de no tener un servicio corriendo se puede levantar uno usando docker compose para levantar el contenedor

```
docker compose up -d
```
3. Correr el siguiente comando
```
go mod tidy
```
4. Levantar el proyecto 
```
go run main.go
```