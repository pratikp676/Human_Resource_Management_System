# Human Resource Management System(HRMS)

* **Technologies used:**
  * Front-end: Vue,js, Javascript,HTML,CSS
  * Back-end: Go(Gin web framework)
  * Database: MongoDB


* **Requirements:**
  * Go (go1.20 windows/amd64)
  * Vue.js version 2.X
  * MongoDB (v5.0.16)


**Instruction on how to run the project:**
1. Open the project folder in a Visual Studio code.
2. Open a new terminal and change directory to '/client'.
3. run 'npm install' to install the required dependencies(If you get any error while installing the dependencies, try to run 'npm install --force').
4. Run client using 'npm run serve'.
5. Open another terminal and change directory to '/server'.
6. run 'go get' to install go dependencies.
7. Run server using 'go run main.go'


**Instruction on how to import dummy data:**
1. Create a database called 'EmployeeDetails'.
2. Create collections called 'Attendance', 'CompanyData', 'EmployeeData' and 'LeavesData'.
3. Import dummy data provided in the 'Data' folder.
