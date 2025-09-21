# RSS简易搜索器

这是一个基于Go语言实现的RSS源搜索器项目，使用到了Go语言中并发编程、接口设计、包管理等核心概念。

## 项目结构

```
goinaction/
├── main.go              # 主程序入口
├── go.mod              # Go模块文件
├── data/
│   └── data.json       # RSS源配置文件
├── search/             # 搜索核心包
│   ├── search.go       # 主搜索逻辑和注册器
│   ├── match.go        # 匹配器接口和结果处理
│   ├── feed.go         # 数据源结构定义
│   └── defalut.go      # 默认匹配器实现
└── matchers/           # 匹配器实现包
    └── rss.go          # RSS匹配器实现
```

## 功能特性

- **并发搜索**: 使用goroutine并发处理多个RSS源
- **接口设计**: 通过Matcher接口实现可扩展的搜索策略
- **RSS解析**: 支持解析RSS XML格式数据
- **正则匹配**: 在标题和描述中搜索指定关键词
- **通道通信**: 使用channel进行goroutine间通信

## 核心组件

### 1. 搜索引擎 (`search/search.go`)
- 管理匹配器注册
- 协调并发搜索任务
- 使用WaitGroup同步goroutine

### 2. 匹配器接口 (`search/match.go`)
- 定义Matcher接口规范
- 处理搜索结果展示
- 管理Result结构体

### 3. RSS匹配器 (`matchers/rss.go`)
- 实现RSS源数据获取
- XML解析和正则匹配
- 支持多个RSS源并发搜索

### 4. 数据源配置 (`data/data.json`)
- 预配置的RSS源列表

## 使用方法

1. **安装依赖**
   ```bash
   go mod tidy
   ```

2. **运行程序**
   ```bash
   go run main.go
   ```

3. **修改搜索词**
   编辑 `main.go` 文件中的搜索词：
   ```go
   search.Run("President")  // 搜索包含"President"的内容
   ```

## 扩展(待完成)

- 添加更多类型的匹配器（如JSON、CSV等）
- 实现搜索结果缓存
- 添加配置文件支持
- 实现搜索结果的排序和过滤
- 添加日志记录功能
