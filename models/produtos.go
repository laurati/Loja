package models

import (
	"produtos-gin/db"
)

type Produto struct {
	Id         int     `json:"id"`
	Nome       string  `json:"nome"`
	Preco      float64 `json:"preco"`
	Quantidade int     `json:"quantidade"`
}

func BuscaTodosOsProdutos() []Produto {

	db := db.ConectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}

	defer db.Close()
	return produtos

}

func CriaNovoProduto(nome string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos (nome, preco, quantidade) values($1, $2, $3)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, preco, quantidade)
	defer db.Close()
}

func BuscaProdutoPorId(id1 string) []Produto {

	db := db.ConectaComBancoDeDados()
	defer db.Close()

	selectProduto, err := db.Query("select * from produtos where id = " + id1)

	if err != nil {
		panic(err.Error())
	}

	prodSelecionado := Produto{}

	prod := []Produto{}

	var id, quantidade int
	var nome string
	var preco float64

	for selectProduto.Next() {
		err = selectProduto.Scan(&id, &nome, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		prodSelecionado.Id = id
		prodSelecionado.Nome = nome
		prodSelecionado.Preco = preco
		prodSelecionado.Quantidade = quantidade

		prod = append(prod, prodSelecionado)
	}

	return prod
}
