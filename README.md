泛在存储测试 ：含有三个分支：客户端、代理服务器、存储节点 
======
大家在上传自己的代码的时候，首先要在variable.go文件里面添加你上传的方法里的用到的全局变量，并给设置一个SET()和GET()方法  且记得注释
------
#下面的方法不能单独实现功能时，衍生出的方法命名为   xxxx_衍生函数名字()  xxxx是你要实现功能的名字  上传到对应分支下的文件夹中

客户端的方法要求：                                          
UpLoaded() //文件上传                        
DownLoaded()//文件下载                        
Updated() //文件修改                        
Deleted()//文件删除                          
Selected()//文件查询                        
Copyed()//文件复制                            
Erasure_Coding()//纠删码                      
Encryption()//文件加密                      
Section()//文件切片                          
Combination()//整合处理后的数据              
Transaction()//交易模块                    
Set_Date()//设置客户端的变量                
Get_Date()//获取客户端的变量                 
Req_proxy()//请求网络代理                   
Res_proxy()//接收网络代理的数据                
Req_node()//请求存储节点                      
Res_node()//接收存储节点的数据                 
Send_proxy()//向代理发送数据                  
Send_node()//向存储节点发送数据                
Breakpoint_con()//断点续传                    
Register_ID()//注册账号                       
Authentication()//客户端认证                
Decrypt()//解密数据                        
Administration()//管理权限、信誉等           
Cache_file()//数据缓存                       
Inter_Node()//加密通信                      

#代理服务器：
Stun_NAT()//NAT穿透   
Chosee_Node()//选择存储节点  
Audit_Node()//审计存储节点   
Datebase_con()//数据库交互   
pre_online()//保存在线分片   
Send_clint()//向客户端发送数据  
Send_Node()//向存储节点发送数据   
Res_clint()//接收客户端数据    
Res_node()//接收存储节点数据    
Bandwidth_Adn()//管理带宽消耗   
Metadata_pre()//元数据保存    
Metadata_rep()//元数据修复   
Metadata_sel()//访问元数据   
Metadata_upd()//修改元数据   
Metadata_del()//删除元数据   
Transaction()//交易模块   
Admin_Register()//管理注册、分配ID   
IPFS()//留给IPFS   
Judge_clint()//判断客户端的权限等   
Adm_upload()//响应上传    
Adm_download()//响应下载   
Adm_update()//响应更新    
Adm_deleted()//响应删除    
Adm_copy()//响应复制     
Adm_select()//响应查询    
Reputation()//管理代理信誉    
Garbage_col()//垃圾回收     
Set() Get()//设置和获取变量        
                                              
#存储节点：    
Storage_clint()//存储客户端的文件    
Res_proxy()//接收代理数据     
Res_clint()//接收客户端数据     
Send_proxy()//向代理发送数据     
Send_clint()//向客户端发数据      
Select_file()//根据节点ID查询数据    
Register_ID()//注册账号         
Choose_register()//留有选择注册   
Bandwidth_set()//跟踪带宽消耗    
Transaction()//费用结算    
Inter_clint()//加密通信     
Res_asudit()//响应审计    
Adm_respution()//管理信誉    
Database_Node()//数据库的交互    
Deleted_file()//响应删除   
Set() Get()//设置和获取变量    
Node_find()//节点发现      








