import sys
import argparse


def perm2mainMethCall(perm):
    ret = "f1" + "["
    args = []
    targs = []
    for x in perm:
        t_s = 'Red' if x == 'a' else 'Blue'
        targs.append(t_s + "[]")
        args.append(t_s + "[]{}")
    ret += ", ".join(targs)
    ret += "](" + ",".join(args) + ")"
    return ret


def gen_tformals(m):
    tformals = []
    for i in range(1, m + 1):
        tformals.append('b' + str(i) + " Color[]")
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
    sig = "g{0}[type ]({1}) Base[{2}]".format(str(i), ",".join(params), ",".join(tparams))
    return sig


def gen_base(m, n):
    ret = "type Base[type " + gen_tformals(m) + "] interface {\n"
    ret += "g_chan[type]() Top[];\n"
    sigs = []
    for i in range(1, n+ 1):
        sigs.append(gen_sig(m,n,i))
    ret += ";\n".join(sigs)
    ret += "};\n"
    return ret


def gen_gi(i, m, n):
    ret = "func (this Derived[type " + gen_tformals(m) + "]) " + gen_sig(m,n,i) + "{\n return this\n};\n"
    return ret


def gen_derived(m,n,c, chanops):
    ret = "type Derived[type " + gen_tformals(m) + "] struct{};\n"
    ret += "func (this Derived[type " + gen_tformals(m) + "]) g_chan[type ]() Top[] {\n"
    if chanops:
        ret += "t := Top[]{}; v2 := Top[]{}; mych := make(chan Top[]);"
        for i in range(0, c):
            ret += "go t.Recv[](mych);\n"
            ret += "mych<-v2;\n"
        ret += "return Top[]{}\n};\n"
    else:
        ret += "return Top[]{}"
        for i in range(0, c):
            ret += ".Op[]()"
        ret += "\n};\n"
    for i in range(1, n + 1):
        ret += gen_gi(i, m, n)
    return ret


def gen_f0(m,n):
    params, tparams = gen_params_tparams(m)
    ret = "func (this Top[type ]) f0[type {0}, d_ Base[{1}]](x d_, {2}) Top[] {{\n"\
        .format(gen_tformals(m), ",".join(tparams), ",".join(params))
    ret += "return this"
    for i in range(1, n + 1):
        ret += ".do[](x.g{0}[]({1}))".format(str(i), ",".join(['p' + str(j) for j in range(1, m + 1)]))
    ret += ".do[](x.g_chan[]())\n};\n"
    return ret


def gen_fp(p, m):
    params, tparams = gen_params_tparams(m)
    params = ",".join(params)
    tparams = ",".join(tparams)
    return """func (this Top[type ]) f{4}[type {0}]({1}) Top[] {{
return this.f0[{2}, Derived[{2}]](Derived[{2}]{{}}, {3})
}};""".format(gen_tformals(m), params, tparams, ",".join(['p' + str(i) for i in range(1, m + 1)]), str(p))


def gen_fi(m, i):
    tformals = []
    params = []
    tparams = []
    args = []
    for j in range(1, m + 1):
        tformals.append('b{0} Color[]'.format(str(j)))
        params.append('p{0} b{0}'.format(str(j)))
        tparams.append('b{0}'.format(str(j)))
        args.append('p{0}'.format(str(j)))
    tformals = ",".join(tformals)
    params = ",".join(params)
    tparams = ",".join(tparams)
    args = ",".join(args)
    ret = "func (this Top[type ]) f" + str(i) + "[type " + tformals + "]( " + params + ") Top[] {\n"
    ret += "return this"
    ret += ".f{0}[{1}]({2})".format(str(i+1), tparams, args)
    ret += "\n};\n"
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
    ret += "_ = Top[]{}."
    ret += ".".join([perm2mainMethCall(perm) for perm in perms])
    return "\n" + ret + "}"


def codegen(m, n, c, p, chanops=False):
    ret = """package main;

type Color[type ] interface{};
type Top[type ] struct {};
//func (this Top[type ]) Recv[type b1 Color[]](x chan b1) b1 {return <-x};
"""
    if chanops:
        ret += "func (this Top[type ]) Recv[type ](x chan Top[]) Top[] {return <-x};\n"
    else:
        ret += "func (this Top[type ]) Op[type ]() Top[] {return Top[]{}};\n"
    ret += """
type Red[type ] struct {};
type Blue[type ] struct {};

func (this Top[type ]) do[type ](x Color[]) Top[] {
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
    return codegen(args.m, args.n, args.c, args.p, args.chan_ops==1)

if __name__ == "__main__":
    print(main())
