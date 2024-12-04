# API SOAP en Go

Este proyecto implementa un servidor SOAP sencillo utilizando Go. El servidor maneja solicitudes SOAP y responde con un mensaje basado en los datos recibidos en el cuerpo de la solicitud. El cliente realiza solicitudes SOAP a este servidor y recibe una respuesta personalizada.

## Requisitos

- **Go 1.23.4** o superior.

## Instalación

1. **Clona el repositorio** y navega al directorio del proyecto:

    ```bash
    git clone <https://github.com/kevinseya/soap_go.git>
    ```

2. **Compila y ejecuta el servidor**:

    ```bash
    go run main.go
    ```

   El servidor SOAP estará disponible en `http://localhost:8080/hello`.

## Uso

### Estructura del Proyecto

- **`main.go`**: Contiene el servidor SOAP que maneja las solicitudes en la ruta `/hello`.
- **Cliente SOAP**: Realiza una solicitud SOAP al servidor y muestra la respuesta.

### 1. Solicitud SOAP del cliente

El cliente realiza una solicitud POST con un cuerpo XML que contiene el nombre. El cuerpo de la solicitud SOAP se ve de la siguiente manera:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://www.example.org/hello">
    <soapenv:Header/>
    <soapenv:Body>
        <web:HelloRequest>
            <web:name>Mundo</web:name>
        </web:HelloRequest>
    </soapenv:Body>
</soapenv:Envelope>
```
### 2. Respuesta del servidor

El servidor responde con un mensaje SOAP que contiene una respuesta personalizada, como:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://www.example.org/hello">
    <soapenv:Header/>
    <soapenv:Body>
        <web:HelloResponse>
            <web:message>¡Hola Mundo, Mundo!</web:message>
        </web:HelloResponse>
    </soapenv:Body>
</soapenv:Envelope>
```
