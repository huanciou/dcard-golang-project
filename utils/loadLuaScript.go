package utils

import (
	"dcard-golang-project/models"
	"fmt"
	"os"
)

var LuaHash string

func LoadLuaScript() {
	scriptByte, err := os.ReadFile("bitmap.lua")
	if err != nil {
		fmt.Println("Error reading Lua script file:", err)
		return
	}
	script := string(scriptByte)

	/* 預將 Lua 腳本加載到 Redis 上，在服務器緩存。每次只發送 SHA1 減少重複送腳本的動作 */
	luaHashCmd := models.Client.ScriptLoad(models.Ctx, script)
	LuaHash = luaHashCmd.Val()
}
