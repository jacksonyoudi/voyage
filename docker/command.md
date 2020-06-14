### 常用命令

1. docker images
2. docker contains
3. docker compose



### 数据卷
作用：
    
    数据持久化
    容器和宿主机数据共享
 
 docker run -d  -v /host_path:/container_path images_name
 docker run -d  -v /host_path:/container_path:ro images_name
 