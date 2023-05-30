package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type cliente struct {
	name  string
	ident string
}

type User struct {
	ID    int
	Name  string
	Ident string
}

type envio struct {
	TypeProduct       string
	Amount            string
	RegistrationDate  string
	DateDelivery      string
	Store             string
	Price             string
	VehicleIdentifier string
	GuideNumber       string
	Free              string
}

type envio2 struct {
	TypeProduct       string
	Amount            string
	RegistrationDate  string
	DateDelivery      string
	Store             string
	Price             string
	VehicleIdentifier string
	GuideNumber       string
	Free              string
	ClienteId         string
	TipoTransporteId  string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/add_maritimo", add_maritimo)
	http.HandleFunc("/add_terrestre", add_terrestre)
	http.HandleFunc("/visua", getUsers)

	fmt.Print("el servidor esta iniciado ")
	fmt.Print(" miralo en el http://localhost:3000/")
	http.ListenAndServe(":3000", nil)
}

func index(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, nil)
	}
}

func add_terrestre(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template, err := template.ParseFiles("templates/add_maritimo.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(rw, nil)
		}
		return
	} else if r.Method == "POST" {

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

		/// veiculo
		match, _ := regexp.MatchString("^[a-zA-Z]{3}[0-9]{3}$", VehicleIdentifier)

		if match == false {
			http.Error(rw, "el identificador del vehículo  no es valido ", http.StatusInternalServerError)
			return
		}

		match, _ = regexp.MatchString("^[A-Za-z0-9]{10}$", GuideNumber)
		if match == false {
			http.Error(rw, "el numero de guia no es valido ", http.StatusInternalServerError)
			return
		}

		num, _ := strconv.Atoi(Amount)
		PriceInt, _ := strconv.Atoi(Price)

		DescAux := 0.0
		if num > 10 {
			DescAux = float64(PriceInt) * 0.95
		}

		Free := strconv.FormatFloat(DescAux, 'f', -1, 64)

		// Almacenar los datos en la estructura correspondiente
		cliente := cliente{
			name:  name,
			ident: ident,
		}

		envio := envio{
			TypeProduct:       TypeProduct,
			Amount:            Amount,
			RegistrationDate:  RegistrationDate,
			DateDelivery:      DateDelivery,
			Store:             Store,
			Price:             Price,
			VehicleIdentifier: VehicleIdentifier,
			GuideNumber:       GuideNumber,
			Free:              Free,
		}

		validateFields(cliente, envio)

		err4 := saveFormDataWare(envio, cliente, 2)

		if err4 != nil {
			fmt.Print("error de carga 1 : ", err4.Error())
			fmt.Print("\n")
			http.Error(rw, "Error al guardar los datos en la base de datos error 2", http.StatusInternalServerError)
			return
		}

		//fmt.Fprint(rw, "Datos guardados exitosamente en la base de datos")
		http.Redirect(rw, r, "/add_terrestre", http.StatusFound)
		return
	}
	http.Error(rw, "Método no permitido", http.StatusMethodNotAllowed)
}

func add_maritimo(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template, err := template.ParseFiles("templates/add_maritimo.html")
		if err != nil {
			panic(err)
		} else {
			template.Execute(rw, nil)
		}
		return
	} else if r.Method == "POST" {

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

		///
		match, _ := regexp.MatchString("^[a-zA-Z]{3}[0-9]{4}$", VehicleIdentifier)

		if match == false {
			http.Error(rw, "el identificador del vehículo  no es valido ", http.StatusInternalServerError)
			return
		}

		match, _ = regexp.MatchString("^[A-Za-z0-9]{10}$", GuideNumber)
		if match == false {
			http.Error(rw, "el numero de guia no es valido ", http.StatusInternalServerError)
			return
		}

		num, _ := strconv.Atoi(Amount)
		PriceInt, _ := strconv.Atoi(Price)

		DescAux := 0.0
		if num > 10 {
			DescAux = float64(PriceInt) * 0.97
		}

		Free := strconv.FormatFloat(DescAux, 'f', -1, 64)

		// Almacenar los datos en la estructura correspondiente
		cliente := cliente{
			name:  name,
			ident: ident,
		}

		envio := envio{
			TypeProduct:       TypeProduct,
			Amount:            Amount,
			RegistrationDate:  RegistrationDate,
			DateDelivery:      DateDelivery,
			Store:             Store,
			Price:             Price,
			VehicleIdentifier: VehicleIdentifier,
			GuideNumber:       GuideNumber,
			Free:              Free,
		}

		validateFields(cliente, envio)

		err4 := saveFormDataWare(envio, cliente, 1)

		if err4 != nil {
			fmt.Print("error de carga 1 : ", err4.Error())
			fmt.Print("\n")
			http.Error(rw, "Error al guardar los datos en la base de datos error 2", http.StatusInternalServerError)
			return
		}

		//fmt.Fprint(rw, "Datos guardados exitosamente en la base de datos")
		http.Redirect(rw, r, "/add_terrestre", http.StatusFound)
		return
	}
	http.Error(rw, "Método no permitido", http.StatusMethodNotAllowed)
}

func saveFormDataClient(data cliente) error {
	// Configurar la conexión a la base de datos
	db, err := sql.Open("mysql", "root:berdugo13@tcp(localhost:3306)/db_empresa")
	if err != nil {
		return err
	}
	defer db.Close()
	// Preparar la consulta SQL
	query := "INSERT INTO cliente (NameClient, IdClient) VALUES (?, ?);"

	// Ejecutar la consulta SQL
	_, err = db.Exec(query, data.name, data.ident)
	if err != nil {
		fmt.Print("error de carga 0 ", err.Error())
		return err
	}

	return nil
}

func saveFormDataWare(data envio, Cl cliente, typeTransport int) error {
	// Configurar la conexión a la base de datos
	db, err := sql.Open("mysql", "root:berdugo13@tcp(localhost:3306)/db_empresa")
	if err != nil {
		fmt.Print("error de conexion")
		return err

	}
	defer db.Close()
	// Preparar la consulta SQL
	id, err := buscarIDPorValor(db, "cliente", "IdClient", Cl.ident)

	if err != nil {
		// verifica errores
		fmt.Println(err)
	} else if id == 0 {
		/// se guarda el cliente si no exite y verifica que no exiatan erroes
		err2 := saveFormDataClient(Cl)
		if err2 != nil {
			return fmt.Errorf("error en la base de datos ")
		}
	}

	// Preparar la consulta SQL
	id2, err2 := buscarIDPorValor(db, "cliente", "IdClient", Cl.ident)

	if err2 != nil {
		return err2
	}

	/// se cargan el resto de datos
	query := "INSERT INTO planentrega (typeProduct,amount,RegistrationDate,DateDelivery,Store,Price,VehicleIdentifier,GuideNumber,reduction,Cliente_Id,TipoTransporte_Id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?,?, ?)"

	_, err = db.Exec(query, data.TypeProduct, data.Amount, data.RegistrationDate, data.DateDelivery, data.Store, data.Price, data.VehicleIdentifier, data.GuideNumber, data.Free, id2, typeTransport)

	if err != nil {
		fmt.Print("error de carga 1 : ", err.Error())
		fmt.Print("\n")
		return err
	}
	return nil
}

func buscarIDPorValor(db *sql.DB, tabla string, columna string, valorBuscado string) (int, error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE %s = ?", tabla, columna)
	var id int
	err := db.QueryRow(query, valorBuscado).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return id, nil
}

func buscarPorValor(db *sql.DB, tabla string, columna string, valorBuscado int, columna2 string) (string, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", columna2, tabla, columna)
	var valor string
	err := db.QueryRow(query, valorBuscado).Scan(&valor)
	if err != nil {
		if err == sql.ErrNoRows {
			return "0", nil
		}
		return "0", err
	}
	return valor, nil
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	//// variables auciliares para reistro de la id
	var IDPedido int
	var Aux1 int
	var Aux2 int

	db, err := sql.Open("mysql", "root:berdugo13@tcp(localhost:3306)/db_empresa")
	if err != nil {
		return
	}
	defer db.Close()
	// Consultar la base de datos
	rows, err := db.Query("SELECT * FROM cliente ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// añade a la lista el plan de entrega
	rows, err = db.Query("SELECT * FROM planentrega ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	envios := []envio2{}
	for rows.Next() {
		var envio2 envio2
		err := rows.Scan(&IDPedido, &envio2.TypeProduct, &envio2.Amount, &envio2.RegistrationDate, &envio2.DateDelivery, &envio2.Store, &envio2.Price, &envio2.VehicleIdentifier, &envio2.GuideNumber, &envio2.Free, &Aux1, &Aux2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err := buscarPorValor(db, "cliente", "Id", Aux1, "NameClient")

		if err != nil {
			return
		}
		envio2.ClienteId = id

		id, err = buscarPorValor(db, "tipotransporte", "Id", Aux2, "typeTransport")

		if err != nil {
			return
		}

		envio2.TipoTransporteId = id

		envios = append(envios, envio2)
	}

	// Cargar la plantilla HTML
	tmpl, err := template.ParseFiles("templates/visualize.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Renderizar la plantilla con los datos
	///err = tmpl.Execute(w, users)
	err = tmpl.Execute(w, envios)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func validateFields(user cliente, packa envio) error {
	if user.name == "" {
		return fmt.Errorf("El campo 'name' es requerido")
	}

	if user.ident == "" {
		return fmt.Errorf("El campo 'identidad' es requerido")
	}

	if packa.TypeProduct == "" {
		return fmt.Errorf("El campo 'tipo de producto ' es requerido")
	}

	if packa.Amount == "" {
		return fmt.Errorf("El campo 'cantidad' es requerido")
	}

	if packa.RegistrationDate == "" {
		return fmt.Errorf("El campo 'fecha de registro' es requerido")
	}

	if packa.DateDelivery == "" {
		return fmt.Errorf("El campo 'fecha de entrega' es requerido")
	}

	if packa.Store == "" {
		return fmt.Errorf("El campo 'bodega' es requerido")
	}

	if packa.Price == "" {
		return fmt.Errorf("El campo 'bodega' es requerido")
	}

	if packa.VehicleIdentifier == "" {
		return fmt.Errorf("El campo 'identificador del vehiculo' es requerido")
	}

	if packa.GuideNumber == "" {
		return fmt.Errorf("El campo 'numero de guia' es requerido")
	}

	return nil
}
