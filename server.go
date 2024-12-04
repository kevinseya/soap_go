package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

// Definimos las estructuras para el mensaje SOAP
type HelloRequest struct {
	XMLName xml.Name `xml:"HelloRequest"`
	Name    string   `xml:"name"`
}

type HelloResponse struct {
	XMLName xml.Name `xml:"HelloResponse"`
	Message string   `xml:"message"`
}

// Estructura para el Envelope SOAP
type SOAPEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    SOAPBody `xml:"Body"`
}

// Estructura para el Body SOAP, que contiene el HelloRequest
type SOAPBody struct {
	HelloRequest HelloRequest `xml:"HelloRequest"`
}

// Esta función maneja las solicitudes SOAP
func soapHandler(w http.ResponseWriter, r *http.Request) {
	// Aseguramos que solo aceptamos solicitudes POST
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	// Decodificamos el mensaje SOAP entrante
	var envelope SOAPEnvelope
	decoder := xml.NewDecoder(r.Body)
	err := decoder.Decode(&envelope)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error leyendo el SOAP: %v", err), http.StatusBadRequest)
		return
	}

	// Creamos la respuesta
	res := HelloResponse{
		Message: fmt.Sprintf("¡Hola Mundo, %s!", envelope.Body.HelloRequest.Name),
	}

	// Preparamos la respuesta SOAP
	response := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://www.example.org/hello">
		<soapenv:Header/>
		<soapenv:Body>
			<web:HelloResponse>
				<web:message>%s</web:message>
			</web:HelloResponse>
		</soapenv:Body>
	</soapenv:Envelope>`, res.Message)

	// Enviamos la respuesta SOAP
	w.Header().Set("Content-Type", "text/xml")
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/hello", soapHandler) // Manejamos las solicitudes en la ruta /hello
	fmt.Println("Servidor SOAP escuchando en http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Iniciamos el servidor en el puerto 8080
}