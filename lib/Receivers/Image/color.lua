-- get data from go
x = getX()
y = getY()
l = getL()
max = getMax()

-- helper funcs
function dist (a_x, a_y, b_x, b_y)
    return math.sqrt((a_x - b_x)^2 + (a_y - b_y)^2)
end

-- calculate values
r = 255
g = 0
b = (dist(x, y, 0, 0)/max) * 255
a = 255

-- return as table
rgba = {r,g,b,a}
return rgba