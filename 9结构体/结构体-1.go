package main

import "fmt"

type Books struct {
	title, author, subject string
	book_id int
}

func (b Books)String()string{
	return `格式化:`+`[`+b.title+`]`
}

// 创建结构题
func test1()  {
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 非法
	// fmt.Println(Books{"Go 语言", "www.runoob.com"})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

}
// 结构体访问
func test2()  {
	var Book1 Books        /* 声明 Book1 为 Books 类型 */
	var Book2 Books        /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf( "Book 1 title : %s\n", Book1.title)
	fmt.Printf( "Book 1 author : %s\n", Book1.author)
	fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
	fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf( "Book 2 title : %s\n", Book2.title)
	fmt.Printf( "Book 2 author : %s\n", Book2.author)
	fmt.Printf( "Book 2 subject : %s\n", Book2.subject)
	fmt.Printf( "Book 2 book_id : %d\n", Book2.book_id)
}

// 结构体传参
func printBook1( book Books ) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
}
func printBook2( book *Books ) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
}
func test3()  {
	book := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	printBook1(book)
	printBook2(&book)
	fmt.Println(book)
}
func main() {
	//test1()
	//test2()
	test3()
}
