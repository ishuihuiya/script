package bash

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	//GetBashRet("")
	//GetBashRet()
	GetBashRet("ping -c 1 www.baidu.com")
	TodoBash("touch /home/kala/Desktop/22222222ds.txt")
}

// 获取cmd的返回值
func GetBashRet(cmd string) {
	c := exec.Command("bash", "-c", cmd)
	output, err := c.Output()
	if err != nil {
		fmt.Println("exec shell error, check please ...")
	}
	fmt.Println(string(output))
}

// 去除换行符号为空格
func GetBashRetNoStyle(cmd string) {
	c := exec.Command("bash", "-c", cmd)
	output, err := c.Output()
	if err != nil {
		fmt.Println("exec shell error, check please ...")
	}
	temp := strings.ReplaceAll(string(output), "\n", " ")
	fmt.Println(temp)
}

// 可接受键盘输入输出的bash
func ExecBashWithInput() {
	// Subsequent updates
}

// 非堵塞的bash
func TodoBash(cmd string) {
	c := exec.Command("bash", "-c", cmd)
	output := c.Start()
	fmt.Println(output)
}
