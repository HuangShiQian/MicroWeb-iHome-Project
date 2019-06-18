package handler

import (
	"context"
    "sss/IhomeWeb/utils"
	"fmt"
	example "sss/PostAvatar/proto/example"
	"path"
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/orm"
	"sss/IhomeWeb/models"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) PostAvatar(ctx context.Context, req *example.Request, rsp *example.Response) error {
	fmt.Println(" 上传头像  PostAvatar  /api/v1.0/user/avatar ")
	/*1 初始化返回值*/
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg =utils.RecodeText(rsp.Errno)

	/*2 数据对比*/
	if req.Filesize != int64(len(req.Buffer)){
		fmt.Println("diyi web")
		fmt.Println("接受数据异常")
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg =utils.RecodeText(rsp.Errno)
		return nil
	}

	/*3 上传图片到fastdfs*/
	//Ext函数返回path文件扩展名。返回值是路径最后一个斜杠分隔出的路径元素的最后一个'.'起始的后缀（包括'.'）。如果该元素没有'.'会返回空字符串。
	ext:=path.Ext(req.Filename)//.jpg
	fileid,err:=utils.Uploadbybuf(req.Buffer,ext[1:])//ext[1:] 意思是去掉 “.jpg” 里的 “.”
	if err!=nil{
		fmt.Println("diyi web")
		fmt.Println("文件上传fastdfs失败")
		rsp.Errno = utils.RECODE_IOERR
		rsp.Errmsg =utils.RecodeText(rsp.Errno)
		return nil
	}
	/*4 连接redis*/
	bm,err:=utils.RedisOpen(utils.G_server_name,utils.G_redis_addr,utils.G_redis_port,utils.G_redis_dbnum)
	if err!=nil{
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg =utils.RecodeText(rsp.Errno)
		return nil
	}

	/*5 拼接key*/
	key:=req.Sessionid+"user_id"

	/*6 查询user_id*/
	value:=bm.Get(key)
	value_int,_:=redis.Int(value,nil)
	/*7 连接mysql*/
	o:=orm.NewOrm()
	user:=models.User{Id:value_int,Avatar_url:fileid}

	/*8 将数据更新到 表中*/
	_,err=o.Update(&user,"avatar_url")
	if err !=nil{
		fmt.Println("更新数据库 文件地址失败 ",err)
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg =utils.RecodeText(rsp.Errno)
		return nil
	}

	/*9  返回fileid*/
	rsp.Fileid=fileid

	return nil
}
