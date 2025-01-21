
# MicroBlog-Services

## 1. Project Overview

**MicroBlog-Services** is a microservices-based project consisting of three separate services:

1. **Auth Service**: Built using Django Rest Framework (DRF), it manages user registration and authentication via JWT.
2. **Blog Service**: Built using Flask, it allows authenticated users to perform CRUD operations on blogs.
3. **Go Comments Service**: Built using Go, it allows authenticated users to manage comments for blogs.

These services are independent but communicate via JWT tokens and REST APIs.

---

## 2. Features

### **Auth Service**
- User registration
- User authentication via JWT (access and refresh tokens)

### **Blog Service**
- Create, read, update, and delete blogs
- Authenticated operations using JWT tokens from the Auth Service

### **Go Comments Service**
- Create, read, and manage comments for blogs
- Validate blog IDs with the Blog Service
- Authenticated operations using JWT tokens from the Auth Service

---

## 3. How to Clone the Repository

Run the following command to clone the repository:

```bash
git clone https://github.com/NoManNayeem/MicroBlog-Services.git
cd MicroBlog-Services
```

---

## 4. How to Run All Services Locally

### **Environment Setup**

#### **Windows**
1. Open a terminal and navigate to the project root directory.
2. Create virtual environments for the Python services:
   ```bash
   python -m venv auth_env
   python -m venv blog_env
   ```
3. Activate the environments:
   - **Auth Service**:
     ```bash
     auth_env\Scripts\activate
     ```
   - **Blog Service**:
     ```bash
     blog_env\Scripts\activate
     ```

#### **Linux/MacOS**
1. Open a terminal and navigate to the project root directory.
2. Create virtual environments for the Python services:
   ```bash
   python3 -m venv auth_env
   python3 -m venv blog_env
   ```
3. Activate the environments:
   - **Auth Service**:
     ```bash
     source auth_env/bin/activate
     ```
   - **Blog Service**:
     ```bash
     source blog_env/bin/activate
     ```

---

### **Install Requirements**

- **Auth Service**:
  ```bash
  pip install -r auth_service/requirements.txt
  ```

- **Blog Service**:
  ```bash
  pip install -r blog_service/requirements.txt
  ```

---

### **Environment Variables**

Ensure each service has its own `.env` file. Use the provided `.env.example` files as templates.

- **Auth Service**:
  ```
  JWT_SECRET_KEY=your-secret-key
  ```
- **Blog Service**:
  ```
  FLASK_ENV=development
  ```

- **Go Comments Service**:
  ```
  FLASK_APP_URL=http://127.0.0.1:5000
  TOKEN_VERIFY_URL=http://127.0.0.1:8000/api/token/verify/
  PORT=8080
  ```

---

### **Run the Services**

- **Auth Service**:
  ```bash
  cd auth_service
  python manage.py runserver
  ```

- **Blog Service**:
  ```bash
  cd blog_service
  python app.py
  ```

- **Go Comments Service**:
  ```bash
  cd go_comments
  go run main.go
  ```

---

## 5. CURLs to Test

### **1. Register a User**
```bash
curl -X POST http://127.0.0.1:8000/api/users/ -H "Content-Type: application/json" -d '{
    "username": "testuser",
    "email": "testuser@example.com",
    "password": "testpassword"
}'
```

### **2. Get Tokens**
```bash
curl -X POST http://127.0.0.1:8000/api/token/ -H "Content-Type: application/json" -d '{
    "username": "testuser",
    "password": "testpassword"
}'
```

### **3. Create a Blog**
```bash
curl -X POST http://127.0.0.1:5000/blogs -H "Authorization: Bearer <access_token>" -H "Content-Type: application/json" -d '{
    "title": "My First Blog",
    "content": "This is the content of my first blog post."
}'
```

### **4. Read All Blogs**
```bash
curl -X GET http://127.0.0.1:5000/blogs -H "Authorization: Bearer <access_token>"
```

### **5. Create a Comment**
```bash
curl -X POST http://127.0.0.1:8080/comments -H "Authorization: Bearer <access_token>" -H "Content-Type: application/json" -d '{
    "post_id": 2,
    "title": "Sample Comment",
    "content": "This is a sample comment for the specified post."
}'
```

### **6. Read All Comments**
```bash
curl -X GET http://127.0.0.1:8080/comments -H "Authorization: Bearer <access_token>"
```

---

## 6. Contributions

Contributions are welcome! Feel free to fork the repository and submit a pull request.

---

## 7. License

This project is licensed under the MIT License. See the LICENSE file for details.
