package store

import "log"

var IdEmpMap map[int]Employee
var DeptEmpMap map[string]*([]Employee)
var LocEmpMap map[int]*([]Employee) 
type Address struct{
	Doorno int
	Street string 
	Locality string 
	PIN int 
	
	}

	//helper methods for Address
	func (a *Address) getPin() int{

		return a.PIN
	  }

	  func (a *Address) GetDoorNo() int{

		return a.Doorno
	  }


	  func (a *Address) GetStreet() string{

		return a.Street
	  }

	  func (a *Address) GetLocality() string{

		return a.Locality
	  }


	type Employee struct {
		EmpID int 
		Name string
		Department []string
		Addresses []Address
		There bool
	  }

//helper methods for Employee
	  func (e *Employee) PrintName(){

		log.Println(e.Name)
	  }

	  func (e *Employee) GetID() int{

		return e.EmpID
	  }

	  func (e *Employee) GetPins() []int{

		var pins []int
		
		for _, addr := range e.Addresses{

				pins = append(pins, addr.getPin())

		}

		return pins
	  }

	  func (e *Employee) GetDoorNos() []int{

		var doornos []int
		
		for _, addr := range e.Addresses{

				doornos = append(doornos, addr.GetDoorNo())

		}

		return doornos
	  }

	  func (e *Employee) GetStreets() []string{

		var streets []string
		
		for _, addr := range e.Addresses{

				streets = append(streets, addr.GetStreet())

		}

		return streets
	  }

	  func (e *Employee) GetLocalities() []string{

		var localities []string
		
		for _, addr := range e.Addresses{

			localities = append(localities, addr.GetLocality())

		}

		return localities
	  }

	  func (e *Employee) GetDept() []string{

		return e.Department
	  }


	  func StoreEmployeesByIdDeptAndLoc(employees *([]Employee)){


	  //stores ids as keys and employees as values
	  IdEmpMap = storeEmployeesByIDMap(employees)

	  //stores department names (strings) as keys and reference to array of employees as values
	 DeptEmpMap = storeEmployeesByDeptMap(employees)
  
	  //stores PIN codes (int) as keys and reference to array of employees as values
  
	 LocEmpMap = storeEmployeesByLocMap(employees)
  
	  }




func storeEmployeesByDeptMap(employees *([]Employee)) map[string]*([]Employee){

	deptEmpMap := make(map[string]*([]Employee))

	for _, empl:= range *employees{

		deptList  := empl.GetDept()


		for _, dept := range deptList{

			_, ok := deptEmpMap[dept]

			var emplswithThisDept *([]Employee)

			if(ok){

			emplswithThisDept = deptEmpMap[dept]

			} else{

				
				emptySlice := make([]Employee, 0)
				emplswithThisDept = &(emptySlice)
			}

			*emplswithThisDept = append(*emplswithThisDept, empl)
			deptEmpMap[dept] = emplswithThisDept

		}
	}

		return deptEmpMap
}

func storeEmployeesByIDMap(employees *([]Employee)) map[int]Employee{

	idEmpMap := make(map[int]Employee)

	for _, empl:= range *employees{

	
		idEmpMap[empl.GetID()] = empl
		

	}


		return idEmpMap
}

func storeEmployeesByLocMap(employees *([]Employee)) map[int]*([]Employee){

	locEmpMap := make(map[int]*[](Employee))

	for _, empl:= range *employees{

		pins := empl.GetPins()

		for _, pin := range pins{

			_, ok := locEmpMap[pin]

			var emplswithThisPin *([]Employee)

			if(ok){

			emplswithThisPin = locEmpMap[pin]

			} else{

				emptySlice := (make([]Employee, 0))

				emplswithThisPin = &emptySlice
			}

			*emplswithThisPin = append(*emplswithThisPin, empl)
			locEmpMap[pin] = emplswithThisPin


		}

	}

		return locEmpMap
}