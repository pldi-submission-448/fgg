ssa() {
  ./ssadump -build=F tmp_go2go.go > tmp_go2go_ssa.txt
  python3 ./count_ssa_ins.py tmp_go2go_ssa.txt | wc -l
}

word_count() {
  cat tmp_go2go.go2 | wc -m
}

PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo "p=" "2..7"
for PARAM_P in {2..30}
do
python3 ./make_go2_input.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp_go2go.go2
~/fg/go2go_goroot/goroot/bin/go tool go2go translate tmp_go2go.go2
echo $PARAM_P "$(ssa)" "$(word_count)"
done


PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo "m=" "2..7"
for PARAM_M in {2..9}
do
python3 ./make_go2_input.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp_go2go.go2
~/fg/go2go_goroot/goroot/bin/go tool go2go translate tmp_go2go.go2
echo $PARAM_M "$(ssa)" "$(word_count)"
done


PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo range c from 2 to 100:
for PARAM_C in $(seq 2 5 200)
do
python3 ./make_go2_input.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp_go2go.go2
~/fg/go2go_goroot/goroot/bin/go tool go2go translate tmp_go2go.go2
echo $PARAM_C "$(ssa)" "$(word_count)"
done



PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo range N from 2 to 40:
for PARAM_N in {2..40}
do
python3 ./make_go2_input.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp_go2go.go2
~/fg/go2go_goroot/goroot/bin/go tool go2go translate tmp_go2go.go2
echo $PARAM_N "$(ssa)" "$(word_count)"
done
