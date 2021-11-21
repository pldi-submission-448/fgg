import sys

no_gc = True

def main():
    ret = []
    with open(sys.argv[1]) as fd:
        for line in fd:
            newline = ""
            if "func main() {" in line:
                newline = line.replace("func main() {", "func _main() {")
            elif "package main" in line:
                if no_gc:
                    newline = line.replace("package main", "package main;import(\"runtime/debug\")")
                else:
                    newline = line.replace("package main", "package main;//import(\"runtime\")")
            else:
                newline = line
            ret.append(newline)
    if no_gc:
        ret.append("""
func main() {
debug.SetGCPercent(-1)
"""
)
    else:
        ret.append("""
func main() {
""")
    ret.append("""
for i := 1; i<10000; i++ {
_main()
if i % 500 == 0 {
//runtime.GC()
}
}
}
    """)
    return "\n".join(ret)

if __name__ == "__main__":
    print(main())