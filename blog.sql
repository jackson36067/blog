/*
 Navicat MySQL Dump SQL

 Source Server         : wsl
 Source Server Type    : MySQL
 Source Server Version : 80044 (8.0.44-0ubuntu0.24.04.2)
 Source Host           : 172.30.41.46:3306
 Source Schema         : gin_blog

 Target Server Type    : MySQL
 Target Server Version : 80044 (8.0.44-0ubuntu0.24.04.2)
 File Encoding         : 65001

 Date: 05/02/2026 21:40:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `abstract` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `category_id` bigint UNSIGNED NULL DEFAULT NULL,
  `tag_list` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `coverage` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `browse_count` bigint NULL DEFAULT NULL,
  `like_count` bigint NULL DEFAULT NULL,
  `comment_count` bigint NULL DEFAULT NULL,
  `collect_count` bigint NULL DEFAULT NULL,
  `public_comment` tinyint(1) NULL DEFAULT NULL,
  `status` tinyint NULL DEFAULT NULL,
  `visibility` tinyint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (1, '2025-09-11 21:03:48.000', '2025-12-03 16:06:13.444', 'Go 基础语法入门', '学习 Go 的基本语法和结构。', '\r\n\r\n## 1. Go 简介\r\n\r\nGo（Golang）是一门由 Google 开发的语言，具有以下特点：\n1. 语法简洁\r\n2. 内置并发 \r\n3. 自动垃圾回收\r\n4. 原生跨平台编译\r\n5. 强大的工具\n\r\n\r\n\r\n---\n\n\r\n## 2. 第一个 Go 程序\r\n\r\n下面是你的第一个 Go 程序：\r\n\r\n```go\r\npackage main\r\n\r\nimport \"fmt\"\r\n\r\nfunc main() {\r\n    fmt.Println(\"Hello Go!\")\r\n}\r\n```\r\n\r\n---\r\n\r\n## 3. 变量与常量\r\n\r\n### 变量声明方式：\r\n\r\n```go\r\nvar a int = 10\r\nvar b = 20\r\nc := 30\r\n```\r\n\r\n### 常量：\r\n\r\n```go\r\nconst PI = 3.14\r\nconst Name string = \"GoLang\"\r\n```\n\r\n\r\n---\n## 4. 基础数据类型\r\n\r\nGo 支持如下常用类型：\r\n\r\n- **整型**：`int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, ...\r\n- **浮点型**：`float32`, `float64`\r\n- **字符串**：`\"hello\"`\r\n- **布尔型**：`true` / `false`\r\n- **字节与字符**：`byte`（相当于 `uint8`），`rune`（相当于 `int32`，用于表示 Unicode 码点）\r\n- **复数**：`complex64`, `complex128`\r\n\r\n### 示例\r\n\r\n```go\r\nvar a int = 10\r\nvar b float64 = 3.14\r\nvar s string = \"hello\"\r\nvar ok bool = true\r\nvar ch byte = \'A\'\r\nvar r rune = \'你\'\r\nvar c complex128 = complex(1, 2)\r\n```\r\n\r\n### 零值（默认值）\r\n\r\n- 整数类型默认 `0`\r\n- 浮点类型默认 `0.0`\r\n- 字符串默认 `\"\"`\r\n- 布尔默认 `false`\r\n- 指针、切片、map、channel 等引用类型默认 `nil`\r\n\r\n### 类型转换\r\n\r\nGo 不会自动进行不同类型之间的隐式转换，需要显式转换：\r\n\r\n```go\r\nvar i int = 42\r\nvar f float64 = float64(i)\r\nvar u uint = uint(i)\r\n```\r\n\r\n### 小结\r\n\r\n- 常用基本类型如上；  \r\n- 注意 `byte` / `rune` 的用途（字节 vs Unicode 码点）；  \r\n- 不同类型之间需要显式转换；  \r\n- 如果你需要更详细的表格或示例（比如各类型位宽、最大最小值、零值表），我可以继续补充。\r\n\r\n\r\n\r\n\r\n## 5. 条件判断 if\r\n\r\n```go\r\nif num > 10 {\r\n    fmt.Println(\"big\")\r\n} else {\r\n    fmt.Println(\"small\")\r\n}\r\n```\r\n\r\n---\r\n\r\n## 6. 循环 for\r\n\r\nGo 只有一种循环关键字：`for`\r\n\r\n```go\r\nfor i := 0; i < 5; i++ {\r\n    fmt.Println(i)\r\n}\r\n\r\ni := 0\r\nfor i < 5 {\r\n    i++\r\n}\r\n\r\nfor {\r\n    fmt.Println(\"loop\")\r\n}\r\n```\r\n\r\n---\r\n\r\n## 7. 数组、切片、map\r\n\r\n### 数组（固定长度）\r\n\r\n```go\r\nvar arr [3]int = [3]int{1, 2, 3}\r\n```\r\n\r\n### 切片（动态）\r\n\r\n```go\r\nnums := []int{1, 2, 3}\r\nnums = append(nums, 4)\r\n```\r\n\r\n### map（键值对）\r\n\r\n```go\r\nm := map[string]int{\r\n    \"apple\": 2,\r\n    \"banana\": 5,\r\n}\r\nm[\"orange\"] = 3\r\n```\r\n\r\n---\r\n\r\n## 8. 函数\r\n\r\n```go\r\nfunc add(a int, b int) int {\r\n    return a + b\r\n}\r\n```\r\n\r\n### 多返回值：\r\n\r\n```go\r\nfunc divide(a, b float64) (float64, error) {\r\n    if b == 0 {\r\n        return 0, fmt.Errorf(\"division by zero\")\r\n    }\r\n    return a / b, nil\r\n}\r\n```\r\n\r\n---\r\n\r\n## 9. 结构体 Struct\r\n\r\n```go\r\ntype User struct {\r\n    Name string\r\n    Age  int\r\n}\r\n```\r\n\r\n---\r\n\r\n## 10. 方法（面向对象）\r\n\r\n```go\r\nfunc (u User) SayHello() {\r\n    fmt.Println(\"Hello\", u.Name)\r\n}\r\n```\r\n\r\n---\r\n\r\n## 11. 指针\r\n\r\n```go\r\na := 10\r\np := &a\r\nfmt.Println(*p)\r\n```\r\n\r\n---\r\n\r\n## 12. 并发 goroutine 与 channel\r\n\r\n### goroutine\r\n\r\n```go\r\ngo func() {\r\n    fmt.Println(\"goroutine running\")\r\n}()\r\n```\r\n\r\n### channel\r\n\r\n```go\r\nch := make(chan int)\r\n\r\ngo func() {\r\n    ch <- 10\r\n}()\r\n\r\nvalue := <-ch\r\nfmt.Println(value)\r\n```\r\n\r\n---\r\n\r\n## 13. 错误处理\r\n\r\n```go\r\nresult, err := divide(10, 2)\r\nif err != nil {\r\n    fmt.Println(\"error:\", err)\r\n}\r\n```\r\n\r\n---\r\n\r\n## 14. Go Modules\r\n\r\n初始化模块：\r\n\r\n```\r\ngo mod init myapp\r\n```\r\n\r\n运行：\r\n\r\n```\r\ngo run main.go\r\n```\r\n\r\n拉取依赖：\r\n\r\n```\r\ngo get xxx\r\n```\r\n\r', 7, '[\"Go\",\"入门\",\"语法\"]', '', 1, 480, 2, 23, 1, 1, 3, 0);
INSERT INTO `article` VALUES (2, '2024-03-14 21:03:48.000', '2025-10-31 21:03:48.012', 'GORM 使用教程', '介绍 GORM 的模型定义与 CRUD 操作。', '', 9, '[\"Go\",\"GORM\",\"数据库\"]', '', 1, 271, 1, 0, 0, 1, 3, 0);
INSERT INTO `article` VALUES (3, '2025-09-11 21:03:48.000', '2026-01-15 20:37:35.354', 'JSON 在 Go 中的应用', '讲解如何在 Go 中解析与生成 JSON。', '', 7, '[\"Go\",\"JSON\",\"教程\"]', '', 1, 241, 1, 0, 1, 1, 3, 0);
INSERT INTO `article` VALUES (4, '2023-03-23 21:03:48.000', '2025-10-31 21:03:48.012', '前后端分离架构', '探讨前后端分离的优缺点。', '', 2, '[\"架构\",\"前端\",\"后端\"]', '', 1, 218, 0, 0, 0, 1, 3, 0);
INSERT INTO `article` VALUES (5, '2023-03-24 21:03:48.000', '2025-10-31 21:03:48.012', 'Vue3 + TypeScript 实战', '使用 Vue3 与 TypeScript 构建前端项目。', '', 5, '[\"Vue3\",\"TypeScript\",\"前端\"]', '', 1, 36, 1, 0, 0, 1, 3, 0);
INSERT INTO `article` VALUES (6, '2025-09-11 21:03:48.000', '2025-11-26 16:35:56.131', 'Spring Boot 权限管理系统', '构建一个基于 Spring Security 的权限系统。', '', 4, '[\"Java\",\"Spring Boot\",\"权限\"]', '', 1, 106, 6, 0, 1, 1, 3, 0);
INSERT INTO `article` VALUES (7, '2025-10-31 21:03:48.012', '2025-12-02 20:27:38.293', 'MySQL 查询优化技巧', '提升数据库性能的常见优化策略。', '', 3, '[\"MySQL\",\"数据库\",\"优化\"]', '', 1, 191, 2, 0, 1, 1, 3, 0);
INSERT INTO `article` VALUES (8, '2024-03-14 21:03:48.000', '2026-01-15 20:54:41.583', 'Redis 缓存应用场景', 'Redis 在高并发系统中的使用案例。', '# Redis 缓存应用场景\r\n\r\nRedis 是一个开源的内存数据结构存储系统，可以用作数据库、缓存和消息中间件，具有以下特点：\r\n\r\n- 高性能，支持每秒数百万级操作\r\n- 支持丰富的数据类型，包括字符串、哈希、列表、集合、有序集合等\r\n- 支持持久化机制，RDB 快照和 AOF 日志\r\n- 支持发布/订阅、事务和 Lua 脚本\r\n- 数据存储在内存中，访问延迟低\r\n\r\n---\r\n\r\n## 1. 热点数据缓存\r\n\r\n热点数据缓存是 Redis 最常见的应用场景之一，通过缓存频繁访问的数据，可以显著减少数据库压力，提高系统响应速度。\r\n\r\n### 1.1 应用示例\r\n\r\n- 商品信息缓存\r\n- 用户资料缓存\r\n- 配置参数缓存\r\n\r\n### 1.2 示例代码\r\n\r\n```shell\r\nSET product:1001 \'{\"name\":\"Laptop\",\"price\":999}\' EX 3600\r\nGET product:1001\r\n```\r\n\r\n## 2. 会话管理\r\n\r\nRedis 可用于存储用户会话信息，如登录状态、Token 或 Session，支持设置过期时间管理会话生命周期。\r\n\r\n### 2.1 应用示例\r\n\r\n- Web 应用用户登录状态管理\r\n- 移动端 Token 缓存\r\n- 单点登录 Session 存储\r\n\r\n### 2.2 示例代码\r\n\r\n```shell\r\nSET session:abcd1234 \"user_id:1001\" EX 1800\r\nGET session:abcd1234\r\n```\r\n## 3. 排行榜与计数器\r\n\r\nRedis 的有序集合和字符串结构非常适合实现高性能排行榜和计数器。\r\n\r\n### 3.1 排行榜\r\n\r\n- 游戏积分排行榜\r\n- 投票排名\r\n- 用户活跃度排名\r\n\r\n### 3.2 排行榜示例代码\r\n\r\n```shell\r\nZINCRBY leaderboard 50 \"user:1001\"\r\nZRANGE leaderboard 0 9 WITHSCORES\r\n```\r\n## 4. 限流控制\r\n\r\nRedis 原子操作可实现接口访问频率限制，防止恶意刷流量或接口滥用。\r\n\r\n### 4.1 应用示例\r\n\r\n- API 请求限流\r\n- 登录尝试次数限制\r\n- 秒杀抢购频率控制\r\n\r\n### 4.2 示例代码\r\n\r\n```shell\r\nINCR user:1001:request_count\r\nEXPIRE user:1001:request_count 60\r\n```\r\n## 5. 消息队列\r\n\r\nRedis 列表可用作简单消息队列，用于异步任务处理和消息通知。\r\n\r\n### 5.1 应用示例\r\n\r\n- 异步邮件发送\r\n- 后台任务处理\r\n- 消息通知推送\r\n\r\n### 5.2 示例代码\r\n\r\n```shell\r\nLPUSH task_queue \"task1\"\r\nRPOP task_queue\r\n```\r\n## 6. 持久化策略\r\n\r\nRedis 支持两种持久化方式：\r\n\r\n- **RDB（快照）**：定期生成数据快照，适合数据恢复场景  \r\n- **AOF（追加日志）**：记录每条写操作，适合高安全性需求\r\n\r\n### 6.1 配置示例\r\n\r\n```shell\r\n# RDB 快照配置\r\nsave 900 1\r\nsave 300 10\r\nsave 60 10000\r\n\r\n# AOF 配置\r\nappendonly yes\r\nappendfsync everysec\r\n```\r\n## 7. 使用建议\r\n\r\n- 高频读写数据适合缓存\r\n- 数据可重建、非关键数据适合缓存\r\n- 根据访问特点选择合适的数据结构\r\n- 设置合理的过期时间，避免缓存雪崩\r\n\r\n\r\n\r\n', 3, '[\"Redis\",\"缓存\",\"高并发\"]', '', 7, 2537, 1, 0, 1, 1, 3, 0);
INSERT INTO `article` VALUES (9, '2025-06-11 21:03:48.000', '2025-10-31 21:03:48.012', '前端性能优化指南', '提高网页加载速度的 10 种方法。', '', 2, '[\"前端\",\"性能\",\"优化\"]', '', 7, 2432, 0, 0, 0, 1, 3, 2);
INSERT INTO `article` VALUES (10, '2025-06-18 21:03:48.000', '2025-10-31 21:03:48.012', 'RESTful API 设计规范', '如何设计清晰一致的接口风格。', '', 1, '[\"API\",\"RESTful\",\"规范\"]', '', 7, 437, 1, 0, 0, 1, 3, 1);
INSERT INTO `article` VALUES (11, '2025-10-31 21:03:48.012', '2025-11-26 11:07:09.537', 'Go 并发编程详解', '学习 goroutine 与 channel 的使用。', '', 7, '[\"Go\",\"并发\",\"高性能\"]', '', 7, 2218, 0, 0, 0, 1, 3, 2);
INSERT INTO `article` VALUES (12, '2025-10-31 21:03:48.012', '2026-01-15 20:54:34.775', '微服务架构实践', '使用 Spring Cloud 实现服务拆分。', '', 4, '[\"Java\",\"微服务\",\"架构\"]', '', 7, 272, 0, 0, 1, 1, 3, 1);
INSERT INTO `article` VALUES (27, '2025-12-01 16:13:24.091', '2025-12-04 08:59:02.490', '环形单链表', '使用环形单链表解决约瑟夫问题', '#  环形单链表\n\n## 定义一个环形单链表\n\n```java\npublic class CircleSingleLinkedList {\n    private Boy first = null; // 第一个节点\n    // 传递一个数字,根据数字大小创建一个环形单链表\n    private void add(int nums) {\n        // 校验nums的合法性\n        if (nums <= 0) {\n            System.out.println(\"Invalid input\");\n            return;\n        }\n        // 辅助指针, 帮助构建环形单链表\n        Boy currBoy = null;\n        for(int i = 0; i < nums; i++) {\n            Boy boy = new Boy(i);\n            // 当设置第一个节点时,让first节点接收, 并且设置first的next节点为自己\n            if(i==1) {\n                first = boy;\n                first.setNext(first);\n                // 使用辅助指针接收第一个节点\n                currBoy = first;\n            } else {\n                // 接收第二个以及后面的节点\n                // 1.将环形单链表最后一个节点的设置为目前添加到节点(作用: 1.让当前节点与链表最后一个节点连接, 2.断开最后一个节点与第一个节点的连接)\n                currBoy.setNext(boy);\n                // 2.让目前添加的节点的next设置为第一个节点\n                boy.setNext(first);\n                // 3.让辅助节点重新接收当前添加的节点\n                currBoy = boy;\n            }\n        }\n    }\n}\n\n```\n\n\n\n## 定义一个boy对象\n\n```java\npublic class Boy {\n    private int id;\n    private Boy next;\n\n    public Boy(int id) {\n        this.id = id;\n    }\n\n    public int getId() {\n        return id;\n    }\n\n    public void setId(int id) {\n        this.id = id;\n    }\n\n    public Boy getNext() {\n        return next;\n    }\n\n    public void setNext(Boy next) {\n        this.next = next;\n    }\n\n    @Override\n    public String toString() {\n        return \"Boy{\" +\n                \"id=\" + id +\n                \'}\';\n    }\n}\n```\n\n\n\n## 解决约瑟夫问题\n\n```java\n	// 约瑟夫问题:\n    // nums个小男孩围成一圈,从startNum个男孩开始报数,从1报到m,报到m的出圈\n    // 问: 男孩的出圈顺序\n    private Boy first = null; // 第一个节点\n    public void outCircleLinkedListOrder(int startNum, int m, int nums) {\n        // 数据校验以及环形单链表是否为空校验\n        if (first == null || startNum < 0 || startNum > nums) {\n            System.out.println(\"Invalid input\");\n            return;\n        }\n        // 定义辅助指针\n        Boy helper = first;\n        // 让辅助指针指向环形单向链表的最后一个节点(便于后续判断链表是否只剩下一个节点)\n        while(helper.getNext() != first) {\n            helper.setNext(helper.getNext())\n        }\n        // 从startNum个男孩开始报数, 让first节点以及辅助节点helper向后移动startNum-1次\n        for(int i = 0; i < startNum; i++) {\n            first = first.getNext();\n            helper = helper.getNext();\n        }\n        // 开始报数\n        // 退出条件, 当环形单向链表中只剩下一个节点后,退出循环, 只剩下一个节点时,辅助指针指向first节点\n        while(first != helper) {\n            // 报数从1报到m,也就是让first以及辅助指针helper移动m-1次\n            for(int i = 0; i < m - 1; i++) {\n                first = first.getNext();\n                helper = helper.getNext();\n            }\n            // 此时first就是报到m的男孩,输出\n            System.out.printf(\"出圈的男孩编号%d\\n\", first.getId());\n            // 从环形单向链表中移除该节点\n            // 让报到数的节点的前一个节点的next节点设置为报到数的节点的next节点\n            helper.setNext(first.getNext());\n            // 继续下一轮报数\n            first = first.getNext();\n        }\n        // 此时环形单向链表还剩下一个节点, 最后报数\n        System.out.printf(\"最后出圈节点编号%d\\n\", first.getId());\n    }\n```\n\n', 10, '[\"数据结构\",\"链表\",\"Java\"]', '', 1, 107, 1, 0, 0, 1, 3, 0);
INSERT INTO `article` VALUES (28, '2025-12-02 21:12:20.668', '2025-12-09 11:03:08.302', 'rabbitMQ基础知识', '学习rabbitmq并在项目中使用', '# rabbitMQ基础知识\n\n## rabbitMQ的作用方式\n\n1. rabbitMQ是典型的异步中间件\n2. rabbitMQ中主要有虚拟主机,交换机,队列\n3. 虚拟主机主要用于隔离数据\n4. 交换机用户向队列路由消息\n5. 队列收到交换机的消息后发送给消费者\n\n## 在java中使用rabbitmq\n\n> 使用springAMQP其中封装好了使用rabbitMQ的api\n>\n> amqp: **advanced message queue protocol**, 是一种消息队列的协议\n\n## 在项目中配置rabbit!MQ\n\n```java\nspring:\n  rabbitmq:\n    host: 192.168.80.101 #主机\n    port: 5672 #端口\n    virtual-host: host1 # 虚拟主机名称\n    username: jackson # 用户名\n    password: 123456 # 密码\n```\n\n\n\n## 使用api发送消息\n\n> 使用RabbitTemplate中的ConvertAndSend发布消息\n\n```java\n  // 首先引入springAMQP的依赖\n  <dependency>\n        <groupId>org.springframework.boot</groupId>\n        <artifactId>spring-boot-starter-amqp</artifactId>\n        <version>替换为最新版本</version>\n  </dependency>\n```\n\n```java\n  // 在测试类中测试发送消息\n  // 1.引入RabbitTemplate\n  @Resource\n  private RabbitTemplate rabbitTemplate\n  \n  // 直接发送消息到queue中\n  @Test    \n  public void sendMessage(){\n      // 参数1为queue名称 参数2为消息\n      rabbitTemplate.convertAndSend(\"\",\"\");\n  } \n```\n\n\n\n## 消费者消费消息\n\n> 使用@RabbitListener注解监听队列获取消息\n\n```java\n    @RabbitListener(queues = \"simple.queue\")\n    private void listenerSimpleQueueMessage(String message) {\n        log.info(\"监听到队列simple.queue中发送的消息: {}\", message);\n        // 执行中间逻辑\n        log.info(\"消息处理完毕\");\n    }\n```\n\n\n\n## rabbitMQ-work模式\n\n> work模式即为多消费者消费模式,需要配置prefetch: 1才能做到多消费者能者多劳的效果\n\n```java\nspring:\n  rabbitmq:\n    listener:\n      simple:\n        # 消费者处理完一条再处理下一条\n        prefetch: 1\n```\n\n## 交换机类型\n\n1. fanout -> 广播\n2. direct -> 通过binding值绑定\n3. topic -> 也是通过binding值绑定,但是binding可以使用多个单词表示,使用.分隔,还可以使用通配符 #**(匹配0个或者多个单词)**, ***(匹配一个单词)**\n\n## 通过api创建队列以及交换机\n\n### 使用配置类创建\n\n> 准备工作: 首先将类声明为配置类\n\n```java\n  	// 声明一个队列\n    @Bean\n    public Queue genQueue(){\n        return QueueBuilder.durable(\"object.queue\").build();\n    }\n\n	// 声明各种类型的交换机\n	@Bean\n	public Excahnge genExahcnge(){\n        // return ExchangeBulider.fanoutExchange(\"\"),bulider();\n        // return ExchangeBulider.directExchange(\"\"),bulider();\n        return ExchangeBulider.topicExchange(\"\"),bulider();\n    }\n\n	// 声明binding\n	@Bean\n	public Binding genBinding(){\n	    // 其中因为fanout类型的交换机为广播模式,所以不需要指定with()方法,即不要指定binding值与queue绑定\n        return BindingBuilder.bind(genQueue()).to(genExchange()).with(\"\"); \n    }\n```\n\n## 消息转换器\n\n> 当我们发送对象格式的消息后,我们会发现消费者消费的消息是一串字符\n>\n> 这是因为rabbitMQ对对象消息进行了默认的序列化转化,导致消费的消息是一个字符串\n>\n> 当需要获取到消息对象具体值后,可以将rabbitMQ消息转换器变为Json消息转换器, 其中我们可以使用springAMQP提供好的json消息转换器类,我们需要做的就是替换\n>\n> 所有消息转换器都实现了MessageConverter接口,只需要修改他的实现就可以替换成我们想要的消息转换器\n\n```java\n  /**\n     * 使用json消息转换器\n     * @return\n     */\n    @Bean\n    public MessageConverter JsonMessageConverter() {\n        // springAMQP提供的Json消息转换器\n        return new Jackson2JsonMessageConverter();\n    }\n```\n\n', 11, '[\"RabbitMQ\",\"Java\"]', '', 1, 4, 0, 0, 0, 1, 3, 0);
INSERT INTO `article` VALUES (29, '2025-12-02 21:18:30.596', '2025-12-28 12:27:53.637', 'java实现快速排序', 'java实现快速排序的方法', '# 快速排序\n\n## 思路: 将数组通过一个中间值分成两部分, 将大于中间值的数都放到他的右边, 将小于他的数都放到他的左边, 然后取左半边以及右半边进行递归,直到最终排序成功\n\n```java\npublic void quickSort(int arr[], int left, int right) {\n    int l = left;\n    int r = right;\n    int pivot = (left + right) / 2;\n    // 如果左索引大于了右索引就结束循环\n    while(l < r) {\n        // 获取左边部分数组比中间值大的\n        while(arr[l] < arr[pivot]) {\n            l++;\n        }\n        // 获取右边数组部分比中间值小的\n        while(arr[r] > arr[pivot]) {\n            r--;\n        }\n        // 假设交换完成后, l以及r都会到达pivot的位置, 这时就要退出循环\n        if(l >= r) {\n             break;\n        }\n        // 交换左右边的值\n        int temp = arr[l];\n        arr[l] = arr[r];\n        arr[r] = temp;\n        \n        // 处理死循环的问题\n        // 出现死循环的原因有: 当左右两边有等于中间值的值是, 且刚好左边和右边交换值都等于中间值, 这时就会一直交换这两个\n        \n        // 如果此时l到达中间值的位置, 让r向前移一位, 防止死循环\n        if(arr[l] == arr[pivot]) {\n            r--;\n        }\n        // 如果此时r到达了中间值的位置, 让l向后移一位, 防止死循环\n        if(arr[r] == arr[pivot]) {\n            l++;\n        }\n    }\n    // 处理堆栈溢出异常问题\n    // 堆栈移除异常问题出现原因有: 当l=r时, 进行递归左部分或者有部分可能是相同部分的递归\n    // 假设一个数组为[5,5,5,5,5], 查找左右两边中间值后发现左右索引都到了中间位置, 那么左右两边递归的还是同样的区间, 这时就会出现递归死循环, 导致堆栈移除异常\n    // 当 l = r 时, 左右两边的递归加起来跟上一次的区间一致, 所以要改变l与r的值防止这种情况\n    if(l == r) {\n        l++;\n        r--;\n    }\n    if(l < right) {\n        quickSort(arr,left,r);\n    }\n    if(r > left) {\n        quickSort(arr,l,right);\n    }\n}\n```\n\n', 12, '[\"数据结构\",\"Java\"]', '', 1, 273, 0, 1, 1, 1, 3, 0);
INSERT INTO `article` VALUES (30, '2025-12-02 21:37:25.534', '2025-12-03 16:12:07.937', '选择排序', '使用java实现选择排序', '# 选择排序\n\n## 思路: 每次获取最小值, 将其换到前面\n\n```java\npublic void selectSort(int[] arr) {\n    for(int i = 0; i < arr.length; i++) {\n        // 定义一个辅助指针保存每次内部循环找到的最小值, 并且假设最小值是外部循环的索引处\n        int minIndex = i;\n        for(int j = i + 1; j < arr.length; j++) {\n            if(arr[j] < arr[i]) {\n                minIndex = j;\n            }\n        }\n        // 如果minIndex变了, 说明在数组中有在该索引后面比该索引出小的值\n        if(minIndex != i) {\n            // 交换位置\n            int temp = arr[j];\n            arr[i] = arr[j];\n            arr[j] = temp;\n        }\n    }\n}\n```\n\n', 12, '[\"数据结构\",\"Java\"]', '', 1, 12, 0, 0, 0, 1, 1, 0);
INSERT INTO `article` VALUES (31, '2025-12-26 15:44:25.000', '2025-12-28 21:25:59.443', '队列', '使用java实现队列', '# 队列\r\n\r\n## 使用数组模拟队列\r\n\r\n```java\r\npublic class ArrayQueue {\r\n    private int maxSize;\r\n    private int front; // 指向队列的第一个数据的前一个数据, 队列的头指针\r\n    private int rear; // 指向队列最后一个数据, 队列的尾指针\r\n    private int[] arr; // 存储数据 -> 模拟队列\r\n    \r\n    public ArrayQueue (int maxSize) {\r\n        this.maxSize = maxSize;\r\n        arr = new int[maxSize];\r\n        front = -1;\r\n        rear = -1;\r\n    }\r\n}\r\n```\r\n\r\n## 给队列添加一些功能 (判断队列是否存满, 判断队列是否为空, 向队列中存入数据, 从队列中去除数据, 遍历打印队列,  显示队列的头信息)\r\n\r\n```java\r\n// 判断队列是否存满\r\npublic boolean isFull() {\r\n    // 当最后一个数据的指针为队列最大长度 - 1时,表示队列存满;\r\n    return rear == maxSize - 1;\r\n}\r\n\r\n// 判断队列是否为空\r\npublic boolean isEmpty() {\r\n    // 当队列最后一个数据的指针与队列第一个数据的指针相等时, 队列为空\r\n    return rear == front;\r\n}\r\n\r\n// 向队列中存入数据, 存入到队列的尾部\r\npublic void add(int value) {\r\n    // 判断队列是否存满\r\n    if(isFull()) {\r\n        System.out.println(\"queue is full\");\r\n        return;\r\n    }\r\n    // 把队列的尾指针向后移动一个位置\r\n    rear++;\r\n    arr[rear] = value;\r\n}\r\n\r\n// 从队列中获取数据\r\npublic int get() {\r\n    // 判断队列是否为空\r\n    if(isEmpty()) {\r\n        throw new RuntimeException(\"queue is empty\");\r\n    }\r\n    // 把队里的头指针向后移动一个位置\r\n    front++;\r\n    return arr[front];\r\n}\r\n\r\n// 遍历队列\r\nfor(int i = 0 ;i < arr.length; i++) {\r\n    // 判断队列是否为空\r\n    if(isEmpty) {\r\n        System.out.println(\"queue is empty\");\r\n        return;\r\n    }\r\n    System.out.println(arr[i]);\r\n}\r\n\r\n// 显示队列的头信息\r\npublic int getQueueHeadInfo() {\r\n    // 判断队列是否为空\r\n    if(isEmpty) {\r\n        throw new RuntimeException(\"Queue is empty\");\r\n    }\r\n    return arr[front + 1];\r\n}\r\n```\r\n\r\n> `这种队列存在的问题, 只能存取一次, 需要改成环形队列优化这个队列`\r\n\r\n## 定义环形队列\r\n\r\n```java\r\npublic class CircleArrayQueue {\r\n    private int maxSize;\r\n    private int front; // 指向队列的第一个数据, 队列的头指针\r\n    private int rear; // 指向队列最后一个数据的后一个位置, 队列的尾指针. (也就是指向队列还没有存数据的第一个位置)\r\n    private int[] arr; // 存储数据 -> 模拟队列\r\n    \r\n    public CircleArrayQueue(int maxSize) {\r\n        this.maxSize = maxSize;\r\n        arr = new int[maxSize];\r\n        // 此时front以及rear默认值为0,不需要设置\r\n    }\r\n}\r\n```\r\n\r\n## 给环形队列添加一些功能 (判断队列是否存满, 判断队列是否为空, 向队列中存入数据, 从队列中去除数据, 遍历打印队列, 显示队列的头信息)\r\n\r\n```java\r\n// 判断队列是否存满\r\npublic boolean isFull() {\r\n    // 由于rear指向队列最后一个元素后一个位置,当rear指向倒数第二个元素时,且没有取值的情况, 也就是front保持原值,那么队列被存满\r\n    return (rear + 1) % maxSize == front;\r\n}\r\n\r\n// 判断队列是否为空\r\npublic boolean isEmpty() {\r\n    // 当队列最后一个数据的指针与队列第一个数据的指针相等时, 队列为空\r\n    return rear == front;\r\n}\r\n\r\n// 向队列中存入数据, 存入到队列的尾部\r\npublic void add(int value) {\r\n    // 判断队列是否存满\r\n    if(isFull()) {\r\n        System.out.println(\"queue is full\");\r\n        return;\r\n    }\r\n    // 由于rear指向队列最后一个数据的后一个位置,这里可以直接赋值\r\n    arr[rear] = value;\r\n    // 把rear向后移动一个位置, 考虑对maxSize取余, 防止rear超过队列最大长度\r\n    rear = ( rear + 1 ) % maxSize;\r\n}\r\n\r\n// 从队列中获取数据\r\npublic int get() {\r\n    // 判断队列是否为空\r\n    if(isEmpty()) {\r\n        throw new RuntimeException(\"queue is empty\");\r\n    }\r\n    int temp = front;\r\n    // 把队里的头指针向后移动一个位置, 考虑对maxSize取余, 防止front超过队列的最大长度\r\n    front = ( front + 1 ) % maxSize;\r\n    return arr[temp];\r\n}\r\n\r\n// 计算队列中数据的个数\r\npublic int getValidDataCount() {\r\n    // 加上maxSize的原因是防止rear小于front的情况出现\r\n    return (rear + maxSize - front) % maxSize;\r\n}\r\n\r\n// 遍历队列\r\nfor(int i = front ;i < front + getValidDataCount(); i++) {\r\n    // 判断队列是否为空\r\n    if(isEmpty) {\r\n        System.out.println(\"queue is empty\");\r\n        return;\r\n    }\r\n    // 注意这里的索引需要对maxSize取余, 防止索引超出, 出现异常\r\n    System.out.println(arr[i % maxSize])\r\n}\r\n\r\n// 显示队列的头信息\r\npublic int getQueueHeadInfo() {\r\n    // 判断队列是否为空\r\n    if(isEmpty) {\r\n        throw new RuntimeException(\"Queue is empty\");\r\n    }\r\n    return arr[front];\r\n}\r\n```\r\n\r\n', 10, '[\"数据结构\",\"Java\"]', NULL, 8, 10, 0, 0, 0, 1, 3, 0);
INSERT INTO `article` VALUES (32, '2025-12-31 09:05:03.000', '2025-12-31 09:05:06.000', 'css背景属性', 'css背景属性汇总', '# 背景属性汇总\r\n\r\n## background-color\r\n\r\n> 用于设置`背景颜色`\r\n\r\n## background-image\r\n\r\n> 用于设置背景图片\r\n\r\n## background-repeat\r\n\r\n>  用于设置背景图片是否平铺,默认为平铺(repeat 水平和垂直方向都平铺)\r\n>\r\n>  属性值还有: \r\n>\r\n>  1.`no-repeat`(不平铺)\r\n>\r\n>  2.`repeat-x`(水平平铺)\r\n>\r\n>  3.repeat-y(垂直平铺)\r\n\r\n## background-position\r\n\r\n> 用于设置背景图片的位置,图片默认位置为`左上`\r\n>\r\n> 可以通过`方位词`或者`像素`指定\r\n\r\n> 使用`方位词`指定背景图片位置\r\n\r\n```html\r\n// 方位词一共有right,left,top,bottom,center\r\n// 其中center,right,left用于设置水平方向的位置,  center,top,bottom用于设置垂直方向的位置\r\nbackground-position: center top;\r\nbackground-position: top center;\r\n// 这两种显示的背景图片的位置的效果一致\r\n// 可以省略一个方位词,省略的方位词默认是center\r\nbackgoround-position: top;\r\n```\r\n\r\n> 使用`像素`指定图片位置\r\n\r\n```html\r\n// 根据左上偏离的像素确定位置\r\n// 使用像素指定背景图片的位置时,第一个为x轴,第二个为y轴\r\nbackground-position: 12px,14px;\r\n// 可以省略一个,省略的只能为y,那么y默认为center\r\nbackground-position: 12px;\r\n```\r\n\r\n> 使用`方位词`以及`像素`指定背景图片的位置\r\n\r\n```html\r\n// 使用方位词以及像素指定背景图片的位置时,第一个为水平,第二个为垂直方向\r\nbackground-position: 12px top;\r\nbackground-position: right 12px;\r\n```\r\n\r\n## background-attachment\r\n\r\n> 控制图片的固定\r\n>\r\n> 主要有两个值\r\n>\r\n> 1.`scroll`: 随着页面滚动,当页面滚动时,图片会被覆盖\r\n>\r\n> 2.`fixed`: 页面滚动时,图片还是会一直固定在一个位置\r\n\r\n## 背景属性复合写法\r\n\r\n> 复合写法: background:\r\n>\r\n> 背景属性复合写法没有要求顺序\r\n>\r\n> 但是, 规范是background-color background-image background-repeat background-attachment background-poisition\r\n>\r\n> 例如: background: pink url(../) no-repaet fixed center top\r\n\r\n## 背景颜色透明度设置\r\n\r\n> 设置背景颜色的透明度:\r\n>\r\n> background: rgba(0,0,0,0.5)\r\n>\r\n> // 其中四个参数, 第一个表示红色 第二个表示绿色, 但三个表示蓝色, 最后一个表示透明度\r\n>\r\n> // 透明度的范围是0~1, 0表示完全透明 1表示不透明\r\n\r\n', 13, '[\"前端\",\"CSS\"]', NULL, 8, 14, 0, 0, 0, 1, 3, 0);
INSERT INTO `article` VALUES (33, '2025-12-31 09:07:50.000', '2025-12-31 09:07:50.000', '冒泡排序', '使用java实现冒泡排序', '# 冒泡排序\r\n\r\n## 思路: 循环比较前后大小, 把最大的交换到最后\r\n\r\n```java\r\npublic void bubbleSort(int[] arr) {\r\n    // 外面循环每循环一次就会将最大值推送到内部循环的最后一个索引位置, 达到排序的效果\r\n    for(int i = 0; i < arr.length; i++) {\r\n        // 多减一个1是为了防止索引超出异常\r\n        for(int j = 0; j < arr.length - i - 1; j++) {\r\n            if(arr[j] > arr[j + 1]) {\r\n                int temp = arr[j];\r\n                arr[j] = arr[j + 1];\r\n                arr[j + 1] = temp;\r\n            }\r\n        }\r\n    }\r\n}\r\n```\r\n\r\n', 12, '[\"数据结构\",\"Java\"]', NULL, 7, 421, 0, 0, 0, 1, 3, 0);
INSERT INTO `article` VALUES (34, '2025-12-31 09:09:16.000', '2025-12-31 09:09:12.000', '插入排序', '使用java实现插入排序', '# 插入排序\r\n\r\n## 思路:  使用该出索引值与前面排好序的前部分数组进行比较, 如果比前面排好顺序的前部分数组小, 就让前部分数组后移,直到该数大于前部分数组的某个值为止\r\n\r\n```java\r\npublic void selectSort(int[] arr) {\r\n    for(int i = 1; i < arr.length; i++) {\r\n        // 记录当前i索引处要插入的值\r\n        int current = arr[i]\r\n        // 从第i-1个开始比较大小\r\n        int j = i - 1;\r\n        // 如果j>=0 以及当前要插入到值小于前面比较的数据\r\n        while(j >= 0 && arr[j] > current) {\r\n            // 后移\r\n            arr[j + 1] = arr[j];\r\n            // j--, 继续比较\r\n            j--;\r\n        }\r\n        // 空出的位置在j + 1处, 让该值插入到这个位置\r\n        arr[j + 1] = current;\r\n    }\r\n}\r\n```\r\n\r\n', 12, '[\"数据结构\",\"Java\"]', NULL, 8, 0, 0, 0, 0, 1, 3, 0);
INSERT INTO `article` VALUES (35, '2025-12-31 09:16:06.000', '2026-01-15 20:31:16.394', '快速排序', '使用java实现快速排序', '# 快速排序\r\n\r\n## 思路: 将数组通过一个中间值分成两部分, 将大于中间值的数都放到他的右边, 将小于他的数都放到他的左边, 然后取左半边以及右半边进行递归,直到最终排序成功\r\n\r\n```java\r\npublic void quickSort(int arr[], int left, int right) {\r\n    int l = left;\r\n    int r = right;\r\n    int pivot = (left + right) / 2;\r\n    // 如果左索引大于了右索引就结束循环\r\n    while(l < r) {\r\n        // 获取左边部分数组比中间值大的\r\n        while(arr[l] < arr[pivot]) {\r\n            l++;\r\n        }\r\n        // 获取右边数组部分比中间值小的\r\n        while(arr[r] > arr[pivot]) {\r\n            r--;\r\n        }\r\n        // 假设交换完成后, l以及r都会到达pivot的位置, 这时就要退出循环\r\n        if(l >= r) {\r\n             break;\r\n        }\r\n        // 交换左右边的值\r\n        int temp = arr[l];\r\n        arr[l] = arr[r];\r\n        arr[r] = temp;\r\n        \r\n        // 处理死循环的问题\r\n        // 出现死循环的原因有: 当左右两边有等于中间值的值是, 且刚好左边和右边交换值都等于中间值, 这时就会一直交换这两个\r\n        \r\n        // 如果此时l到达中间值的位置, 让r向前移一位, 防止死循环\r\n        if(arr[l] == arr[pivot]) {\r\n            r--;\r\n        }\r\n        // 如果此时r到达了中间值的位置, 让l向后移一位, 防止死循环\r\n        if(arr[r] == arr[pivot]) {\r\n            l++;\r\n        }\r\n    }\r\n    // 处理堆栈溢出异常问题\r\n    // 堆栈移除异常问题出现原因有: 当l=r时, 进行递归左部分或者有部分可能是相同部分的递归\r\n    // 假设一个数组为[5,5,5,5,5], 查找左右两边中间值后发现左右索引都到了中间位置, 那么左右两边递归的还是同样的区间, 这时就会出现递归死循环, 导致堆栈移除异常\r\n    // 当 l = r 时, 左右两边的递归加起来跟上一次的区间一致, 所以要改变l与r的值防止这种情况\r\n    if(l == r) {\r\n        l++;\r\n        r--;\r\n    }\r\n    if(l < right) {\r\n        quickSort(arr,left,r);\r\n    }\r\n    if(r > left) {\r\n        quickSort(arr,l,right);\r\n    }\r\n}\r\n```\r\n\r\n', 12, '[\"数据结构\",\"Java\"]', NULL, 8, 5, 0, 1, 0, 1, 3, 0);

-- ----------------------------
-- Table structure for article_category
-- ----------------------------
DROP TABLE IF EXISTS `article_category`;
CREATE TABLE `article_category`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of article_category
-- ----------------------------
INSERT INTO `article_category` VALUES (1, '后端', 1);
INSERT INTO `article_category` VALUES (2, '前端', 1);
INSERT INTO `article_category` VALUES (3, '数据库', 1);
INSERT INTO `article_category` VALUES (4, 'Java', 1);
INSERT INTO `article_category` VALUES (5, 'Vue3', 1);
INSERT INTO `article_category` VALUES (6, 'React', 1);
INSERT INTO `article_category` VALUES (7, 'Go', 1);
INSERT INTO `article_category` VALUES (8, 'Mybatis', 1);
INSERT INTO `article_category` VALUES (9, 'Gorm', 1);
INSERT INTO `article_category` VALUES (10, '数据结构', 1);
INSERT INTO `article_category` VALUES (11, 'RabbitMQ', 1);
INSERT INTO `article_category` VALUES (12, '排序', 1);
INSERT INTO `article_category` VALUES (13, 'css', 1);

-- ----------------------------
-- Table structure for article_like
-- ----------------------------
DROP TABLE IF EXISTS `article_like`;
CREATE TABLE `article_like`  (
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `article_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  UNIQUE INDEX `idx_article_likes`(`user_id` ASC, `article_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of article_like
-- ----------------------------
INSERT INTO `article_like` VALUES (7, 1, '2025-11-03 10:29:28.000');
INSERT INTO `article_like` VALUES (7, 2, '2025-11-03 10:29:39.000');
INSERT INTO `article_like` VALUES (7, 3, '2025-11-03 10:30:05.000');
INSERT INTO `article_like` VALUES (1, 8, '2025-11-03 12:53:21.000');
INSERT INTO `article_like` VALUES (1, 10, '2025-11-03 12:53:38.000');
INSERT INTO `article_like` VALUES (1, 1, '2025-11-23 15:01:51.765');
INSERT INTO `article_like` VALUES (1, 27, '2025-12-01 20:43:42.674');

-- ----------------------------
-- Table structure for article_tag
-- ----------------------------
DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `browse_count` bigint NULL DEFAULT NULL,
  `p_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 338 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of article_tag
-- ----------------------------
INSERT INTO `article_tag` VALUES (33, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Python', 178, 0);
INSERT INTO `article_tag` VALUES (34, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', '爬虫开发', 758, 33);
INSERT INTO `article_tag` VALUES (35, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', '数据分析', 258, 33);
INSERT INTO `article_tag` VALUES (36, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', '机器学习', 16, 33);
INSERT INTO `article_tag` VALUES (37, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', '深度学习', 306, 33);
INSERT INTO `article_tag` VALUES (38, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', '自动化脚本', 481, 33);
INSERT INTO `article_tag` VALUES (39, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Flask', 488, 33);
INSERT INTO `article_tag` VALUES (40, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Django', 999, 33);
INSERT INTO `article_tag` VALUES (41, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'FastAPI', 532, 33);
INSERT INTO `article_tag` VALUES (42, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Web开发', 663, 33);
INSERT INTO `article_tag` VALUES (43, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', '科学计算', 720, 33);
INSERT INTO `article_tag` VALUES (44, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', '协程', 613, 33);
INSERT INTO `article_tag` VALUES (45, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Java', 903, 0);
INSERT INTO `article_tag` VALUES (46, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'SpringBoot', 679, 45);
INSERT INTO `article_tag` VALUES (47, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Spring Cloud', 686, 45);
INSERT INTO `article_tag` VALUES (48, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'MyBatis', 393, 45);
INSERT INTO `article_tag` VALUES (49, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'JVM', 908, 45);
INSERT INTO `article_tag` VALUES (50, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', '微服务', 361, 45);
INSERT INTO `article_tag` VALUES (51, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Java Web', 80, 45);
INSERT INTO `article_tag` VALUES (52, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'JUC 并发', 318, 45);
INSERT INTO `article_tag` VALUES (53, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Netty', 351, 45);
INSERT INTO `article_tag` VALUES (54, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', '设计模式', 801, 45);
INSERT INTO `article_tag` VALUES (55, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', '编程语言', 952, 0);
INSERT INTO `article_tag` VALUES (56, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'C', 359, 55);
INSERT INTO `article_tag` VALUES (57, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'C++', 938, 55);
INSERT INTO `article_tag` VALUES (58, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Go', 615, 55);
INSERT INTO `article_tag` VALUES (59, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Rust', 262, 55);
INSERT INTO `article_tag` VALUES (60, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'JavaScript', 465, 55);
INSERT INTO `article_tag` VALUES (61, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'TypeScript', 540, 55);
INSERT INTO `article_tag` VALUES (62, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Kotlin', 306, 55);
INSERT INTO `article_tag` VALUES (63, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'PHP', 911, 55);
INSERT INTO `article_tag` VALUES (64, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Swift', 636, 55);
INSERT INTO `article_tag` VALUES (65, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Dart', 447, 55);
INSERT INTO `article_tag` VALUES (66, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'R', 329, 55);
INSERT INTO `article_tag` VALUES (67, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'Shell', 306, 55);
INSERT INTO `article_tag` VALUES (68, '2024-05-01 12:00:00.000', '2024-05-01 12:00:00.000', 'SQL', 541, 55);
INSERT INTO `article_tag` VALUES (69, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '开发工具', 787, 0);
INSERT INTO `article_tag` VALUES (70, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', 'Git', 315, 69);
INSERT INTO `article_tag` VALUES (71, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', 'VS Code', 212, 69);
INSERT INTO `article_tag` VALUES (72, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', 'IDEA', 117, 69);
INSERT INTO `article_tag` VALUES (73, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', 'Docker', 950, 69);
INSERT INTO `article_tag` VALUES (74, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', 'Postman', 401, 69);
INSERT INTO `article_tag` VALUES (75, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', 'Jenkins', 152, 69);
INSERT INTO `article_tag` VALUES (76, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', 'Kubernetes', 559, 69);
INSERT INTO `article_tag` VALUES (77, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', 'Nginx', 340, 69);
INSERT INTO `article_tag` VALUES (78, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', 'GitHub', 24, 69);
INSERT INTO `article_tag` VALUES (79, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', 'GitLab', 98, 69);
INSERT INTO `article_tag` VALUES (80, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '数据结构与算法', 420, 0);
INSERT INTO `article_tag` VALUES (81, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '链表', 807, 80);
INSERT INTO `article_tag` VALUES (82, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '树', 774, 80);
INSERT INTO `article_tag` VALUES (83, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '二叉树', 451, 80);
INSERT INTO `article_tag` VALUES (84, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '图', 933, 80);
INSERT INTO `article_tag` VALUES (85, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '动态规划', 313, 80);
INSERT INTO `article_tag` VALUES (86, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '贪心', 765, 80);
INSERT INTO `article_tag` VALUES (87, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '排序算法', 889, 80);
INSERT INTO `article_tag` VALUES (88, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '回溯', 151, 80);
INSERT INTO `article_tag` VALUES (89, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '堆', 88, 80);
INSERT INTO `article_tag` VALUES (90, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '栈', 987, 80);
INSERT INTO `article_tag` VALUES (91, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '队列', 673, 80);
INSERT INTO `article_tag` VALUES (92, '2026-01-05 14:00:00.000', '2026-01-05 14:00:00.000', '大数据', 405, 0);
INSERT INTO `article_tag` VALUES (93, '2026-01-05 14:00:01.000', '2026-01-05 14:00:01.000', 'Hadoop', 4, 92);
INSERT INTO `article_tag` VALUES (94, '2026-01-05 14:00:01.000', '2026-01-05 14:00:01.000', 'Spark', 805, 92);
INSERT INTO `article_tag` VALUES (95, '2026-01-05 14:00:01.000', '2026-01-05 14:00:01.000', 'Flink', 16, 92);
INSERT INTO `article_tag` VALUES (96, '2026-01-05 14:00:01.000', '2026-01-05 14:00:01.000', 'Hive', 663, 92);
INSERT INTO `article_tag` VALUES (97, '2026-01-05 14:00:01.000', '2026-01-05 14:00:01.000', 'HBase', 268, 92);
INSERT INTO `article_tag` VALUES (98, '2026-01-05 14:00:01.000', '2026-01-05 14:00:01.000', 'Kafka', 351, 92);
INSERT INTO `article_tag` VALUES (99, '2026-01-05 14:00:01.000', '2026-01-05 14:00:01.000', '数据仓库', 953, 92);
INSERT INTO `article_tag` VALUES (100, '2026-01-05 14:00:01.000', '2026-01-05 14:00:01.000', '数仓建模', 713, 92);
INSERT INTO `article_tag` VALUES (112, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '前端', 675, 0);
INSERT INTO `article_tag` VALUES (113, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'HTML', 253, 112);
INSERT INTO `article_tag` VALUES (114, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'CSS', 242, 112);
INSERT INTO `article_tag` VALUES (115, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'JavaScript', 450, 112);
INSERT INTO `article_tag` VALUES (116, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Vue', 525, 112);
INSERT INTO `article_tag` VALUES (117, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'React', 273, 112);
INSERT INTO `article_tag` VALUES (118, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Next.js', 793, 112);
INSERT INTO `article_tag` VALUES (119, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'TailwindCSS', 148, 112);
INSERT INTO `article_tag` VALUES (120, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '小程序', 360, 112);
INSERT INTO `article_tag` VALUES (121, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Vite', 355, 112);
INSERT INTO `article_tag` VALUES (122, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '前端工程化', 697, 112);
INSERT INTO `article_tag` VALUES (123, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '后端', 420, 0);
INSERT INTO `article_tag` VALUES (124, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'RESTful API', 8, 123);
INSERT INTO `article_tag` VALUES (125, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'JWT 鉴权', 782, 123);
INSERT INTO `article_tag` VALUES (126, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '微服务', 886, 123);
INSERT INTO `article_tag` VALUES (127, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'RPC', 83, 123);
INSERT INTO `article_tag` VALUES (128, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '分布式', 757, 123);
INSERT INTO `article_tag` VALUES (129, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '数据库设计', 539, 123);
INSERT INTO `article_tag` VALUES (130, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '消息队列', 423, 123);
INSERT INTO `article_tag` VALUES (131, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'API 网关', 500, 123);
INSERT INTO `article_tag` VALUES (132, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '云原生', 229, 0);
INSERT INTO `article_tag` VALUES (133, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Docker', 648, 132);
INSERT INTO `article_tag` VALUES (134, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Kubernetes', 553, 132);
INSERT INTO `article_tag` VALUES (135, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Istio', 823, 132);
INSERT INTO `article_tag` VALUES (136, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Helm', 454, 132);
INSERT INTO `article_tag` VALUES (137, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Prometheus', 802, 132);
INSERT INTO `article_tag` VALUES (138, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Grafana', 649, 132);
INSERT INTO `article_tag` VALUES (139, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Serverless', 838, 132);
INSERT INTO `article_tag` VALUES (140, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'DevOps', 245, 132);
INSERT INTO `article_tag` VALUES (141, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '移动开发', 711, 0);
INSERT INTO `article_tag` VALUES (142, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Android', 821, 141);
INSERT INTO `article_tag` VALUES (143, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'iOS', 973, 141);
INSERT INTO `article_tag` VALUES (144, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Flutter', 402, 141);
INSERT INTO `article_tag` VALUES (145, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'React Native', 90, 141);
INSERT INTO `article_tag` VALUES (146, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', 'Kotlin', 244, 141);
INSERT INTO `article_tag` VALUES (147, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '移动端适配', 954, 141);
INSERT INTO `article_tag` VALUES (148, '2026-01-05 14:03:25.000', '2026-01-05 14:03:25.000', '性能优化', 35, 141);
INSERT INTO `article_tag` VALUES (149, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '人工智能', 314, 0);
INSERT INTO `article_tag` VALUES (150, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '深度学习', 468, 149);
INSERT INTO `article_tag` VALUES (151, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '机器学习', 397, 149);
INSERT INTO `article_tag` VALUES (152, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'NLP', 580, 149);
INSERT INTO `article_tag` VALUES (153, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'CV', 710, 149);
INSERT INTO `article_tag` VALUES (154, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'Transformer', 811, 149);
INSERT INTO `article_tag` VALUES (155, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'AI 应用开发', 924, 149);
INSERT INTO `article_tag` VALUES (156, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '语音识别', 187, 149);
INSERT INTO `article_tag` VALUES (157, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '大模型', 165, 149);
INSERT INTO `article_tag` VALUES (158, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '网络与通信', 264, 0);
INSERT INTO `article_tag` VALUES (159, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'TCP/IP', 827, 158);
INSERT INTO `article_tag` VALUES (160, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'HTTP', 343, 158);
INSERT INTO `article_tag` VALUES (161, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'HTTPS', 235, 158);
INSERT INTO `article_tag` VALUES (162, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'WebSocket', 146, 158);
INSERT INTO `article_tag` VALUES (163, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '网络抓包', 25, 158);
INSERT INTO `article_tag` VALUES (164, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'DNS', 690, 158);
INSERT INTO `article_tag` VALUES (165, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'VPN', 374, 158);
INSERT INTO `article_tag` VALUES (166, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '嵌入式', 799, 0);
INSERT INTO `article_tag` VALUES (167, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '单片机', 875, 166);
INSERT INTO `article_tag` VALUES (168, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'STM32', 980, 166);
INSERT INTO `article_tag` VALUES (169, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'Arduino', 275, 166);
INSERT INTO `article_tag` VALUES (170, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '物联网', 433, 166);
INSERT INTO `article_tag` VALUES (171, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '嵌入式 Linux', 342, 166);
INSERT INTO `article_tag` VALUES (172, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '硬件开发', 410, 0);
INSERT INTO `article_tag` VALUES (173, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'PCB 设计', 25, 172);
INSERT INTO `article_tag` VALUES (174, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '电路基础', 895, 172);
INSERT INTO `article_tag` VALUES (175, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '数字电路', 402, 172);
INSERT INTO `article_tag` VALUES (176, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'FPGA', 325, 172);
INSERT INTO `article_tag` VALUES (177, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '传感器', 417, 172);
INSERT INTO `article_tag` VALUES (178, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '游戏', 114, 0);
INSERT INTO `article_tag` VALUES (179, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'Unity', 318, 178);
INSERT INTO `article_tag` VALUES (180, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'Unreal（UE）', 248, 178);
INSERT INTO `article_tag` VALUES (181, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', 'Cocos', 288, 178);
INSERT INTO `article_tag` VALUES (182, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '游戏引擎', 694, 178);
INSERT INTO `article_tag` VALUES (183, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '游戏策划', 606, 178);
INSERT INTO `article_tag` VALUES (184, '2026-01-05 14:04:32.000', '2026-01-05 14:04:32.000', '游戏建模', 951, 178);
INSERT INTO `article_tag` VALUES (185, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', 'HarmonyOS', 937, 0);
INSERT INTO `article_tag` VALUES (186, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', 'ArkUI', 834, 185);
INSERT INTO `article_tag` VALUES (187, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '鸿蒙应用开发', 356, 185);
INSERT INTO `article_tag` VALUES (188, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '分布式能力', 282, 185);
INSERT INTO `article_tag` VALUES (189, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', 'DevEco Studio', 342, 185);
INSERT INTO `article_tag` VALUES (190, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '微软技术', 863, 0);
INSERT INTO `article_tag` VALUES (191, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '.NET', 292, 190);
INSERT INTO `article_tag` VALUES (192, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', 'C#', 868, 190);
INSERT INTO `article_tag` VALUES (193, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', 'Azure', 467, 190);
INSERT INTO `article_tag` VALUES (194, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', 'WPF', 732, 190);
INSERT INTO `article_tag` VALUES (195, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', 'Blazor', 257, 190);
INSERT INTO `article_tag` VALUES (196, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '操作系统', 89, 0);
INSERT INTO `article_tag` VALUES (197, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', 'Linux', 675, 196);
INSERT INTO `article_tag` VALUES (198, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', 'Windows 内核', 107, 196);
INSERT INTO `article_tag` VALUES (199, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '系统调优', 513, 196);
INSERT INTO `article_tag` VALUES (200, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', 'Shell 编程', 243, 196);
INSERT INTO `article_tag` VALUES (201, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '文件系统', 679, 196);
INSERT INTO `article_tag` VALUES (202, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '搜索', 663, 0);
INSERT INTO `article_tag` VALUES (203, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', 'Elasticsearch', 282, 202);
INSERT INTO `article_tag` VALUES (204, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '搜索引擎', 418, 202);
INSERT INTO `article_tag` VALUES (205, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '倒排索引', 246, 202);
INSERT INTO `article_tag` VALUES (206, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '搜索优化', 978, 202);
INSERT INTO `article_tag` VALUES (207, '2026-01-05 14:05:25.000', '2026-01-05 14:05:25.000', '设计模式', 152, 0);
INSERT INTO `article_tag` VALUES (208, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', '工厂模式', 826, 207);
INSERT INTO `article_tag` VALUES (209, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', '单例模式', 676, 207);
INSERT INTO `article_tag` VALUES (210, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', '策略模式', 903, 207);
INSERT INTO `article_tag` VALUES (211, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', '观察者模式', 486, 207);
INSERT INTO `article_tag` VALUES (212, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', '装饰器模式', 721, 207);
INSERT INTO `article_tag` VALUES (213, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', 'MVC/MVVM', 147, 207);
INSERT INTO `article_tag` VALUES (214, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', '测试', 571, 0);
INSERT INTO `article_tag` VALUES (215, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', '单元测试', 418, 214);
INSERT INTO `article_tag` VALUES (216, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', '接口测试', 375, 214);
INSERT INTO `article_tag` VALUES (217, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', '自动化测试', 621, 214);
INSERT INTO `article_tag` VALUES (218, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', '性能测试', 980, 214);
INSERT INTO `article_tag` VALUES (219, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', 'pytest', 40, 214);
INSERT INTO `article_tag` VALUES (220, '2026-01-05 14:05:26.000', '2026-01-05 14:05:26.000', 'JMeter', 259, 214);
INSERT INTO `article_tag` VALUES (221, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '云平台', 174, 0);
INSERT INTO `article_tag` VALUES (222, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '阿里云', 94, 221);
INSERT INTO `article_tag` VALUES (223, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '腾讯云', 948, 221);
INSERT INTO `article_tag` VALUES (224, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '华为云', 461, 221);
INSERT INTO `article_tag` VALUES (225, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'AWS', 459, 221);
INSERT INTO `article_tag` VALUES (226, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'GCP', 913, 221);
INSERT INTO `article_tag` VALUES (227, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '软件工程', 187, 0);
INSERT INTO `article_tag` VALUES (228, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '架构设计', 198, 227);
INSERT INTO `article_tag` VALUES (229, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '敏捷开发', 431, 227);
INSERT INTO `article_tag` VALUES (230, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'Scrum', 561, 227);
INSERT INTO `article_tag` VALUES (231, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'CI/CD', 514, 227);
INSERT INTO `article_tag` VALUES (232, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '项目管理', 886, 227);
INSERT INTO `article_tag` VALUES (233, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '区块链', 889, 0);
INSERT INTO `article_tag` VALUES (234, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'Web3', 787, 233);
INSERT INTO `article_tag` VALUES (235, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'Solidity', 270, 233);
INSERT INTO `article_tag` VALUES (236, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '智能合约', 990, 233);
INSERT INTO `article_tag` VALUES (237, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '以太坊', 137, 233);
INSERT INTO `article_tag` VALUES (238, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '区块链安全', 719, 233);
INSERT INTO `article_tag` VALUES (239, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '数学', 185, 0);
INSERT INTO `article_tag` VALUES (240, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '高等数学', 767, 239);
INSERT INTO `article_tag` VALUES (241, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '线性代数', 280, 239);
INSERT INTO `article_tag` VALUES (242, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '概率论', 102, 239);
INSERT INTO `article_tag` VALUES (243, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '统计学', 668, 239);
INSERT INTO `article_tag` VALUES (244, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '运筹学', 35, 239);
INSERT INTO `article_tag` VALUES (245, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '离散数学', 173, 239);
INSERT INTO `article_tag` VALUES (246, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '运维', 759, 0);
INSERT INTO `article_tag` VALUES (247, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'Linux 运维', 279, 246);
INSERT INTO `article_tag` VALUES (248, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '自动化运维', 116, 246);
INSERT INTO `article_tag` VALUES (249, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '监控告警', 745, 246);
INSERT INTO `article_tag` VALUES (250, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '故障排查', 376, 246);
INSERT INTO `article_tag` VALUES (251, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '网络空间安全', 647, 0);
INSERT INTO `article_tag` VALUES (252, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '渗透测试', 105, 251);
INSERT INTO `article_tag` VALUES (253, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '漏洞分析', 588, 251);
INSERT INTO `article_tag` VALUES (254, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'CTF', 622, 251);
INSERT INTO `article_tag` VALUES (255, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '逆向工程', 350, 251);
INSERT INTO `article_tag` VALUES (256, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'Web 安全', 882, 251);
INSERT INTO `article_tag` VALUES (257, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '加密算法', 361, 251);
INSERT INTO `article_tag` VALUES (258, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '服务器', 161, 0);
INSERT INTO `article_tag` VALUES (259, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'Nginx', 722, 258);
INSERT INTO `article_tag` VALUES (260, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'Apache', 128, 258);
INSERT INTO `article_tag` VALUES (261, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'Redis', 473, 258);
INSERT INTO `article_tag` VALUES (262, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'MySQL', 982, 258);
INSERT INTO `article_tag` VALUES (263, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'PostgreSQL', 493, 258);
INSERT INTO `article_tag` VALUES (264, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'MongoDB', 520, 258);
INSERT INTO `article_tag` VALUES (265, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '学习和成长', 123, 0);
INSERT INTO `article_tag` VALUES (266, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '自律', 52, 265);
INSERT INTO `article_tag` VALUES (267, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '高效学习', 892, 265);
INSERT INTO `article_tag` VALUES (268, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '时间管理', 303, 265);
INSERT INTO `article_tag` VALUES (269, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '效率工具', 842, 265);
INSERT INTO `article_tag` VALUES (270, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '教育培训', 301, 0);
INSERT INTO `article_tag` VALUES (271, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '在线课程', 981, 270);
INSERT INTO `article_tag` VALUES (272, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', 'IT 培训', 1, 270);
INSERT INTO `article_tag` VALUES (273, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '技术分享', 64, 270);
INSERT INTO `article_tag` VALUES (274, '2026-01-05 14:06:13.000', '2026-01-05 14:06:13.000', '考证', 317, 270);
INSERT INTO `article_tag` VALUES (275, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '用户体验设计', 395, 0);
INSERT INTO `article_tag` VALUES (276, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '交互设计', 22, 275);
INSERT INTO `article_tag` VALUES (277, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'UI 设计', 928, 275);
INSERT INTO `article_tag` VALUES (278, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '原型设计', 574, 275);
INSERT INTO `article_tag` VALUES (279, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '用户研究', 85, 275);
INSERT INTO `article_tag` VALUES (280, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '信息架构', 705, 275);
INSERT INTO `article_tag` VALUES (281, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '音视频', 269, 0);
INSERT INTO `article_tag` VALUES (282, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'FFmpeg', 230, 281);
INSERT INTO `article_tag` VALUES (283, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'WebRTC', 343, 281);
INSERT INTO `article_tag` VALUES (284, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '视频编码', 27, 281);
INSERT INTO `article_tag` VALUES (285, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '音频处理', 107, 281);
INSERT INTO `article_tag` VALUES (286, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '直播技术', 453, 281);
INSERT INTO `article_tag` VALUES (287, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '行业数字化', 946, 0);
INSERT INTO `article_tag` VALUES (288, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '电商系统', 371, 287);
INSERT INTO `article_tag` VALUES (289, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '金融科技', 17, 287);
INSERT INTO `article_tag` VALUES (290, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '智慧城市', 975, 287);
INSERT INTO `article_tag` VALUES (291, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '物流系统', 823, 287);
INSERT INTO `article_tag` VALUES (292, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '医疗信息化', 189, 287);
INSERT INTO `article_tag` VALUES (293, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '非 IT 技术', 476, 0);
INSERT INTO `article_tag` VALUES (294, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '心理学', 816, 293);
INSERT INTO `article_tag` VALUES (295, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '摄影', 653, 293);
INSERT INTO `article_tag` VALUES (296, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '绘画', 816, 293);
INSERT INTO `article_tag` VALUES (297, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '音乐', 122, 293);
INSERT INTO `article_tag` VALUES (298, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '健身', 165, 293);
INSERT INTO `article_tag` VALUES (299, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '前沿技术', 456, 0);
INSERT INTO `article_tag` VALUES (300, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'AI Agent', 785, 299);
INSERT INTO `article_tag` VALUES (301, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '数字孪生', 559, 299);
INSERT INTO `article_tag` VALUES (302, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '边缘计算', 439, 299);
INSERT INTO `article_tag` VALUES (303, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'VR/AR', 519, 299);
INSERT INTO `article_tag` VALUES (304, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '量子计算', 279, 299);
INSERT INTO `article_tag` VALUES (305, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'IT 工具', 838, 0);
INSERT INTO `article_tag` VALUES (306, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'ChatGPT', 352, 305);
INSERT INTO `article_tag` VALUES (307, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'Copilot', 247, 305);
INSERT INTO `article_tag` VALUES (308, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'Cursor', 179, 305);
INSERT INTO `article_tag` VALUES (309, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '飞书', 154, 305);
INSERT INTO `article_tag` VALUES (310, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'Notion', 234, 305);
INSERT INTO `article_tag` VALUES (311, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '开发组件', 711, 0);
INSERT INTO `article_tag` VALUES (312, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'UI 组件库', 851, 311);
INSERT INTO `article_tag` VALUES (313, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '工具类', 122, 311);
INSERT INTO `article_tag` VALUES (314, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '自定义 Hook', 58, 311);
INSERT INTO `article_tag` VALUES (315, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '封装组件', 923, 311);
INSERT INTO `article_tag` VALUES (316, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '开源', 444, 0);
INSERT INTO `article_tag` VALUES (317, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '开源项目', 449, 316);
INSERT INTO `article_tag` VALUES (318, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '代码审查', 917, 316);
INSERT INTO `article_tag` VALUES (319, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'GitHub 贡献', 236, 316);
INSERT INTO `article_tag` VALUES (320, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'Issue/PR', 428, 316);
INSERT INTO `article_tag` VALUES (321, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '其他', 436, 0);
INSERT INTO `article_tag` VALUES (322, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '科技新闻', 893, 321);
INSERT INTO `article_tag` VALUES (323, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '博客搭建', 160, 321);
INSERT INTO `article_tag` VALUES (324, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '个人网站', 120, 321);
INSERT INTO `article_tag` VALUES (325, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '效率提升', 120, 321);
INSERT INTO `article_tag` VALUES (326, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '3C 硬件', 240, 0);
INSERT INTO `article_tag` VALUES (327, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '键盘', 839, 326);
INSERT INTO `article_tag` VALUES (328, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '鼠标', 479, 326);
INSERT INTO `article_tag` VALUES (329, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '显卡', 875, 326);
INSERT INTO `article_tag` VALUES (330, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '电脑 DIY', 942, 326);
INSERT INTO `article_tag` VALUES (331, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '主机搭建', 85, 326);
INSERT INTO `article_tag` VALUES (332, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'AIGC', 598, 0);
INSERT INTO `article_tag` VALUES (333, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '提示词工程', 737, 332);
INSERT INTO `article_tag` VALUES (334, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '图像生成', 893, 332);
INSERT INTO `article_tag` VALUES (335, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'AI 编程', 252, 332);
INSERT INTO `article_tag` VALUES (336, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', '模型微调', 581, 332);
INSERT INTO `article_tag` VALUES (337, '2026-01-05 14:07:07.000', '2026-01-05 14:07:07.000', 'AI 创作', 151, 332);

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `coverage` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `href` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of banner
-- ----------------------------

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `article_id` bigint UNSIGNED NULL DEFAULT NULL,
  `parent_id` bigint UNSIGNED NULL DEFAULT NULL,
  `root_parent_id` bigint UNSIGNED NULL DEFAULT NULL,
  `like_count` bigint UNSIGNED NULL DEFAULT NULL,
  `reply_to_user_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, '2025-11-03 10:31:55.000', '2025-11-03 10:31:58.000', '恭喜你开始了博客创作的旅程！正则表达式的运用是一个很有趣的话题，相信你一定会写得很棒。接下来，我建议你可以尝试写一些实际案例，或者结合其他编程语言来展示正则表达式的应用，这样可以让读者更容易理解和运用。希望你能继续坚持写下去，期待看到更多精彩的内容！', 7, 1, NULL, NULL, 0, NULL);
INSERT INTO `comment` VALUES (2, '2025-11-24 21:34:08.000', '2025-11-24 21:34:06.000', '恭喜你开始了博客创作的旅程！正则表达式的运用是一个很有趣的话题。', 1, 1, 1, 1, 0, 7);
INSERT INTO `comment` VALUES (4, '2025-11-24 21:51:21.000', '2025-11-24 21:51:23.000', '恭喜你开始了博客创作的旅程！正则表达式的运用是一个很有趣的话题。', 8, 1, 1, 1, 0, 1);
INSERT INTO `comment` VALUES (5, '2025-11-25 10:37:12.362', '2025-11-25 15:53:01.014', '欢迎大家评论', 1, 1, NULL, NULL, 1, NULL);
INSERT INTO `comment` VALUES (8, '2025-11-25 16:14:12.892', '2025-11-25 21:41:24.277', 'Golang基础知识🫠', 1, 1, NULL, NULL, 1, NULL);
INSERT INTO `comment` VALUES (9, '2025-11-25 16:23:11.605', '2025-11-25 16:23:11.605', '111', 1, 1, NULL, NULL, 0, NULL);
INSERT INTO `comment` VALUES (10, '2025-11-25 16:23:46.384', '2025-11-25 16:23:46.384', '222', 7, 1, NULL, NULL, 0, NULL);
INSERT INTO `comment` VALUES (11, '2025-11-25 16:26:20.603', '2025-11-25 16:26:20.603', '333', 8, 1, NULL, NULL, 0, NULL);
INSERT INTO `comment` VALUES (12, '2025-11-25 16:26:23.483', '2025-11-25 16:26:23.483', '444', 8, 1, NULL, NULL, 0, NULL);
INSERT INTO `comment` VALUES (13, '2025-11-25 16:53:11.166', '2025-11-25 22:02:50.129', '555', 7, 1, NULL, NULL, 1, NULL);
INSERT INTO `comment` VALUES (14, '2025-11-25 16:53:15.723', '2025-11-25 22:03:17.347', '666', 8, 1, NULL, NULL, 2, NULL);
INSERT INTO `comment` VALUES (16, '2025-11-25 16:56:04.205', '2025-11-25 16:56:04.205', '777', 1, 1, NULL, NULL, 0, NULL);
INSERT INTO `comment` VALUES (17, '2025-11-25 16:57:41.652', '2025-11-25 16:57:41.652', '来了🙃', 1, 1, 5, 5, 0, 1);
INSERT INTO `comment` VALUES (18, '2025-11-25 16:58:06.135', '2025-11-26 16:09:03.419', '第一🤪', 1, 1, NULL, NULL, 1, NULL);
INSERT INTO `comment` VALUES (19, '2025-11-25 17:37:09.179', '2025-11-25 17:37:09.179', '1', 1, 1, 14, 14, 0, 8);
INSERT INTO `comment` VALUES (24, '2025-12-16 09:24:25.536', '2025-12-16 09:24:25.536', '很好的排序方法', 1, 29, NULL, NULL, 0, NULL);
INSERT INTO `comment` VALUES (25, '2026-01-15 20:31:16.373', '2026-01-15 20:31:16.373', '11', 1, 35, NULL, NULL, 0, NULL);

-- ----------------------------
-- Table structure for comment_like
-- ----------------------------
DROP TABLE IF EXISTS `comment_like`;
CREATE TABLE `comment_like`  (
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `comment_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  UNIQUE INDEX `idx_comment_likes`(`user_id` ASC, `comment_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of comment_like
-- ----------------------------
INSERT INTO `comment_like` VALUES (1, 5, '2025-11-25 15:53:01.014');
INSERT INTO `comment_like` VALUES (1, 14, '2025-11-25 17:07:04.959');
INSERT INTO `comment_like` VALUES (1, 21, '2025-11-25 21:10:39.222');
INSERT INTO `comment_like` VALUES (7, 21, '2025-11-25 21:40:50.864');
INSERT INTO `comment_like` VALUES (7, 8, '2025-11-25 21:41:24.276');
INSERT INTO `comment_like` VALUES (7, 13, '2025-11-25 22:02:50.129');
INSERT INTO `comment_like` VALUES (7, 14, '2025-11-25 22:03:17.346');
INSERT INTO `comment_like` VALUES (1, 18, '2025-11-26 16:09:03.418');

-- ----------------------------
-- Table structure for conversation
-- ----------------------------
DROP TABLE IF EXISTS `conversation`;
CREATE TABLE `conversation`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `type` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `last_message` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `last_message_time` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of conversation
-- ----------------------------

-- ----------------------------
-- Table structure for conversation_user
-- ----------------------------
DROP TABLE IF EXISTS `conversation_user`;
CREATE TABLE `conversation_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `conversation_id` bigint UNSIGNED NOT NULL,
  `user_id` bigint UNSIGNED NOT NULL,
  `is_pinned` tinyint(1) NULL DEFAULT NULL,
  `is_muted` tinyint(1) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_conversation_user`(`conversation_id` ASC, `user_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of conversation_user
-- ----------------------------

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `abstract` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `coverage` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `is_default` tinyint(1) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of favorite
-- ----------------------------
INSERT INTO `favorite` VALUES (1, '2025-11-03 10:30:43.000', '2025-11-04 15:59:12.675', '默认收藏夹', 'jackson的默认收藏夹', NULL, 1, 0);
INSERT INTO `favorite` VALUES (2, '2025-11-03 22:54:34.540', '2025-11-04 16:09:32.724', 'java', '保存Java相关知识的博文', NULL, 1, 1);
INSERT INTO `favorite` VALUES (5, '2025-11-03 23:05:03.308', '2025-11-04 15:37:09.256', '前端', '用于收集前端知识blog', '', 1, 0);
INSERT INTO `favorite` VALUES (9, '2025-11-24 10:18:54.048', '2025-11-24 10:18:54.048', 'golang', '4324', '', 1, 0);
INSERT INTO `favorite` VALUES (10, '2025-12-18 11:18:10.282', '2025-12-18 11:18:10.282', 'java', '收藏Java的文章', '', 7, 1);

-- ----------------------------
-- Table structure for favorite_articles
-- ----------------------------
DROP TABLE IF EXISTS `favorite_articles`;
CREATE TABLE `favorite_articles`  (
  `favorite_id` bigint UNSIGNED NOT NULL,
  `article_id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`favorite_id`, `article_id`) USING BTREE,
  UNIQUE INDEX `favorite_articles_key`(`favorite_id` ASC, `article_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of favorite_articles
-- ----------------------------
INSERT INTO `favorite_articles` VALUES (1, 1, '2025-11-23 22:20:00.335');
INSERT INTO `favorite_articles` VALUES (1, 6, '2025-11-24 11:40:47.672');
INSERT INTO `favorite_articles` VALUES (1, 12, '2025-11-24 11:39:52.349');
INSERT INTO `favorite_articles` VALUES (2, 1, '2025-11-23 22:09:16.537');
INSERT INTO `favorite_articles` VALUES (2, 3, '2025-11-24 11:44:29.151');
INSERT INTO `favorite_articles` VALUES (2, 7, '2025-11-24 11:34:20.267');
INSERT INTO `favorite_articles` VALUES (2, 8, '2025-11-24 11:32:48.527');
INSERT INTO `favorite_articles` VALUES (5, 1, '2025-11-23 22:20:06.716');
INSERT INTO `favorite_articles` VALUES (5, 3, '2025-11-24 11:44:33.356');
INSERT INTO `favorite_articles` VALUES (5, 8, '2025-11-24 11:32:52.327');
INSERT INTO `favorite_articles` VALUES (9, 1, '2025-11-24 11:47:09.969');
INSERT INTO `favorite_articles` VALUES (9, 8, '2025-11-24 11:34:07.187');
INSERT INTO `favorite_articles` VALUES (10, 29, '2025-12-18 11:18:12.979');

-- ----------------------------
-- Table structure for global_notification
-- ----------------------------
DROP TABLE IF EXISTS `global_notification`;
CREATE TABLE `global_notification`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `href` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of global_notification
-- ----------------------------

-- ----------------------------
-- Table structure for image
-- ----------------------------
DROP TABLE IF EXISTS `image`;
CREATE TABLE `image`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `filename` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `size` bigint NULL DEFAULT NULL,
  `hash` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of image
-- ----------------------------

-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `type` tinyint NULL DEFAULT NULL,
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `level` tinyint NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `addr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_read` tinyint(1) NULL DEFAULT NULL,
  `login_status` tinyint(1) NULL DEFAULT NULL,
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `login_type` tinyint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of log
-- ----------------------------
INSERT INTO `log` VALUES (38, '2025-10-26 23:06:58.363', '2025-10-26 23:06:58.363', 2, '根新站点信息', '{\r\n    \"name\": \"jackson\"\r\n}\n{\"code\":0,\"msg\":\"修改站点信息\"}', 1, 1, '127.0.0.1', '内网IP', 0, 0, '', '', 0);

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `type` tinyint NULL DEFAULT NULL,
  `send_user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `receive_user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `extra` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `is_read` tinyint(1) NULL DEFAULT NULL,
  `read_time` datetime(3) NULL DEFAULT NULL,
  `send_time` datetime(3) NULL DEFAULT NULL,
  `plan_push_time` datetime(3) NULL DEFAULT NULL,
  `real_push_time` datetime(3) NULL DEFAULT NULL,
  `status` tinyint NULL DEFAULT NULL,
  `action_message` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `content_type` tinyint NULL DEFAULT NULL,
  `conversation_id` bigint UNSIGNED NULL DEFAULT NULL,
  `session_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 112 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (1, '2025-12-15 22:51:15.000', '2025-12-28 17:18:28.961', 1, 1, 7, 'hello', NULL, 1, NULL, '2025-12-15 22:52:27.000', '2025-12-15 22:52:33.000', '2025-12-15 22:52:37.000', 0, NULL, 1, NULL, 1);
INSERT INTO `message` VALUES (2, '2025-12-15 23:02:03.000', '2025-12-28 21:27:15.235', 1, 7, 1, 'hello, what\'s up', NULL, 1, NULL, '2025-12-15 23:02:22.000', '2025-12-15 23:02:24.000', '2025-12-15 23:02:26.000', 0, NULL, 1, NULL, 1);
INSERT INTO `message` VALUES (3, '2025-12-18 10:29:16.000', '2025-12-18 10:29:18.000', 2, 7, 1, '恭喜你开始了博客创作的旅程！正则表达式的运用是一个很有趣的话题，相信你一定会写得很棒。接下来，我建议你可以尝试写一些实际案例，或者结合其他编程语言来展示正则表达式的应用，这样可以让读者更容易理解和运用。希望你能继续坚持写下去，期待看到更多精彩的内容！', '1', 0, NULL, '2025-12-18 10:30:23.000', '2025-12-18 10:30:28.000', '2025-12-18 10:30:32.000', 0, '评论了你的文章', 1, NULL, NULL);
INSERT INTO `message` VALUES (4, '2025-12-18 11:11:27.000', '2025-12-18 11:11:30.000', 2, 8, 1, '恭喜你开始了博客创作的旅程！正则表达式的运用是一个很有趣的话题。', '1', 0, NULL, '2025-12-18 11:12:02.000', '2025-12-18 11:12:04.000', '2025-12-18 11:12:08.000', 0, '回复了你的评论', 1, NULL, NULL);
INSERT INTO `message` VALUES (5, '2025-12-23 20:07:09.000', '2025-12-23 20:07:12.000', 3, 7, 1, NULL, '29', 0, NULL, '2025-12-23 20:08:52.000', '2025-12-23 20:08:55.000', '2025-12-23 20:08:57.000', 0, '收藏了你的博客', 1, NULL, NULL);
INSERT INTO `message` VALUES (6, '2025-12-23 20:10:20.000', '2025-12-23 20:10:22.000', 3, 7, 1, NULL, '1', 0, NULL, '2025-12-23 20:10:32.000', '2025-12-23 20:10:35.000', '2025-12-23 20:10:39.000', 0, '点赞了你的博文', 1, NULL, NULL);
INSERT INTO `message` VALUES (8, '2025-12-21 17:05:21.000', '2025-12-21 17:05:18.000', 4, 8, 1, NULL, '8', 0, NULL, '2025-12-21 17:05:25.000', '2025-12-21 17:05:28.000', '2025-12-21 17:05:30.000', 0, '通过 博客 关注了你', 1, NULL, NULL);
INSERT INTO `message` VALUES (9, '2025-12-23 20:11:01.000', '2025-12-23 20:11:03.000', 3, 7, 1, NULL, '2', 0, NULL, '2025-12-23 20:11:10.000', '2025-12-23 20:11:12.000', '2025-12-23 20:11:14.000', 0, '点赞了你的博文', 1, NULL, NULL);
INSERT INTO `message` VALUES (10, '2025-12-23 20:12:13.000', '2025-12-23 20:12:00.000', 3, 7, 1, NULL, '3', 0, NULL, '2025-12-23 20:12:02.000', '2025-12-23 20:12:04.000', '2025-12-23 20:12:06.000', 0, '点赞了你的博文', 1, NULL, NULL);
INSERT INTO `message` VALUES (11, '2025-12-23 20:23:41.000', '2025-12-23 20:23:43.000', 4, 7, 1, NULL, '7', 0, NULL, '2025-12-23 20:24:02.000', '2025-12-23 20:24:05.000', '2025-12-23 20:24:08.000', 0, '通过 博客 关注了你', 1, NULL, NULL);
INSERT INTO `message` VALUES (14, '2025-12-27 14:41:57.964', '2025-12-28 12:54:23.704', 1, 1, 8, '111', '', 1, NULL, '2025-12-27 14:41:57.964', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (15, '2025-12-27 14:42:31.320', '2025-12-28 12:54:23.704', 1, 1, 8, '222', '', 1, NULL, '2025-12-27 14:42:31.320', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (16, '2025-12-27 14:44:26.135', '2025-12-28 12:54:23.704', 1, 1, 8, '333', '', 1, NULL, '2025-12-27 14:44:26.135', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (17, '2025-12-27 14:44:53.374', '2025-12-28 12:54:23.704', 1, 1, 8, '444', '', 1, NULL, '2025-12-27 14:44:53.374', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (18, '2025-12-27 14:45:07.860', '2026-01-15 20:31:01.683', 1, 8, 1, '555', '', 1, NULL, '2025-12-27 14:45:07.860', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (19, '2025-12-27 14:45:17.724', '2026-01-15 20:31:01.683', 1, 8, 1, '666', '', 1, NULL, '2025-12-27 14:45:17.724', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (20, '2025-12-27 14:45:32.317', '2026-01-15 20:31:01.683', 1, 8, 1, '777', '', 1, NULL, '2025-12-27 14:45:32.317', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (21, '2025-12-27 14:56:28.524', '2025-12-28 12:54:23.704', 1, 1, 8, '888', '', 1, NULL, '2025-12-27 14:56:28.524', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (22, '2025-12-27 14:56:44.561', '2026-01-15 20:31:01.683', 1, 8, 1, '999', '', 1, NULL, '2025-12-27 14:56:44.561', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (23, '2025-12-27 14:57:35.976', '2025-12-28 12:54:23.704', 1, 1, 8, '110', '', 1, NULL, '2025-12-27 14:57:35.976', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (24, '2025-12-27 14:58:04.978', '2025-12-28 12:54:23.704', 1, 1, 8, '112', '', 1, NULL, '2025-12-27 14:58:04.978', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (25, '2025-12-27 14:58:16.880', '2025-12-28 12:54:23.704', 1, 1, 8, '113', '', 1, NULL, '2025-12-27 14:58:16.880', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (26, '2025-12-27 15:00:02.596', '2026-01-15 20:31:01.683', 1, 8, 1, '444', '', 1, NULL, '2025-12-27 15:00:02.596', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (27, '2025-12-27 15:00:13.349', '2026-01-15 20:31:01.683', 1, 8, 1, '555', '', 1, NULL, '2025-12-27 15:00:13.349', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (28, '2025-12-27 15:05:54.358', '2025-12-28 12:54:23.704', 1, 1, 8, '666', '', 1, NULL, '2025-12-27 15:05:54.358', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (29, '2025-12-27 15:06:14.974', '2026-01-15 20:31:01.683', 1, 8, 1, '777', '', 1, NULL, '2025-12-27 15:06:14.974', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (30, '2025-12-27 15:06:32.571', '2025-12-28 12:54:23.704', 1, 1, 8, '888', '', 1, NULL, '2025-12-27 15:06:32.571', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (31, '2025-12-27 15:07:31.032', '2025-12-28 12:54:23.704', 1, 1, 8, '999', '', 1, NULL, '2025-12-27 15:07:31.032', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (32, '2025-12-27 15:07:38.282', '2026-01-15 20:31:01.683', 1, 8, 1, '1000', '', 1, NULL, '2025-12-27 15:07:38.282', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (33, '2025-12-27 15:17:32.151', '2025-12-28 12:54:23.704', 1, 1, 8, '2222', '', 1, NULL, '2025-12-27 15:17:32.151', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (34, '2025-12-27 15:24:31.788', '2025-12-28 12:54:23.704', 1, 1, 8, '2241', '', 1, NULL, '2025-12-27 15:24:31.787', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (35, '2025-12-28 12:30:27.533', '2026-01-15 20:31:01.683', 1, 8, 1, '222', '', 1, NULL, '2025-12-28 12:30:27.533', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (36, '2025-12-28 12:31:01.991', '2026-01-15 20:31:01.683', 1, 8, 1, '333', '', 1, NULL, '2025-12-28 12:31:01.991', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (37, '2025-12-28 12:31:08.077', '2025-12-28 12:54:23.704', 1, 1, 8, '444', '', 1, NULL, '2025-12-28 12:31:08.077', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (38, '2025-12-28 12:31:20.693', '2025-12-28 12:54:23.704', 1, 1, 8, '555', '', 1, NULL, '2025-12-28 12:31:20.693', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (39, '2025-12-28 12:31:26.041', '2025-12-28 12:54:23.704', 1, 1, 8, '333', '', 1, NULL, '2025-12-28 12:31:26.041', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (40, '2025-12-28 12:31:55.732', '2026-01-15 20:31:01.683', 1, 8, 1, '666', '', 1, NULL, '2025-12-28 12:31:55.732', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (41, '2025-12-28 12:32:18.326', '2026-01-15 20:31:01.683', 1, 8, 1, 'hello', '', 1, NULL, '2025-12-28 12:32:18.326', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (42, '2025-12-28 12:32:56.847', '2025-12-28 12:54:23.704', 1, 1, 8, 'hello, what\'s up', '', 1, NULL, '2025-12-28 12:32:56.847', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (43, '2025-12-28 12:33:17.654', '2026-01-15 20:31:01.683', 1, 8, 1, '请问?', '', 1, NULL, '2025-12-28 12:33:17.654', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (44, '2025-12-28 12:33:24.890', '2025-12-28 12:54:23.704', 1, 1, 8, 'who are you', '', 1, NULL, '2025-12-28 12:33:24.890', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (45, '2025-12-28 12:33:38.026', '2026-01-15 20:31:01.683', 1, 8, 1, 'i am zs, what\'s your name', '', 1, NULL, '2025-12-28 12:33:38.026', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (46, '2025-12-28 12:33:48.797', '2025-12-28 12:54:23.704', 1, 1, 8, 'i am jackson', '', 1, NULL, '2025-12-28 12:33:48.797', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (47, '2025-12-28 12:33:54.447', '2026-01-15 20:31:01.683', 1, 8, 1, 'ok', '', 1, NULL, '2025-12-28 12:33:54.447', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (48, '2025-12-28 12:34:28.293', '2026-01-15 20:31:01.683', 1, 8, 1, 'jackson, what are you doing', '', 1, NULL, '2025-12-28 12:34:28.293', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (49, '2025-12-28 12:35:12.271', '2025-12-28 12:54:23.704', 1, 1, 8, 'i am chatting with you', '', 1, NULL, '2025-12-28 12:35:12.271', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (50, '2025-12-28 12:35:47.221', '2026-01-15 20:31:01.683', 1, 8, 1, 'it\'s so funny', '', 1, NULL, '2025-12-28 12:35:47.221', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (51, '2025-12-28 12:36:00.011', '2025-12-28 12:54:23.704', 1, 1, 8, 'hhh', '', 1, NULL, '2025-12-28 12:36:00.011', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (52, '2025-12-28 12:37:48.402', '2025-12-28 15:19:08.623', 1, 8, 7, 'hello', '', 1, NULL, '2025-12-28 12:37:48.402', NULL, NULL, 0, '', 1, NULL, 4);
INSERT INTO `message` VALUES (53, '2025-12-28 12:39:50.789', '2025-12-28 12:41:40.707', 1, 7, 8, 'hello, i from China', '', 1, NULL, '2025-12-28 12:39:50.789', NULL, NULL, 0, '', 1, NULL, 4);
INSERT INTO `message` VALUES (54, '2025-12-28 12:41:17.274', '2025-12-28 15:19:08.623', 1, 8, 7, 'i from China too', '', 1, NULL, '2025-12-28 12:41:17.274', NULL, NULL, 0, '', 1, NULL, 4);
INSERT INTO `message` VALUES (55, '2025-12-28 12:42:01.040', '2025-12-28 12:42:01.040', 1, 7, 8, 'oh, is crazy', '', 0, NULL, '2025-12-28 12:42:01.040', NULL, NULL, 0, '', 1, NULL, 4);
INSERT INTO `message` VALUES (56, '2025-12-28 12:56:35.327', '2025-12-28 17:18:28.961', 1, 1, 7, 'fdsa', '', 1, NULL, '2025-12-28 12:56:35.327', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (57, '2025-12-28 12:57:01.884', '2025-12-28 21:27:15.235', 1, 7, 1, 'ffa', '', 1, NULL, '2025-12-28 12:57:01.884', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (58, '2025-12-28 12:57:05.905', '2025-12-28 17:18:28.961', 1, 1, 7, 'fsda', '', 1, NULL, '2025-12-28 12:57:05.905', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (59, '2025-12-28 13:57:53.384', '2025-12-28 17:18:28.961', 1, 1, 7, 'fdsa', '', 1, NULL, '2025-12-28 13:57:53.384', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (60, '2025-12-28 14:03:21.098', '2025-12-28 17:18:28.961', 1, 1, 7, '11', '', 1, NULL, '2025-12-28 14:03:21.098', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (61, '2025-12-28 14:06:10.994', '2025-12-28 17:18:28.961', 1, 1, 7, '333', '', 1, NULL, '2025-12-28 14:06:10.994', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (62, '2025-12-28 14:06:34.125', '2025-12-28 21:27:15.235', 1, 7, 1, '444', '', 1, NULL, '2025-12-28 14:06:34.125', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (63, '2025-12-28 14:07:31.262', '2025-12-28 21:27:15.235', 1, 7, 1, '555', '', 1, NULL, '2025-12-28 14:07:31.262', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (64, '2025-12-28 14:11:25.120', '2025-12-28 21:27:15.235', 1, 7, 1, '555', '', 1, NULL, '2025-12-28 14:11:25.120', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (65, '2025-12-28 14:12:54.449', '2025-12-28 17:18:28.961', 1, 1, 7, '666', '', 1, NULL, '2025-12-28 14:12:54.449', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (66, '2025-12-28 14:13:00.868', '2025-12-28 17:18:28.961', 1, 1, 7, '777', '', 1, NULL, '2025-12-28 14:13:00.868', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (67, '2025-12-28 14:13:05.587', '2025-12-28 17:18:28.961', 1, 1, 7, '888', '', 1, NULL, '2025-12-28 14:13:05.587', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (68, '2025-12-28 14:14:17.163', '2025-12-28 17:18:28.961', 1, 1, 7, '999', '', 1, NULL, '2025-12-28 14:14:17.163', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (69, '2025-12-28 14:15:28.357', '2025-12-28 21:27:15.235', 1, 7, 1, '100', '', 1, NULL, '2025-12-28 14:15:28.357', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (70, '2025-12-28 14:17:02.529', '2025-12-28 21:27:15.235', 1, 7, 1, '182', '', 1, NULL, '2025-12-28 14:17:02.529', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (71, '2025-12-28 14:18:37.921', '2025-12-28 21:27:15.235', 1, 7, 1, '143', '', 1, NULL, '2025-12-28 14:18:37.921', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (72, '2025-12-28 14:18:52.452', '2025-12-28 21:27:15.235', 1, 7, 1, '4321', '', 1, NULL, '2025-12-28 14:18:52.452', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (73, '2025-12-28 14:19:33.862', '2025-12-28 17:18:28.961', 1, 1, 7, '32142', '', 1, NULL, '2025-12-28 14:19:33.862', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (74, '2025-12-28 14:19:46.636', '2025-12-28 17:18:28.961', 1, 1, 7, '4321', '', 1, NULL, '2025-12-28 14:19:46.636', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (75, '2025-12-28 14:19:51.690', '2025-12-28 17:18:28.961', 1, 1, 7, '4312', '', 1, NULL, '2025-12-28 14:19:51.690', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (76, '2025-12-28 14:21:01.047', '2025-12-28 17:18:28.961', 1, 1, 7, 'f3214ds', '', 1, NULL, '2025-12-28 14:21:01.047', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (77, '2025-12-28 14:21:45.180', '2025-12-28 21:27:15.235', 1, 7, 1, 'fdsafds', '', 1, NULL, '2025-12-28 14:21:45.180', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (78, '2025-12-28 14:22:01.106', '2025-12-28 21:27:15.235', 1, 7, 1, 'dfadfda', '', 1, NULL, '2025-12-28 14:22:01.106', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (79, '2025-12-28 14:22:07.713', '2025-12-28 21:27:15.235', 1, 7, 1, 'fadsf', '', 1, NULL, '2025-12-28 14:22:07.713', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (80, '2025-12-28 14:22:15.182', '2025-12-28 21:27:15.235', 1, 7, 1, 'fasfds', '', 1, NULL, '2025-12-28 14:22:15.182', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (81, '2025-12-28 14:25:22.382', '2025-12-28 17:18:28.961', 1, 1, 7, '43214d', '', 1, NULL, '2025-12-28 14:25:22.382', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (82, '2025-12-28 14:25:30.981', '2025-12-28 17:18:28.961', 1, 1, 7, 'fdsafsd', '', 1, NULL, '2025-12-28 14:25:30.981', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (83, '2025-12-28 14:26:08.754', '2025-12-28 21:27:15.235', 1, 7, 1, 'fdsafsd', '', 1, NULL, '2025-12-28 14:26:08.754', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (84, '2025-12-28 14:26:58.484', '2025-12-28 14:26:58.484', 1, 1, 8, 'fsdafsd', '', 0, NULL, '2025-12-28 14:26:58.484', NULL, NULL, 0, '', 1, NULL, 3);
INSERT INTO `message` VALUES (85, '2025-12-28 14:27:33.670', '2025-12-28 17:18:28.961', 1, 1, 7, 'yes', '', 1, NULL, '2025-12-28 14:27:33.670', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (86, '2025-12-28 14:33:57.213', '2025-12-28 21:27:15.235', 1, 7, 1, 'no', '', 1, NULL, '2025-12-28 14:33:57.213', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (87, '2025-12-28 14:34:13.862', '2025-12-28 17:18:28.961', 1, 1, 7, 'ok, no', '', 1, NULL, '2025-12-28 14:34:13.862', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (88, '2025-12-28 14:35:01.728', '2025-12-28 21:27:15.235', 1, 7, 1, 'no, ok yes', '', 1, NULL, '2025-12-28 14:35:01.728', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (89, '2025-12-28 14:37:01.855', '2025-12-28 21:27:15.235', 1, 7, 1, 'fdsa', '', 1, NULL, '2025-12-28 14:37:01.855', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (90, '2025-12-28 14:58:23.964', '2025-12-28 21:27:15.235', 1, 7, 1, '4321', '', 1, NULL, '2025-12-28 14:58:23.964', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (91, '2025-12-28 14:58:35.096', '2025-12-28 17:18:28.961', 1, 1, 7, 'fdsa', '', 1, NULL, '2025-12-28 14:58:35.095', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (98, '2025-12-28 15:15:08.734', '2025-12-28 17:18:28.961', 1, 1, 7, 'fdsa', '', 1, NULL, '2025-12-28 15:15:08.734', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (99, '2025-12-28 15:15:16.683', '2025-12-28 21:27:15.235', 1, 7, 1, 'rewqr', '', 1, NULL, '2025-12-28 15:15:16.683', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (103, '2025-12-28 15:18:49.310', '2025-12-28 21:27:15.235', 1, 7, 1, 'fsafd', '', 1, NULL, '2025-12-28 15:18:49.310', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (104, '2025-12-28 15:19:06.320', '2025-12-28 17:18:28.961', 1, 1, 7, 'jkjl', '', 1, NULL, '2025-12-28 15:19:06.320', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (105, '2025-12-28 15:19:13.652', '2025-12-28 17:18:28.961', 1, 1, 7, 'fdsfdsa', '', 1, NULL, '2025-12-28 15:19:13.652', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (108, '2025-12-28 16:11:35.134', '2025-12-28 17:18:28.961', 1, 1, 7, 'fsdaf', '', 1, NULL, '2025-12-28 16:11:35.134', NULL, NULL, 0, '', 1, NULL, 1);
INSERT INTO `message` VALUES (110, '2025-12-28 17:00:17.573', '2025-12-28 17:18:28.961', 1, 1, 7, 'fdsafd', '', 1, NULL, '2025-12-28 17:00:17.573', NULL, NULL, 0, '', 1, NULL, 1);

-- ----------------------------
-- Table structure for session
-- ----------------------------
DROP TABLE IF EXISTS `session`;
CREATE TABLE `session`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `user_id_a` bigint UNSIGNED NULL DEFAULT NULL,
  `user_id_b` bigint UNSIGNED NULL DEFAULT NULL,
  `latest_chat_time` datetime(3) NULL DEFAULT NULL,
  `latest_message` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `is_pinned_a` tinyint(1) NULL DEFAULT NULL,
  `is_pinned_b` tinyint(1) NULL DEFAULT NULL,
  `is_muted_a` tinyint(1) NULL DEFAULT NULL,
  `is_muted_b` tinyint(1) NULL DEFAULT NULL,
  `deleted_at_a` datetime(3) NULL DEFAULT NULL,
  `deleted_at_b` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of session
-- ----------------------------
INSERT INTO `session` VALUES (1, '2025-12-09 09:12:52.000', '2025-12-28 17:00:17.578', 1, 7, '2025-12-28 17:00:17.573', 'fdsafd', 0, 0, 0, 0, NULL, NULL);
INSERT INTO `session` VALUES (3, '2025-12-26 15:59:59.388', '2025-12-28 14:26:58.485', 1, 8, '2025-12-28 14:26:58.484', 'fsdafsd', 0, 0, 1, 0, NULL, NULL);
INSERT INTO `session` VALUES (4, '2025-12-28 11:59:48.741', '2025-12-28 12:42:01.041', 8, 7, '2025-12-28 12:42:01.040', 'oh, is crazy', 0, 0, 1, 0, NULL, NULL);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `abstract` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `register_source` tinyint NULL DEFAULT NULL,
  `password` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `code_age` bigint NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `open_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `role` tinyint NULL DEFAULT NULL,
  `sex` tinyint NULL DEFAULT NULL,
  `birthday` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_user_username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2025-10-27 16:49:25.000', '2025-11-20 14:55:29.520', 'jackson', 'jackson', 'http://jackson1.oss-cn-beijing.aliyuncs.com/1763215266_1738.jpg', 'hello world', 0, '$2a$10$dv/ZzgwxE5HLz5HbulvDoevem3/ClELKUk4BV5QcTxU47RMpchJl.', 1, 'jacksonn36067@gmail.com', NULL, 1, 0, '2007-02-02 00:00:00.000');
INSERT INTO `user` VALUES (7, '2025-10-29 20:54:12.983', '2025-10-29 20:54:12.983', 'wendy', 'wendy', 'http://jackson1.oss-cn-beijing.aliyuncs.com/a61be1b7-f7b3-4027-8e18-930fab3bdeef.jpg', '', 0, '$2a$10$JjiEdx8Cl62wnDAn5b1r/.iMsKzJn3/0ZQtW9OA0LNgGEty0pIlGu', 0, '2252559105@qq.com', '', 0, 1, '2025-11-11 09:22:44.000');
INSERT INTO `user` VALUES (8, '2025-11-05 17:26:34.000', '2025-11-05 17:26:37.000', 'zs', 'zs', 'http://jackson1.oss-cn-beijing.aliyuncs.com/a61be1b7-f7b3-4027-8e18-930fab3bdeef.jpg', NULL, 0, '$2a$10$JjiEdx8Cl62wnDAn5b1r/.iMsKzJn3/0ZQtW9OA0LNgGEty0pIlGu', 1, 'ajcksfsd@qq.com', NULL, 1, 1, '2025-11-11 09:22:46.000');

-- ----------------------------
-- Table structure for user_article_browse_history
-- ----------------------------
DROP TABLE IF EXISTS `user_article_browse_history`;
CREATE TABLE `user_article_browse_history`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `article_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4314 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_article_browse_history
-- ----------------------------
INSERT INTO `user_article_browse_history` VALUES (4263, '2025-12-02 21:43:21.078', '2025-12-02 21:43:21.078', 1, 27);
INSERT INTO `user_article_browse_history` VALUES (4264, '2025-12-02 20:28:05.506', '2025-12-02 20:28:05.506', 1, 1);
INSERT INTO `user_article_browse_history` VALUES (4265, '2025-12-02 20:27:38.293', '2025-12-02 20:27:38.293', 1, 7);
INSERT INTO `user_article_browse_history` VALUES (4266, '2025-12-02 21:13:36.921', '2025-12-02 21:13:36.921', 1, 28);
INSERT INTO `user_article_browse_history` VALUES (4268, '2025-12-02 21:57:59.298', '2025-12-02 21:57:59.298', 1, 29);
INSERT INTO `user_article_browse_history` VALUES (4269, '2025-12-02 21:57:38.134', '2025-12-02 21:57:38.134', 1, 30);
INSERT INTO `user_article_browse_history` VALUES (4270, '2025-12-03 21:22:27.945', '2025-12-03 21:22:27.946', 1, 29);
INSERT INTO `user_article_browse_history` VALUES (4271, '2025-12-03 16:04:10.664', '2025-12-03 16:04:10.664', 1, 29);
INSERT INTO `user_article_browse_history` VALUES (4272, '2025-12-03 16:05:27.438', '2025-12-03 16:05:27.439', 1, 27);
INSERT INTO `user_article_browse_history` VALUES (4273, '2025-12-03 16:06:13.444', '2025-12-03 16:06:13.445', 1, 1);
INSERT INTO `user_article_browse_history` VALUES (4274, '2025-12-03 16:12:07.937', '2025-12-03 16:12:07.937', 1, 30);
INSERT INTO `user_article_browse_history` VALUES (4275, '2025-12-04 08:59:02.487', '2025-12-04 08:59:02.487', 1, 27);
INSERT INTO `user_article_browse_history` VALUES (4276, '2025-12-04 08:59:02.491', '2025-12-04 08:59:02.491', 1, 27);
INSERT INTO `user_article_browse_history` VALUES (4277, '2025-12-04 09:29:27.689', '2025-12-04 09:29:27.689', 1, 29);
INSERT INTO `user_article_browse_history` VALUES (4278, '2025-12-04 11:22:09.994', '2025-12-04 11:22:09.994', 1, 12);
INSERT INTO `user_article_browse_history` VALUES (4279, '2025-12-08 19:49:33.784', '2025-12-08 19:49:33.785', 1, 12);
INSERT INTO `user_article_browse_history` VALUES (4280, '2025-12-08 18:57:37.718', '2025-12-08 18:57:37.718', 1, 12);
INSERT INTO `user_article_browse_history` VALUES (4281, '2025-12-08 19:47:35.472', '2025-12-08 19:47:35.472', 1, 8);
INSERT INTO `user_article_browse_history` VALUES (4282, '2025-12-08 19:40:17.047', '2025-12-08 19:40:17.047', 1, 8);
INSERT INTO `user_article_browse_history` VALUES (4283, '2025-12-08 19:49:14.521', '2025-12-08 19:49:14.521', 1, 29);
INSERT INTO `user_article_browse_history` VALUES (4284, '2025-12-09 11:03:08.303', '2025-12-09 11:03:08.303', 1, 28);
INSERT INTO `user_article_browse_history` VALUES (4285, '2025-12-09 11:03:08.303', '2025-12-09 11:03:08.303', 1, 28);
INSERT INTO `user_article_browse_history` VALUES (4286, '2025-12-09 11:05:36.510', '2025-12-09 11:05:36.510', 1, 29);
INSERT INTO `user_article_browse_history` VALUES (4287, '2025-12-11 21:51:31.756', '2025-12-11 21:51:31.757', 1, 29);
INSERT INTO `user_article_browse_history` VALUES (4288, '2025-12-16 11:10:06.576', '2025-12-16 11:10:06.576', 1, 29);
INSERT INTO `user_article_browse_history` VALUES (4289, '2025-12-16 09:23:51.931', '2025-12-16 09:23:51.931', 1, 29);
INSERT INTO `user_article_browse_history` VALUES (4290, '2025-12-18 11:18:13.003', '2025-12-18 11:18:13.003', 7, 29);
INSERT INTO `user_article_browse_history` VALUES (4291, '2025-12-18 11:17:54.210', '2025-12-18 11:17:54.210', 7, 29);
INSERT INTO `user_article_browse_history` VALUES (4292, '2025-12-21 17:24:52.672', '2025-12-21 17:24:52.672', 1, 12);
INSERT INTO `user_article_browse_history` VALUES (4293, '2025-12-21 17:24:52.673', '2025-12-21 17:24:52.673', 1, 12);
INSERT INTO `user_article_browse_history` VALUES (4294, '2025-12-26 15:20:13.854', '2025-12-26 15:20:13.854', 1, 12);
INSERT INTO `user_article_browse_history` VALUES (4295, '2025-12-26 15:20:13.855', '2025-12-26 15:20:13.855', 1, 12);
INSERT INTO `user_article_browse_history` VALUES (4296, '2025-12-26 15:47:23.868', '2025-12-26 15:47:23.868', 1, 31);
INSERT INTO `user_article_browse_history` VALUES (4297, '2025-12-26 15:47:23.876', '2025-12-26 15:47:23.876', 1, 31);
INSERT INTO `user_article_browse_history` VALUES (4298, '2025-12-26 16:06:35.204', '2025-12-26 16:06:35.204', 8, 29);
INSERT INTO `user_article_browse_history` VALUES (4299, '2025-12-26 16:06:35.209', '2025-12-26 16:06:35.209', 8, 29);
INSERT INTO `user_article_browse_history` VALUES (4300, '2025-12-27 14:30:01.543', '2025-12-27 14:30:01.544', 1, 31);
INSERT INTO `user_article_browse_history` VALUES (4301, '2025-12-27 12:01:53.408', '2025-12-27 12:01:53.408', 1, 31);
INSERT INTO `user_article_browse_history` VALUES (4302, '2025-12-28 11:59:30.705', '2025-12-28 11:59:30.705', 8, 8);
INSERT INTO `user_article_browse_history` VALUES (4303, '2025-12-28 11:59:30.715', '2025-12-28 11:59:30.715', 8, 8);
INSERT INTO `user_article_browse_history` VALUES (4304, '2025-12-28 12:27:53.640', '2025-12-28 12:27:53.640', 8, 29);
INSERT INTO `user_article_browse_history` VALUES (4305, '2025-12-28 12:12:58.974', '2025-12-28 12:12:58.974', 8, 29);
INSERT INTO `user_article_browse_history` VALUES (4306, '2025-12-28 21:25:59.438', '2025-12-28 21:25:59.438', 1, 31);
INSERT INTO `user_article_browse_history` VALUES (4307, '2025-12-28 21:25:59.445', '2025-12-28 21:25:59.445', 1, 31);
INSERT INTO `user_article_browse_history` VALUES (4308, '2026-01-15 20:31:16.395', '2026-01-15 20:31:16.395', 1, 35);
INSERT INTO `user_article_browse_history` VALUES (4309, '2026-01-15 20:30:41.196', '2026-01-15 20:30:41.196', 1, 35);
INSERT INTO `user_article_browse_history` VALUES (4310, '2026-01-15 20:54:34.776', '2026-01-15 20:54:34.776', 1, 12);
INSERT INTO `user_article_browse_history` VALUES (4311, '2026-01-15 20:37:35.354', '2026-01-15 20:37:35.354', 1, 3);
INSERT INTO `user_article_browse_history` VALUES (4312, '2026-01-15 20:54:41.574', '2026-01-15 20:54:41.574', 1, 8);
INSERT INTO `user_article_browse_history` VALUES (4313, '2026-01-15 20:54:41.583', '2026-01-15 20:54:41.583', 1, 8);

-- ----------------------------
-- Table structure for user_article_collect
-- ----------------------------
DROP TABLE IF EXISTS `user_article_collect`;
CREATE TABLE `user_article_collect`  (
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `article_id` bigint UNSIGNED NULL DEFAULT NULL,
  `favorite_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  UNIQUE INDEX `idx_user_article_collect`(`user_id` ASC, `article_id` ASC, `favorite_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_article_collect
-- ----------------------------
INSERT INTO `user_article_collect` VALUES (1, 1, 2, '2025-11-23 22:16:23.000');
INSERT INTO `user_article_collect` VALUES (1, 1, 1, '2025-11-23 22:20:00.333');
INSERT INTO `user_article_collect` VALUES (1, 1, 5, '2025-11-23 22:20:06.715');
INSERT INTO `user_article_collect` VALUES (1, 8, 2, '2025-11-24 11:32:48.527');
INSERT INTO `user_article_collect` VALUES (1, 8, 5, '2025-11-24 11:32:52.327');
INSERT INTO `user_article_collect` VALUES (1, 8, 9, '2025-11-24 11:34:07.186');
INSERT INTO `user_article_collect` VALUES (1, 7, 2, '2025-11-24 11:34:20.266');
INSERT INTO `user_article_collect` VALUES (1, 12, 1, '2025-11-24 11:39:52.348');
INSERT INTO `user_article_collect` VALUES (1, 6, 1, '2025-11-24 11:40:47.670');
INSERT INTO `user_article_collect` VALUES (1, 3, 2, '2025-11-24 11:44:29.151');
INSERT INTO `user_article_collect` VALUES (1, 3, 5, '2025-11-24 11:44:33.356');
INSERT INTO `user_article_collect` VALUES (1, 1, 9, '2025-11-24 11:47:09.969');
INSERT INTO `user_article_collect` VALUES (7, 29, 10, '2025-12-18 11:18:12.978');

-- ----------------------------
-- Table structure for user_config
-- ----------------------------
DROP TABLE IF EXISTS `user_config`;
CREATE TABLE `user_config`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint UNSIGNED NOT NULL,
  `hobby_tags` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `update_username_date` datetime(3) NULL DEFAULT NULL,
  `public_collect_list` tinyint(1) NULL DEFAULT NULL,
  `public_follow_list` tinyint(1) NULL DEFAULT NULL,
  `public_fan_list` tinyint(1) NULL DEFAULT NULL,
  `home_style_id` bigint UNSIGNED NULL DEFAULT NULL,
  `public_browse_history` tinyint(1) NULL DEFAULT NULL,
  `public_personal_list` tinyint(1) NULL DEFAULT NULL,
  `public_like_list` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_user_config_user_id`(`user_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_config
-- ----------------------------
INSERT INTO `user_config` VALUES (1, 1, '[\"TypeScript\",\"Spring Boot\",\"Go\",\"爬虫开发\"]', '2025-11-20 11:34:50.375', 1, 1, 1, 1, 1, 1, 1);
INSERT INTO `user_config` VALUES (2, 7, NULL, '2025-09-25 11:41:32.000', 0, 1, 1, 1, 0, 0, 0);
INSERT INTO `user_config` VALUES (3, 8, NULL, '2025-11-11 11:41:54.000', 1, 1, 1, 1, 1, 1, 1);

-- ----------------------------
-- Table structure for user_follow
-- ----------------------------
DROP TABLE IF EXISTS `user_follow`;
CREATE TABLE `user_follow`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `follower_id` bigint UNSIGNED NOT NULL,
  `followed_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_follower_followed`(`follower_id` ASC, `followed_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_follow
-- ----------------------------
INSERT INTO `user_follow` VALUES (3, '2025-11-05 17:29:20.000', '2025-11-05 17:29:12.000', 7, 1);
INSERT INTO `user_follow` VALUES (4, '2025-11-05 17:29:32.000', '2025-11-05 17:29:30.000', 8, 1);
INSERT INTO `user_follow` VALUES (16, '2025-12-08 19:39:32.760', '2025-12-08 19:39:32.760', 1, 1);
INSERT INTO `user_follow` VALUES (17, '2025-12-08 19:47:35.464', '2025-12-08 19:47:35.464', 1, 7);
INSERT INTO `user_follow` VALUES (18, '2025-12-23 20:24:30.763', '2025-12-23 20:24:30.763', 1, 8);

-- ----------------------------
-- Table structure for user_login
-- ----------------------------
DROP TABLE IF EXISTS `user_login`;
CREATE TABLE `user_login`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `addr` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `ua` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_login
-- ----------------------------
INSERT INTO `user_login` VALUES (1, '2025-11-14 12:54:22.061', '2025-11-14 12:54:22.061', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (2, '2025-11-15 20:49:14.798', '2025-11-15 20:49:14.798', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (3, '2025-11-15 22:09:55.838', '2025-11-15 22:09:55.838', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (4, '2025-11-19 13:56:00.265', '2025-11-19 13:56:00.265', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (5, '2025-11-19 13:56:30.517', '2025-11-19 13:56:30.517', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (6, '2025-11-23 13:54:20.625', '2025-11-23 13:54:20.625', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (7, '2025-11-23 14:50:17.262', '2025-11-23 14:50:17.262', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (8, '2025-11-25 17:02:12.653', '2025-11-25 17:02:12.653', 7, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (9, '2025-11-25 17:06:52.236', '2025-11-25 17:06:52.236', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (10, '2025-11-25 21:25:44.558', '2025-11-25 21:25:44.558', 7, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (11, '2025-11-26 11:39:20.253', '2025-11-26 11:39:20.253', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (12, '2025-11-26 16:07:54.827', '2025-11-26 16:07:54.827', 8, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (13, '2025-11-26 16:08:32.240', '2025-11-26 16:08:32.240', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (14, '2025-11-26 16:34:56.277', '2025-11-26 16:34:56.277', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (15, '2025-12-03 16:38:23.565', '2025-12-03 16:38:23.565', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (16, '2025-12-08 15:01:17.959', '2025-12-08 15:01:17.959', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (17, '2025-12-11 21:51:23.610', '2025-12-11 21:51:23.610', 1, '::1', '', '用户名密码登录');
INSERT INTO `user_login` VALUES (18, '2025-12-15 22:17:53.464', '2025-12-15 22:17:53.464', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (19, '2025-12-18 11:17:36.883', '2025-12-18 11:17:36.883', 7, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (20, '2025-12-18 11:25:24.865', '2025-12-18 11:25:24.865', 7, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (21, '2025-12-18 11:25:50.473', '2025-12-18 11:25:50.473', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (22, '2025-12-25 12:54:58.739', '2025-12-25 12:54:58.739', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (23, '2025-12-25 17:19:00.662', '2025-12-25 17:19:00.662', 8, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (24, '2025-12-27 14:29:49.375', '2025-12-27 14:29:49.375', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (25, '2025-12-28 12:09:30.234', '2025-12-28 12:09:30.234', 7, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (26, '2025-12-28 12:10:38.648', '2025-12-28 12:10:38.648', 8, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (27, '2025-12-28 12:12:54.560', '2025-12-28 12:12:54.560', 8, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (28, '2025-12-28 12:14:37.125', '2025-12-28 12:14:37.125', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (29, '2025-12-28 12:39:19.483', '2025-12-28 12:39:19.483', 7, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (30, '2025-12-28 12:55:41.999', '2025-12-28 12:55:41.999', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (31, '2025-12-28 21:32:37.403', '2025-12-28 21:32:37.403', 7, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (32, '2025-12-29 22:24:11.509', '2025-12-29 22:24:11.509', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (33, '2026-01-01 15:51:51.550', '2026-01-01 15:51:51.550', 1, '127.0.0.1', '本机', '用户名密码登录');
INSERT INTO `user_login` VALUES (34, '2026-01-11 16:08:51.022', '2026-01-11 16:08:51.022', 1, '127.0.0.1', '本机', '用户名密码登录');

-- ----------------------------
-- Table structure for user_top_article
-- ----------------------------
DROP TABLE IF EXISTS `user_top_article`;
CREATE TABLE `user_top_article`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `article_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user_top_article`(`user_id` ASC, `article_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_top_article
-- ----------------------------
INSERT INTO `user_top_article` VALUES (1, 1, 1, '2025-11-01 11:06:19.000');
INSERT INTO `user_top_article` VALUES (2, 1, 4, '2025-11-01 11:12:46.000');

SET FOREIGN_KEY_CHECKS = 1;
