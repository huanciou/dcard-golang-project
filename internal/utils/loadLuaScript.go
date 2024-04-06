package utils

import (
	"dcard-golang-project/models"
	"fmt"
	"os"
)

var LuaHash1 string
var LuaHash2 string

func LoadLuaScript() {
	scriptByte, err := os.ReadFile("setBitmaps.lua")
	if err != nil {
		fmt.Println("Error reading Lua script file:", err)
		return
	}
	scriptByte2, err := os.ReadFile("mget.lua")
	if err != nil {
		fmt.Println("Error reading Lua script file:", err)
		return
	}
	script := string(scriptByte)
	script2 := string(scriptByte2)

	/* 預將 Lua 腳本加載到 Redis 上，在服務器緩存 SHA1 減少重複送腳本的動作 */
	luaHashCmd := models.Client.ScriptLoad(models.Ctx, script)
	LuaHash1 = luaHashCmd.Val()

	luaHashCmd2 := models.Client.ScriptLoad(models.Ctx, script2)
	LuaHash2 = luaHashCmd2.Val()
}
