# Architecture Overview 架构概述

EShop, A application, demonstrates modern cloud-native practices including *microservices*,event-drive communication,containerization, and AI integration. 

This application follows a distributed architecture pattern where functionality is split into multiple autonomous services（自治服务） that communicate via APIs and message passing.（服务之间通过 API 和 消息传递进行通信）

![eshop_architecture](D:\文档\学习\Golang\微服务\go-zero\eShop\eshop_architecture.png)

## Microservices Architecture 微服务架构

| Service           | Description        | Database   | Communication Methods            |
| ----------------- | ------------------ | ---------- | -------------------------------- |
| Catalog API       | 管理产品目录和库存 | PostgreSQL | gRPC，Event Publishing           |
| Basket API        | 管理用户购物车     | Redis      | gRPC，Event Publishing           |
| Ordering API      | 处理订单           | MySQL      | gRPC，Event Publishing           |
| Identity API      | 管理身份验证和授权 | MySQL      | Gateway HTTP API、OpenID connect |
| Order Processor   | 异步处理订单事件   | MySQL      | Event Subscription               |
| Payment Processor | 异步处理付款事件   | MySQL      | Event Subscription               |

**Event Bus** 使用 go-queue 促进服务之间的异步通信

# Application Structure

* Backend Services 后端服务，将核心业务功能实现为 微服务
* Processor 处理器，通过事件总线处理异步业务流程

## Backend Services and Communication

后端服务通过 HTTP/gRPC 调用和通过事件总线的异步消息传递进行通信。每个服务都有子集的数据库并实现特定的业务功能。

# Backend Services

## Catalog API

管理产品目录信息并提供产品搜索功能，包括AI驱动的语义搜索功能（把 token 映射为高维向量）

**主要特点**：

* 具有向量支持的 PostgreSQL 数据库，用于语义搜索
* 与 AI 服务集成 以生成文本嵌入
* 基于事件的订单状态更改（通过事件总线监听订单状态变化事件，并根据这些事件执行相应的库存管理操作）

## Basket API

为用户管理购物篮数据，提供在购物篮中添加、更新和删除商品的功能

**主要特点**：

* Redis 使用高速内存存储购物篮数据
* gRPC 用于高性能购物篮服务
* 与事件总线集成，用于订单创建事件

## Ordering API

处理订单，包括订单创建、状态更新和历史记录跟踪

**主要特点**

* 用于订单数据持久性的 SQL 数据库
* 具有独立域和基础设施层的领域驱动设计就（DDD）方法
* 基于事件的订单状态更改通信

## Identity API

为 eShop 应用程序提供身份验证和授权服务，包括用户管理

**主要特点**

* OAuth 2.0 和 JWT
* 用于用户数据存储的PostgreSQL数据库

## Order Processor 

处理订单业务逻辑，对来自事件总线的事件进行响应

**主要特点**

* 订阅订单相关事件
* 使用 ordering 数据库
* 实现订单处理的业务逻辑

## Payment Processor

处理支付业务逻辑，响应来自事件总线的事件

**主要特点**

* 订阅与支付相关的事件
* 实现支付处理的业务逻辑

## Event-Driven Communication 事件驱动

eShop 应用程序使用RabbitMQ作为事件总线实现事件驱动框架，从而支持服务之间的异步通信，增强了服务之间订单解耦，并支持以下场景：

* 订单处理管道
* 付款处理
* 商品清单更新

### 集成事件：

用于在服务之间传达状态更改

* 订单状态更改
* 购物篮结账事件
* 库存验证事件

## AI Integration

集成AI提供语义搜索功能，使用 AI 进行文本嵌入，并使用 pgvector 在 PostgreSQL 中进行向量搜索

## Chatbot

* 响应产品查询

* 搜索目录

* 帮助管理购物车

* 提供产品推荐

## Accessing Application Component

| Component 元件                 | Default URL 默认 URL                              | Description 描述                                             |
| ------------------------------ | ------------------------------------------------- | ------------------------------------------------------------ |
| Aspire Dashboard Aspire 仪表板 | [http://localhost:19888](http://localhost:19888/) | Central dashboard for monitoring all services 用于监控所有服务的中央仪表板 |
| Web Application Web 应用程序   | [https://localhost:5001](https://localhost:5001/) | Main eShop web interface 主 eShop Web 界面                   |
| Catalog API 目录 API           | [https://localhost:5201](https://localhost:5201/) | API for product catalog management 用于产品目录管理的 API    |
| Basket API 购物篮 API          | [https://localhost:5202](https://localhost:5202/) | API for shopping basket operations 用于购物篮作的 API        |
| Ordering API 订购 API          | [https://localhost:5203](https://localhost:5203/) | API for order processing 用于订单处理的 API                  |
| Identity API 身份 API          | [https://localhost:5204](https://localhost:5204/) | API for authentication and user management 用于身份验证和用户管理的 API |

### Catelog API

### Endpoints Overview 终端节点概述

| Method 方法 | Endpoint 端点                                                | Version 版本  | Description 描述                                             |
| ----------- | ------------------------------------------------------------ | ------------- | ------------------------------------------------------------ |
| GET         | /api/catalog/items                                           | v1, v2 v1、v2 | List catalog items with pagination 列出带有分页的目录项      |
| GET         | /api/catalog/items/by                                        | v1, v2 v1、v2 | Batch get items by IDs 按 ID 批量获取项目                    |
| GET         | /api/catalog/items/{id}                                      | v1, v2 v1、v2 | Get item by ID 按 ID 获取项目                                |
| GET         | /api/catalog/items/{id}/pic                                  | v1, v2 v1、v2 | Get item picture 获取项目图片                                |
| GET         | /api/catalog/items/by/{name}                                 | v1 第 1 版    | Get items by name 按名称获取项                               |
| GET         | /api/catalog/items/withsemanticrelevance/{text} /api/catalog/items/withsemanticrelevance/{文本} | v1 第 1 版    | Search items with semantic relevance 搜索具有语义相关性的项目 |
| GET         | /api/catalog/items/withsemanticrelevance                     | v2 第 2 版    | Search items with semantic relevance (query param) 搜索具有语义相关性的项目 （query param） |
| GET         | /api/catalog/items/type/{typeId}/brand/{brandId}             | v1 第 1 版    | Get items by type and brand 按类型和品牌获取商品             |
| GET         | /api/catalog/items/type/all/brand/{brandId}                  | v1 第 1 版    | Get items by brand 按品牌获取商品                            |
| GET         | /api/catalog/catalogtypes                                    | v1, v2 v1、v2 | Get all catalog types 获取所有目录类型                       |
| GET         | /api/catalog/catalogbrands                                   | v1, v2 v1、v2 | Get all catalog brands 获取所有目录品牌                      |
| PUT         | /api/catalog/items                                           | v1 第 1 版    | Update item (ID in body) 更新项（正文中的 ID）               |
| PUT         | /api/catalog/items/{id}                                      | v2 第 2 版    | Update item (ID in path) 更新项（路径中的 ID）               |
| POST        | /api/catalog/items                                           | v1, v2 v1、v2 | Create new item 创建新项                                     |
| DELETE      | /api/catalog/items/{id}                                      | v1, v2 v1、v2 | Delete item by ID 按 ID 删除项目                             |

``` go
CatalogItem {
    id: integer
    name: string (required)
    description: string
    price: decimal
    pictureFileName: string
    catalogTypeId: integer
    catalogType: CatalogType
    catalogBrandId: integer
    catalogBrand: CatalogBrand
    availableStock: integer
    restockThreshold: integer
    maxStockThreshold: integer
    onReorder: boolean
    embedding: Vector (for AI search)
}

CatalogBrand {
    id: integer
    brand: string (required)
}

CatalogType {
    id: integer
    type: string (required)
}
```

### Vector Embeddings 向量嵌入

目录项与根据其名称和描述生成的矢量嵌入一起存储：

### Semantic Search 语义搜索
当用户使用语义相关性进行搜索时，API 会：

1. Generates an embedding for the search text
   为搜索文本生成嵌入
2. Computes cosine distance between the search embedding and all item embeddings
   计算搜索嵌入向量和所有项目嵌入向量之间的余弦距离
3. Returns items sorted by relevance (smallest distance first)
   返回按相关性排序的项目（最小距离优先）

### Basket API 

购物篮 API 管理 eShop 应用程序中的购物车作。它提供用于创建、检索、更新和删除用户购物篮以及启动结帐流程的功能

| Endpoint 端点              | HTTP Method HTTP 方法 | Description 描述                                     |
| -------------------------- | --------------------- | ---------------------------------------------------- |
| `/api/v1/basket/{buyerId}` | GET                   | Get basket for a specific buyer 为特定买家购买购物车 |
| `/api/v1/basket`           | POST                  | Update basket with items 使用商品更新购物篮          |
| `/api/v1/basket/{buyerId}` | DELETE                | Delete a basket 删除购物篮                           |
| `/api/v1/basket/checkout`  | POST                  | Process basket checkout 流程购物篮结账               |

``` go
BasketItem {
    productID int
    ProductName string
    UnitPrice int
    Quantity int
}

BasketCheckoutInfo : 结账所需信息，送货地址、付款详细信息
```

当用户将商品添加到购物车时

* 购物篮 API 将数据存储在Redis
* 获取购物篮时，调用 CataAPI 获取详细的产品信息

结账时

* BasketCheckoutEvent 发布到 EventBus
* 订单 API 订阅此事件并创建新订单
* 购物篮 API 在结账后删除购物篮

结账流程：

* 用户提交发货和付款信息
* BasketState.CheckoutAsync() 获取结账信息
* 调用 Ordering API 创建新订单
* 成功创建订单后，购物车被删除

### Ordering API

| Method 方法 | Endpoint 端点           | Description 描述                                             | Request Body 请求正文                      | Response 响应                                      |
| ----------- | ----------------------- | ------------------------------------------------------------ | ------------------------------------------ | -------------------------------------------------- |
| GET         | `/api/orders`           | Gets all orders for the authenticated user 获取已验证用户的所有订单 | -                                          | List of OrderSummary objects OrderSummary 对象列表 |
| GET         | `/api/orders/{orderId}` | Gets a specific order by ID 按 ID 获取特定订单               | -                                          | Order object Order 对象                            |
| GET         | `/api/orders/cardtypes` | Gets all supported card types 获取所有支持的卡类型           | -                                          | List of CardType objects CardType 对象列表         |
| POST        | `/api/orders`           | Creates a new order 创建新订单                               | CreateOrderRequest CreateOrderRequest 请求 | 200 OK 200 确定                                    |
| POST        | `/api/orders/draft`     | Creates a draft order from basket 从购物篮创建草稿订单       | CreateOrderDraftCommand                    | OrderDraftDTO                                      |
| PUT         | `/api/orders/cancel`    | Cancels an existing order 取消现有订单                       | Order number 订单号                        | 200 OK 200 确定                                    |
| PUT         | `/api/orders/ship`      | Ships an existing order 运送现有订单                         | Order number 订单号                        | 200 OK 200 确定                                    |

Order 实体包含所有与订单相关的信息，包括商品、送货地址和付款详细信息

### Identity API

它管理用户身份、处理身份验证流程，并颁发其他服务用于验证用户身份和授权的安全令牌。



# 实施

## catalog.api

### ① 用户功能（暴露给 Gateway）

| 功能         | 描述                     | 是否暴露到网关 |
| ------------ | ------------------------ | -------------- |
| 获取所有商品 | 展示商品列表             | ✅              |
| 查询商品详情 | 根据商品 ID 获取详细信息 | ✅              |
| 查询库存     | 查询某个商品当前库存     | ✅              |

| 功能             | 描述                             | 是否暴露到网关    |
| ---------------- | -------------------------------- | ----------------- |
| 添加商品         | 添加新商品                       | ✅（需鉴权）       |
| 修改商品         | 更新商品名称、描述、价格等信息   | ✅（需鉴权）       |
| 删除商品         | 删除商品                         | ✅（需鉴权）       |
| 扣减库存（内部） | 下单后调用，减少库存（RPC 调用） | ❌（RPC 内部调用） |
| 回滚库存（内部） | 订单取消后恢复库存               | ❌（RPC 内部调用） |

## basket.api

| 功能             | 描述                   |
| ---------------- | ---------------------- |
| 添加商品到购物车 | 用户将商品添加到购物车 |
| 获取购物车       | 查询用户当前购物车内容 |
| 修改购物车项数量 | 更新某一商品的数量     |
| 删除购物车项     | 删除购物车中的某一商品 |
| 清空购物车       | 一键清空购物车         |
| 支持 JWT 鉴权    | 所有接口默认开启鉴权   |

## ordering.api

| 功能          | 描述                                                 |
| ------------- | ---------------------------------------------------- |
| 创建订单      | 用户提交购物车后生成订单（异步通知 Order Processor） |
| 查询订单列表  | 用户查询自己的所有订单                               |
| 查询订单详情  | 用户查看某个订单的详细信息                           |
| 取消订单      | 用户主动取消订单                                     |
| 管理订单状态  | 后台异步更新订单状态（由 OrderProcessor 监听消息）   |
| 支持 JWT 鉴权 | 所有接口默认开启鉴权                                 |

用户提交订单 --> PENDING
支付成功 --> PAID
取消订单 --> CANCELLED
商家发货 --> SHIPPED

`placeOrder` 会将订单存入数据库后发送异步事件（通过 `go-queue`）给 `Order Processor`。

`updateOrderStatus` 由 `Order Processor` 或 `Payment Processor` 监听事件后调用。

所有订单都绑定 `userId`（从 JWT 中提取）。

## identity.api

| 功能                 | 描述                                |
| -------------------- | ----------------------------------- |
| 用户注册             | 用户通过邮箱和密码注册              |
| 用户登录             | 用户凭账号密码登录，获取 JWT token  |
| 获取当前用户信息     | 登录用户查看自己的信息              |
| 用户登出（可选）     | 清除 refresh token 或客户端本地状态 |
| 更新用户信息（可选） | 修改邮箱、昵称等                    |
| 支持 JWT 鉴权        | 除注册和登录接口外，均需要登录鉴权  |

`goctl api go -api eshop.api -dir . --style=goZero --jwt --jwt-auth "makabaka!123#"`

购物篮成功创建订单后，rpc 删除订单

订单支付成功事件：catalog 订阅 用于减少内存，order订阅用于 减少库存 rpc

取消支付成功的订单 回滚内存 rpc