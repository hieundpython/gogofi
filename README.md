* Design module User

- Create api for user
- Using Database create table for user


- Create docker DB: 
docker run --name postgres16 -e POSTGRES_USER=hieund -e POSTGRES_PASSWORD=It123456@ -p 5432:5432 -d postgres


- Công việc


* triển khai thành công restful api, kết nối, call database với hiệu suất tốt nhất
* có cơ chế triển khai migration database hiệu quả
* triển khai hệ thống routing hợp lý
* triển khai các middleware cho hệ thống



* tách nó ra thành các tầng 


Controller (Routing, tầng này có nhiệm vụ xử lý với request, response) -> Service (Kết nối đến database và giao tiếp với db)