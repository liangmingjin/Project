#!/bin/sh
mkdir -p asspath

function searchAndcp() {
	filenamemp4=`ls -a | grep "\[${1}\]"`
	echo -e "$filenamemp4"
	filename=${filenamemp4%*.mp4}
	assFile=`ls ./142814399318415/ | grep " ${1}.ass"`
	echo $assFile
	newass=$filename".ass"
	cp ./142814399318415/"$assFile" asspath/"$newass"
}

searchAndcp "01"
searchAndcp "02"
searchAndcp "03"
searchAndcp "04"
searchAndcp "05"
searchAndcp "06"
searchAndcp "07"
searchAndcp "08"
searchAndcp "09"
for count in {10..175}
do
	searchAndcp "$count"
done
ls asspath/



