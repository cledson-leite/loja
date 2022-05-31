package models

import "loja/db"

type Produto struct {
	Nome, Descricao string
	Preco float64
	Quantidade, id int
}

func GetAll() []Produto {
	db := db.ConectaDB()

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
		// defer executa depois de tudo seja executado
	defer db.Close()

	return produtos
}

func PostProduto(nome, descricao string, preco float64, quantidade int)  {
		db := db.ConectaDB()

		insert, err := db.Prepare(
			"insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)",
		)
		if err != nil {
			panic(err.Error())
		}
		insert.Exec(nome, descricao, preco, quantidade)
		defer db.Close()
}