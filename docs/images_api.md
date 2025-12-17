## 图片生成&编辑 API 实现参考

**说明**: 使用豆包 Seedream API 服务。

- **基础 URL**: `https://api.cozex.cn/v1`
- **端点**: `/images/generations`
- **认证**: Bearer Token (ARK_API_KEY)

### 1 文生图 (Text-to-Image) 协议

后端需构造如下 HTTP 请求:

```bash
curl https://api.cozex.cn/v1/images/generations \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ARK_API_KEY" \
  -d '{
    "model": "doubao-seedream-4-5-251128",
    "prompt": "A cute baby sea otter",
    "size": "2K",
    "stream": false,
    "response_format": "url",
    "watermark": false
  }'
```

### 2 图文生图 (Image+Text-to-Image) 协议

后端需构造如下 HTTP 请求 (JSON 格式):

```bash
curl https://api.cozex.cn/v1/images/generations \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ARK_API_KEY" \
  -d '{
    "model": "doubao-seedream-4-5-251128",
    "prompt": "保持模特姿势和液态服装的流动形状不变。将服装材质从银色金属改为完全透明的清水。",
    "image": "data:image/png;base64,iVBORw0KGgoAAAA...",
    "size": "2K",
    "stream": false,
    "response_format": "url",
    "watermark": false
  }'
```

**注意**:

- `image` 字段支持图片 URL 或 base64 编码的图片数据

### 3 多图融合 (Multi-Image-Fusion) 协议

后端需构造如下 HTTP 请求 (JSON 格式):

```bash
curl https://api.cozex.cn/v1/images/generations \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ARK_API_KEY" \
  -d '{
    "model": "doubao-seedream-4-5-251128",
    "prompt": "融合所有图片的风格，创造出一幅新的艺术作品",
    "image": [
      "https://example.com/image1.png",
      "https://example.com/image2.png",
      "https://example.com/image3.png"
    ],
    "size": "2K",
    "stream": false,
    "response_format": "url",
    "watermark": false
  }'
```

**注意**:

- `image` 字段为数组，最多包含 10 张图片
- 每张图片支持 URL 或 base64 格式
- 不使用 `sequential_image_generation` 参数，默认输出单张融合图

### 4 组图输出 (Batch-Generation) 协议

#### 4.1 文生组图

```bash
curl https://api.cozex.cn/v1/images/generations \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ARK_API_KEY" \
  -d '{
    "model": "doubao-seedream-4-5-251128",
    "prompt": "生成一组共4张连贯插画，核心为同一庭院一角的四季变迁，以统一风格展现四季独特色彩、元素与氛围",
    "size": "2K",
    "sequential_image_generation": "auto",
    "sequential_image_generation_options": {
      "max_images": 4
    },
    "stream": false,
    "response_format": "url",
    "watermark": false
  }'
```

**注意**:

- `sequential_image_generation`: 设置为 "auto" 启用组图生成
- `sequential_image_generation_options.max_images`: 设置生成数量 (1-4)

### 4.2 单图生组图

```bash
curl https://api.cozex.cn/v1/images/generations \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ARK_API_KEY" \
  -d '{
    "model": "doubao-seedream-4-5-251128",
    "prompt": "参考这个LOGO，做一套户外运动品牌视觉设计，品牌名称为"GREEN"，包括包装袋、帽子、卡片、挂绳等。绿色视觉主色调，趣味、简约现代风格",
    "image": "https://ark-project.tos-cn-beijing.volces.com/doc_image/seedream4_imageToimages.png",
    "size": "2K",
    "sequential_image_generation": "auto",
    "sequential_image_generation_options": {
      "max_images": 4
    },
    "stream": false,
    "response_format": "url",
    "watermark": false
  }'
```

#### 4.3 多图生组图

```bash
curl https://api.cozex.cn/v1/images/generations \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ARK_API_KEY" \
  -d '{
    "model": "doubao-seedream-4-5-251128",
    "prompt": "生成3张女孩和奶牛玩偶在游乐园开心地坐过山车的图片，涵盖早晨、中午、晚上",
    "image": [
      "https://ark-project.tos-cn-beijing.volces.com/doc_image/seedream4_imagesToimages_1.png",
      "https://ark-project.tos-cn-beijing.volces.com/doc_image/seedream4_imagesToimages_2.png"
    ],
    "sequential_image_generation": "auto",
    "sequential_image_generation_options": {
      "max_images": 3
    },
    "size": "2K",
    "watermark": false
  }'
```

**注意**:

- 所有组图模式使用 `sequential_image_generation` 参数启用
- `sequential_image_generation_options.max_images` 设置生成数量 (1-4)
- 所有模式使用统一的 `/images/generations` 端点
- **单图输入**: `image` 为字符串 (URL 或 base64)
- **多图输入**: `image` 为数组 (URL 或 base64 列表)
- 通过 `image` 字段类型和 `sequential_image_generation` 参数的不同组合实现不同功能

## 支持的模型名称

- doubao-seedream-4-0-250828
- doubao-seedream-4-5-251128
