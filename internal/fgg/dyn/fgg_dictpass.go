package dyn

import (
	"github.com/pldi-submission-448/fgg/internal/fg"
	"github.com/pldi-submission-448/fgg/internal/fgg"
	"strconv"
)

const chan_wrapper_name = "chan_wrapper"

func MakeAnyInterface() fg.ITypeLit {
	return fg.NewITypeLit(fg.Type("Any"), []fg.Spec{})
}

func AuxAsParam(psi fgg.BigPsi) []fg.FieldOrParamDecl {
	ret := []fg.FieldOrParamDecl{}
	for i, tFormal := range psi.TFormals {
		fd := fg.NewFieldDecl("dict_"+strconv.Itoa(i), AuxTypeDict(AuxRootType(tFormal.U_I)))
		ret = append(ret, fd)
	}
	return ret
}

func NamingDictOf(tFormal fgg.TFormal) fg.Name {
	return "DictOf" + AuxBracket(tFormal.Name.String())
}

func ToFieldDecl(fpdecls []fg.FieldOrParamDecl) []fg.FieldDecl {
	ret := []fg.FieldDecl{}
	for _, decl := range fpdecls {
		ret = append(ret, fg.NewFieldDecl(decl.GetName(), decl.GetType()))
	}
	return ret
}

func ToParamDecl(fpdecls []fg.FieldOrParamDecl) []fg.ParamDecl {
	ret := []fg.ParamDecl{}
	for _, decl := range fpdecls {
		ret = append(ret, fg.NewParamDecl(decl.GetName(), decl.GetType()))
	}
	return ret
}

//I-DICT
func TranslateSigDict(spec fgg.Sig) fg.FieldDecl {
	name, functy := TranslateSigDictInner(spec)
	return fg.NewFieldDecl(AuxBracket(name), fg.Type(functy.String())) //!!tmp hack for only supporting name as types
}

func TranslateSigDictInner(spec fgg.Sig) (fgg.Name, fg.FuncType) {
	newPDecls := []fg.ParamDecl{fg.NewParamDecl("receiver", fg.Type("Any"))}
	newPDecls = append(newPDecls, ToParamDecl(AuxAsParam(spec.Psi))...)
	pdecls := spec.GetParamDecls()
	for _, pdecl := range pdecls {
		newpd := fg.NewParamDecl(AuxBracket(pdecl.Name), fg.Type("Any"))
		newPDecls = append(newPDecls, newpd)
	}
	return spec.GetMethod(), fg.FuncType{
		ParamDecls: newPDecls,
		RetType:    fg.Type("Any"),
	} //!!tmp hack for only supporting name as types
}

//I-DICTS
func MakeDictForLists(ds []fgg.Decl, eta map[fgg.TParam]fg.FGExpr, delta fgg.Delta, tau []fgg.Type, psi fgg.BigPsi) []fg.FGExpr {
	ret := []fg.FGExpr{}
	l := len(tau)
	for i := 0; i < l; i++ {
		tau_i := tau[i]
		psi_i := psi.TFormals[i].U_I
		ret = append(ret, makeDict(ds, eta, delta, tau_i, psi_i))
	}
	return ret
}

func makeDict(ds []fgg.Decl, eta map[fgg.TParam]fg.FGExpr, delta fgg.Delta, tau fgg.Type, smalldelta fgg.Type) fg.FGExpr {
	switch tau_ := tau.(type) {
	case fgg.TNamed: //I-NEWDICT
		meths, order := fgg.MethodsDelta(ds, delta, smalldelta)
		t_S := AuxTypeDict(AuxRootType(smalldelta))
		es := []fg.FGExpr{}
		for _, meth := range order {
			_, functy := TranslateSigDictInner(meths[meth])
			args := []fg.FGExpr{}
			for _, item := range functy.ParamDecls[1:] {
				args = append(args, fg.NewVariable(item.GetName()))
			}
			es = append(es, fg.FuncDecl{
				Name:       "",
				ParamDecls: functy.ParamDecls,
				RetType:    functy.RetType,
				BodyExpr:   fg.NewCall(fg.NewAssert(fg.NewVariable("receiver"), fg.Type(AuxBracket(tau_.TName))), AuxBracket(meths[meth].GetMethod()), args),
			})
		}
		es = append(es, AuxTypeMeta(composeEtaToZeta(eta), tau_))
		return fg.NewStructLit(t_S, es)
	case fgg.TParam:
		if fgg.IsNamedIfaceType(ds, smalldelta) { //I-VARDICT
			if AuxRootType(smalldelta) == AuxRootType(delta[tau_]) {
				return eta[tau_]
			} else {
				panic("AuxRootType(smalldelta) != AuxRootType(delta[tau]). This shouldn't happen!")
			}
		} else if fgg.IsStructType(ds, smalldelta) { //I-VARDICTSUB
			meths := fgg.Methods(ds, smalldelta)
			t_S := AuxTypeDict(AuxRootType(smalldelta))
			es := []fg.FGExpr{}
			for _, meth := range meths {
				es = append(es, fg.NewSelect(eta[tau_], meth.GetMethod()))
			}
			es = append(es, fg.NewSelect(eta[tau_], fg.Name("_type")))
			return fg.NewStructLit(t_S, es)
		}
	case fgg.ChannelType:
		meths, _ := fgg.MethodsDelta(ds, delta, smalldelta)
		if len(meths) != 0 {
			panic("Channels should only be used on interfaces with no methods.")
		}
		zeta := composeEtaToZeta(eta)
		return fg.NewStructLit(AuxTypeDict(fg.Type(smalldelta.(fgg.TNamed).TName)), []fg.FGExpr{AuxTypeMeta(zeta, tau_)})
	}
	panic("!!!")
}

func (this DictPassTranslator) TranslateMethBody(delta fgg.Delta, eta map[fgg.TParam]fg.FGExpr, gamma fgg.Gamma, ds []fgg.Decl, stmts fgg.MethBody, paramsToSub map[string]struct{}) fg.MethBody {
	tmp_gamma := fgg.Gamma{}
	for k, v := range gamma {
		tmp_gamma[k] = v
	}

	ret := []fg.FGStmt{}
	for _, stmt_ := range stmts.StmtBody {
		switch stmt := stmt_.(type) {
		case fgg.ReturnStmt:
			zeta := composeEtaToZeta(eta)
			expr, isStructLit := this.TranslateExpr(stmt.ExprBody, ds, delta, eta, tmp_gamma)
			expr2 := AuxRechan(zeta, expr, isStructLit, stmt.ExprBody.Typing(ds, delta, tmp_gamma, false))
			ret = append(ret, fg.FGReturnStmt{ExprBody: expr2})
		case fgg.FGGCaseSelectStmt:
			new_case_select := fg.FGCaseSelectStmt{
				Cases:       []fg.FGCaseBody{},
				DefaultCase: nil,
			}
			for _, casebody := range stmt.Cases {
				new_case := fg.FGCaseBody{
					CaseGuard: nil,
					Body:      nil,
				}
				new_stmt_seq := []fgg.FGGStmt{casebody.CaseGuard.ToFGGStmt()}
				new_stmt_seq = append(new_stmt_seq, casebody.Body...)
				new_mb := this.TranslateMethBody(delta, eta, tmp_gamma, ds, fgg.MethBody{new_stmt_seq}, paramsToSub).StmtBody
				switch first_stmt := new_mb[0].(type) {
				case fg.FGChDispatchStmt:
					new_case.CaseGuard = fg.FGDispatchCaseGuard{LeftExpr: first_stmt.LeftExpr, RightExpr: first_stmt.RightExpr}
				case fg.FGAssignmentStmt:
					new_case.CaseGuard = fg.FGReceiveCaseGuard{
						LeftName:  first_stmt.VarName,
						RightExpr: first_stmt.Body,
					}
				}
				new_case.Body = new_mb[1:]
				new_case_select.Cases = append(new_case_select.Cases, new_case)
			}
			if stmt.DefaultCase != nil && len(stmt.DefaultCase) > 0 {
				new_mb := this.TranslateMethBody(delta, eta, tmp_gamma, ds, fgg.MethBody{stmt.DefaultCase}, paramsToSub).StmtBody
				new_case_select.DefaultCase = new_mb
			}
			ret = append(ret, new_case_select)
		case fgg.FGGChCloseStmt:
			sub_expr, isStructLit := this.TranslateExpr(stmt.Body, ds, delta, eta, tmp_gamma)
			new_assert := fg.NewAssert(sub_expr, fg.Type(chan_wrapper_name))
			if isStructLit {
				new_ch_close := fg.FGChCloseStmt{Body: fg.NewSelect(sub_expr, fg.Name("ch"))}
				ret = append(ret, new_ch_close)
			} else {
				new_expr := fg.NewSelect(new_assert, fg.Name("ch"))
				new_ch_close := fg.FGChCloseStmt{Body: new_expr}
				ret = append(ret, new_ch_close)
			}
		case fgg.FGGChDispatchStmt:
			sub_expr, isSLit := this.TranslateExpr(stmt.LeftExpr, ds, delta, eta, tmp_gamma)

			new_assert := fg.NewAssert(sub_expr, fg.Type(chan_wrapper_name))
			new_expr := fg.NewSelect(new_assert, fg.Name("ch"))
			if isSLit {
				new_expr = fg.NewSelect(sub_expr, fg.Name("ch"))
			}
			right_expr, _ := this.TranslateExpr(stmt.RightExpr, ds, delta, eta, tmp_gamma)
			new_ch_dispatch := fg.FGChDispatchStmt{
				LeftExpr:  new_expr,
				RightExpr: right_expr,
			}
			ret = append(ret, new_ch_dispatch)
		case fgg.FGGAssignmentStmt:
			/*if stmt.VarName == "rangerPair" {
				fmt.Println("debug point")
			}*/
			body, isSLit := this.TranslateExpr(stmt.Body, ds, delta, eta, tmp_gamma)
			if isSLit {
				this.env.StructVariables[stmt.VarName] = struct{}{}
			}
			new_assign := fg.FGAssignmentStmt{
				VarName: AuxBracket(stmt.VarName),
				Body:    body,
			}
			ret = append(ret, new_assign)
			tmp_gamma[stmt.VarName] = stmt.Body.Typing(ds, delta, tmp_gamma, false)
		case fgg.GoroutineStmt:
			callexpr, _ := this.TranslateExpr(stmt.CallExpr, ds, delta, eta, tmp_gamma)
			new_call := fg.GoroutineStmt{CallExpr: callexpr.(fg.Call)}
			ret = append(ret, new_call)
		}
	}
	return fg.MethBody{StmtBody: ret}
}

//I-METH
func (this DictPassTranslator) TranslateMethod(ds []fgg.Decl, decl fgg.MethDecl) []fg.Decl {
	//rule out Phi != Phi'
	td := fgg.GetTDecl(ds, decl.TRecv)
	tfs_td := td.GetBigPsi().TFormals
	for i := 0; i < len(tfs_td); i++ {
		subs_md := fgg.MakeParamIndexSubs(decl.Psi_recv)
		subs_td := fgg.MakeParamIndexSubs(td.GetBigPsi())
		if !decl.Psi_recv.TFormals[i].U_I.TSubs(subs_md). // Canonicalised
			Equals(tfs_td[i].U_I.TSubs(subs_td)) {
			panic("Receiver parameter upperbound is not the same as type decl upperbound:" +
				"\n\tmdecl=" + decl.Psi_recv.TFormals[i].String() + ", tdecl=" +
				tfs_td[i].String())
		}
	}
	delta := decl.Psi_recv.ToDelta()
	for _, v := range decl.Psi_meth.TFormals {
		delta[v.Name] = v.U_I
	}
	eta := map[fgg.TParam]fg.FGExpr{}
	//alphas := []fgg.Type{}
	for i, tformal := range decl.Psi_recv.TFormals {
		eta[tformal.Name] = fg.NewSelect(fg.NewVariable(AuxBracket(decl.XRecv)), "dict_"+strconv.Itoa(i))
		//alphas = append(alphas, tformal.U_I)
	}
	for i, tformal := range decl.Psi_meth.TFormals {
		eta[tformal.Name] = fg.NewVariable("dict_" + strconv.Itoa(i))
	}
	//bigphi_dagger := ToParamDecl(AuxAsParam(decl.Psi_meth))
	new_paramdecls := []fg.ParamDecl{}
	//new_paramdecls = append(new_paramdecls, bigphi_dagger...)

	gamma := fgg.Gamma{}
	gamma[decl.XRecv] = fgg.NewTName(decl.TRecv, decl.Psi_recv.Hat())
	for _, paramdecl := range decl.GetParamDecls() {
		gamma[paramdecl.Name] = paramdecl.U
		new_paramdecls = append(new_paramdecls, fg.NewParamDecl(AuxBracket(paramdecl.Name), fg.Type("Any")))
	}
	//e_dagger := TranslateExpr(decl.E_body, ds, delta, eta, gamma)
	this.env.StructVariables = make(map[string]struct{})
	this.env.StructVariables[decl.XRecv] = struct{}{}
	body_dagger := this.TranslateMethBody(delta, eta, gamma, ds, decl.BodyStmt, nil)
	fg_methdecl := fg.NewMDecl(fg.NewParamDecl(AuxBracket(decl.XRecv), fg.Type(AuxBracket(decl.TRecv))), AuxBracket(decl.GetName()), new_paramdecls, fg.Type("Any"), nil, body_dagger.StmtBody)
	if this.env.NoUnusedVar {
		UnusedVarRemover{}.VisitMethodDecl(fg_methdecl)
	}
	ret := []fg.Decl{fg_methdecl}
	return ret
	zeta := make(map[fgg.Name]fg.FGExpr)
	//alpha
	for i, tfml := range decl.Psi_recv.TFormals {
		zeta[tfml.Name.String()] = fg.NewSelect(fg.NewSelect(fg.NewVariable(AuxBracket(decl.XRecv)), fg.Name("dict_"+strconv.Itoa(i))), fg.Name("_type"))
	}
	//beta
	for i, tformal := range decl.Psi_meth.TFormals {
		this.env.param_index_set[i] = struct{}{}
		zeta[tformal.Name.String()] = fg.NewStructLit(fg.Type("param_index_"+strconv.Itoa(i)), []fg.FGExpr{})
	}
	spec_struct_lit_es := []fg.FGExpr{}
	for _, ty := range decl.Psi_meth.TFormals {
		spec_struct_lit_es = append(spec_struct_lit_es, AuxTypeMeta(zeta, ty.U_I))
	}
	for _, ty := range decl.GetParamDecls() {
		spec_struct_lit_es = append(spec_struct_lit_es, AuxTypeMeta(zeta, ty.U))
	}
	spec_struct_lit_es = append(spec_struct_lit_es, AuxTypeMeta(zeta, decl.U_ret))
	spec_len := len(decl.Psi_meth.TFormals) + len(decl.GetParamDecls())
	this.env.n_bar[spec_len] = struct{}{}
	spec_struct_type := spec_metadata_name + strconv.Itoa(spec_len)
	spec_body_expr := fg.NewStructLit(fg.Type(spec_struct_type), spec_struct_lit_es)
	spec_ret := fg.FGReturnStmt{ExprBody: spec_body_expr}
	spec_methbody := fg.MethBody{}
	spec_methbody.StmtBody = []fg.FGStmt{spec_ret}
	ret = append(ret, fg.NewMDecl(
		fg.NewParamDecl(AuxBracket(decl.XRecv), fg.Type(AuxBracket(decl.TRecv))),
		AuxSpecName(decl.GetName()),
		[]fg.ParamDecl{},
		fg.Type(spec_struct_type), nil, spec_methbody.StmtBody),
	)
	return ret
}

func (this DictPassTranslator) TranslateDecls(decls []fgg.Decl) []fg.Decl {
	ret := []fg.Decl{MakeAnyInterface()}
	for _, decl := range decls {
		switch decl_ := decl.(type) {
		case fgg.MethDecl:
			ret = append(ret, this.TranslateMethod(decls, decl_)...)
		case fgg.ITypeLit:
			ret = append(ret, this.TranslateITypeLit(decl_)...)
		case fgg.STypeLit:
			ret = append(ret, this.TranslateSTypeLit(decl_)...)
		}
	}
	return ret
}
