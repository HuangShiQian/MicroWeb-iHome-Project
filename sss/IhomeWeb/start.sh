#redis 启动

redis-server  ./conf/redis.conf

#以下两步启动fastdfs
#启动trackerd
fdfs_trackerd /home/itcast/workspace/go/src/sss/IhomeWeb/conf/tracker.conf  restart
#启动storaged
fdfs_storaged /home/itcast/workspace/go/src/sss/IhomeWeb/conf/storage.conf  restart

#nginx 启动
sudo nginx
