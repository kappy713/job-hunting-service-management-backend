# 開発者用ドキュメント

## API実装手順（SampleUserを参考にした新エンティティの実装）

このドキュメントでは、既存の`SampleUser`を参考にして新しいエンティティのAPI実装を行う手順を説明します。

### 📋 実装手順概要

クリーンアーキテクチャに従って、以下の順序で実装します：

1. **Entity（エンティティ）**: データ構造の定義
2. **Repository（リポジトリ）**: データアクセス層
3. **Usecase（ユースケース）**: ビジネスロジック層
4. **Handler（ハンドラー）**: HTTPリクエスト/レスポンス層
5. **Router（ルーター）**: エンドポイント設定
6. **Migration（マイグレーション）**: データベーステーブル作成
7. **DI（依存性注入）**: main.goでの組み立て

### 🚀 実装手順詳細

#### 1. Entity（エンティティ）の作成

エンティティ=テーブル、と一旦は認識していただいて構いません。<br>
新しいテーブル定義やモデルを追加する場合にエンティティを作成します。

**ファイル**: `app/internal/entity/{entity_name}.go`

```go
...
// 例：Task エンティティの場合
type Task struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	...
}

func (Task) TableName() string {
	return "tasks"
}
```

**ポイント**:
- GORMタグでDB制約を定義
- JSONタグでAPIレスポンス形式を指定
- `TableName()`メソッドでテーブル名を明示

#### 2. Repository（リポジトリ）の作成

**ファイル**: `app/internal/repository/{entity_name}_repository.go`

```go
...
type TaskRepository interface {
	GetAllTasks(c *gin.Context) ([]entity.Task, error)
	CreateTask(c *gin.Context, task *entity.Task) error
    ...
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetAllTasks(c *gin.Context) ([]entity.Task, error) {
	var tasks []entity.Task
	result := r.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *taskRepository) CreateTask(c *gin.Context, task *entity.Task) error {
	result := r.db.Create(task)
	return result.Error
}
...
```

**ポイント**:
- インターフェースでAPIを定義
- コンストラクタ関数でDIを実現
- CRUD操作を実装

#### 3. Usecase（ユースケース）の作成

**ファイル**: `app/internal/usecase/{entity_name}_usecase.go`

```go
...
type TaskUsecase interface {
	GetAllTasks(c *gin.Context) (*[]entity.Task, error)
	CreateTask(c *gin.Context, task *entity.Task) error
    ...
}

type taskUsecase struct {
	tr repository.TaskRepository
}

func NewTaskUsecase(r repository.TaskRepository) TaskUsecase {
	return &taskUsecase{tr: r}
}

func (u *taskUsecase) GetAllTasks(c *gin.Context) (*[]entity.Task, error) {
	tasks, err := u.tr.GetAllTasks(c)
	if err != nil {
		return nil, err
	}
	
	// ここでビジネスロジックを追加
	// 例：フィルタリング、ソート、データ変換など
	
	return &tasks, nil
}

func (u *taskUsecase) CreateTask(c *gin.Context, task *entity.Task) error {
	// バリデーションやビジネスルールをここに追加
	return u.tr.CreateTask(c, task)
}
...
```

**ポイント**:
- ビジネスロジックを実装
- Repositoryに依存
- バリデーションやデータ変換を担当

#### 4. Handler（ハンドラー）の作成

**ファイル**: `app/internal/handler/{entity_name}_handler.go`

```go
...
type TaskHandler interface {
	GetAllTasks(c *gin.Context)
	CreateTask(c *gin.Context)
}

type taskHandler struct {
	tu usecase.TaskUsecase
}

func NewTaskHandler(u usecase.TaskUsecase) TaskHandler {
	return &taskHandler{tu: u}
}

func (h *taskHandler) GetAllTasks(c *gin.Context) {
	tasks, err := h.tu.GetAllTasks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *taskHandler) CreateTask(c *gin.Context) {
	var task entity.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.tu.CreateTask(c, &task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, task)
}
...
```

**ポイント**:
- HTTP固有の処理を担当
- JSONバインディングとレスポンス
- エラーハンドリング

#### 5. Router（ルーター）の更新

**ファイル**: `app/internal/router/router.go`

```go
func NewRouter(
	suh handler.SampleUserHandler,
	th handler.TaskHandler, // 新しいハンドラーを追加
) *gin.Engine {
	...
	// 新しいエンドポイント
	taskRoutes := r.Group("/api/tasks")
	{
		taskRoutes.GET("", th.GetAllTasks)
		taskRoutes.POST("", th.CreateTask)
        ...
	}

	return r
}
```

#### 7. DI（依存性注入）の更新

**ファイル**: `app/cmd/server/main.go`

```go
func main() {
    ...
	// 新しいエンティティのDI
	taskRepository := repository.NewTaskRepository(database)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	taskHandler := handler.NewTaskHandler(taskUsecase)

	// ルーター設定
	r := router.NewRouter(sampleUserHandler, taskHandler)
    ...
}
```

### 🧪 テスト方法

#### 1. サーバー起動
```bash
go run app/cmd/server/main.go
```

#### 3. API動作確認

**GET /api/tasks**
```bash
curl http://localhost:8080/api/tasks
```

**POST /api/tasks**
```bash
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"新しいタスク","description":"タスクの説明","priority":1}'
```


### 📝 注意点

1. **import文のパス**: モジュール名に注意
2. **エラーハンドリング**: 適切なHTTPステータスコードを返す
3. **バリデーション**: 必要に応じてUsecaseで実装
4. **トランザクション**: 複雑な処理では考慮する
5. **ログ出力**: デバッグ用のログを適切に配置

### 🎯 ベストプラクティス

- インターフェースファーストで設計
- 各層の責任を明確に分離
- エラーハンドリングを適切に実装
- テスト可能な構造を維持
- ドキュメントを適切に更新

この手順に従うことで、SampleUserと同様の品質でAPIを実装できます。
