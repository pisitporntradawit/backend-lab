git clone https://github.com/YOUR_USERNAME/api-user-management.git
cd api-user-management

# 3. ตั้งค่า Go
go mod init github.com/YOUR_USERNAME/api-user-management

# 4. สร้าง README.md
cat > README.md << 'EOF'
# API User Management - Lab Project

Learning project for Go + Gin + PostgreSQL

## Tech Stack
- Go 1.21
- Gin Framework
- PostgreSQL
- pgx Driver

## Setup
```bash
docker-compose up -d
go run main.go
```

## Author
[Your Name]
EOF

# 5. Commit & Push
git add .
git commit -m "Initial lab project setup"
git push origin main