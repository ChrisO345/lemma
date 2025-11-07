package core

import "github.com/chriso345/lemma/common"

type LemmaType struct {
	TypeName string
	Factory  func() common.Lemma
}

// TODO: Add more types as needed
var (
	Int = LemmaType{
		TypeName: "int",
		Factory:  func() common.Lemma { return &intLemma{} },
	}
	Float = LemmaType{
		TypeName: "float",
		Factory:  func() common.Lemma { return &floatLemma{} },
	}

	Custom = LemmaType{ // Must provide custom generator in corollary
		TypeName: "custom",
		Factory:  nil,
	}

	Undefined = LemmaType{ // For testing unsupported types
		TypeName: "undefined",
		Factory:  nil,
	}
)
