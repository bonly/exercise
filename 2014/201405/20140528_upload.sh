#/bin/bash
reg_file=$1
dir_id=$2
des_name=$3
for i in `ls $reg_file*` 
do
	#echo 20140527_upload_rom -f $i -n ${i%.*} -i $dir_id
	./20140527_upload_rom -f $i -n ${i%.*} -i $dir_id
done

#echo ${filename#*.}
#echo ${filename%.*}
