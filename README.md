# MES1 - 制造执行系统 (Manufacturing Execution System)

这是一个基于 Go 和 Vue 3 开发的制造执行系统 (MES)。

## 项目背景
该项目旨在实现一个轻量级、高效的制造执行系统，涵盖从原材料入库到成品出库的全过程管理。

## GitHub 仓库
[https://github.com/silly-zero/mes1.git](https://github.com/silly-zero/mes1.git)

## 技术栈
- **后端**: [Go](https://go.dev/) (Gin/Gorm)
- **前端**: [Vue 3](https://vuejs.org/) (Vite + Pinia + Vue Router)
- **数据库**: MySQL / PostgreSQL (待定)

## 项目结构
```text
mes1/
├── backend/    # Go 后端代码
└── frontend/   # Vue 3 前端代码
```

## 核心功能模块
1. **登录系统**: 用户身份验证与权限管理。
2. **库存管理**: 实时监控库存状态、库存预警。
3. **收料管理**: 供应商到货确认与记录。
4. **来料检验 (IQC)**: 质量控制，对到货原材料进行检验。
5. **入库管理**: 检验合格后的物料入库操作。
6. **出库管理**: 生产领料或销售出库操作。

## 快速开始
### 后端启动
```bash
cd backend
# 待完善
```

### 前端启动
```bash
cd frontend
# 待完善
```

## 开发规范
- 提交代码请遵循 Git Commit 规范。
- 后端接口遵循 RESTful API 设计。
