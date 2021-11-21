ssa_dictpass() {
  ./ssadump -build=F tmp_dict.go |wc -l
}

ssa_monom() {
  ./ssadump -build=F tmp_monom.go |wc -l
}

ssa_dyn() {
  ./ssadump -build=F tmp_dynamic.go |wc -l
}

get_LOC_new() {
    impl_type=$1
    ./ssadump -build=F tmp_"$impl_type".go > tmp_"$impl_type"_ssa.txt
    python3 ./count_ssa_ins.py tmp_"$impl_type"_ssa.txt | wc -l
}

get_LOC_no_assertion() {
    impl_type=$1
    ./ssadump -build=F tmp_"$impl_type".go > tmp_"$impl_type"_ssa.txt
    python3 check_no_assertion_loc.py tmp_"$impl_type"_ssa.txt | wc -l
}

word_count() {
  cat tmp.fgg | wc -m
}

PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo range p from 2 to 40:
for PARAM_P in {2..40}
do
python3 ./make_fcgg_input_ordinary_call_chain.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp.fgg
~/go/bin/fgg -fgg -dictpass -- ./tmp.fgg > tmp_dict.go
~/go/bin/fgg -fgg -monomc -- ./tmp.fgg > tmp_monom.go
~/go/bin/fgg -fgg -dictpassnoassertion -- ./tmp.fgg > tmp_dict_no_assertion.go
~/go/bin/fgg -fgg -dynamic -- ./tmp.fgg > tmp_dyn.go
echo $PARAM_P "$(get_LOC_new dict)" "$(get_LOC_new monom)" "$(get_LOC_new dict_no_assertion)" "$(get_LOC_new dyn)"
done
exit

PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo "m=" "2..7"
for PARAM_M in {2..9}
do
python3 ./make_fcgg_input_ordinary_call_chain.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp.fgg
~/go/bin/fgg -fgg -dictpass -- ./tmp.fgg > tmp_dict.go
~/go/bin/fgg -fgg -monomc -- ./tmp.fgg > tmp_monom.go
~/go/bin/fgg -fgg -dictpassnoassertion -- ./tmp.fgg > tmp_dict_no_assertion.go
~/go/bin/fgg -fgg -dynamic -- ./tmp.fgg > tmp_dyn.go
echo $PARAM_M "$(get_LOC_new dict)" "$(get_LOC_new monom)" "$(get_LOC_new dict_no_assertion)" "$(get_LOC_new dyn)"
done

PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo range c from 2 to 100:
for PARAM_C in $(seq 2 5 200)
do
python3 ./make_fcgg_input_ordinary_call_chain.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp.fgg
~/go/bin/fgg -fgg -dictpass -- ./tmp.fgg > tmp_dict.go
~/go/bin/fgg -fgg -monomc -- ./tmp.fgg > tmp_monom.go
~/go/bin/fgg -fgg -dictpassnoassertion -- ./tmp.fgg > tmp_dict_no_assertion.go
~/go/bin/fgg -fgg -dynamic -- ./tmp.fgg > tmp_dyn.go
echo $PARAM_C "$(get_LOC_new dict)" "$(get_LOC_new monom)" "$(get_LOC_new dict_no_assertion)" "$(get_LOC_new dyn)"
done




PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo range n from 2 to 40:
for PARAM_N in {2..40}
do
python3 ./make_fcgg_input_ordinary_call_chain.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp.fgg
~/go/bin/fgg -fgg -dictpass -- ./tmp.fgg > tmp_dict.go
~/go/bin/fgg -fgg -monomc -- ./tmp.fgg > tmp_monom.go
~/go/bin/fgg -fgg -dictpassnoassertion -- ./tmp.fgg > tmp_dict_no_assertion.go
~/go/bin/fgg -fgg -dynamic -- ./tmp.fgg > tmp_dyn.go
echo $PARAM_N "$(get_LOC_new dict)" "$(get_LOC_new monom)" "$(get_LOC_new dict_no_assertion)" "$(get_LOC_new dyn)"
done

PARAM_P=2
PARAM_M=2
PARAM_N=2
PARAM_C=2
echo "p=" $PARAM_P " m=" $PARAM_M " n="  $PARAM_N " c=" $PARAM_C
echo range p from 2 to 40:
for PARAM_P in {2..40}
do
python3 ./make_fcgg_input_ordinary_call_chain.py --p=$PARAM_P --m=$PARAM_M --n=$PARAM_N --c="$PARAM_C" -chanops 0 > tmp.fgg
~/go/bin/fgg -fgg -dictpass -- ./tmp.fgg > tmp_dict.go
~/go/bin/fgg -fgg -monomc -- ./tmp.fgg > tmp_monom.go
~/go/bin/fgg -fgg -dictpassnoassertion -- ./tmp.fgg > tmp_dict_no_assertion.go
~/go/bin/fgg -fgg -dynamic -- ./tmp.fgg > tmp_dyn.go
echo $PARAM_P "$(get_LOC_new dict)" "$(get_LOC_new monom)" "$(get_LOC_new dict_no_assertion)" "$(get_LOC_new dyn)"
done