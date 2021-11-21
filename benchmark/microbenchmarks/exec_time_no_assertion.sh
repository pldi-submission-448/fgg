
exec_time_mono() {
  ts=$(date +%s%N)
  ./tmp_monom
  echo $((($(date +%s%N) - $ts)/1000))
}

exec_time_dict() {
  ts=$(date +%s%N)
  ./tmp_dictnoassertion
  echo $((($(date +%s%N) - $ts)/1000))
}

exec_time_dyn() {
  ts=$(date +%s%N)
  ./tmp_dyn
  echo $((($(date +%s%N) - $ts)/1000))
}

exec_time_dyn_no_gc() {
  ts=$(date +%s%N)
  GOGC=off ./tmp_dyn
  echo $((($(date +%s%N) - $ts)/1000))
}

doit() {
~/go/bin/fgg -fgg -dictpassnoassertion -- ./tmp.fgg 2>/dev/null 1> tmp_dictnoassertion.go
~/go/bin/fgg -fgg -monomc -- ./tmp.fgg 2>/dev/null 1> tmp_monom.go
~/go/bin/fgg -fgg -dynamic -- ./tmp.fgg 2>/dev/null 1> tmp_dynamic.go
python3 rewrite_src.py tmp_dictnoassertion.go > tmp_dictnoassertion_rewrite.go
python3 rewrite_src.py tmp_monom.go > tmp_monom_rewrite.go
python3 rewrite_src.py tmp_dynamic.go > tmp_dynamic_rewrite.go
go build -gcflags '-N -l' -o tmp_dictnoassertion tmp_dictnoassertion_rewrite.go
go build -gcflags '-N -l' -o tmp_monom tmp_monom_rewrite.go
go build -gcflags '-N -l' -o tmp_dyn tmp_dynamic_rewrite.go
}

PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo "p=" "2..40"
for PARAM_P in {2..40}
do
python3 ./make_fcgg_input_ordinary_call_chain.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp.fgg
doit
echo -n "dyn "
for loopvar in {1..10}
do
echo -n "$(exec_time_dyn) "
done
echo -n "dict "
for loopvar in {1..10}
do
echo -n "$(exec_time_dict) "
done
echo
done


PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo "m=" "2..10"
for PARAM_M in {2..9}
do
python3 ./make_fcgg_input_ordinary_call_chain.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp.fgg
doit
echo -n "mono "
for loopvar in {1..10}
do
echo -n "$(exec_time_mono) "
done
echo -n "dict "
for loopvar in {1..10}
do
echo -n "$(exec_time_dict) "
done
echo
done
exit

PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "chanops=0" "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" "$PARAM_C"
echo range c from 2 to 2000:
for PARAM_C in $(seq 2 5 200)
do
python3 ./make_fcgg_input_ordinary_call_chain.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp.fgg
doit
echo -n "dyn "
for loopvar in {1..10}
do
echo -n "$(exec_time_dyn) "
done
echo -n "dict "
for loopvar in {1..10}
do
echo -n "$(exec_time_dict) "
done
echo
done


PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" "$PARAM_C"
echo range n from 2 to 40:
for PARAM_N in {2..40}
do
python3 ./make_fcgg_input_ordinary_call_chain.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp.fgg
doit
echo -n "dyn "
for loopvar in {1..10}
do
echo -n "$(exec_time_dyn) "
done
echo -n "dict "
for loopvar in {1..10}
do
echo -n "$(exec_time_dict) "
done
echo
done
