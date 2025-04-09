// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Base struct {
	name string
}


type Child struct {
	Base
	lastName string
}
type Family interface{
	Say()
}


func main() {
	b1 :=&Base{}
	b1.name="Parent"

	c1 :=&Child{}
	c1.name="Child"
	c1.lastName="Inherited"

	b1.Say()
	c1.Say()

	slc:=make([]Family,0)
	slc = append(slc, b1,c1)
	for k,_:= range slc{
		slc[k].Say()
	}
	v1,err:=NewObject(b1)
	if err!=nil {
		fmt.Printf("NewObject errr")
	}
	v2,err:=NewObject(c1)
	if err!=nil {
		fmt.Printf("NewObject errr")
	}
	v1.Say()
	v2.Say()
}

func(b *Base) Say(){
	fmt.Printf("Hello, %v!\n",b.name)
}

func(c *Child) Say(){
	fmt.Printf(" Hello, %v %v!\n",c.lastName,c.name)
}

func NewObject(f Family) (Family,error) {
	switch f.(type){
	case *Base:
		b1 :=&Base{}
	b1.name="Parent"
		return b1,nil;
	case *Child:
		c1 :=&Child{}
	c1.name="Child"
	c1.lastName="Inherited"
		return c1,nil;
	}
	return nil,fmt.Errorf("error NewObject ")
}

/*
   1. Сделать структуры Base и Child.
   2. Структура Base должна содержать строковое поле name.
   3. Структура Child должна содержать строковое поле lastName.
   4. Сделать функцию Say у структуры Base, которая распечатывает на экране: Hello, %name!
   5. Пронаследовать Child от Base.
   6. Инициализировать экземпляр b1 Base.
       присвоить name значение Parent
   7. Инициализировать экземпляр c1 Сhild.
       присвоить name значение Child
       присвоить lastName значение Inherited
   8. Вызвать у обоих экземпляров метод Say.
   9. Переопределить метод Say для стурктуры Child, чтобы он выводил на экран: Hello, %lastName %name!
   10. Сделать массив, содержащий b1 и c1.
   11. Вызвать Say у всех элементов массива из шага 10.
   12. Сделать метод NewObject для создания экземпляров Base и Child в зависимости от входного параметра.
   13. Написать юнит тесты для метода NewObject.
   14. Сделать генератор объетов Base и Child такой, чтобы:
       объекты Base создавались в фоновом потоке с задержкой 1 секунда;
       объекты Child создавались в фоновом потоке с задержкой 2 секунды;
       общее время генерации объектов не превышало 11 секунд;
   15. Сделать асинхронный обработчик сгенерированных объектов такой, чтобы:
       метод Say вызывался в порядке генерации объектов;
       не приводил к утечкам памяти;
*/
