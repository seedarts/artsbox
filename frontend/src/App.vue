<script setup>
import { ref } from 'vue';
import Sidebar from './components/Sidebar.vue';
import GeneratePanel from './components/GeneratePanel.vue';
import ImageGallery from './components/ImageGallery.vue';
import SettingsPanel from './components/SettingsPanel.vue';

const currentView = ref('generate');
const generatedImages = ref([]);

const handleNavigate = (view) => {
  currentView.value = view;
};

const handleGenerated = (images) => {
  generatedImages.value = images;
};
</script>

<template>
  <div class="app-container">
    <Sidebar @navigate="handleNavigate" />
    
    <div class="main-content">
      <GeneratePanel 
        v-if="currentView === 'generate'"
        @generated="handleGenerated"
      />
      <SettingsPanel v-else-if="currentView === 'settings'" />
    </div>

    <ImageGallery 
      v-if="currentView === 'generate'"
      :images="generatedImages" 
    />
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  height: 100vh;
  width: 100vw;
  overflow: hidden;
}

.main-content {
  flex: 1;
  display: flex;
  overflow: hidden;
}
</style>
