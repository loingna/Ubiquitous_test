/*
  当因网络中断造成上传失败时，此方法应该记录上传成功的分片
  失败的分片,开始记录节点总数(数组接收)，开始都为1，每成功一个节点这边的记录就减一
  恢复网络后就遍历还为1的节点
*/
type Breakpoint struct{

}

func (break *Breakpoint) Set_num(){

}

func (break *Breakpoint) Reduce_point(){

}

func (break *Breakpoint) Get_point(){

}
