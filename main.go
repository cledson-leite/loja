package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaDB() *sql.DB  {
	conect := "user=postgres dbname=loja password=cursogo host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conect)

	if err !=nil {
		panic(err.Error())
	}

	return db
}

type Produto struct {
	Nome, Descricao string
	Preco float64
	Quantidade, id int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main()  {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000",nil)
}

func index (writer http.ResponseWriter, request *http.Request) {
	db := conectaDB()

	selectAll, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	produtos := []Produto{}

	for selectAll.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAll.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produto.id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)

	}
	templates.ExecuteTemplate(writer, "Index", produtos)

	// defer executa depois de tudo seja executado
	defer db.Close()
}