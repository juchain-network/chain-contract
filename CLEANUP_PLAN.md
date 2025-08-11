# sys-contract 目录整理方案

## 🎯 目标
将混合的工具链整理为干净的项目结构，以 Foundry 为主要开发工具。

## 📊 当前状况分析

### 工具链分类
1. **Go 工具链**: cmd/, go.mod, go.sum, main.go, Makefile, bin/
2. **Foundry 工具链** (主要): forge-scripts/, forge-tests/, foundry.toml, out/, contracts/
3. **Node.js 工具链** (历史遗留): hardhat.config.js, package.json, truffle-config.js, etc.
4. **项目文档**: docs/, *.md, generate-contracts.js

## 📋 整理方案

### 方案A: Foundry 主导 + 最小化 Node.js
**保留**:
- ✅ Foundry 完整工具链
- ✅ 必要的 Node.js 工具 (generate-contracts.js, scripts/)
- ✅ Go 工具链 (如果需要)
- ✅ 项目文档

**移除**:
- ❌ Hardhat 配置
- ❌ Truffle 配置  
- ❌ 历史测试文件
- ❌ 编译缓存

### 方案B: 纯 Foundry
**保留**:
- ✅ Foundry 完整工具链
- ✅ 项目文档
- ✅ Go 工具链 (如果需要)

**移除**:
- ❌ 所有 Node.js 工具链

### 方案C: 多工具链共存
**保留**:
- ✅ 所有工具链
- ✅ 清理编译缓存
- ✅ 整理目录结构

## 🎯 推荐方案: 方案A

基于以下考虑:
1. Foundry 是当前主要开发工具
2. generate-contracts.js 和部分脚本仍然有用
3. 保持必要的灵活性
4. 清理历史遗留文件

## 📁 目标目录结构

```
sys-contract/
├── contracts/           # Solidity 源码
├── forge-scripts/       # Foundry 脚本
├── forge-tests/         # Foundry 测试
├── scripts/            # Node.js 脚本 (generate-contracts, update_genesis)
├── docs/               # 项目文档
├── cmd/                # Go 命令行工具 (如果需要)
├── foundry.toml        # Foundry 配置
├── package.json        # 最小化 Node.js 依赖
├── generate-contracts.js
├── go.mod              # Go 依赖 (如果需要)
├── README.md
└── .gitignore
```

## 🗑️ 待清理文件

### 立即删除
- hardhat.config.js
- truffle-config.js  
- migrations/
- test/ (Hardhat 测试)
- artifacts/
- cache/
- node_modules/ (重新安装最小依赖)

### 评估删除
- cmd/, go.mod, main.go (确认是否还需要 Go 工具)
- package-lock.json (重新生成)

## ⚡ 执行步骤

1. 备份当前状态
2. 删除历史文件
3. 更新 package.json (最小化依赖)
4. 重新安装依赖
5. 测试 Foundry 工具链
6. 测试必要脚本
7. 更新文档
