 package main

 import (
         "fmt"
         "os"
         "syscall"
 )

 func main() {
         r, w, err := os.Pipe()

         if err != nil {
                 panic(err)
         }

         defer r.Close()

         process, err := os.StartProcess("/bin/ps", []string{"-ef"}, &os.ProcAttr{Files: []*os.File{nil, w, os.Stderr}})

         if err != nil {
                 panic(err)
         }

         processState, err := process.Wait()

         if err != nil {
                 panic(err)
         }

         err = process.Release()

         if err != nil {
                 panic(err)
         }

         fmt.Println("Did the child process exited? : ", processState.Exited())
         fmt.Println("So the child pid is? : ", processState.Pid())
         fmt.Println("Exited successfully? : ", processState.Success())

         fmt.Println("Exited system CPU time ? : ", processState.SystemTime())
         fmt.Println("Exited user CPU time ? : ", processState.UserTime())

         // just to be sure, let's kill again
         err = process.Signal(syscall.SIGKILL)

         if err != nil {
                 fmt.Println(err) // see what the serial killer has to say
                 return
         }

         w.Close()

 }