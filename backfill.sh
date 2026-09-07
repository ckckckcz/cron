#!/bin/bash
START_DATE="2026-01-01"
END_DATE=$(date +%Y-%m-%d)
current="$START_DATE"

while [[ "$current" < "$END_DATE" || "$current" == "$END_DATE" ]]; do
  if (( RANDOM % 3 == 0 )); then
    current=$(date -I -d "$current + 1 day")
    continue
  fi

  commits=$(( (RANDOM % 4) + 1 ))
  for ((i=0; i<commits; i++)); do
    echo "log $current #$i" >> activity.log
    git add activity.log
    GIT_AUTHOR_DATE="$current 12:00:00" GIT_COMMITTER_DATE="$current 12:00:00" \
      git commit -m "chore: update ($current)" --quiet
  done

  current=$(date -I -d "$current + 1 day")
done

git push origin main
