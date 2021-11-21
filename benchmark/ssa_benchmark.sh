get_LOC() {
    impl_type=$1
    ./ssadump -build=F tmp_"$impl_type".go |wc -l
}

get_LOC_no_assertion() {
    impl_type=$1
    ./ssadump -build=F tmp_"$impl_type".go > tmp_"$impl_type"_ssa.txt
    python3 check_no_assertion_loc.py tmp_"$impl_type"_ssa.txt | wc -l
}

doit() {
  path=$1
  impl_type=$2
  cp ./"$path"/"$impl_type"/main.go ./tmp_"$impl_type".go
}

for path_i in miniboxing-benchmarks/List miniboxing-benchmarks/ResizableArray pizza-benchmarks/cell pizza-benchmarks/list_reverse pizza-benchmarks/vector_reverse pizza-benchmarks/hashtable
do
  echo $path_i
  doit $path_i go2go
  get_LOC "go2go"
  doit $path_i mono
  get_LOC "mono"
  doit $path_i dict
  get_LOC "dict"
  get_LOC_no_assertion "dict"
  doit $path_i dyn
  get_LOC "dyn"
done
