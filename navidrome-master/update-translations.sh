#!/bin/sh

set -e

download_lang() {
  url=$(curl -s -X POST https://poeditor.com/api/ \
    -d api_token="${POEDITOR_APIKEY}" \
    -d action="export" \
    -d id="${POEDITOR_PROJECTID}" \
    -d language="$1" \
    -d type="key_value_json" | jq -r .item)
  if [ -z "$url" ]; then
    echo "Failed to export $1"
    return 1
  fi
  curl -sSL "$url" > resources/i18n/"$1".json
}

for file in resources/i18n/*.json; do
  name=$(basename "$file")
  code=$(echo "$name" | cut -f1 -d.)
  lang=$(jq -r .languageName < "$file")
  echo "Downloading $lang ($code)"
  download_lang "$code"
done
