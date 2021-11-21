import sys


def main():
    ret = []
    with open(sys.argv[1]) as fd:
        for line in fd:
            newline = ""
            if "func main() {" in line:
                newline = line.replace("func main() {", "func _main() {")
            else:
                newline = line
            ret.append(newline)
    ret.append("""
func main() {
for i := 1; i<10000; i++ {
_main()
}
}
    """)
    return "\n".join(ret)

if __name__ == "__main__":
    print(main())