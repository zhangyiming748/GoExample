package aboutCMD

import (
	"fmt"
	"os/exec"
)

func pid(pid string) {
	cmd := exec.Command("netstat", "-nap", "|", "grep", pid)
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		fmt.Printf("cmd.StdoutPipe产生的错误:%v", err)
	}
	if err = cmd.Start(); err != nil {
		fmt.Printf("cmd.Run产生的错误:%v", err)
	}

	// 从管道中实时获取输出并打印到终端
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		fmt.Println(string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		fmt.Sprintf("命令运行期间产生的错误:%v\n", err)

	}
}
