package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 打开文件
	file, err := os.Open("adsl.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 读取文件内容
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return
	}
	// 将字节数组转换为字符串
	contentStr := string(content)

	// 使用分割字符串函数获取参数（这里只是一个示例，你需要根据你的实际需求进行修改）
	parameters := strings.Split(contentStr, " ")
	if len(parameters) != 4 {
		fmt.Println("文件内容格式不正确,依次顺序应该为'宽带名称' '用户名' '密码' '拨号间隔'")
		return
	}

	// 打印参数（这里只是一个示例，你需要根据你的实际需求进行修改）
	fmt.Println("宽带名称为:", parameters[0])
	fmt.Println("用户名:", parameters[1])
	fmt.Println("密码:", parameters[2])
	fmt.Println("拨号间隔时间:", parameters[3])
	pppoeCmd := fmt.Sprintf("rasdial %s %s %s", parameters[0], parameters[1], parameters[2])
	// 循环执行PPPoE拨号和断开操作
	for {
		// 执行PPPoE拨号命令
		cmd := exec.Command("cmd", "/c", pppoeCmd)
		err := cmd.Run()
		if err != nil {
			fmt.Println("PPPoE拨号失败:", err)
		} else {
			fmt.Println("PPPoE拨号成功")
		}

		// 等待一定时间后断开PPPoE连接（可选）
		t, _ := strconv.Atoi(parameters[3])
		time.Sleep(time.Duration(t) * time.Second)
		cmd = exec.Command("cmd", "/c", "rasdial", (parameters[0]), "/DISCONNECT")
		err = cmd.Run()
		if err != nil {
			fmt.Println("PPPoE断开失败:", err)
		} else {
			fmt.Println("PPPoE断开成功")
		}
	}
}
