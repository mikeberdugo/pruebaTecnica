package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Teste() {
	http.HandleFunc("/add_maritimo_test", formHandler)
	http.HandleFunc("/add_terrestre_test", formHandler)
	http.ListenAndServe(":8080", nil)
}

func TestFormHandler1(t *testing.T) {
	formData := strings.NewReader("name=alex&ident=12345&typeProduct=carga&amount=12&RegistrationDate=2023-05-10&DateDelivery=2023-05-10&Store=cali&Price=12300&VehicleIdentifier=ABC1234&GuideNumber=1234567890")
	req := httptest.NewRequest("POST", "/add_maritimo", formData)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(formHandler)

	handler.ServeHTTP(recorder, req)

	expectedBody := "¡Formulario enviado!"
	if recorder.Body.String() != expectedBody {
		t.Errorf("El cuerpo de la respuesta no coincide. Se esperaba '%s', se obtuvo '%s'", expectedBody, recorder.Body.String())
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	ident := r.FormValue("ident")
	// datos
	TypeProduct := r.FormValue("typeProduct")
	Amount := r.FormValue("amount")
	RegistrationDate := r.FormValue("RegistrationDate")
	DateDelivery := r.FormValue("DateDelivery")
	Store := r.FormValue("Store")
	Price := r.FormValue("Price")
	VehicleIdentifier := r.FormValue("VehicleIdentifier")
	GuideNumber := r.FormValue("GuideNumber")

	fmt.Fprintf(w, "¡Formulario enviado!\n")
	fmt.Printf("Datos recibidos:\nNombre: %s", name)
	fmt.Printf("Datos recibidos:\nide: %s", ident)
	fmt.Printf("Datos recibidos:\nNombre: %s", TypeProduct)
	fmt.Printf("Datos recibidos:\nNombre: %s", Amount)
	fmt.Printf("Datos recibidos:\nNombre: %s", RegistrationDate)
	fmt.Printf("Datos recibidos:\nNombre: %s", DateDelivery)
	fmt.Printf("Datos recibidos:\nNombre: %s", Store)
	fmt.Printf("Datos recibidos:\nNombre: %s", Price)
	fmt.Printf("Datos recibidos:\nNombre: %s", VehicleIdentifier)
	fmt.Printf("Datos recibidos:\nNombre: %s", GuideNumber)

	// Aquí puedes realizar la lógica de manejo del formulario, como validar los datos, almacenarlos, etc.

}
