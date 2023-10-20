
## Products RESTful API (BACKEND) 
Aplicación MVC con Go, Gin, Gorm y MySQL
Esta es una aplicación web que utiliza el patrón de diseño MVC (Modelo-Vista-Controlador) para estructurar el código y facilitar el mantenimiento y la escalabilidad. La aplicación está escrita en Go, un lenguaje de programación de alto rendimiento y concurrente. Para el desarrollo web, se utiliza el framework Gin, que ofrece una API sencilla y rápida para crear rutas, manejar peticiones y respuestas, validar datos, etc. Para la persistencia de datos, se utiliza el ORM Gorm, que permite interactuar con diferentes bases de datos relacionales de forma abstracta y eficiente. En este caso, se usa MySQL como base de datos.

## Instrucciones de construcción y ejecución
  Para construir y ejecutar esta aplicación, se necesita tener instalado Go, Gin, Gorm y MySQL en el sistema. También se necesita crear una base de datos llamada store_table y configurar las credenciales de acceso en el archivo config/db.go. Los pasos son los siguientes:

  -Clonar el repositorio de código fuente desde GitHub
  -Entrar en el directorio del proyecto: cd mi_app
  -Instalar las dependencias: go mod tidy
  -Ejecutar las migraciones para crear las tablas en la base de datos: go run migrations/migrate.go
  -Ejecutar la aplicación: go run main.go or go run . 
  -Abrir un navegador web y acceder a la dirección http://localhost:8080/ap1/v1/products para ver el endpoint GetProducts funcionando.