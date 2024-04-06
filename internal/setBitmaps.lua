--[[
    使用 Lua 腳本進行 Redis IO 原因:
    1. 保持原子性. Lua 將整個腳本視為一個整體來操作, 包含使用的 Get, MGet method
       因此當其一在操作中失敗, 將會 rollback, 避免只有單一動作完成, 另一失敗的問題
    
    2. 減少 Redis IO. 如果使用 go-redis 直接操作 Redis, 多次的 query 會造成多次
       IO Redis 進而造成多餘的網路開銷, 使用 Lua 將整個事務視為一體可減少 Redis 負載
]]

local m = tonumber(ARGV[3])
local n = tonumber(ARGV[4])

local queryConditions = {}
for condition in string.gmatch(ARGV[1], "[^,]+") do
    table.insert(queryConditions, condition)
end

local queryAge = {}
for age in string.gmatch(ARGV[2], "[^,]+") do
    table.insert(queryAge, age)
end

redis.call("BITOP", "OR", "queryAge", unpack(queryAge))
redis.call("BITOP", "AND", "queryConditions", unpack(queryConditions))
redis.call("BITOP", "AND", "result", "queryAge", "queryConditions")

local bitMapData = redis.call('GET', "result")

return bitMapData