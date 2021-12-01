package models

import "produtos-gin/db"

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