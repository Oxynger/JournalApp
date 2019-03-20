package model

type imageURL = string

// Singatures godoc
type Singatures struct {
	Signature imageURL `bson:"signature" json:"signature" example:"https://qph.fs.quoracdn.net/main-qimg-2248bdd01f82b9fb9becdc4bd9a92c53" `
}
