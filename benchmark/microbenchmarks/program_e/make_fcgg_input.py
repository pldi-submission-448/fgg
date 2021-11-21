import sys
import argparse


def perm2mainMethCall(perm, c):
    ret = "f" + str(c) + "["
    args = []
    for x in perm:
        args.append('a' + str(ord(x) - ord('a') + 1) + '[]')
    ret += ", ".join(args)
    ret += "]()"
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
    ret += "g_chan[type](value b1) Top[];\n"
    sigs = []
    for i in range(1, n+ 1):
        sigs.append(gen_sig(m,n,i))
    ret += ";\n".join(sigs)
    ret += "};\n"
    return ret


def gen_gi(i, m, n):
    ret = "func (this Derived[type " + gen_tformals(m) + "]) " + gen_sig(m,n,i) + "{\n return this\n};\n"
    return ret


def gen_derived(m,n,c):
    ret = "type Derived[type " + gen_tformals(m) + "] struct{};\n"
    ret += "func (this Derived[type " + gen_tformals(m) + "]) g_chan[type ](value b1) Top[] {\n"
    # ret += "t := Top[]{}; v2 := Top[]{}; mych := make(chan Top[]);"
    # for i in range(0, c):
    #     ret += "go t.Recv[](mych);\n"
    #     ret += "mych<-v2;\n"
    # ret += "return Top[]{}\n};\n"
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
    #ret += "mych := make(chan b1);\n"
    ret += "return this"
    for i in range(1, n + 1):
        ret += ".do[](x.g{0}[]({1}))".format(str(i), ",".join(['p' + str(j) for j in range(1, m + 1)]))
    ret += ".do[](x.g_chan[](p1))\n};\n"
    return ret


def gen_f1(m):
    params, tparams = gen_params_tparams(m)
    params = ",".join(params)
    tparams = ",".join(tparams)
    return """func (this Top[type ]) f1[type {0}]({1}) Top[] {{
return this.f0[{2}, Derived[{2}]](Derived[{2}]{{}}, {3})
}};""".format(gen_tformals(m), params, tparams, ",".join(['p' + str(i) for i in range(1, m + 1)]))


def gen_fi(m, i):
    upper_bound = m - i + 1
    tformals = []
    params = []
    tparams = []
    args = []
    for j in range(1, upper_bound + 1):
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
    ret += ".f{0}[{1}, Black[]]({2}, Black[]{{}})".format(str(i-1), tparams, args)
    ret += ".f{0}[{1}, White[]]({2}, White[]{{}})".format(str(i - 1), tparams, args)
    ret += "\n};\n"
    return ret


def codegen_main(m):
    return """func main(){{
_ = Top[]{{}}.f{0}[Black[]](Black[]{{}}).f{0}[White[]](White[]{{}})
}}""".format(str(m))


def codegen(m, n, c):
    ret = """
package main;

type Color[type ] interface{};
type Top[type ] struct {};
//func (this Top[type ]) Recv[type b1 Color[]](x chan b1) b1 {return <-x};
//func (this Top[type ]) Recv[type ](x chan Top[]) Top[] {return <-x};
func (this Top[type ]) Op[type ]() Top[] {return Top[]{}};

type Black[type ] struct {};
type White[type ] struct {};

func (this Top[type ]) do[type ](x Color[]) Top[] {
return this
};
"""
    ret += gen_base(m,n)
    ret += gen_derived(m,n,c)
    ret += gen_f0(m, n)
    ret += gen_f1(m)
    for i in range(2, m + 1):
        ret += gen_fi(m, i)
    ret += codegen_main(m)
    return ret


def main():
    argParser = argparse.ArgumentParser()
    argParser.add_argument("-n", "--n", help="dictionary size", type=int)
    argParser.add_argument("-m", "--m", help="the number of type parameters", type=int)
    argParser.add_argument("-c", "--c", help="the number of channel operations", type=int)
    args = argParser.parse_args()
    #print(args.p, args.m, args.n, args.c)
    return codegen(args.m, args.n, args.c)

if __name__ == "__main__":
    print(main())
