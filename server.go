package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Creamos el cuerpo XML para la solicitud SOAP
	soapRequest := `<?xml version="1.0" encoding="UTF-8"?>
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://www.example.org/hello">
		<soapenv:Header/>
		<soapenv:Body>
			<web:HelloRequest>
				<web:name>Mundo</web:name>
			</web:HelloRequest>
		</soapenv:Body>
	</soapenv:Envelope>`

	// Hacemos la solicitud SOAP al servidor
	url := "http://localhost:8080/hello"
	req, err := http.NewRequest("POST", url, strings.NewReader(soapRequest))
	if err != nil {
		log.Fatalf("Error creando la solicitud: %v", err)
	}

	// Configuramos los encabezados HTTP
	req.Header.Set("Content-Type", "text/xml")

	// Enviamos la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error al enviar la solicitud SOAP: %v", err)
	}
	defer resp.Body.Close()

	// Leemos la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error al leer la respuesta SOAP: %v", err)
	}

	// Mostramos la respuesta en consola
	fmt.Printf("Respuesta del servidor:\n%s\n", string(body))
}
