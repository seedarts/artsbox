<template>
  <div class="settings-panel">
    <h2>设置</h2>

    <div class="settings-form">
      <div class="form-group">
        <label>API Key</label>
        <input 
          v-model="settings.api_key" 
          type="password"
          placeholder="请输入您的 API Key"
          class="form-input"
        />
        <span class="form-hint">用于调用 AI 图像生成服务</span>
      </div>

      <div class="form-group">
        <label>API 地址</label>
        <input 
          v-model="settings.base_url" 
          type="text"
          placeholder="https://api.cozex.cn/v1"
          class="form-input"
        />
        <span class="form-hint">兼容 OpenAI 格式的 API 服务地址</span>
      </div>

      <div class="form-group">
        <label>默认模型</label>
        <select v-model="settings.model" class="form-select">
          <option value="doubao-seedream-4-0-250828">豆包 Seedream 4.0</option>
          <option value="doubao-seedream-4-5-251128">豆包 Seedream 4.5</option>
        </select>
      </div>

      <div class="form-group">
        <label>保存目录</label>
        <div class="dir-input-group">
          <input 
            v-model="settings.save_dir" 
            type="text"
            readonly
            class="form-input"
          />
          <button @click="selectDir" class="dir-select-btn">选择</button>
        </div>
      </div>

      <div class="form-actions">
        <button @click="handleSave" class="save-btn">
          💾 保存配置
        </button>
        <button @click="handleReset" class="reset-btn">
          🔄 恢复默认
        </button>
      </div>

      <div v-if="message" :class="['message', messageType]">
        {{ message }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { GetSettings, SaveSettings, SelectSaveDirectory } from '../../wailsjs/go/main/App';

const settings = ref({
  api_key: '',
  base_url: '',
  model: '',
  save_dir: ''
});

const message = ref('');
const messageType = ref('');

onMounted(async () => {
  await loadSettings();
});

const loadSettings = async () => {
  try {
    const data = await GetSettings();
    settings.value = { ...data };
  } catch (error) {
    showMessage('加载配置失败: ' + error, 'error');
  }
};

const handleSave = async () => {
  try {
    await SaveSettings(settings.value);
    showMessage('配置保存成功！', 'success');
  } catch (error) {
    showMessage('保存失败: ' + error, 'error');
  }
};

const handleReset = async () => {
  if (confirm('确定要恢复默认配置吗？')) {
    settings.value = {
      api_key: '',
      base_url: 'https://api.cozex.cn/v1',
      model: 'doubao-seedream-4-5-251128',
      save_dir: ''
    };
    await handleSave();
  }
};

const selectDir = async () => {
  try {
    const dir = await SelectSaveDirectory();
    if (dir) {
      settings.value.save_dir = dir;
    }
  } catch (error) {
    showMessage('选择目录失败: ' + error, 'error');
  }
};

const showMessage = (msg, type) => {
  message.value = msg;
  messageType.value = type;
  setTimeout(() => {
    message.value = '';
  }, 3000);
};
</script>

<style scoped>
.settings-panel {
  flex: 1;
  padding: 40px;
  overflow-y: auto;
}

.settings-panel h2 {
  margin: 0 0 32px 0;
  font-size: 24px;
  color: var(--text-primary);
}

.settings-form {
  max-width: 600px;
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-group label {
  color: var(--text-primary);
  font-weight: 500;
  font-size: 14px;
}

.form-input,
.form-select {
  padding: 12px 16px;
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-primary);
  font-size: 14px;
  outline: none;
  transition: border-color 0.3s;
}

.form-input:focus,
.form-select:focus {
  border-color: var(--accent-color);
}

.form-hint {
  font-size: 12px;
  color: var(--text-secondary);
}

.dir-input-group {
  display: flex;
  gap: 8px;
}

.dir-input-group .form-input {
  flex: 1;
}

.dir-select-btn {
  padding: 12px 20px;
  background: var(--bg-tertiary);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  color: var(--text-primary);
  cursor: pointer;
  transition: all 0.2s;
}

.dir-select-btn:hover {
  background: var(--accent-color);
  border-color: var(--accent-color);
}

.form-actions {
  display: flex;
  gap: 12px;
  margin-top: 12px;
}

.save-btn,
.reset-btn {
  flex: 1;
  padding: 14px 24px;
  border: none;
  border-radius: 8px;
  font-size: 15px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.save-btn {
  background: linear-gradient(135deg, var(--accent-color), var(--accent-hover));
  color: var(--text-primary);
}

.save-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 20px rgba(99, 102, 241, 0.3);
}

.reset-btn {
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  color: var(--text-primary);
}

.reset-btn:hover {
  background: var(--bg-tertiary);
}

.message {
  padding: 12px 16px;
  border-radius: 8px;
  font-size: 14px;
}

.message.success {
  background: rgba(16, 185, 129, 0.1);
  border: 1px solid var(--success-color);
  color: var(--success-color);
}

.message.error {
  background: rgba(239, 68, 68, 0.1);
  border: 1px solid var(--error-color);
  color: var(--error-color);
}
</style>
