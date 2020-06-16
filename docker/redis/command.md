docker run -p 16379:6379
-v  /Users/youdi/Project/GoProject/voyage/docker/redis/data:/data
-v /Users/youdi/Project/GoProject/voyage/docker/redis/redis.conf:/usr/local/etc/redis/redis.conf
-d redis redis-server /usr/local/etc/redis/redis.conf
--appendonly yes
