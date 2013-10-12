package main

import (
	"github.com/jmoiron/modl"
)

/* This data was generated at http://www.generatedata.com/ */

var Departments = []Department{
	{DepartmentID: 1, Name: "Interdum Enim Non Foundation"},
	{DepartmentID: 2, Name: "Pede Nunc Ltd"},
	{DepartmentID: 3, Name: "Et Libero Foundation"},
	{DepartmentID: 4, Name: "Nascetur Ridiculus Mus Associates"},
	{DepartmentID: 5, Name: "Nam Industries"},
	{DepartmentID: 6, Name: "Feugiat Non Inc."},
	{DepartmentID: 7, Name: "Eget Mollis Lectus Limited"},
	{DepartmentID: 8, Name: "Duis Incorporated"},
	{DepartmentID: 9, Name: "Eu Placerat Institute"},
	{DepartmentID: 10, Name: "Sociis Natoque Limited"},
	{DepartmentID: 11, Name: "Nonummy PC"},
	{DepartmentID: 12, Name: "Mattis Semper Dui Foundation"},
	{DepartmentID: 13, Name: "Pharetra Sed PC"},
	{DepartmentID: 14, Name: "Curae; Incorporated"},
}

var Employees = []Employee{
	{EmployeeID: 1, DepartmentID: 8, BossID: 0, Name: "Jade Buchanan", Salary: 253789},
	{EmployeeID: 2, DepartmentID: 9, BossID: 1, Name: "Gretchen Benton", Salary: 232873},
	{EmployeeID: 3, DepartmentID: 4, BossID: 1, Name: "Desiree Reid", Salary: 190374},
	{EmployeeID: 4, DepartmentID: 2, BossID: 2, Name: "Deborah Mills", Salary: 203148},
	{EmployeeID: 5, DepartmentID: 4, BossID: 4, Name: "Claudia House", Salary: 293193},
	{EmployeeID: 6, DepartmentID: 7, BossID: 5, Name: "Mechelle Kramer", Salary: 164555},
	{EmployeeID: 7, DepartmentID: 8, BossID: 2, Name: "Barrett Pruitt", Salary: 248567},
	{EmployeeID: 8, DepartmentID: 4, BossID: 1, Name: "Kyra Harrell", Salary: 88734},
	{EmployeeID: 9, DepartmentID: 4, BossID: 2, Name: "Shad Chapman", Salary: 284941},
	{EmployeeID: 10, DepartmentID: 3, BossID: 5, Name: "Dahlia Wiggins", Salary: 217431},
	{EmployeeID: 11, DepartmentID: 5, BossID: 7, Name: "Oliver Pena", Salary: 88653},
	{EmployeeID: 12, DepartmentID: 10, BossID: 4, Name: "Gavin Cardenas", Salary: 227669},
	{EmployeeID: 13, DepartmentID: 6, BossID: 3, Name: "Leo Chen", Salary: 273135},
	{EmployeeID: 14, DepartmentID: 3, BossID: 1, Name: "Montana Clayton", Salary: 227003},
	{EmployeeID: 15, DepartmentID: 8, BossID: 2, Name: "Uma Daugherty", Salary: 176748},
	{EmployeeID: 16, DepartmentID: 5, BossID: 13, Name: "Lunea Cole", Salary: 147576},
	{EmployeeID: 17, DepartmentID: 5, BossID: 11, Name: "Germane Bartlett", Salary: 186167},
	{EmployeeID: 18, DepartmentID: 6, BossID: 2, Name: "Adrian Cleveland", Salary: 211932},
	{EmployeeID: 19, DepartmentID: 8, BossID: 13, Name: "Tasha Reyes", Salary: 213184},
	{EmployeeID: 20, DepartmentID: 3, BossID: 15, Name: "Teagan Alston", Salary: 159882},
	{EmployeeID: 21, DepartmentID: 5, BossID: 15, Name: "Channing York", Salary: 151167},
	{EmployeeID: 22, DepartmentID: 9, BossID: 14, Name: "Cecilia Preston", Salary: 296823},
	{EmployeeID: 23, DepartmentID: 1, BossID: 13, Name: "Wyatt Richmond", Salary: 151951},
	{EmployeeID: 24, DepartmentID: 6, BossID: 1, Name: "Willa Hayes", Salary: 72705},
	{EmployeeID: 25, DepartmentID: 4, BossID: 22, Name: "Ishmael Pace", Salary: 256135},
	{EmployeeID: 26, DepartmentID: 5, BossID: 5, Name: "Stuart White", Salary: 133126},
	{EmployeeID: 27, DepartmentID: 8, BossID: 1, Name: "Odette Underwood", Salary: 102573},
	{EmployeeID: 28, DepartmentID: 7, BossID: 4, Name: "Ursa Ruiz", Salary: 179185},
	{EmployeeID: 29, DepartmentID: 2, BossID: 28, Name: "Reuben Snider", Salary: 227786},
	{EmployeeID: 30, DepartmentID: 10, BossID: 2, Name: "Levi Mcleod", Salary: 235048},
	{EmployeeID: 31, DepartmentID: 6, BossID: 4, Name: "Hall Morrow", Salary: 118401},
	{EmployeeID: 32, DepartmentID: 4, BossID: 20, Name: "Amethyst Cervantes", Salary: 273802},
	{EmployeeID: 33, DepartmentID: 5, BossID: 31, Name: "Shelly Le", Salary: 126729},
	{EmployeeID: 34, DepartmentID: 8, BossID: 17, Name: "John Carrillo", Salary: 221224},
	{EmployeeID: 35, DepartmentID: 3, BossID: 29, Name: "Aquila Lane", Salary: 239008},
	{EmployeeID: 36, DepartmentID: 10, BossID: 4, Name: "Ferdinand Webb", Salary: 175668},
	{EmployeeID: 37, DepartmentID: 2, BossID: 4, Name: "Roanna Benton", Salary: 232190},
	{EmployeeID: 38, DepartmentID: 5, BossID: 36, Name: "Mia Cash", Salary: 128522},
	{EmployeeID: 39, DepartmentID: 1, BossID: 15, Name: "Rhonda Bowman", Salary: 254525},
	{EmployeeID: 40, DepartmentID: 9, BossID: 14, Name: "Dawn Alston", Salary: 291871},
	{EmployeeID: 41, DepartmentID: 2, BossID: 34, Name: "Kennan Flowers", Salary: 65410},
	{EmployeeID: 42, DepartmentID: 1, BossID: 10, Name: "Isadora Gilliam", Salary: 78907},
	{EmployeeID: 43, DepartmentID: 6, BossID: 21, Name: "Sigourney Black", Salary: 259164},
	{EmployeeID: 44, DepartmentID: 2, BossID: 13, Name: "Alexa Mccormick", Salary: 55258},
	{EmployeeID: 45, DepartmentID: 9, BossID: 11, Name: "Marvin Hansen", Salary: 255638},
	{EmployeeID: 46, DepartmentID: 4, BossID: 17, Name: "Kennedy Contreras", Salary: 220784},
	{EmployeeID: 47, DepartmentID: 4, BossID: 9, Name: "Edan Preston", Salary: 71886},
	{EmployeeID: 48, DepartmentID: 6, BossID: 30, Name: "Petra Levy", Salary: 249277},
	{EmployeeID: 49, DepartmentID: 3, BossID: 16, Name: "Samson Jenkins", Salary: 280726},
	{EmployeeID: 50, DepartmentID: 3, BossID: 39, Name: "Rose Macdonald", Salary: 197879},
	{EmployeeID: 51, DepartmentID: 7, BossID: 43, Name: "Magee Espinoza", Salary: 141836},
	{EmployeeID: 52, DepartmentID: 7, BossID: 25, Name: "Sierra Carpenter", Salary: 281803},
	{EmployeeID: 53, DepartmentID: 6, BossID: 43, Name: "Keely Peterson", Salary: 275559},
	{EmployeeID: 54, DepartmentID: 1, BossID: 22, Name: "Ulric Colon", Salary: 58198},
	{EmployeeID: 55, DepartmentID: 3, BossID: 40, Name: "Vaughan Mckay", Salary: 155698},
	{EmployeeID: 56, DepartmentID: 7, BossID: 54, Name: "Bruce Guzman", Salary: 285929},
	{EmployeeID: 57, DepartmentID: 10, BossID: 14, Name: "Shana Finch", Salary: 252513},
	{EmployeeID: 58, DepartmentID: 7, BossID: 20, Name: "Selma Alvarez", Salary: 229558},
	{EmployeeID: 59, DepartmentID: 5, BossID: 7, Name: "Edward Collins", Salary: 261885},
	{EmployeeID: 60, DepartmentID: 2, BossID: 39, Name: "Gail Mckinney", Salary: 231011},
	{EmployeeID: 61, DepartmentID: 9, BossID: 57, Name: "Fletcher Jensen", Salary: 217784},
	{EmployeeID: 62, DepartmentID: 3, BossID: 28, Name: "Melyssa Glenn", Salary: 85336},
	{EmployeeID: 63, DepartmentID: 6, BossID: 18, Name: "Wade Cline", Salary: 213351},
	{EmployeeID: 64, DepartmentID: 1, BossID: 3, Name: "Kitra Dotson", Salary: 140959},
	{EmployeeID: 65, DepartmentID: 5, BossID: 5, Name: "Kelsey Everett", Salary: 124117},
	{EmployeeID: 66, DepartmentID: 4, BossID: 55, Name: "Elijah Head", Salary: 187126},
	{EmployeeID: 67, DepartmentID: 6, BossID: 48, Name: "Lester Wells", Salary: 281162},
	{EmployeeID: 68, DepartmentID: 10, BossID: 54, Name: "Kylan Guerra", Salary: 235590},
	{EmployeeID: 69, DepartmentID: 9, BossID: 9, Name: "Seth Pope", Salary: 242200},
	{EmployeeID: 70, DepartmentID: 6, BossID: 8, Name: "August Maxwell", Salary: 98902},
	{EmployeeID: 71, DepartmentID: 1, BossID: 7, Name: "Phillip Hansen", Salary: 84435},
	{EmployeeID: 72, DepartmentID: 4, BossID: 51, Name: "Malachi Campbell", Salary: 216728},
	{EmployeeID: 73, DepartmentID: 2, BossID: 40, Name: "Howard Roberts", Salary: 183538},
	{EmployeeID: 74, DepartmentID: 8, BossID: 21, Name: "Charde Hanson", Salary: 128179},
	{EmployeeID: 75, DepartmentID: 8, BossID: 53, Name: "Garth Young", Salary: 270469},
	{EmployeeID: 76, DepartmentID: 10, BossID: 59, Name: "Griffin Kirkland", Salary: 102531},
	{EmployeeID: 77, DepartmentID: 3, BossID: 55, Name: "Daniel Valdez", Salary: 129869},
	{EmployeeID: 78, DepartmentID: 4, BossID: 26, Name: "Joan Schroeder", Salary: 134816},
	{EmployeeID: 79, DepartmentID: 3, BossID: 43, Name: "Cleo Holland", Salary: 290718},
	{EmployeeID: 80, DepartmentID: 9, BossID: 1, Name: "Lana Eaton", Salary: 165740},
	{EmployeeID: 81, DepartmentID: 10, BossID: 56, Name: "Amaya Lynn", Salary: 194349},
	{EmployeeID: 82, DepartmentID: 3, BossID: 2, Name: "Vernon Whitehead", Salary: 77064},
	{EmployeeID: 83, DepartmentID: 3, BossID: 17, Name: "Wylie Paul", Salary: 139793},
	{EmployeeID: 84, DepartmentID: 4, BossID: 46, Name: "Colorado Whitley", Salary: 278346},
	{EmployeeID: 85, DepartmentID: 5, BossID: 65, Name: "Susan Dunlap", Salary: 80447},
	{EmployeeID: 86, DepartmentID: 4, BossID: 53, Name: "Ava Britt", Salary: 75522},
	{EmployeeID: 87, DepartmentID: 3, BossID: 13, Name: "Kelsie Carver", Salary: 209091},
	{EmployeeID: 88, DepartmentID: 2, BossID: 10, Name: "Hunter Cameron", Salary: 251679},
	{EmployeeID: 89, DepartmentID: 5, BossID: 61, Name: "Myles Rivers", Salary: 85918},
	{EmployeeID: 90, DepartmentID: 7, BossID: 81, Name: "Palmer Oneil", Salary: 240424},
	{EmployeeID: 91, DepartmentID: 8, BossID: 5, Name: "Henry Blanchard", Salary: 165394},
	{EmployeeID: 92, DepartmentID: 7, BossID: 9, Name: "Amela Martinez", Salary: 262512},
	{EmployeeID: 93, DepartmentID: 10, BossID: 28, Name: "Rylee Shaw", Salary: 62180},
	{EmployeeID: 94, DepartmentID: 10, BossID: 20, Name: "Violet Sykes", Salary: 104213},
	{EmployeeID: 95, DepartmentID: 5, BossID: 56, Name: "Amela Galloway", Salary: 146457},
	{EmployeeID: 96, DepartmentID: 5, BossID: 88, Name: "James Nolan", Salary: 129572},
	{EmployeeID: 97, DepartmentID: 2, BossID: 75, Name: "Ivory Mckenzie", Salary: 260720},
	{EmployeeID: 98, DepartmentID: 2, BossID: 87, Name: "Eliana Nichols", Salary: 125917},
	{EmployeeID: 99, DepartmentID: 7, BossID: 32, Name: "Patricia Vargas", Salary: 253456},
	{EmployeeID: 100, DepartmentID: 8, BossID: 71, Name: "Jocelyn Guy", Salary: 191937},
}

func LoadData(dbm *modl.DbMap) {
	txn, err := dbm.Begin()
	if err != nil {
		panic(err)
	}
	for _, department := range Departments {
		txn.Insert(&department)
	}

	for _, employee := range Employees {
		txn.Insert(&employee)
	}
	txn.Commit()
}
