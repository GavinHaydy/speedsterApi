merchant.api
```text


Merchant                  // 商家管理
├── CreateMerchant        // 创建商家
├── UpdateMerchant        // 修改商家
├── DeleteMerchant        // 删除商家（逻辑删除）
├── GetMerchant           // 获取商家详情
├── ListMerchant          // 获取商家分页列表
├── UpdateMerchantStatus  // 修改商家状态（营业/停业）
├── AuditMerchant         // 审核商家（通过/拒绝）
└── GetMerchantStatistics // 获取商家统计数据

MerchantStore             // 门店管理
├── CreateMerchantStore   // 创建门店
├── UpdateMerchantStore   // 修改门店
├── DeleteMerchantStore   // 删除门店
├── GetMerchantStore      // 获取门店详情
├── ListMerchantStore     // 获取门店分页列表
└── UpdateStoreStatus     // 修改门店营业状态

BusinessTime              // 营业时间管理
├── GetBusinessTime       // 获取门店营业时间
└── SaveBusinessTime      // 保存门店营业时间

Settlement                // 商家结算管理
├── GetSettlement         // 获取商家结算配置
└── UpdateSettlement      // 修改商家结算配置

OperateLog                // 商家操作日志
└── ListOperateLog        // 获取商家操作日志列表

MerchantApply
├── CreateMerchantApply      提交入驻申请
├── GetMyMerchantApply       获取我的申请
├── UpdateMerchantApply      修改申请资料
├── CancelMerchantApply      取消申请 (下方为admin功能)
├── ListMerchantApply        入驻申请列表
├── GetMerchantApply         申请详情
├── ApproveMerchantApply     审核通过
└── RejectMerchantApply      审核拒绝


用户注册
│
▼
CreateMerchantApply
│
▼
merchant_apply(status=0)
│
▼
后台审核
│
├── 拒绝 → status=2
│
└── 通过
│
├── 创建 merchant
├── 创建 merchant_store
├── 创建 merchant_settlement
└── merchant_apply.status=1
```