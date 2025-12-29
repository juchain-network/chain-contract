#!/usr/bin/env python3

import subprocess
import json
import os
import sys
import argparse
from datetime import datetime
from jinja2 import Template

# 测试脚本列表
test_scripts = [
    {
        "name": "Deployment",
        "script": "script/Deployment.s.sol:DeploymentScript",
        "description": "测试合约部署流程"
    },
    {
        "name": "Validator Management",
        "script": "script/ValidatorManagement.s.sol:ValidatorManagementScript",
        "description": "测试验证节点管理功能"
    },
    {
        "name": "Validator Lifecycle",
        "script": "script/ValidatorLifecycleTest.s.sol:ValidatorLifecycleTest",
        "description": "测试验证器完整生命周期"
    },
    {
        "name": "Delegator Lifecycle",
        "script": "script/DelegatorLifecycleTest.s.sol:DelegatorLifecycleTest",
        "description": "测试委托人完整生命周期"
    },
    {
        "name": "Staking Mechanism",
        "script": "script/StakingMechanism.s.sol:StakingMechanismScript",
        "description": "测试质押机制"
    },
    {
        "name": "Proposal System",
        "script": "script/ProposalSystem.s.sol:ProposalSystemScript",
        "description": "测试提案系统"
    },
    {
        "name": "Punishment Mechanism",
        "script": "script/PunishmentMechanism.s.sol:PunishmentMechanismScript",
        "description": "测试惩罚机制"
    },
    {
        "name": "Integration Test",
        "script": "script/PoSAIntegrationTest.s.sol:PoSAIntegrationTest",
        "description": "整体集成测试"
    }
]

# HTML报告模板
html_template = '''
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>PoSA测试报告</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
            background-color: white;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        h1 {
            color: #333;
            text-align: center;
            border-bottom: 2px solid #3498db;
            padding-bottom: 10px;
        }
        .summary {
            display: flex;
            justify-content: space-around;
            margin: 20px 0;
        }
        .summary-item {
            text-align: center;
            padding: 15px;
            background-color: #f8f9fa;
            border-radius: 5px;
            flex: 1;
            margin: 0 10px;
        }
        .summary-item h3 {
            margin: 0;
            color: #555;
        }
        .summary-item .number {
            font-size: 2em;
            font-weight: bold;
            margin: 10px 0;
        }
        .passed { color: #27ae60; }
        .failed { color: #e74c3c; }
        .duration { color: #3498db; }
        .test-results {
            margin-top: 30px;
        }
        .test-item {
            margin: 20px 0;
            padding: 15px;
            border: 1px solid #ddd;
            border-radius: 5px;
            transition: box-shadow 0.3s;
        }
        .test-item:hover {
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        .test-item h2 {
            margin: 0 0 10px 0;
            color: #2c3e50;
            font-size: 1.2em;
        }
        .test-info {
            font-size: 0.9em;
            color: #7f8c8d;
            margin-bottom: 10px;
        }
        .test-status {
            display: inline-block;
            padding: 5px 10px;
            border-radius: 3px;
            font-weight: bold;
            margin-right: 10px;
        }
        .status-passed {
            background-color: #d4edda;
            color: #155724;
        }
        .status-failed {
            background-color: #f8d7da;
            color: #721c24;
        }
        .test-output {
            background-color: #f8f9fa;
            padding: 10px;
            border-radius: 3px;
            font-family: monospace;
            font-size: 0.8em;
            max-height: 300px;
            overflow-y: auto;
            margin-top: 10px;
        }
        pre {
            margin: 0;
            white-space: pre-wrap;
            word-break: break-all;
        }
        .footer {
            text-align: center;
            margin-top: 30px;
            padding: 10px;
            border-top: 1px solid #ddd;
            color: #7f8c8d;
            font-size: 0.9em;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>PoSA本地测试系统 - 测试报告</h1>
        <div class="summary">
            <div class="summary-item">
                <h3>测试总数</h3>
                <div class="number">{{ total_tests }}</div>
            </div>
            <div class="summary-item">
                <h3>通过测试</h3>
                <div class="number passed">{{ passed_tests }}</div>
            </div>
            <div class="summary-item">
                <h3>失败测试</h3>
                <div class="number failed">{{ failed_tests }}</div>
            </div>
            <div class="summary-item">
                <h3>测试时长</h3>
                <div class="number duration">{{ duration }}s</div>
            </div>
        </div>
        <div class="test-results">
            {% for test in test_results %}
            <div class="test-item">
                <h2>{{ test.name }}</h2>
                <div class="test-info">
                    {{ test.description }} | 
                    脚本: {{ test.script }} | 
                    开始时间: {{ test.start_time }} | 
                    结束时间: {{ test.end_time }} | 
                    时长: {{ test.duration }}s
                </div>
                <div class="test-status status-{{ test.status }}">
                    {{ test.status }}
                </div>
                <div class="test-output">
                    <pre>{{ test.output }}</pre>
                </div>
            </div>
            {% endfor %}
        </div>
        <div class="footer">
            测试报告生成时间: {{ report_time }} | 
            测试环境: {{ environment }}
        </div>
    </div>
</body>
</html>
'''

def run_test(script, broadcast=False):
    """运行单个测试脚本"""
    cmd = ["forge", "script", script, "--fork-url", "http://localhost:8545"]
    if broadcast:
        cmd.append("--broadcast")
    cmd.append("--silent")
    
    start_time = datetime.now()
    result = subprocess.run(cmd, capture_output=True, text=True)
    end_time = datetime.now()
    
    duration = (end_time - start_time).total_seconds()
    
    return {
        "status": "passed" if result.returncode == 0 else "failed",
        "output": result.stdout + result.stderr,
        "returncode": result.returncode,
        "start_time": start_time.strftime("%Y-%m-%d %H:%M:%S"),
        "end_time": end_time.strftime("%Y-%m-%d %H:%M:%S"),
        "duration": round(duration, 2)
    }

def main():
    parser = argparse.ArgumentParser(description="PoSA测试结果收集和报告生成脚本")
    parser.add_argument("--output", "-o", default="./test-results", help="测试结果输出目录")
    parser.add_argument("--broadcast", action="store_true", help="是否广播交易")
    parser.add_argument("--scripts", nargs="*", help="指定要运行的测试脚本名称")
    args = parser.parse_args()
    
    # 创建输出目录
    os.makedirs(args.output, exist_ok=True)
    
    # 过滤要运行的测试脚本
    scripts_to_run = test_scripts
    if args.scripts:
        scripts_to_run = [script for script in test_scripts if script["name"] in args.scripts]
    
    print(f"开始运行 {len(scripts_to_run)} 个测试脚本...")
    print("=" * 60)
    
    # 运行所有测试
    test_results = []
    total_start_time = datetime.now()
    
    for test in scripts_to_run:
        print(f"\n运行测试: {test['name']}")
        print(f"脚本: {test['script']}")
        print("-" * 60)
        
        # 运行测试
        result = run_test(test["script"], args.broadcast)
        
        # 合并测试结果
        test_result = {
            **test,
            **result
        }
        
        test_results.append(test_result)
        
        # 打印测试结果
        status = "✓ 通过" if result["status"] == "passed" else "✗ 失败"
        print(f"结果: {status} | 时长: {result['duration']}s")
        if result["status"] == "failed":
            print(f"错误信息: {result['output'][:500]}...")
    
    total_end_time = datetime.now()
    total_duration = round((total_end_time - total_start_time).total_seconds(), 2)
    
    # 统计结果
    passed_tests = sum(1 for test in test_results if test["status"] == "passed")
    failed_tests = len(test_results) - passed_tests
    
    print("\n" + "=" * 60)
    print("测试总结:")
    print(f"测试总数: {len(test_results)}")
    print(f"通过测试: {passed_tests}")
    print(f"失败测试: {failed_tests}")
    print(f"总时长: {total_duration}s")
    
    # 生成测试报告
    report_data = {
        "test_results": test_results,
        "total_tests": len(test_results),
        "passed_tests": passed_tests,
        "failed_tests": failed_tests,
        "duration": total_duration,
        "report_time": total_end_time.strftime("%Y-%m-%d %H:%M:%S"),
        "environment": "Anvil本地测试环境"
    }
    
    # 生成JSON报告
    json_report_path = os.path.join(args.output, "test-report.json")
    with open(json_report_path, "w", encoding="utf-8") as f:
        json.dump(report_data, f, ensure_ascii=False, indent=2)
    
    print(f"\nJSON报告已生成: {json_report_path}")
    
    # 生成HTML报告
    template = Template(html_template)
    html_report = template.render(**report_data)
    html_report_path = os.path.join(args.output, "test-report.html")
    with open(html_report_path, "w", encoding="utf-8") as f:
        f.write(html_report)
    
    print(f"HTML报告已生成: {html_report_path}")
    
    # 退出状态码
    sys.exit(0 if failed_tests == 0 else 1)

if __name__ == "__main__":
    main()
