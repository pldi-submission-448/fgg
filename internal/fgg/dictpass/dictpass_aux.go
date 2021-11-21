package dictpass

import (
	"github.com/pldi-submission-448/fgg/internal/fg"
	"github.com/pldi-submission-448/fgg/internal/fgg"
)

func AuxBracket(s fgg.Name) fg.Name {
	return s + "ǂ"
}

//fgg.TParam <: fgg.Type
//fgg.TNamed <: fgg.Type
func AuxRootType(ty fgg.Type) fg.TypeBase {
	switch ty_ := ty.(type) {
	case fgg.TNamed:
		return fg.Type(AuxBracket(ty_.TName))
	case fgg.TParam:
		return fg.Type("Any")
	case fgg.ChannelType:
		return fg.ChannelType{
			ElementType: AuxRootType(ty_.ElementType),
			ChType:      ty_.ChType,
		}
	default:
		panic("undefined fgg.Type")
	}
}

func AuxTypeDict(ty fg.TypeBase) fg.Type {
	switch ty_ := ty.(type) {
	case fg.ChannelType:
		panic("no dict for channel types")
	case fg.Type:
		return fg.Type(AuxBracket(ty_.String() + "Dict")) //AuxRootType(ty)
	default:
		panic("not implemented type.")
	}

}

func AuxRequires(tFormals []fgg.TFormal) []fg.Name {
	ret := []fg.Name{}
	for _, formal := range tFormals {
		if tnamed, ok := formal.U_I.(fgg.TNamed); ok {
			ret = append(ret, fg.Name(tnamed.TName))
		}
	}
	return ret
}

func addRequires(formals []fgg.TFormal, requires map[fg.Name]struct{}) {
	for _, name := range AuxRequires(formals) {
		requires[name] = struct{}{}
	}
}

func getDict(ProgDecls []fgg.Decl) map[fg.Name]struct{} {
	ret := make(map[fg.Name]struct{})
	for _, decl := range ProgDecls {
		switch decl_ := decl.(type) {
		case fgg.MethDecl:
			addRequires(decl_.Psi_recv.TFormals, ret)
			addRequires(decl_.Psi_meth.TFormals, ret)
		case fgg.ITypeLit:
			addRequires(decl_.Psi.TFormals, ret)
			for _, spec := range decl_.GetSpecs() {
				if sig, ok := spec.(fgg.Sig); ok {
					addRequires(sig.Psi.TFormals, ret)
				} else {
					panic("should be a fgg.Sig here.")
				}
			}
		case fgg.STypeLit:
			break
		}
	}
	return ret
}

var dictCache map[fg.Name]struct{} = nil

//pre: dictCache is initialized by getDict()
func InDictionaries(name fg.Name) bool {
	_, ok := dictCache[name]
	return ok
}
