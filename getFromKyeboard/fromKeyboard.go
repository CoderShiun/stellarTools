package getFromKyeboard

import "fmt"

func FromKeyboard()(publicKey string){
	fmt.Scanln(&publicKey) //扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行
	// fmt.Scanf("%s %s", &firstName, &lastName)    //Scanf与其类似，除了 Scanf 的第一个参数用作格式字符串，用来决定如何读取。
	return  publicKey
}