ssa() {
  ./ssadump -build=F tmp_go2go.go > tmp_go2go_ssa.txt
  python3 ../count_ssa_ins.py tmp_go2go_ssa.txt | wc -l
}

word_count() {
  cat tmp_go2go.go2 | wc -m
}

PARAM_M=2
PARAM_N=2
PARAM_C=2
echo " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo "m=" "2..7"
for PARAM_M in {2..9}
do
python3 ./make_go2_input.py --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" > tmp_go2go.go2
./go2go translate tmp_go2go.go2
echo $PARAM_M "$(ssa)" #"$(word_count)"
done
exit

PARAM_M=2
PARAM_N=2
PARAM_C=3
echo "m=" $PARAM_M " n="  $PARAM_N " c=" "$PARAM_C"
echo range c from 2 to 100:
for PARAM_C in $(seq 2 5 100)
do
python3 ./make_go2_input.py --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" > tmp_go2go.go2
./go2go translate tmp_go2go.go2
echo $PARAM_C "$(ssa)" "$(word_count)"
done

PARAM_M=2
PARAM_N=2
PARAM_C=3
echo " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo range n from 2 to 20:
for PARAM_N in {2..20}
do
python3 ./make_go2_input.py --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" > tmp_go2go.go2
./go2go translate tmp_go2go.go2
echo $PARAM_N "$(ssa)" "$(word_count)"
done
