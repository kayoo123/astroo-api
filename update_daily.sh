#!/bin/bash
###########
cd docs
FILE_WORK='website'
FILE='jour_tmp.json'

curl -s http://www.astroo.com/horoscope.php > ${FILE_WORK}

echo "{" > $FILE
echo "\"date\": '\"$()\""
for SIGNE in BELIER TAUREAU GEMEAUX CANCER LION VIERGE BALANCE SCORPION SAGITTAIRE CAPRICORNE VERSEAU POISSONS; do 
  echo "  \"${SIGNE,,}\": \"$(cat ${FILE_WORK} |grep "$SIGNE</a>" |awk -F'</span>' '{ print $2 }' |recode html..utf8)\"," >> $FILE
done
sed -i '$ s/.$//' $FILE
echo "}" >> $FILE
