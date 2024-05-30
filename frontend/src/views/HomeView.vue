<template>
  <main class="text-center bg-gray-100 h-screen flex justify-center items-center">
    <var-paper elevation="2" radius="6" class="p-6 min-w-[360px] max-w-[540px]">
      <!-- 标题 -->
      <div class="font-bold text-2xl mb-2">AP Biology AI Randomized Quiz</div>
      <div class="font-medium">A quiz fulfilled with questions related to the following topics:</div>

      <!-- 当前涵盖的主题列表 -->
      <div class="flex gap-2 my-2">
        <TopicCard v-for="topic in ['Mitosis', 'Meiosis']" :topic />
      </div>

      <!-- 按钮栏 -->
      <div class="mt-5 flex flex-col items-center">
        <var-button type="info" @click="$router.push('/leaderboard')" class="my-3">Leaderboard</var-button>
        <var-button type="primary" @click="$router.push('/quiz')" class="font-bold">Click to Begin</var-button>
      </div>
    </var-paper>
  </main>
</template>

<script setup lang="ts">
import TopicCard from "@/components/TopicCard.vue";
</script>
