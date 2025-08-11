# 🔄 合约生成脚本合并改进

## 📋 改进概要

将原来的两个独立脚本合并为一个更灵活的统一脚本：
- ❌ `generate-contracts.js` (只生成生产合约)
- ❌ `generate-mock-contracts.js` (只生成测试合约)  
- ✅ `generate-contracts.js` (统一脚本，支持两种模式)

## 🚀 新功能特性

### 1. 命令行参数支持
```bash
# 生成生产合约 (默认)
node generate-contracts.js

# 生成测试合约
node generate-contracts.js --mock

# 显示帮助信息
node generate-contracts.js --help
```

### 2. 环境变量支持
```bash
# 通过环境变量生成测试合约
MOCK=true node generate-contracts.js
```

### 3. 智能模式识别
- 自动检测 `--mock` 参数
- 支持 `MOCK=true` 环境变量
- 清晰的模式提示信息

### 4. 增强的日志输出
```
Generating contracts in MOCK mode...
Generated: contracts/Params.sol (mock version)
Generate mock contracts success
```

## 📝 使用说明

### 开发环境 (测试合约)
```bash
node generate-contracts.js --mock
```

### 生产环境 (系统合约)
```bash
node generate-contracts.js
```

### 获取帮助
```bash
node generate-contracts.js --help
```

## 🔧 技术改进

### 原始实现 (两个文件)
```javascript
// generate-contracts.js
const config = { mock: false }

// generate-mock-contracts.js  
const config = { mock: true }
```

### 合并后实现 (一个文件)
```javascript
// 动态配置
const isMock = process.argv.includes('--mock') || process.env.MOCK === 'true';
const config = { mock: isMock }

// 智能日志
console.log(`Generating contracts in ${isMock ? 'MOCK' : 'PRODUCTION'} mode...`);
```

## 📚 更新的文档

### README.md
```bash
# 原来
node generate-mock-contracts.js

# 现在  
node generate-contracts.js --mock
```

### package.json
```json
{
  "main": "generate-contracts.js"
}
```

## ✅ 验证测试

所有功能都已验证正常工作：
- ✅ 生产模式: `node generate-contracts.js`
- ✅ 测试模式: `node generate-contracts.js --mock`  
- ✅ 环境变量: `MOCK=true node generate-contracts.js`
- ✅ 帮助信息: `node generate-contracts.js --help`

## 🎯 优势总结

1. **减少重复**: 一个文件替代两个文件
2. **更灵活**: 命令行参数和环境变量支持
3. **更清晰**: 明确的模式提示和日志
4. **更易用**: 内置帮助文档
5. **更易维护**: 单一代码源，减少维护负担

**合并完成！脚本功能更强大，使用更简便！** 🎉
