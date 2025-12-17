<template>
  <div class="generate-panel">
    <!-- 模式切换 -->
    <div class="mode-tabs">
      <div 
        v-for="mode in modes" 
        :key="mode.value"
        class="mode-tab"
        :class="{ active: currentMode === mode.value }"
        @click="switchMode(mode.value)"
      >
        <div class="mode-label">{{ mode.label }}</div>
        <div class="mode-desc">{{ mode.desc }}</div>
      </div>
    </div>

    <!-- 组图子模式切换 -->
    <div v-if="currentMode === 'batch'" class="batch-submodes">
      <div 
        v-for="subMode in batchSubModes" 
        :key="subMode.value"
        class="batch-submode"
        :class="{ active: batchSubMode === subMode.value }"
        @click="switchBatchSubMode(subMode.value)"
      >
        {{ subMode.label }}
      </div>
    </div>

    <!-- 图片选择区域 -->
    <div v-if="needImageUpload" class="image-upload-section">
      <div class="upload-header">
        <h3>{{ allowMultipleImages ? '选择多张图片' : '选择参考图片' }}</h3>
        <span class="upload-hint">
          {{ allowMultipleImages ? '最多选择 10 张图片' : '支持 PNG, JPG, JPEG 格式' }}
        </span>
      </div>
      
      <div class="upload-area">
        <div v-if="selectedImages.length === 0" class="upload-placeholder" @click="selectImages">
          <div class="upload-icon">📁</div>
          <p>点击选择图片</p>
        </div>
        
        <div v-else class="selected-images">
          <div 
            v-for="(img, index) in selectedImages" 
            :key="index"
            class="image-preview"
          >
            <img :src="img.preview" :alt="img.name" />
            <button class="remove-btn" @click="removeImage(index)">×</button>
            <div class="image-name">{{ img.name }}</div>
          </div>
          
          <div 
            v-if="allowMultipleImages && selectedImages.length < 10" 
            class="add-more"
            @click="selectImages"
          >
            <div class="add-icon">+</div>
            <p>添加更多</p>
          </div>
        </div>
      </div>
    </div>

    <!-- 提示词输入 -->
    <div class="prompt-section">
      <textarea
        v-model="prompt"
        class="prompt-input"
        placeholder="输入您的创意描述，例如：一只戴着贝雷帽的可爱海獭宝宝..."
        rows="4"
      ></textarea>

      <div class="prompt-footer">
        <select v-model="selectedModel" class="model-select">
          <option value="doubao-seedream-4-0-250828">豆包 Seedream 4.0</option>
          <option value="doubao-seedream-4-5-251128">豆包 Seedream 4.5</option>
        </select>

        <button class="generate-btn" @click="handleGenerate" :disabled="isGenerating">
          <span v-if="!isGenerating">✨ 开始生成</span>
          <span v-else>⏳ 生成中...</span>
        </button>
      </div>
    </div>

    <!-- 参数设置 -->
    <div class="params-section">
      <!-- 图片质量选择 -->
      <div class="param-group">
        <label>图片质量</label>
        <div class="quality-options">
          <div 
            v-for="size in sizes" 
            :key="size.value"
            class="quality-option"
            :class="{ active: selectedSize === size.value }"
            @click="selectedSize = size.value"
          >
            {{ size.label }}
          </div>
        </div>
      </div>

      <!-- 组图数量设置 (仅组图输出模式显示) -->
      <div v-if="showBatchCount" class="param-group">
        <label>生成数量 (1-4张)</label>
        <div class="count-slider">
          <input 
            type="range" 
            v-model.number="maxImages" 
            min="1" 
            max="4" 
            step="1"
            class="slider"
          />
          <span class="count-value">{{ maxImages }} 张</span>
        </div>
      </div>
    </div>

    <!-- 错误/成功提示 -->
    <div v-if="errorMessage" :class="['message', errorMessage.startsWith('✅') ? 'success-message' : 'error-message']">
      {{ errorMessage }}
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue';
import { 
  GenerateImage, 
  GenerateBatchImages,
  GenerateFromImage, 
  GenerateFromMultiImages,
  GenerateBatchFromImage,
  GenerateBatchFromMultiImages,
  SaveFile,
  GetSettings,
  SaveSettings
} from '../../wailsjs/go/main/App';

const currentMode = ref('text'); // text, image, multi, batch
const batchSubMode = ref('text'); // text, image, multi (用于组图输出子模式)
const prompt = ref('');
const selectedModel = ref('doubao-seedream-4-5-251128');
const selectedSize = ref('2K');
const maxImages = ref(4); // 组图数量 (1-4)
const isGenerating = ref(false);
const errorMessage = ref('');
const selectedImages = ref([]);

const emit = defineEmits(['generated']);

// 加载配置
onMounted(async () => {
  try {
    const settings = await GetSettings();
    if (settings.model) {
      selectedModel.value = settings.model;
    }
  } catch (error) {
    console.error('加载配置失败:', error);
  }
});

// 监听模型选择，自动保存到配置
watch(selectedModel, async (newModel) => {
  try {
    const settings = await GetSettings();
    settings.model = newModel;
    await SaveSettings(settings);
    console.log('模型已更新:', newModel);
  } catch (error) {
    console.error('保存模型失败:', error);
  }
});

const modes = [
  { label: '文生图', value: 'text', desc: '纯文本→单图' },
  { label: '图文生图', value: 'image', desc: '单图+文本→单图' },
  { label: '多图融合', value: 'multi', desc: '多图+文本→单图' },
  { label: '组图输出', value: 'batch', desc: '批量生成1-4张' },
];

const batchSubModes = [
  { label: '文生组图', value: 'text' },
  { label: '单图生组图', value: 'image' },
  { label: '多图生组图', value: 'multi' },
];

const sizes = [
  { label: '2K', value: '2K' },
  { label: '4K', value: '4K' },
];

// 是否需要显示图片上传区
const needImageUpload = computed(() => {
  if (currentMode.value === 'image' || currentMode.value === 'multi') {
    return true;
  }
  if (currentMode.value === 'batch') {
    return batchSubMode.value !== 'text';
  }
  return false;
});

// 是否允许多图上传
const allowMultipleImages = computed(() => {
  if (currentMode.value === 'multi') return true;
  if (currentMode.value === 'batch' && batchSubMode.value === 'multi') return true;
  return false;
});

// 是否显示组图数量设置
const showBatchCount = computed(() => {
  return currentMode.value === 'batch';
});

const switchMode = (mode) => {
  currentMode.value = mode;
  selectedImages.value = [];
  errorMessage.value = '';
  if (mode === 'batch') {
    batchSubMode.value = 'text';
  }
};

const switchBatchSubMode = (subMode) => {
  batchSubMode.value = subMode;
  selectedImages.value = [];
  errorMessage.value = '';
};

const selectImages = () => {
  const input = document.createElement('input');
  input.type = 'file';
  input.accept = 'image/png,image/jpeg,image/jpg';
  input.multiple = allowMultipleImages.value;
  
  input.onchange = (e) => {
    const files = Array.from(e.target.files);
    
    if (!allowMultipleImages.value) {
      // 单图模式：只选择一张
      if (files.length > 0) {
        processFiles([files[0]]);
      }
    } else {
      // 多图模式：最多10张
      const remainingSlots = 10 - selectedImages.value.length;
      const filesToAdd = files.slice(0, remainingSlots);
      processFiles(filesToAdd);
      
      if (files.length > remainingSlots) {
        errorMessage.value = `最多只能选择 10 张图片，已自动限制为 ${remainingSlots} 张`;
        setTimeout(() => errorMessage.value = '', 3000);
      }
    }
  };
  
  input.click();
};

const processFiles = (files) => {
  if (!allowMultipleImages.value) {
    selectedImages.value = [];
  }
  
  files.forEach(file => {
    const reader = new FileReader();
    reader.onload = (e) => {
      selectedImages.value.push({
        file: file,
        name: file.name,
        preview: e.target.result,
        path: file.path || ''
      });
    };
    reader.readAsDataURL(file);
  });
};

const removeImage = (index) => {
  selectedImages.value.splice(index, 1);
};

const handleGenerate = async () => {
  if (!prompt.value.trim()) {
    errorMessage.value = '请输入提示词';
    return;
  }

  if (needImageUpload.value && selectedImages.value.length === 0) {
    errorMessage.value = '请先选择图片';
    return;
  }

  isGenerating.value = true;
  errorMessage.value = '';

  try {
    let result;
    
    // 根据模式调用不同的 API
    if (currentMode.value === 'text') {
      // 文生图（单图）
      result = await GenerateImage(prompt.value, selectedSize.value, 1);
    } else if (currentMode.value === 'image') {
      // 图文生图（单图→单图）
      const imageBase64 = selectedImages.value[0].preview;
      result = await GenerateFromImage(prompt.value, imageBase64);
    } else if (currentMode.value === 'multi') {
      // 多图融合（多图→单图）
      const imageBase64List = selectedImages.value.map(img => img.preview);
      result = await GenerateFromMultiImages(prompt.value, imageBase64List);
    } else if (currentMode.value === 'batch') {
      // 组图输出
      if (batchSubMode.value === 'text') {
        // 文生组图
        result = await GenerateBatchImages(prompt.value, selectedSize.value, maxImages.value);
      } else if (batchSubMode.value === 'image') {
        // 单图生组图
        const imageBase64 = selectedImages.value[0].preview;
        result = await GenerateBatchFromImage(prompt.value, imageBase64, maxImages.value);
      } else {
        // 多图生组图
        const imageBase64List = selectedImages.value.map(img => img.preview);
        result = await GenerateBatchFromMultiImages(prompt.value, imageBase64List, maxImages.value);
      }
    }

    if (result.success) {
      // 生成成功后，自动保存所有图片
      const savedPaths = [];
      const timestamp = Date.now();
      
      for (let i = 0; i < result.images.length; i++) {
        const imageUrl = result.images[i];
        const filename = `image_${timestamp}_${i + 1}.png`;
        try {
          const savePath = await SaveFile(imageUrl, filename);
          if (!savePath.startsWith('保存失败')) {
            savedPaths.push(savePath);
          }
        } catch (err) {
          console.error(`保存图片 ${i + 1} 失败:`, err);
        }
        // 每保存一张图片间隔100ms，避免文件名冲突
        if (i < result.images.length - 1) {
          await new Promise(resolve => setTimeout(resolve, 100));
        }
      }
      
      // 显示保存成功提示
      if (savedPaths.length > 0) {
        errorMessage.value = `✅ 生成成功！已自动保存 ${savedPaths.length} 张图片`;
        setTimeout(() => errorMessage.value = '', 5000);
      }
      
      emit('generated', result.images);
    } else {
      errorMessage.value = result.message;
    }
  } catch (error) {
    errorMessage.value = '生成失败: ' + error;
  } finally {
    isGenerating.value = false;
  }
};
</script>

<style scoped>
.generate-panel {
  flex: 1;
  padding: 30px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.mode-tabs {
  display: flex;
  gap: 12px;
  background: var(--bg-secondary);
  padding: 6px;
  border-radius: 12px;
}

.mode-tab {
  flex: 1;
  padding: 16px 20px;
  text-align: center;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s;
  color: var(--text-secondary);
  font-weight: 500;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.mode-tab:hover {
  background: var(--bg-tertiary);
  color: var(--text-primary);
}

.mode-tab.active {
  background: var(--accent-color);
  color: var(--text-primary);
}

.mode-label {
  font-size: 15px;
  font-weight: 600;
}

.mode-desc {
  font-size: 11px;
  opacity: 0.8;
}

.mode-tab.active .mode-desc {
  opacity: 1;
}

/* 组图子模式切换 */
.batch-submodes {
  display: flex;
  gap: 8px;
  background: var(--bg-secondary);
  padding: 4px;
  border-radius: 8px;
}

.batch-submode {
  flex: 1;
  padding: 10px 16px;
  text-align: center;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.3s;
  color: var(--text-secondary);
  font-size: 13px;
  font-weight: 500;
}

.batch-submode:hover {
  background: var(--bg-tertiary);
  color: var(--text-primary);
}

.batch-submode.active {
  background: linear-gradient(135deg, rgba(99, 102, 241, 0.2), rgba(168, 85, 247, 0.2));
  color: var(--text-primary);
  border: 1px solid var(--accent-color);
}

.prompt-section {
  background: var(--bg-secondary);
  border-radius: 16px;
  padding: 20px;
  border: 1px solid var(--border-color);
}

.prompt-input {
  width: 100%;
  background: var(--bg-tertiary);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 16px;
  color: var(--text-primary);
  font-size: 15px;
  line-height: 1.6;
  resize: none;
  font-family: inherit;
  outline: none;
  transition: border-color 0.3s;
}

.prompt-input:focus {
  border-color: var(--accent-color);
}

.prompt-input::placeholder {
  color: var(--text-secondary);
}

.prompt-footer {
  display: flex;
  gap: 12px;
  margin-top: 16px;
}

.model-select {
  flex: 1;
  padding: 12px 16px;
  background: var(--bg-tertiary);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-primary);
  font-size: 14px;
  cursor: pointer;
  outline: none;
}

.generate-btn {
  padding: 12px 32px;
  background: linear-gradient(135deg, var(--accent-color), var(--accent-hover));
  border: none;
  border-radius: 8px;
  color: var(--text-primary);
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.3s;
}

.generate-btn:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(99, 102, 241, 0.3);
}

.generate-btn:active:not(:disabled) {
  transform: translateY(0);
}

.generate-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.params-section {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.param-group label {
  display: block;
  margin-bottom: 12px;
  color: var(--text-primary);
  font-weight: 500;
  font-size: 14px;
}

/* 质量选项 */
.quality-options {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
}

.quality-option {
  padding: 14px 20px;
  background: var(--bg-secondary);
  border: 2px solid var(--border-color);
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s;
  text-align: center;
  font-size: 15px;
  font-weight: 600;
  color: var(--text-secondary);
}

.quality-option:hover {
  border-color: var(--accent-color);
  background: var(--bg-tertiary);
  color: var(--text-primary);
}

.quality-option.active {
  border-color: var(--accent-color);
  background: var(--accent-color);
  color: var(--text-primary);
}

/* 数量滑块 */
.count-slider {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px 16px;
  background: var(--bg-secondary);
  border-radius: 12px;
  border: 1px solid var(--border-color);
}

.slider {
  flex: 1;
  height: 6px;
  -webkit-appearance: none;
  appearance: none;
  background: var(--bg-tertiary);
  border-radius: 3px;
  outline: none;
}

.slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 18px;
  height: 18px;
  background: var(--accent-color);
  border-radius: 50%;
  cursor: pointer;
  transition: all 0.2s;
}

.slider::-webkit-slider-thumb:hover {
  transform: scale(1.2);
  box-shadow: 0 0 0 4px rgba(99, 102, 241, 0.2);
}

.slider::-moz-range-thumb {
  width: 18px;
  height: 18px;
  background: var(--accent-color);
  border-radius: 50%;
  cursor: pointer;
  border: none;
  transition: all 0.2s;
}

.slider::-moz-range-thumb:hover {
  transform: scale(1.2);
  box-shadow: 0 0 0 4px rgba(99, 102, 241, 0.2);
}

.count-value {
  min-width: 50px;
  text-align: right;
  font-weight: 600;
  color: var(--accent-color);
  font-size: 15px;
}

.message {
  padding: 12px 16px;
  border-radius: 8px;
  font-size: 14px;
}

.success-message {
  background: rgba(16, 185, 129, 0.1);
  border: 1px solid var(--success-color);
  color: var(--success-color);
}

.error-message {
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid var(--error-color);
  color: var(--error-color);
}

.error-message::before {
  content: '⚠️ ';
}

/* 图片上传区域 */
.image-upload-section {
  background: var(--bg-secondary);
  border-radius: 16px;
  padding: 20px;
  border: 1px solid var(--border-color);
}

.upload-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.upload-header h3 {
  margin: 0;
  font-size: 16px;
  color: var(--text-primary);
}

.upload-hint {
  font-size: 12px;
  color: var(--text-secondary);
}

.upload-area {
  min-height: 200px;
}

.upload-placeholder {
  height: 200px;
  border: 2px dashed var(--border-color);
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s;
}

.upload-placeholder:hover {
  border-color: var(--accent-color);
  background: var(--bg-tertiary);
}

.upload-icon {
  font-size: 48px;
  margin-bottom: 12px;
  opacity: 0.5;
}

.upload-placeholder p {
  color: var(--text-secondary);
  margin: 0;
}

.selected-images {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 12px;
}

.image-preview {
  position: relative;
  aspect-ratio: 1;
  border-radius: 12px;
  overflow: hidden;
  background: var(--bg-tertiary);
  border: 2px solid var(--border-color);
}

.image-preview img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.remove-btn {
  position: absolute;
  top: 8px;
  right: 8px;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: rgba(239, 68, 68, 0.9);
  border: none;
  color: white;
  font-size: 20px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
  line-height: 1;
}

.remove-btn:hover {
  background: var(--error-color);
  transform: scale(1.1);
}

.image-name {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 8px;
  background: rgba(0, 0, 0, 0.7);
  color: white;
  font-size: 11px;
  text-align: center;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.add-more {
  aspect-ratio: 1;
  border: 2px dashed var(--border-color);
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s;
  background: var(--bg-tertiary);
}

.add-more:hover {
  border-color: var(--accent-color);
  background: var(--bg-secondary);
}

.add-icon {
  font-size: 36px;
  color: var(--text-secondary);
  margin-bottom: 8px;
}

.add-more p {
  margin: 0;
  font-size: 12px;
  color: var(--text-secondary);
}
</style>
