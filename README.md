# ToDoアプリの設計（GO）

### 機能
- ユーザー登録・認証
- TodoのCRUD
- 状態の更新

### 技術スタック
- バックエンド：Gin(GO)
- データベース：PostgreSQL
- ORM：GORM（DB操作）
- 認証：JWT
- フロントエンド：HTML/CSS + JS

### API設計
- POST /register 新規ユーザー登録
- POST /login ログイン（JWTトークンの発行）
- GET /todos ToDo一覧取得（要認証）
- POST /todos 新しいToDoを作成
- GET /todos/:id 特定のToDoを取得
- PUT /todos/:id ToDoを更新
- DELETE /todos/:id ToDoを削除
