package bingoserver

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	// "strconv"
	"fmt"
	//"time"

	//"github.com/fsnotify/fsnotify"

)

func Start(){

	for idx,args := range os.Args{
		// if len(os.)
		if len(os.Args) > 1{
			if idx == 1{
		
			switch {
			case args == "new":
				fmt.Println("'bingo new project' is under develpoment.")
			case args == "help":
				fmt.Println("'bingo help' is under develpoment.")
			default:
				// fmt.Println(args)
				StartServer(args)
				Watch(args)
		
			}
		}
		}else{
			fmt.Println("welcome use bingo! you can input 'bingo path' to start your project.")
			return
		}



		// if idx >= 1{
		// 	// fmt.Println("参数"+strconv.Itoa(idx)+":",args)

		// 	// fmt.Println(os.Args[1])
		// 	str,err := exec_shell(os.Args[1])
		// 	if err != nil {
		// 		fmt.Println("err = ",err)
		// 	}
		// 	fmt.Println(str)
		// }

	}


}

//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func exec_shell(s string) (string, error){
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("go", "run", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	checkErr(err)


	return out.String(), err
}

func checkErr(err error){
	if err != nil {
		fmt.Println("err = ",err)
	}
}

func GetFileModTime(path string)int64 {
	f,err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		//return time.Now().Unix()
		return 0
	}
	defer f.Close()

	fi,err := f.Stat()
	if err != nil {
		log.Println("stat file info error")
		//return time.Now().Unix()
		return 0
	}
	return fi.ModTime().Unix()
}
