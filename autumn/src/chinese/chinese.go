package main
 
import (
	"bufio"
	"fmt"
	"io"
	"flag"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
)



func getFilelist(path string) []string{
    mySlice1 := []string{}
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
        if f == nil {
			return err
		}
        if f.IsDir() {
			return nil
		}
        mySlice1 = append(mySlice1,path)
        return nil
        })
    if err != nil {
		// fmt.Println(path)	
     }
	 return mySlice1
}



func check(e error) {
    if e != nil {
        panic(e)
    }
}

//19968  40869
func main() {
	flag.Parse()
    root := flag.Arg(0)
	output := flag.Arg(1)
	file := getFilelist(root)

	// for i,_ := range(file){
		// fmt.Println(file[i])
	// }
	strSlice := []string{}
	cnstr := ""
	lines := ""
	for _,v := range(file) {
	fmt.Println(v)

	fi, err := os.Open(v)
	
    if err != nil {
        panic(err)		
    }
	defer fi.Close()
	
	rd := bufio.NewReader(fi)
    for {
        line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
        
        if err != nil || io.EOF == err {
            break
        }
			i := strings.Contains(line,"//")
			k := strings.Contains(line,"<!--")
			j := strings.Contains(line,"/*")
			if k == false{
				if i == false && j == false{
					lines += line
				}else{
					linesplit := strings.Split(line,"//")[0]
					linesplit = strings.Split(linesplit,"/*")[0]
					lines += linesplit
				}
			}
            // fmt.Println(line)
              }      
	
	// b := make([]byte,15)
	// fi.Read(b)
	// check(err)
	// fmt.Printf("%d bytes: %s\n", n1, string(b))
	
    // str := string(b)
    // fmt.Println(str)
	
	
	fi.Close()
	
	}
	r := []rune(lines)
	//fmt.Println("rune=", r)
	
	
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			cnstr = cnstr + string(r[i])
			
			if (i != len(r)-1) && (r[i+1] > 40869 || r[i+1] < 19968){
				cnstr = cnstr + "\n"
				// fmt.Println(cnstr)
			}
			
			strSlice = append(strSlice, cnstr)
		}
	}
	fmt.Println(len(strSlice))
	d1 := []byte(strSlice[len(strSlice)-1])
	// d1 := []byte(strSlice[len(strSlice)-1])
    err2 := ioutil.WriteFile(output, d1, 0644)
    check(err2)
}