exec_time_mono() {
  ts=$(date +%s%N)
  ./tmp_monom
  echo $((($(date +%s%N) - $ts)/1000))
}

exec_time_dict() {
  ts=$(date +%s%N)
  ./tmp_dict
  echo $((($(date +%s%N) - $ts)/1000))
}

doit() {
~/go/bin/fgg -fgg -dictpass -- ./tmp.fgg > tmp_dict.go
~/go/bin/fgg -fgg -monomc -- ./tmp.fgg > tmp_monom.go
python3 rewrite_src.py tmp_dict.go > tmp_dict_rewrite.go
python3 rewrite_src.py tmp_monom.go > tmp_monom_rewrite.go
go build -gcflags '-N -l' -o tmp_dict tmp_dict_rewrite.go
go build -gcflags '-N -l' -o tmp_monom tmp_monom_rewrite.go
}


PARAM_M=2
PARAM_N=2
PARAM_C=2
echo " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo "m=" "2..7"
for PARAM_M in {7..7}
do
python3 ./make_fcgg_input.py -m=$PARAM_M -n=$PARAM_N -c="$PARAM_C" > tmp.fgg
doit
echo -n "$PARAM_M mono "
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

PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "m=" $PARAM_M " n="  $PARAM_N " c=" "$PARAM_C"
echo range c from 2 to 100:
for PARAM_C in $(seq 2 5 100)
do
python3 ./make_fcgg_input.py -m=$PARAM_M -n=$PARAM_N -c="$PARAM_C" > tmp.fgg
echo $PARAM_C "$(exec_time_mono)" "$(exec_time_dict)"
doit
done



PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" "$PARAM_C"
echo range n from 2 to 40:
for PARAM_N in {2..40}
do
python3 ./make_fcgg_input.py -m=$PARAM_M -n=$PARAM_N -c="$PARAM_C" > tmp.fgg
doit
echo $PARAM_N "$(exec_time_mono)" "$(exec_time_dict)"
done
