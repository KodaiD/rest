# rest

掲示板用apiです。
ユーザデータと投稿データをjsonで返します。  

## データ構造
user {"id", "name"}  
post {"id", "title", "content", "author_id"}  

## 使用技術
・go  
・gin, gorm  
・postgresSQL
