<script setup lang="ts">
import type { Question } from "@/pkg/services/api_types";
import { ref } from "vue";

const questions = ref<Question[]>(JSON.parse(localStorage.getItem("prevQuestions") ?? "[]"));
const selectedOptions = ref<number[]>(JSON.parse(localStorage.getItem("selectedOptions") ?? "[]"));

const currentIdx = ref(0);

const colorClasses = ["bg-white", "bg-green-300", "bg-red-300"];
function computeColorClass(i: number): string {
  const answer = questions.value[currentIdx.value].answer;
  if (i == answer) return colorClasses[1];

  const selected = selectedOptions.value[currentIdx.value];
  if (selected == i && selected != answer) return colorClasses[2];
  return colorClasses[0];
}
</script>

<template>
  <div class="flex h-screen items-center justify-center bg-gray-100">
    <div class="bg-white flex-none container relative overflow-hidden shadow-lg rounded-lg px-12 py-6">
      <!-- 返回按钮 -->
      <var-button text round @click="$router.push('/')" class="text-red-300 absolute top-2.5 right-2.5 w-8 h-8">×</var-button>

      <div class="relative z-20 mt-8">
        <!-- 问题 -->
        <div class="rounded-lg bg-gray-100 p-2 mt-5 neo-morph-1 text-center font-bold">
          <div class="bg-white p-5 text-gray-800">{{ questions[currentIdx].question }}</div>
        </div>

        <!-- 选项 -->
        <div class="mt-8">
          <div class="option" v-for="(option, i) in questions[currentIdx].choices">
            <div class="rounded-lg p-3 w-ful" :class="computeColorClass(i)">{{ option }}</div>
          </div>
        </div>

        <!-- 进度条 -->
        <div class="mt-8 flex justify-center gap-3">
          <var-button text size="mini" class="text-base" :disabled="currentIdx <= 0" @click="() => currentIdx--">←</var-button>
          <div class="text-center">
            <div class="h-1 w-12 bg-gray-800 rounded-full mx-auto"></div>
            <div class="font-bold text-sm text-gray-800">{{ currentIdx + 1 }} / {{ questions.length }}</div>
          </div>
          <var-button text size="mini" class="text-base" :disabled="currentIdx >= questions.length - 1" @click="currentIdx++">→</var-button>
        </div>
      </div>
    </div>
  </div>
</template>
