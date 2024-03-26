# Dcard-Golang-Project 1.0.0

## Error Handler

將 `errorHandler` 作為 global middleware 置於 routes 之前，當在 routes 中出現
一些預期之外的 error 時，製造 `panic`，此時的 error 會追溯 stack 中最近的 `recovery`
處理問題。於是 `errorHandler` 作為全局的 middleware 便能起到作用。

而我們利用 switch case 將 error 分為幾類：

- `Validation Error`:
  - Status Code: 400 處理用戶驗證錯誤，回傳給用戶 Validation Error。
- `ServerInternalError`
  - Status Code: 500 處理伺服器端異常錯誤，回傳給用戶 Server Internal Error。
- `Error`
  - Status Code: 500 處理非預期的 Runtime Error，回傳給用戶 Unexpected Error。

## Logger

將 `Logger` 作為 global middleware 置於 rotues 之前，並且當 `errorHandler` 在獲取
panic 時，除了回傳給用戶對應的錯誤 status code 以外，並將`實際錯誤內容`寫入當日的 log 檔之中。
同時將 Logger Level 設定為 `Info` 級別，紀錄所有 Info 等級以上至日誌當中。

而我們對應不同種類的 Error 給予不同級別的分級，以下提及本專案實際使用的級別：

- `Info`: 紀錄用戶正常運行狀態
  - 當 GET, POST 廣告輸入 invalid params 時，紀錄。
- `Warning`
  - 當某些程序未依照期待執行時，如 db 連線失敗、redis 連接失敗等，紀錄。
- `Error`
  - 當出現非預期的 Runtime Error 時，紀錄。

## Validator

預先定義 GET, POST method 傳入的 Struct format。當用戶傳入非定義的格式時，將回傳
Validation Error 給予用戶。

## Query N+1 Problem?

Q: 確認 Orm 對 DB 操作的 query 語句 是 join 多張 Table 來達到 query
而非對 DB 多次 io 來達到 query 結果

db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
Logger: logger.Default.LogMode(logger.Info), // 设置日志级别为 Info，以便打印 SQL 语句
})
