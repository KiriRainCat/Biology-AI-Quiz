<script setup lang="ts">
import { RouterView } from "vue-router";
</script>

<template>
  <div class="bg-gray-50 h-screen">
    <router-view v-slot="{ Component }">
      <transition name="fade" mode="out-in">
        <component :is="Component" />
      </transition>
    </router-view>
  </div>
</template>

<style scoped></style>
