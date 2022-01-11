package main

import (
    "fmt"
)

//编写一个学生考试系统
type Student struct {
	Name string
	Age int
	Score int
}

//将Pupil和Graduate共有的方法绑定到 *Student
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名：%v 年龄：%v 成绩：%v \n",stu.Name,stu.Age,stu.Score)
}
func (stu *Student) SetScore (score int){
	//业务判断
	stu.Score = score
}

//小学生
type Pupil struct {
	Student //嵌套了Student匿名结构体
}
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中..")
}

//大学生
type Graduate struct {
	Student //嵌套了Student匿名结构体
}
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}

func main() {
	//当我们对结构体嵌入了匿名结构体，使用方法会发生变化
	pupil := &Pupil{}
	pupil.Name = "tom"
	pupil.Age = 8
	pupil.testing()
	pupil.SetScore(80)
	pupil.ShowInfo()

	graduate := &Graduate{}
	graduate.Name = "mary"
	graduate.Age = 24
	graduate.testing()
	graduate.SetScore(95)
	graduate.ShowInfo()
}