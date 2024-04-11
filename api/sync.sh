#!/bin/bash
# 同步api文档脚本
#apifox项目ID
ProjectId=$1
#apifox开放token
Token=$2
# 生成json文件
for file in *.api; do
    echo $file
    if [ -f "$file" ]; then
        goctl api plugin -plugin goctl-swagger="swagger -filename $file.json" -api $file -dir .
        echo "$json_data"
    fi
done

# 遍历当前文件夹中的所有文件，并输出扩展名为 ".json" 的文件名
for file in *.json; do
    if [ -f "$file" ]; then
        json_data=$(cat $file)
        # 对字符串进行替换(apifox bug)
        new_data=$(echo "$json_data" | sed 's/"requestBody": {},/ /')
        echo "{
          \"importFormat\":\"openapi\",
          \"apiOverwriteMode\":\"merge\",
          \"schemaOverwriteMode\":\"name\",
          \"data\":$json_data
        }" > temp_data.json

        # 请求apifox开放接口 为projectId Bearer 访问令牌
        curl --location -g --request POST "https://api.apifox.com/api/v1/projects/$ProjectId/import-data?locale=zh-CN" \
        --header 'X-Apifox-Version: 2024-01-20' \
        --header "Authorization: Bearer $Token" \
        --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
        --header 'Content-Type: application/json' \
        --data @temp_data.json
        rm temp_data.json
        rm $file
    fi
done