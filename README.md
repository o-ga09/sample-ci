# Sample CI/CD Repository

<img src="./doc/image.png">
<div align="center">CI/CDの参考リポジトリ</div>

## ✅ GitHub Actions 用です

### ci_backend.yml

GoのCI用パイプライン

### ci_frontend_bun.yml

BunのCI用パイプライン

### ci_frontend_node.yml

NodeのCI用パイプライン
テストカバレッジレポート付き

package.jsonの以下の部分を追記すること

script
```json
    "test": "jest",
    "test:coverage": "npm run test -- --coverage --silent --testLocationInResults --ci --json --outputFile=\"report.json\""
```


```json
"jest": {
    "transform": {
      "^.+\\.(t|j)sx?$": "@swc/jest"
    },
    "coverageReporters": [
      [
        "text",
        {
          "file": "report.json"
        }
      ]
    ]
  }
```

### tag_release.yml

tagリリース用のCDパイプライン

### deploy_aws.yml

AWSにデプロイ用のパイプライン

### deploy_gcp.yml

GCPのCloud Runにデプロイ用のパイプライン
