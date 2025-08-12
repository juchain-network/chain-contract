# Congress CLI v1.1.0 - 快速测试指南

## 验证改进后的功能

### 1. 基本命令测试

```bash
# 查看版本
./build/congress-cli version

# 查看帮助
./build/congress-cli --help

# 查看示例
./build/congress-cli examples
```

### 2. 验证输入验证

```bash
# 测试无效 Chain ID
./build/congress-cli create_proposal -c 0 -l https://testnet-rpc.juchain.org -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 -o add
# 应该输出：❌ Validation Error: chain ID is required for command 'create_proposal'

# 测试无效 RPC URL
./build/congress-cli create_proposal -c 202599 -l invalid-url -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 -o add
# 应该输出：❌ Validation Error: invalid RPC URL format: invalid-url

# 测试无效地址
./build/congress-cli create_proposal -c 202599 -l https://testnet-rpc.juchain.org -p invalid-address -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 -o add
# 应该输出：❌ Validation Error: invalid address format: invalid-address
```

### 3. 测试命令语法

```bash
# 创建验证者添加提案 - 测试网
./build/congress-cli create_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \
  -o add

# 创建配置更新提案
./build/congress-cli create_config_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i 0 -v 86400

# 投票赞成（使用 -a 标志）
./build/congress-cli vote_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf \
  -a

# 投票反对（省略 -a 标志）
./build/congress-cli vote_proposal -c 202599 -l https://testnet-rpc.juchain.org \
  -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \
  -i b2be7f3cc702c7a24962df6aed188edbcfdebe20fd55f5670efaedace0e4bcdf

# 提取收益
./build/congress-cli withdraw_profits -c 202599 -l https://testnet-rpc.juchain.org \
  -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
```

## 与 congress.md 的主要差异对比

### 1. 投票语法改进 ✅
- **旧版 (congress.md)**: `-a true` 或 `-a false`
- **新版 (v1.1.0)**: `-a` (赞成) 或省略 `-a` (反对)

### 2. 参数验证增强 ✅
- 自动验证地址格式
- 自动验证 RPC URL 格式
- 自动验证 Chain ID
- 自动验证提案ID格式

### 3. 输出改进 ✅
- 彩色状态指示 (✅ ❌ ℹ️ ⚠️)
- 更清晰的错误消息
- 结构化的成功确认

### 4. 命令一致性 ✅
- 全局参数支持短标志 (`-c`, `-l`)
- 参数名称与 congress.md 完全一致
- 支持测试网和主网配置

## 主要改进总结

1. **用户体验**：更直观的投票语法，更清晰的错误提示
2. **安全性**：全面的输入验证，防止无效参数
3. **便利性**：内置示例命令，支持短标志
4. **兼容性**：完全兼容 congress.md 中的流程
5. **可维护性**：统一的配置管理，更好的代码结构

## 测试验证结果

✅ 所有基本命令工作正常  
✅ 输入验证按预期工作  
✅ 投票语法简化且直观  
✅ 创建提案命令正确  
✅ 示例命令准确反映实际语法  
✅ 与 congress.md 流程兼容
