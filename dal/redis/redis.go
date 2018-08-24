package redis

func HGet(name,key string)(result interface{},err error){
	rd:=Pool.Get();
	defer rd.Close();
	result,err=rd.Do("HGET",name,key)
	return
}
func HSet(name string,key string,val interface{})  error{
	 rd:=Pool.Get();
	 defer rd.Close();
	 _,err:=rd.Do("HSET",name,key,val)
	 return err
}
func HDel(name,key string)  error{
	rd:=Pool.Get();
	defer rd.Close();
	_,err:=rd.Do("HDEL",name,key)
	return err
}