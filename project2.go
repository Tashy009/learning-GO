package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
	
	var filename string
    var filemap = make(map[string]bool)
	fmt.Scan(&filename)
	// write your code here
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    scanner.Split(bufio.ScanWords)  // split each scanned line into words
    
    
   
    for scanner.Scan() {
        filemap[string(scanner.Text())] = true
        
        
        //fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
    
    for true{
        var key string
        fmt.Scan(&key)
        keymap := strings.Split(key, " ")
        if key == "exit"{
            break
        }else{
            for index, element := range keymap {
                if _, ok := filemap[strings.ToLower(element)]; ok {
                    out := []rune(element)
            
                    for i := range out {
                
                        out[i] = '*'
                    } 
                    keymap[index] = string(out)
                
                    
                    
                    
                }
                
            }  
            str2 := strings.Join(keymap, " ") 
            fmt.Println(str2)
            
            

        /* val := filemap[strings.ToLower(key)]
        out := []rune(key)
        if val == true{
            for i := range out {
            
                out[i] = '*'
            } 
            fmt.Println(string(out))
            
            continue
                
        }else if key == "exit"{
            break
        }else{
            fmt.Println(key)
            continue
 */
        }
    }
    fmt.Println("Bye")
    
}
