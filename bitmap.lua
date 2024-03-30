
-- io redis 多次，運用 bitop 找 1

--[[
    使用 Lua 腳本進行 Redis IO 原因:
    1. 保持原子性. Lua 將整個腳本視為一個整體來操作, 包含使用的 Get, MGet method
       因此當其一在操作中失敗, 將會 rollback, 避免只有單一動作完成, 另一失敗的問題
    
    2. 減少 Redis IO. 如果使用 go-redis 直接操作 Redis, 多次的 query 會造成多次
       IO Redis 進而造成多餘的網路開銷, 使用 Lua 將整個事務視為一體可減少 Redis 負載
]]

local key = KEYS[1]
local indexes = {}
local m = tonumber(ARGV[1])
local n = tonumber(ARGV[2])

local startPos = 0

while true do
    local pos = redis.call('BITPOS', key, 1, startPos, -1, "bit")
    if pos == -1 then
        break
    end
    table.insert(indexes, tostring(pos))
    startPos = pos + 1
end

local slicedIndexes = {}
for i =  1+3*(m-1), math.min(#indexes, (1+3*(m-1))+2) do
    table.insert(slicedIndexes, indexes[i])
end

local values = redis.call('MGET', unpack(slicedIndexes))

return values