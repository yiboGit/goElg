package main

import brisk "eglass.com/brisk/msa_rpc"

// 临时入口类 帮助生成 bind/client 文件
func main() {
	// brisk.BuildService((*UserServer)(nil), "UserInstace", "instance", ".")
	brisk.BuildClient((*UserServer)(nil), "uClient", "u_service", ".")
}
