#!/bin/bash
COMMITS=$(( (RANDOM % 4) + 5 ))

for ((i=1; i<=COMMITS; i++)); do
  echo "update $(date -u +%Y-%m-%dT%H:%M:%S) - #$i" >> activity.log
  git add activity.log
  git commit -m "chore: daily update #$i ($(date -u +%Y-%m-%d))"
done

git push
