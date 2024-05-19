-- load the random module
math.randomseed(os.time())

-- define the different paths you want to test
local paths = {
	"/",
	"/login",
	"/score/",
	"/api/scoresearch?query=lel&from=&to=",
	"/me",
}

-- select a random path for each request
request = function()
	local path = paths[math.random(#paths)]
	if path == "/score/" then
		-- If the path is "/score/", generate a random score from 1 to 1000
		path = path .. math.random(1, 1000)
	end
	return wrk.format(nil, path)
end
