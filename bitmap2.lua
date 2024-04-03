local m = tonumber(ARGV[2])
local n = tonumber(ARGV[3])

local indexes = {}
for condition in string.gmatch(ARGV[1], "[^,]+") do
    table.insert(indexes, condition)
end

local slicedIndexes = {}
-- for i = 1+3*(m-1), math.min(#indexes, 1+3*(m-1)+2, n) do
--     table.insert(slicedIndexes, indexes[i])
-- end

for i = 1, n do
    table.insert(slicedIndexes, indexes[i])
end

local values = redis.call('MGET', unpack(slicedIndexes))

return values
