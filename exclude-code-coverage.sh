#!/bin/sh
while read p || [ -n "$p" ] 
do  
sed -i '' "/${p//\//\\/}/d" ./cover.out 
done < ./exclude_code_coverage.txt