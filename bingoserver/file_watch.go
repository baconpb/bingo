package bingoserver

import (
	//"time"
	"os"

	//"context"
	"log"
	"os/exec"

	//"syscall"
	//"time"

	//"time"

	//"os/exec"
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var (
	Pid int64
)

// func main() {
// 	watch()


// }
func Watch(path string){
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {

		for {
			flag := 1
			select {
			case event := <-watcher.Events:
				// log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					ignoreFile := []string{".idea",".git","vendor","debug","error","log",".log"}
					// ignorePost := []string{".log"}
					for i := 0;i<len(ignoreFile);i++ {
						// if event.Name != ignoreFile[i] {
						if strings.Contains(event.Name,ignoreFile[i]) == true {
							flag  = 0
						}
					}
					// for i := 0;i<len(ignorePost);i++ {
					// 	// ret := event.Name.IndexOf(ignorePost[i])
					// 	if strings.Contains(event.Name,ignorePost[i]) ==true {
					// 		flag  = 0
					// 		log.Println(event.Op)
					// 	}
					// }

					
					// log.Println("modified file:", event.Name)
				}
				if flag == 1{
					RestartServer(path)
				}
	
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	//err = watcher.Add("/path/to/file2")//也可以监听文件夹
	//if err != nil {
	//	log.Fatal(err)
	//}
	<-done
}

func StartServer(path string){
	cmd := exec.Command("go", "run", path) //重启命令根据自己的需要自行调整
	cmd.Stdout = os.Stdout
	cmd.Start()

	Pid = int64(cmd.Process.Pid)
}
func RestartServer(path string) {
	
	if Pid != 0 {
		kill := exec.Command("taskkill", "/T","/F","/PID", fmt.Sprint(Pid))
		kill.Stderr = os.Stderr
		kill.Stdout = os.Stdout
		kill.Run()
	}
	cmd := exec.Command("go", "run", path) //重启命令根据自己的需要自行调整
	cmd.Stdout = os.Stdout
	cmd.Start()

	Pid = int64(cmd.Process.Pid)

}



