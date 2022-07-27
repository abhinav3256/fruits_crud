package pkg

func FruitInsertService(reqbody Fruit) bool {

	res := FruitInsert(reqbody)

	return res
}

func GetFruitsService() []Fruit {

	data := GetFruits()
	return data
}

func UpdateFruitsService(reqbody Fruit) bool {

	res := FruitUpdate(reqbody)

	return res
}

func DeleteFruitsService(id int) bool {

	res := FruitDelete(id)

	return res
}
