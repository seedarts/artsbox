<template>
  <div class="image-gallery">
    <div class="gallery-header">
      <h3>生成结果</h3>
      <div class="save-dir">
        <span>保存至: {{ saveDir }}</span>
        <button @click="selectDirectory" class="dir-btn">📁</button>
        <button @click="refreshImages" class="dir-btn" title="刷新">🔄</button>
      </div>
    </div>

    <!-- 加载中提示 -->
    <div v-if="isLoading" class="loading-state">
      <div class="loading-spinner"></div>
      <p>加载中...</p>
    </div>

    <!-- 空状态 -->
    <div v-else-if="savedImages.length === 0" class="empty-state">
      <div class="empty-icon">🖼️</div>
      <p>暂无图片</p>
      <p class="empty-hint">开始创作您的第一张 AI 图片吧！</p>
    </div>

    <!-- 图片网格 -->
    <div v-else class="gallery-grid">
      <div 
        v-for="(image, index) in savedImages" 
        :key="index"
        class="image-card"
        @click="openImage(image)"
      >
        <img :src="getImageUrl(image)" :alt="'Image ' + (index + 1)" @error="handleImageError" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { SelectSaveDirectory, GetSettings, GetSavedImages, OpenImageFile, GetImageAsDataURL } from '../../wailsjs/go/main/App';

const props = defineProps({
  images: {
    type: Array,
    default: () => []
  }
});

const saveDir = ref('');
const savedImages = ref([]); // 保存目录中的所有图片路径
const imageDataUrls = ref({}); // 图片的 base64 Data URLs
const isLoading = ref(false);

// 加载保存目录中的图片
const loadSavedImages = async () => {
  isLoading.value = true;
  try {
    const images = await GetSavedImages();
    savedImages.value = images;
    
    // 为每张图片加载 base64 数据
    imageDataUrls.value = {};
    for (const imagePath of images) {
      try {
        const dataUrl = await GetImageAsDataURL(imagePath);
        if (dataUrl) {
          imageDataUrls.value[imagePath] = dataUrl;
        }
      } catch (error) {
        console.error('加载图片失败:', imagePath, error);
      }
    }
  } catch (error) {
    console.error('加载图片失败:', error);
    savedImages.value = [];
  } finally {
    isLoading.value = false;
  }
};

onMounted(async () => {
  const settings = await GetSettings();
  saveDir.value = settings.save_dir;
  await loadSavedImages();
});

// 监听生成结果变化，重新加载图片（因为生成后会自动保存）
watch(() => props.images, async (newImages) => {
  if (newImages && newImages.length > 0) {
    // 延迟 1 秒再加载，等待文件保存完成
    setTimeout(async () => {
      await loadSavedImages();
    }, 1000);
  }
});

// 获取图片 URL
const getImageUrl = (path) => {
  return imageDataUrls.value[path] || '';
};

// 打开图片文件
const openImage = async (imagePath) => {
  try {
    await OpenImageFile(imagePath);
  } catch (error) {
    console.error('打开图片失败:', error);
  }
};

const selectDirectory = async () => {
  try {
    const dir = await SelectSaveDirectory();
    if (dir) {
      saveDir.value = dir;
      // 更改目录后重新加载图片
      await loadSavedImages();
    }
  } catch (error) {
    console.error('选择目录失败:', error);
  }
};

// 刷新图片列表
const refreshImages = async () => {
  await loadSavedImages();
};

// 处理图片加载错误
const handleImageError = (e) => {
  console.error('图片加载失败:', e.target.src);
  e.target.style.display = 'none';
};
</script>

<style scoped>
.image-gallery {
  width: 400px;
  height: 100%;
  background: var(--bg-secondary);
  border-left: 1px solid var(--border-color);
  display: flex;
  flex-direction: column;
}

.gallery-header {
  padding: 24px;
  border-bottom: 1px solid var(--border-color);
}

.gallery-header h3 {
  margin: 0 0 12px 0;
  font-size: 18px;
  color: var(--text-primary);
}

.save-dir {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 12px;
  color: var(--text-secondary);
}

.dir-btn {
  padding: 4px 8px;
  background: var(--bg-tertiary);
  border: 1px solid var(--border-color);
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.2s;
}

.dir-btn:hover {
  background: var(--bg-primary);
}

.empty-state {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px;
  text-align: center;
}

.empty-icon {
  font-size: 64px;
  margin-bottom: 16px;
  opacity: 0.3;
}

.empty-state p {
  margin: 8px 0;
  color: var(--text-secondary);
}

.empty-hint {
  font-size: 13px;
}

/* 加载状态 */
.loading-state {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px;
  text-align: center;
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid var(--bg-tertiary);
  border-top-color: var(--accent-color);
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 16px;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.loading-state p {
  margin: 0;
  color: var(--text-secondary);
}

.gallery-grid {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  gap: 12px;
  align-content: start;
}

.image-card {
  position: relative;
  aspect-ratio: 1;
  background: var(--bg-tertiary);
  border-radius: 8px;
  overflow: hidden;
  border: 2px solid var(--border-color);
  transition: all 0.2s;
  cursor: pointer;
}

.image-card:hover {
  transform: scale(1.08);
  border-color: var(--accent-color);
  box-shadow: 0 4px 16px rgba(99, 102, 241, 0.4);
  z-index: 10;
}

.image-card img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  display: block;
}
</style>
