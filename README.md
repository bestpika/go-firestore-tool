# 使用說明

## 登入

- <https://firebase.google.com/docs/firestore/server/setup-go?hl=zh-cn> `沒有正體中文說明`

```bash
gcloud auth application-default login
```

## 路徑格式

```bash
WORKDIR/projects/PROJECT/databases/(default)/documents/COLLECTION/DOCUMENT.json
```

## 指令

- `-p=PROJECT`: 要操作的專案*
- `-e=PATH`: 匯出路徑
- `-i=PATH`: 匯入路徑
- `-m=yes/no`: 是否合併

### 範例

```bash
.\firestore.exe -p=project -e=backup # 匯出
.\firestore.exe -p=project -i=backup -m=no # 匯入
```

## 注意

- 目前僅支援 root 底下的 collection
