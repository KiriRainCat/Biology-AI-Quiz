<script setup lang="ts">
import TopicCard from "@/components/TopicCard.vue";

function resultExists(): boolean {
  const storedResult = localStorage.getItem("selectedOptions");
  const prevQuestions = localStorage.getItem("prevQuestions");
  return storedResult != null && prevQuestions != null && storedResult.length > 0 && prevQuestions.length > 0;
}
</script>

<template>
  <div class="text-center h-screen flex justify-center items-center">
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
        <var-button text v-if="resultExists()" @click="() => $router.push('/result')" size="mini" class="text-sm text-blue-500">
          View Previous Results
        </var-button>

        <var-button type="info" @click="() => $router.push('/leaderboard')" class="my-3">Leaderboard</var-button>
        <var-button type="primary" @click="() => $router.push('/quiz')" class="font-bold">Click to Begin</var-button>
      </div>
    </var-paper>
  </div>
</template>
