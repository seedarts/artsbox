<template>
  <div class="sidebar">
    <div class="logo">
      <div class="logo-icon">🎨</div>
      <div class="logo-text">ArtsBox</div>
    </div>

    <div class="nav-items">
      <div 
        class="nav-item active"
        @click="$emit('navigate', 'generate')"
      >
        <div class="nav-icon">✨</div>
        <span>图片生成</span>
      </div>
    </div>

    <div class="nav-bottom">
      <div 
        class="nav-item"
        @click="$emit('navigate', 'settings')"
      >
        <div class="nav-icon">⚙️</div>
        <span>设置</span>
      </div>
    </div>
  </div>
</template>

<script setup>
defineEmits(['navigate']);
</script>

<style scoped>
.sidebar {
  width: 80px;
  height: 100%;
  background: var(--bg-secondary);
  border-right: 1px solid var(--border-color);
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px 0;
}

.logo {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  margin-bottom: 40px;
}

.logo-icon {
  width: 48px;
  height: 48px;
  font-size: 28px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 12px;
  background: linear-gradient(135deg, var(--accent-color), var(--accent-hover));
}

.logo-text {
  font-size: 12px;
  font-weight: 600;
  color: var(--text-primary);
}

.nav-items {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.nav-bottom {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.nav-item {
  width: 60px;
  height: 60px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 4px;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s;
  font-size: 11px;
  color: var(--text-secondary);
}

.nav-item:hover {
  background: var(--bg-tertiary);
  color: var(--text-primary);
}

.nav-item.active {
  background: var(--accent-color);
  color: var(--text-primary);
}

.nav-icon {
  font-size: 24px;
}
</style>
