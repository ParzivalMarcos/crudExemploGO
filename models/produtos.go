package models

import (
	"crudExemploGO/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaProdutos() []Produto {
	db := db.ConectaBanco()

	selectProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	defer db.Close()
	return produtos
}

func InsereNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()

	insertProduto, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertProduto.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaBanco()

	deleteProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduto.Exec(id)

	defer db.Close()

}

func EditaProduto(id string) Produto {
	db := db.ConectaBanco()

	selectProduto, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for selectProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Quantidade = quantidade
		produto.Preco = preco

	}

	defer db.Close()
	return produto

}

func UpdateProduto(id string, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()

	updateQuery, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateQuery.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
