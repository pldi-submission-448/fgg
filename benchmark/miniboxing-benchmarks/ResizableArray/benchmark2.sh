go build -gcflags '-N -l' -o go2go_exe go2go/main.go
go build -gcflags '-N -l' -o mono_exe mono/main.go

exec_time_go2go() {
  ts=$(date +%s%N)
  ./go2go_exe | ts '%.S' #ts '[%Y-%m-%d %H:%M:%.S]'
  echo $((($(date +%s%N) - $ts)/1000))
}

exec_time_dict() {
  ts=$(date +%s%N)
  ./dict_exe | ts '%.S' #| ts '[%Y-%m-%d %H:%M:%.S]'
  echo $((($(date +%s%N) - $ts)/1000))
}

echo "go2go:"
exec_time_go2go
echo "dict:"
exec_time_dict
