#go build -gcflags '-N -l' -o dict_exe dict/main.go
#go build -gcflags '-N -l' -o mono_exe mono/main.go
#
#exec_time_mono() {
#  ts=$(date +%s%N)
#  ./mono_exe | ts '%.S' #ts '[%Y-%m-%d %H:%M:%.S]'
#  echo $((($(date +%s%N) - $ts)/1000))
#}
#
#exec_time_dict() {
#  ts=$(date +%s%N)
#  ./dict_exe | ts '%.S' #| ts '[%Y-%m-%d %H:%M:%.S]'
#  echo $((($(date +%s%N) - $ts)/1000))
#}
#
#echo "mono:"
#exec_time_mono
#echo "dict:"
#exec_time_dict



exec_time() {
  ts=$(date +%s%N)
  for i in {1..50}
  do
  ./"$1"_exe > /dev/null #| ts '%.S' #ts '[%Y-%m-%d %H:%M:%.S]'
  done
  echo $((($(date +%s%N) - $ts)/1000))
}


for mode in go2go mono dict dict_no_cast dyn
do
echo $mode
go build -gcflags '-N -l' -o "$mode"_exe "$mode"/main.go
exec_time $mode
done
