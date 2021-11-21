import sys
import argparse


def perm2mainMethCall(perm):
    ret = "f1" + "["
    args = []
    targs = []
    for x in perm:
        t_s = 'Red' if x == 'a' else 'Blue'
        targs.append(t_s)
        args.append(t_s + "{}")
    ret += ", ".join(targs)
    ret += "](" + ",".join(args) + ")"
    return ret


def gen_tformals(m):
    tformals = []
    for i in range(1, m + 1):
        tformals.append('b' + str(i) + " Color")
    return ",".join(tformals)


def gen_params_tparams(m):
    params = []
    tparams = []
    for j in range(1, m + 1):
        params.append("p{0} b{0}".format(str(j)))
        tparams.append("b{0}".format(str(j)))
    return params, tparams


def gen_sig(m, n, i):
    params, tparams = gen_params_tparams(m)
    sig = "g{0}({1}) Base[{2}]".format(str(i), ",".join(params), ",".join(tparams))
    return sig


def gen_base(m, n):
    ret = "type Base[" + gen_tformals(m) + "] interface {\n"
    ret += "g_chan() Top;\n"
    sigs = []
    for i in range(1, n+ 1):
        sigs.append(gen_sig(m,n,i))
    ret += ";\n".join(sigs)
    ret += "};\n"
    return ret


def gen_gi(i, m, n):
    p, tp = gen_params_tparams(m)
    ret = "func (this Derived[" + ",".join(tp) + "]) " + gen_sig(m,n,i) + "{\n return this\n};\n"
    return ret


def gen_derived(m,n,c, chanops):
    params, tparams = gen_params_tparams(m)
    ret = "type Derived[" + gen_tformals(m) + "] struct{};\n"
    ret += "func (this Derived[" + ",".join(tparams) + "]) g_chan() Top {\n"
    if chanops:
        ret += "v2 := Top{}; mych := make(chan Top);"
        for i in range(0, c):
            ret += "go Recv(mych);\n"
            ret += "mych<-v2;\n"
        ret += "return Top{}\n};\n"
    else:
        for i in range(0, c):
            ret += "Op();"
        ret += "return Top{}\n};\n"
    for i in range(1, n + 1):
        ret += gen_gi(i, m, n)
    return ret


def gen_f0(m,n):
    params, tparams = gen_params_tparams(m)
    ret = "func f0[{0}, d_ Base[{1}]](x d_, {2}) Top {{\n t:=Top{{}}\n"\
        .format(gen_tformals(m), ",".join(tparams), ",".join(params))
    for i in range(1, n + 1):
        ret += "t.do(x.g{0}({1}))\n".format(str(i), ",".join(['p' + str(j) for j in range(1, m + 1)]))
    ret += "t.do(x.g_chan())\n"
    ret += "return t\n};\n"
    return ret


def gen_fp(p, m):
    params, tparams = gen_params_tparams(m)
    params = ",".join(params)
    tparams = ",".join(tparams)
    return """func f{4}[{0}]({1}) Top {{
    this := Top{{}}
    f0[{2}, Derived[{2}]](Derived[{2}]{{}}, {3})
    return this
}};\n""".format(gen_tformals(m), params, tparams, ",".join(['p' + str(i) for i in range(1, m + 1)]), str(p))


def gen_fi(m, i):
    tformals = []
    params = []
    tparams = []
    args = []
    for j in range(1, m + 1):
        tformals.append('b{0} Color'.format(str(j)))
        params.append('p{0} b{0}'.format(str(j)))
        tparams.append('b{0}'.format(str(j)))
        args.append('p{0}'.format(str(j)))
    tformals = ",".join(tformals)
    params = ",".join(params)
    tparams = ",".join(tparams)
    args = ",".join(args)
    ret = "func f" + str(i) + "[" + tformals + "]( " + params + ") Top {\n"
    #ret += "return this"
    ret += "f{0}[{1}]({2})\n".format(str(i + 1), tparams, args)
    ret += "return Top{}\n};\n"
    return ret


def codegen_main(m):
    perms = []
    def dfs(i, perm):
        if i == m + 1:
            #print(i)
            perms.append(perm)
            return
        new_perm = perm + 'a'
        dfs(i+1, new_perm)
        new_perm = perm + 'b'
        dfs(i+1, new_perm)
    dfs(1, '')
    ret = "func main(){\n"
    #ret += "_ = Top{}."
    ret += "\n".join([perm2mainMethCall(perm) for perm in perms])
    return "\n" + ret + "}"


def codegen(m, n, c, p, chanops=False):
    ret = """package main;

type Color interface{};
type Top struct {};
"""
    if chanops:
        ret += "func Recv(x chan Top) Top {return <-x};\n"
    else:
        ret += "func Op() Top {return Top{}};\n"
    ret += """
type Red struct {};
type Blue struct {};

func (this Top) do(x Color) Top {
return this
};
"""

    ret += gen_base(m,n)
    ret += gen_derived(m,n,c, chanops)
    ret += gen_f0(m, n)
    ret += gen_fp(p, m)
    for i in range(1, p):
        ret += gen_fi(m, i)
    ret += codegen_main(m)
    return ret


def main():
    argParser = argparse.ArgumentParser()
    argParser.add_argument("-n", "--n", help="dictionary size", type=int)
    argParser.add_argument("-m", "--m", help="the number of type parameters", type=int)
    argParser.add_argument("-c", "--c", help="the number of channel operations", type=int)
    argParser.add_argument("-p", "--p", help="the length of the callchain", type=int)
    argParser.add_argument("-chanops", "--chan-ops", help="use channel operations. 1 indicates True.", type=int)
    args = argParser.parse_args()

    #print(args.p, args.m, args.n, args.c)
    return codegen(args.m, args.n, args.c, args.p, args.chan_ops == 1)

if __name__ == "__main__":
    print(main())
