# Todo Application

แอปพลิเคชัน Todo ที่พัฒนาด้วย Golang (Backend) และ Vue3 (Frontend)

## โครงสร้างโปรเจค

```
todo_app/
├── backend/         # โค้ดส่วน Backend (Golang)
├── frontend/        # โค้ดส่วน Frontend (Vue3)
```

## การติดตั้ง

### Backend (Golang)

1. ติดตั้ง Golang จาก [golang.org](https://golang.org)

2. เข้าไปที่โฟลเดอร์ backend:
   ```bash
   cd backend
   ```
3. ติดตั้ง dependencies:
   ```bash
   go mod download
   ```
4. รันเซิร์ฟเวอร์:
   ```bash
   go run main.go
   ```

### Frontend (Vue3)

1. ติดตั้ง Node.js จาก [nodejs.org](https://nodejs.org)

2. เข้าไปที่โฟลเดอร์ frontend:
   ```bash
   cd frontend
   ```
3. ติดตั้ง dependencies:
   ```bash
   npm install
   ```
4. รันแอปพลิเคชัน:
   ```bash
   npm run dev
   ```

## ฟีเจอร์

- สร้าง Todo 
- แสดงรายการ Todo ทั้งหมด
- แก้ไขรายการ Todo 
- ลบ Todo 

## เทคโนโลยีที่ใช้

- Backend: Golang
- Frontend: Vue3
- Database: PostgreSQL
