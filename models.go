package models

type Pedido struct {
	ID       string  `json:"id,omitempty" bson:"id"`
	SystemID int     `json:"system_id,omitempty" bson:"system_id"`
	Sabor    string  `json:"sabor,omitempty" bson:"sabor"`
	Tamanho  string  `json:"tamanho,omitempty" bson:"tamanho"`
	Valor    float32 `json:"valor,omitempty" bson:"valor"`
}
type Evento struct {
	Tipo   string `json:"tipo,omitempty" bson:"tipo"`
	Pedido Pedido `json:"pedido,omitempty" bson:"pedido"`
}
