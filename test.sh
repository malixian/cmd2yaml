cat  dockerrun.cmd | while read line
do
  echo $line
  IFS="\t"
  read -ra arr <<< $line
  echo ./c2y -i "\""${arr[0]}"\"" -n ${arr[1]}
done
