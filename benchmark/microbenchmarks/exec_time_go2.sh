exec_go2go() {
  ts=$(date +%s%N)
  ./tmp_go2go
  echo $((($(date +%s%N) - $ts)/1000))
}

doit() {
./go2go translate tmp_go2go.go2
python3 rewrite_src.py tmp_go2go.go > tmp_go2go_rewrite.go
go build -gcflags '-N -l' -o tmp_go2go tmp_go2go_rewrite.go
}


#PARAM_P=2
#PARAM_M=2
#PARAM_N=2
#PARAM_C=2
#echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
#echo "p=" "2..30"
#for PARAM_P in {2..30}
#do
#python3 ./make_go2_input.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp_go2go.go2
#doit
#echo -n "$PARAM_P "
#for loopvar in {1..10}
#do
#echo -n "$(exec_go2go) "
#done
#echo
#done
#exit


PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo range c from 2 to 100:
for PARAM_C in $(seq 2 5 200)
do
python3 ./make_go2_input.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp_go2go.go2
doit
echo -n "$PARAM_C "
for loopvar in {1..10}
do
echo -n "$(exec_go2go) "
done
echo
done



PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo range c from 2 to 100:
for PARAM_C in $(seq 2 5 200)
do
python3 ./make_go2_input.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 1 > tmp_go2go.go2
doit
echo -n "$PARAM_C "
for loopvar in {1..10}
do
echo -n "$(exec_go2go) "
done
echo
done



PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo range n from 2 to 40:
for PARAM_N in {2..40}
do
python3 ./make_go2_input.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp_go2go.go2
doit
echo -n "$PARAM_N "
for loopvar in {1..10}
do
echo -n "$(exec_go2go) "
done
echo
done


PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo range m from 2 to 7:
for PARAM_M in {2..9}
do
python3 ./make_go2_input.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp_go2go.go2
doit
echo -n "$PARAM_M "
for loopvar in {1..10}
do
echo -n "$(exec_go2go) "
done
echo
done
exit